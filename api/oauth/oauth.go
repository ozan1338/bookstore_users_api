package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	resError "users_api/utils/errors"

	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 200 * time.Millisecond,
	}
)

const (
	headerXPublic = "X-Public"
	headerXClientId = "X-Client-Id"
	headerXUserId = "X-User-Id"
	paramterAccessToken = "access_token"
)

type accessToken struct {
	Id string `json:"id"`
	UserId int64 `json:"user_id"`
	ClientId int64 `json:"client_id"`
}


func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}

	return request.Header.Get(headerXPublic) == "true"
}

func GetUserId(request *http.Request) int64 {
	if request == nil {
		return 0
	}

	userId, err := strconv.ParseInt(request.Header.Get(headerXUserId), 10, 64)
	if err != nil {
		return 0
	}

	return userId
}

func GetCienId(request *http.Request) int64 {
	if request == nil {
		return 0
	}

	clientId, err := strconv.ParseInt(request.Header.Get(headerXClientId), 10, 64)
	if err != nil {
		return 0
	}

	return clientId
}

func AuthenticateRequest(request *http.Request) *resError.RestError {
	if request == nil {
		return nil
	}

	fmt.Println("THIS IT")

	cleanRequest(request)

	//http://localhost:8081/resource?access_token=abc123
	accessTokenId := strings.TrimSpace(request.URL.Query().Get(paramterAccessToken))

	if accessTokenId == "" {
		return nil
	}

	at, err := getAccessToken(accessTokenId)
	if err != nil {
		// fmt.Println("NIH EROR",err)
		if err.Status == http.StatusNotFound{
			return nil
		}
		return err
	}

	// fmt.Println("WOW",at)

	request.Header.Add(headerXUserId, fmt.Sprintf("%v",at.UserId))
	request.Header.Add(headerXClientId, fmt.Sprintf("%v",at.ClientId))

	return nil
}

func cleanRequest(request *http.Request) {
	if request == nil {
		return
	}

	request.Header.Del(headerXClientId)
	request.Header.Del(headerXUserId)
}

func getAccessToken(accessTokenId string) (*accessToken, *resError.RestError) {
	response := oauthRestClient.Get(fmt.Sprintf("/oauth/access_token/%s",accessTokenId))

	if response == nil || response.Response == nil {
		return nil, resError.NewInternalServerError("invalid rest client response when to try get access token")
	}

	if response.StatusCode > 399 {
		// fmt.Println("HERERERE")
		var respError resError.RestError
		if err := json.Unmarshal(response.Bytes(), &respError); err != nil {
			return nil, resError.NewInternalServerError("invalid err interface when trying to get access token")
		}
		return nil ,&respError
	}

	var at accessToken
	if err := json.Unmarshal(response.Bytes(), &at); err != nil {
		return nil, resError.NewInternalServerError("error when trying to unmarshall access token response")
	}
	return &at,nil
}