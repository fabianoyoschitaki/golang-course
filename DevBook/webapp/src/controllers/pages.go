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
)

// LoadLoginPage renders login page LOL
func LoadLoginPage(rw http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(rw, "login.html", nil)
}

// LoadSignUpPage renders create user page
func LoadSignUpPage(rw http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(rw, "signup.html", nil)
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
