package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

// User represents a user in the social network
type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Nick      string    `json:"nick"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	Followers []User    `json:"followers"` // different from User of API. we'll get them via concurrent tasks
	Following []User    `json:"following"` // different from User of API. we'll get them via concurrent tasks
	Posts     []Post    `json:"posts"`     // different from User of API. we'll get them via concurrent tasks
}

// #IMPORTANT
// FetchCompleteUser fetches 1. user data, 2. followers, 3. following and 4. posts concurrently
// we'll have 4 goroutines running in paralell and we'll communicate to them using channels
func FetchCompleteUser(userID uint64, webAppRequestWithCookies *http.Request) (User, error) {
	// create 4 channels
	userDataChannel := make(chan User)
	userFollowersChannel := make(chan []User)
	userFollowingChannel := make(chan []User)
	userPostsChannel := make(chan []Post)

	// creating goroutines in parallel
	go FetchUserData(userDataChannel, userID, webAppRequestWithCookies)
	go FetchUserFollowers(userFollowersChannel, userID, webAppRequestWithCookies)
	go FetchUserFollowing(userFollowingChannel, userID, webAppRequestWithCookies)
	go FetchUserPosts(userPostsChannel, userID, webAppRequestWithCookies)

	// we'll wait the requests finish (regardless the order) to build the complete user
	var (
		userData      User
		userFollowers []User
		userFollowing []User
		userPosts     []Post
	)

	// 4 times because we have 4 goroutines
	for i := 0; i < 4; i++ {
		select { // it's like a switch, but for concurrency
		// USER DATA
		case userDataLoaded := <-userDataChannel: // in case we have a value ready to be received by the channel
			log.Println("User data response is ready")
			if userDataLoaded.ID == 0 { // this means we had error in the request (remember User{}. 0 is value zero for uint64)
				return User{}, errors.New("error when fetching user data") // we return empty user and an error
			}
			userData = userDataLoaded

		// USER FOLLOWERS
		case userFollowersLoaded := <-userFollowersChannel:
			log.Println("User followers response is ready")
			if userFollowersLoaded == nil {
				return User{}, errors.New("error when fetching user followers")
			}
			userFollowers = userFollowersLoaded

		// USER FOLLOWING
		case userFollowingLoaded := <-userFollowingChannel:
			log.Println("User following response is ready")
			if userFollowingLoaded == nil {
				return User{}, errors.New("error when fetching user following")
			}
			userFollowing = userFollowingLoaded

		// USER POSTS
		case userPostsLoaded := <-userPostsChannel:
			log.Println("User posts response is ready")
			if userPostsLoaded == nil {
				return User{}, errors.New("error when fetching user posts")
			}
			userPosts = userPostsLoaded
		}
	}

	// if we reached this point this means everything was ok (all 4 requests) so we just need to assemble the following, followers and posts to the user basic data
	userData.Followers = userFollowers
	userData.Following = userFollowing
	userData.Posts = userPosts
	return userData, nil
}

// FetchUserData calls API to fetch user data and writes it back to a channel
func FetchUserData(userDataChannel chan<- User, userID uint64, webAppRequestWithCookies *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, error := requests.MakeRequestWithAuthenticationData(webAppRequestWithCookies, http.MethodGet, url, nil)
	if error != nil {
		// if we have an error, we just return an empty user. We check for it later at FetchCompleteUser
		userDataChannel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if error = json.NewDecoder(response.Body).Decode(&user); error != nil {
		userDataChannel <- User{}
		return
	}
	// if everything is alright, send user back to channel
	userDataChannel <- user
}

// FetchUserFollowers calls API to fetch user followers
func FetchUserFollowers(userFollowersChannel chan<- []User, userID uint64, webAppRequestWithCookies *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	response, error := requests.MakeRequestWithAuthenticationData(webAppRequestWithCookies, http.MethodGet, url, nil)
	if error != nil {
		userFollowersChannel <- nil // value 0 for slice is nil.
		return
	}
	defer response.Body.Close()

	var userFollowers []User
	if error = json.NewDecoder(response.Body).Decode(&userFollowers); error != nil {
		userFollowersChannel <- nil
		return
	}

	// #IMPORTANT we cannot push nil to a channel, so we push an empty slice in case API response is nil. Another way is to return in the API.
	if userFollowers == nil {
		userFollowersChannel <- make([]User, 0)
		return
	}
	userFollowersChannel <- userFollowers

}

// FetchUserFollowing calls API to fetch user following
func FetchUserFollowing(userFollowingChannel chan<- []User, userID uint64, webAppRequestWithCookies *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	response, error := requests.MakeRequestWithAuthenticationData(webAppRequestWithCookies, http.MethodGet, url, nil)
	if error != nil {
		userFollowingChannel <- nil // value 0 for slice is nil.
		return
	}
	defer response.Body.Close()

	var userFollowing []User
	if error = json.NewDecoder(response.Body).Decode(&userFollowing); error != nil {
		userFollowingChannel <- nil
		return
	}

	// #IMPORTANT we cannot push nil to a channel, so we push an empty slice in case API response is nil. Another way is to return in the API.
	if userFollowing == nil {
		userFollowingChannel <- make([]User, 0)
		return
	}
	userFollowingChannel <- userFollowing
}

// FetchUserPosts calls API to fetch user posts
func FetchUserPosts(userPostsChannel chan<- []Post, userID uint64, webAppRequestWithCookies *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.APIURL, userID)
	response, error := requests.MakeRequestWithAuthenticationData(webAppRequestWithCookies, http.MethodGet, url, nil)
	if error != nil {
		userPostsChannel <- nil // value 0 for slice is nil.
		return
	}
	defer response.Body.Close()

	var userPosts []Post
	if error = json.NewDecoder(response.Body).Decode(&userPosts); error != nil {
		userPostsChannel <- nil
		return
	}
	// #IMPORTANT we cannot push nil to a channel, so we push an empty slice in case API response is nil. Another way is to return in the API.
	if userPosts == nil {
		userPostsChannel <- make([]Post, 0)
		return
	}
	userPostsChannel <- userPosts
}
