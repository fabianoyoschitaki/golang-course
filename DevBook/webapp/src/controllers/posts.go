package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

// LikePost calls API to like a post
func LikePost(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	// creating URL to our like route in the API
	likePostAPIUrl := fmt.Sprintf("%s/posts/%d/likes", config.APIURL, postID)
	log.Printf("Adding like to post %d (url: %s)\n", postID, likePostAPIUrl)

	// making actual request
	response, error := requests.MakeRequestWithAuthenticationData(r, http.MethodPost, likePostAPIUrl, nil)
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, response)
		return
	}

	responses.JSON(rw, response.StatusCode, nil)
}

// UnlikePost calls API to like a post
func UnlikePost(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	// creating URL to our unlike route in the API
	unlikePostAPIUrl := fmt.Sprintf("%s/posts/%d/likes", config.APIURL, postID)
	log.Printf("Removing like to post %d (url: %s)\n", postID, unlikePostAPIUrl)

	// making actual request
	response, error := requests.MakeRequestWithAuthenticationData(r, http.MethodDelete, unlikePostAPIUrl, nil)
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, response)
		return
	}

	responses.JSON(rw, response.StatusCode, nil)
}

// CreatePost calls API to create a new post
func CreatePost(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// getting post data from request form values
	newPost, error := json.Marshal(map[string]string{
		"title":   r.FormValue("newPostTitle"),
		"content": r.FormValue("newPostContent"),
	})
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	// we could get the new post variables, let's prepare to call the API
	newPostUrl := fmt.Sprintf("%s/posts", config.APIURL)

	// authenticated request
	response, error := requests.MakeRequestWithAuthenticationData(r, http.MethodPost, newPostUrl, bytes.NewBuffer(newPost))
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, response)
		return
	}

	// we don't need to return anything (the new post), this is because our front-end javascript will reload the posts again (posts.js) window.location = "/home"
	responses.JSON(rw, response.StatusCode, nil)
}

// UpdatePost updates an existing post
func UpdatePost(rw http.ResponseWriter, r *http.Request) {

	// getting postId to update
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	// getting post data from request form values
	r.ParseForm()
	postUpdate, error := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	// making request to our API
	updatePostUrl := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
	response, error := requests.MakeRequestWithAuthenticationData(r, http.MethodPut, updatePostUrl, bytes.NewBuffer(postUpdate))
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, response)
		return
	}

	responses.JSON(rw, response.StatusCode, nil)
}

// DeletePost updates an existing post
func DeletePost(rw http.ResponseWriter, r *http.Request) {

	// getting postId to delete
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	// making request to our API
	deletePostUrl := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
	response, error := requests.MakeRequestWithAuthenticationData(r, http.MethodDelete, deletePostUrl, nil)
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, response)
		return
	}

	responses.JSON(rw, response.StatusCode, nil)
}
