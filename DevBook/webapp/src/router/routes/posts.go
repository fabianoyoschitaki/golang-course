package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postsRoutes = []Route{
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}/likes",
		Method:                 http.MethodPost,
		Function:               controllers.LikePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}/unlikes",
		Method:                 http.MethodPost,
		Function:               controllers.UnlikePost,
		RequiresAuthentication: true,
	},
}
