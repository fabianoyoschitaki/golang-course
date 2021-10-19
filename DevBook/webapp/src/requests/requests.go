package requests

import (
	"io"
	"log"
	"net/http"
	"webapp/src/cookies"
)

// MakeRequestWithAuthenticationData is used to put the token in the request to the backend API
func MakeRequestWithAuthenticationData(webappRequestWithCookie *http.Request, method, url string, data io.Reader) (*http.Response, error) {

	// create the actual request to be made to the authenticated API route
	authenticatedAPIRequest, error := http.NewRequest(method, url, data)
	if error != nil {
		return nil, error
	}

	// we don't need to check the error since the middleware already checked it for us
	cookieValues, _ := cookies.ReadCookie(webappRequestWithCookie)

	// we get the token and add it as a Authorization Bearer
	authenticatedAPIRequest.Header.Add("Authorization", "Bearer "+cookieValues["token"])
	log.Printf("Creating authenticated request for URL: %s, method: %s for user %s and token %s",
		url, method, cookieValues["id"], cookieValues["token"])

	// create http client and make the request
	httpClient := &http.Client{}
	authenticatedAPIResponse, error := httpClient.Do(authenticatedAPIRequest)
	if error != nil {
		return nil, error
	}

	// return response
	return authenticatedAPIResponse, nil
}
