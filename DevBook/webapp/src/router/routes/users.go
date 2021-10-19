package routes

import (
	"net/http"
	"webapp/src/controllers"
)

// either / or /login points to login
var userRoutes = []Route{
	{
		URI:                    "/signup",
		Method:                 http.MethodGet,
		Function:               controllers.LoadSignUpPage,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/fetch-users",
		Method:                 http.MethodGet,
		Function:               controllers.LoadFetchUsersPage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUserProfilePage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/unfollow-user",
		Method:                 http.MethodPost,
		Function:               controllers.UnfollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/follow-user",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		RequiresAuthentication: true,
	},
}
