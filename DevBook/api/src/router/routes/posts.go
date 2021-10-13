package routes

import (
	"api/src/controllers"
	"net/http"
)

var routePosts = []Route{
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts",
		Method:                 http.MethodGet,
		Function:               controllers.FetchPosts,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodGet,
		Function:               controllers.FetchPost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/posts",
		Method:                 http.MethodGet,
		Function:               controllers.FetchPostsByUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}/likes",
		Method:                 http.MethodPost,
		Function:               controllers.LikePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}/likes",
		Method:                 http.MethodDelete,
		Function:               controllers.UnlikePost,
		RequiresAuthentication: true,
	},
}
