package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// LoadLoginPage renders login page LOL
func LoadLoginPage(rw http.ResponseWriter, r *http.Request) {
	// if user is authenticated, we redirect him to /home page
	cookie, _ := cookies.ReadCookie(r)
	if cookie["token"] != "" {
		http.Redirect(rw, r, "/home", http.StatusFound)
		return
	}

	// if user is not authenticated (cookies["token"] is empty, then load login.html)
	utils.RenderTemplate(rw, "login.html", nil)
}

// LoadSignUpPage renders create user page
func LoadSignUpPage(rw http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(rw, "signup.html", nil)
}

// LoadUpdatePostPage fetches post through API and loads update page
func LoadUpdatePostPage(rw http.ResponseWriter, r *http.Request) {
	log.Println("LoadUpdatePostPage pages.go controller")

	// get post id
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	// we need to fetch the post before updating it
	getPostByIDUrl := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
	response, error := requests.MakeRequestWithAuthenticationData(r, http.MethodGet, getPostByIDUrl, nil)
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, response)
		return
	}

	// we get the Post being returned and return it to FE
	var post models.Post
	if error := json.NewDecoder(response.Body).Decode(&post); error != nil {
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.APIError{Error: error.Error()})
		return
	}

	utils.RenderTemplate(rw, "update-post.html", post)
}

// LoadFetchUsersPage loads page with users based on the search term
func LoadFetchUsersPage(rw http.ResponseWriter, r *http.Request) {
	log.Println("LoadFetchUsersPage pages.go controller")

	// getting query parameter
	nameOrNick := r.URL.Query().Get("user-search")
	fetchUsersAPIUrl := fmt.Sprintf("%s/users?user=%s", config.APIURL, nameOrNick)
	response, error := requests.MakeRequestWithAuthenticationData(r, http.MethodGet, fetchUsersAPIUrl, nil)
	if error != nil {
		log.Printf("Error: %s", error)
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	// if http status code is not success
	if response.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, response)
		return
	}

	log.Printf("LoadFetchUsersPage response code is %d", response.StatusCode)

	// users to be returned
	var usersFound []models.User
	if error = json.NewDecoder(response.Body).Decode(&usersFound); error != nil {
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.APIError{Error: error.Error()})
		return
	}

	log.Printf("LoadFetchUsersPage found %d users with search term: %s", len(usersFound), nameOrNick)
	// now lets execute the templace passing the users
	utils.RenderTemplate(rw, "users.html", struct {
		UsersFound []models.User
		SearchTerm string
	}{
		UsersFound: usersFound,
		SearchTerm: nameOrNick,
	})
}

// LoadHomepage renders main homepage with posts
func LoadHomepage(rw http.ResponseWriter, r *http.Request) {
	log.Println("LoadHomepage pages.go controller")
	// we need to get the posts of the authenticated user
	postsUrl := fmt.Sprintf("%s/posts", config.APIURL)

	// instead of "postsResponse, error := http.Get(postsUrl)", we use the authenticated request
	postsResponse, error := requests.MakeRequestWithAuthenticationData(r, http.MethodGet, postsUrl, nil)
	// fmt.Println(postsResponse.StatusCode, error)

	// let's handle the possible errors: error itself and statusCode >= 400
	if error != nil {
		// #IMPORTANT we cannot use response.statusCode because if error != nil, response doesn't have it! (nil) we'll have a PANIC
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer postsResponse.Body.Close()

	// if http status code is not success
	if postsResponse.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, postsResponse)
		return
	}

	// now we can read the response body (the user posts)
	var userPosts []models.Post
	if error := json.NewDecoder(postsResponse.Body).Decode(&userPosts); error != nil {
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.APIError{Error: error.Error()})
		return
	}

	// we don't need the error because if we're here, we've passed through the middleware.
	cookie, _ := cookies.ReadCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	log.Printf("User ID from cookie is: %d\n", userID)

	// we render the page passing the userPosts. this is a slice
	// utils.RenderTemplate(rw, "home.html", userPosts) // this could be used, but let's make a way to pass more parameters
	utils.RenderTemplate(rw, "home.html", struct {
		Posts        []models.Post
		LoggedUserID uint64
		Now          time.Time
	}{
		Posts:        userPosts,
		LoggedUserID: userID,
		Now:          time.Now(),
	})
}

// #IMPORTANT: we get userID and make 4 concurrent requests to the API: user data, posts, followers, following
// LoadUserProfilePage loads a user profile page
func LoadUserProfilePage(rw http.ResponseWriter, r *http.Request) {
	// get query parameter userId
	parameters := mux.Vars(r)
	userID, error := strconv.ParseUint(parameters["userId"], 10, 64)
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	// we want to see if userComplete.followers "contains" loggedUser. if not, then we create a "Follow" button or "Unfollow" otherwise
	cookie, _ := cookies.ReadCookie(r)
	loggedUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	log.Printf("Logged user is: %d", loggedUserID)

	// if user being fetched is the same as the logged user ID, then we redirect to /profile instead of /user
	if loggedUserID == userID {
		http.Redirect(rw, r, "/profile", http.StatusFound)
		return
	}

	// we search for the complete user by userID and the Request r. We use r because we need to use the authenticated token to call the API
	userComplete, error := models.FetchCompleteUser(userID, r)
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}

	// we return the complete user of the profile we're visiting and also the logged User ID to create the "follow/unfollow" feature
	utils.RenderTemplate(rw, "user.html", struct {
		UserComplete models.User
		LoggedUserID uint64
	}{
		UserComplete: userComplete,
		LoggedUserID: loggedUserID,
	})
}

// LoadLoggedUserProfilePage loads profile page
func LoadLoggedUserProfilePage(rw http.ResponseWriter, r *http.Request) {

	// get logged user ID
	cookie, _ := cookies.ReadCookie(r)
	loggedUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	log.Printf("Logged user is: %d", loggedUserID)

	// we search for the complete user by userID and the Request r. We use r because we need to use the authenticated token to call the API
	userComplete, error := models.FetchCompleteUser(loggedUserID, r)
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}

	// we return the complete user of the profile
	utils.RenderTemplate(rw, "profile.html", userComplete)
}

// LoadEditProfilePage loads edit profile page
func LoadEditProfilePage(rw http.ResponseWriter, r *http.Request) {

	// get logged user ID
	cookie, _ := cookies.ReadCookie(r)
	loggedUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	log.Printf("Logged user is: %d", loggedUserID)

	// let's reuse channel to fetch user data. It's not common to use
	// a channel if there are no other concurrent task, but let's use it instead of creating another function
	userDataChannel := make(chan models.User)
	go models.FetchUserData(userDataChannel, loggedUserID, r)
	user := <-userDataChannel

	// it means we had an error
	if user.ID == 0 {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: "error when fetching user data"})
		return
	}

	// now we can load the page with the user from channel
	utils.RenderTemplate(rw, "edit-profile.html", user)
}
