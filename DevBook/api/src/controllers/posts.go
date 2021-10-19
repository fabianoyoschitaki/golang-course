package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePost creates a new post
func CreatePost(rw http.ResponseWriter, r *http.Request) {
	// ID - comes from database
	// AutorID - comes from token
	// Likes - 0
	// we need only Title and Content
	userID, error := authentication.ExtractUserId(r)
	if error != nil {
		responses.Error(rw, http.StatusUnauthorized, error)
		return
	}

	// get Post from request body
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(rw, http.StatusUnprocessableEntity, error)
		return
	}

	// Create Post to be saved to database
	var newPost models.Post
	if error = json.Unmarshal(requestBody, &newPost); error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}
	newPost.AuthorID = userID // authorID comes from token

	// validate Post
	if error = newPost.Prepare(); error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	// open connection to database
	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	// create repository to save new Post
	postRepository := repositories.NewPostsRepository(db)

	// assigns post ID to object being returned to client
	newPost.ID, error = postRepository.Create(newPost)
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(rw, http.StatusCreated, newPost)
}

// FetchPosts fetches all user's posts
func FetchPosts(rw http.ResponseWriter, r *http.Request) {
	// get logged user ID
	userID, error := authentication.ExtractUserId(r)
	if error != nil {
		responses.Error(rw, http.StatusUnauthorized, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostsRepository(db)
	posts, error := postRepository.FetchPosts(userID)
	log.Printf("Fetching posts for user %d: %d posts", userID, len(posts))
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(rw, http.StatusOK, posts)
}

// FetchPost fetches a specific post
func FetchPost(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostsRepository(db)
	post, error := postRepository.FetchPostByID(postID)
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(rw, http.StatusOK, post)
}

// UpdatePost updates an existing post
func UpdatePost(rw http.ResponseWriter, r *http.Request) {
	// get logged user ID
	userID, error := authentication.ExtractUserId(r)
	if error != nil {
		responses.Error(rw, http.StatusUnauthorized, error)
		return
	}

	// get post ID
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	// we need to open database connection to make sure the post ID comes from user ID
	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostsRepository(db)
	postFromDatabase, error := postRepository.FetchPostByID(postID)
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	// checking if it's different from user
	if postFromDatabase.AuthorID != userID {
		responses.Error(rw, http.StatusForbidden, errors.New("it's not possible to update a post you don't own"))
		return
	}

	// now that we know user is the owner of post, we can read the request body and update the post
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(rw, http.StatusUnprocessableEntity, error)
		return
	}

	var postToBeUpdated models.Post
	if error = json.Unmarshal(requestBody, &postToBeUpdated); error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	// validating business rules
	if error = postToBeUpdated.Prepare(); error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	if error = postRepository.UpdatePost(postID, postToBeUpdated); error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(rw, http.StatusNoContent, nil)

}

// DeletePost deletes an existing post
func DeletePost(rw http.ResponseWriter, r *http.Request) {
	// get logged user ID
	userID, error := authentication.ExtractUserId(r)
	if error != nil {
		responses.Error(rw, http.StatusUnauthorized, error)
		return
	}

	// get post ID
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	// we need to open database connection to make sure the post ID comes from user ID
	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostsRepository(db)
	postFromDatabase, error := postRepository.FetchPostByID(postID)
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	// checking if it's different from user
	if postFromDatabase.AuthorID != userID {
		responses.Error(rw, http.StatusForbidden, errors.New("it's not possible to delete a post you don't own"))
		return
	}

	if error = postRepository.DeletePost(postID); error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(rw, http.StatusNoContent, nil)
}

// FetchPostsByUser fetches all posts by user
func FetchPostsByUser(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, error := strconv.ParseUint(parameters["userId"], 10, 64)
	if error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostsRepository(db)
	userPosts, error := postRepository.FetchPostByUserID(userID)
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	log.Println(userPosts == nil)
	responses.JSON(rw, http.StatusOK, userPosts)
}

// LikePost adds a like to a post
func LikePost(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostsRepository(db)
	if error = postRepository.LikePost(postID); error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(rw, http.StatusNoContent, nil)
}

// UnlikePost removes a like to a post
func UnlikePost(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostsRepository(db)
	if error = postRepository.UnlikePost(postID); error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(rw, http.StatusNoContent, nil)
}
