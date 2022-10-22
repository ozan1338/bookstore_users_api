package oauth

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	rest.StartMockupServer()
	os.Exit(m.Run())
}
func TestOauthConstant(t *testing.T) {
	assert.EqualValues(t, headerXPublic, "X-Public")
	assert.EqualValues(t, headerXClientId, "X-Client-Id")
	assert.EqualValues(t, headerXUserId, "X-User-Id")
	assert.EqualValues(t, paramterAccessToken, "access_token")
}

func TestIsPublicNilRequest(t *testing.T) {
	assert.True(t, IsPublic(nil))
}

func TestIsPublicError(t *testing.T) {
	request := http.Request{
		Header:make(http.Header),
	}
	assert.False(t,IsPublic(&request))

	request.Header.Add("X-Public", "true")
	assert.True(t,IsPublic(&request))
}

func TestGetUserIdNil(t *testing.T) {
	assert.Zero(t, GetUserId(nil))
}

func TestGetUserIdError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}

	request.Header.Add(headerXUserId, "")

	assert.Zero(t, GetUserId(&request))
}

func TestGetUserNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}

	request.Header.Add(headerXUserId, "12")

	assert.NotZero(t, GetUserId(&request))
	assert.Equal(t, GetUserId(&request), int64(12))
}

func TestGetClientIdNil(t *testing.T) {
	assert.Zero(t, GetCienId(nil))
}

func TestGetClientIdError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}

	request.Header.Add(headerXClientId, "")
	assert.Zero(t, GetCienId(&request))
}

func TestGetClientIdNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}

	request.Header.Add(headerXClientId, "23")

	assert.NotZero(t, GetCienId(&request))
	assert.Equal(t, GetCienId(&request), int64(23))
}

func TestCleanRequestNil(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}

	request.Header.Add(headerXClientId, "23")
	request.Header.Add(headerXUserId, "12")

	cleanRequest(nil)

	assert.NotEmpty(t,request.Header.Get(headerXClientId))
	assert.Equal(t,request.Header.Get(headerXClientId), "23")
	assert.NotEmpty(t,request.Header.Get(headerXUserId))
	assert.Equal(t,request.Header.Get(headerXUserId), "12")
}

func TestCleanRequestNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	request.Header.Add(headerXClientId, "23")
	request.Header.Add(headerXUserId, "12")

	cleanRequest(&request)

	assert.Empty(t,request.Header.Get(headerXClientId))
	assert.NotEqual(t,request.Header.Get(headerXClientId), "23")
	assert.Equal(t,request.Header.Get(headerXClientId), "")
	assert.Empty(t,request.Header.Get(headerXUserId))
	assert.NotEqual(t,request.Header.Get(headerXUserId), "12")
	assert.Equal(t,request.Header.Get(headerXUserId), "")


}

func TestGetAccessTokenInvalidClientResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodGet,
		URL: "http://localhost:8081/oauth/access_token/123",
		ReqBody: ``,
		RespHTTPCode: 0,
		RespBody: `{}`,
	})

	access_token,err := getAccessToken("123")
	assert.Nil(t,access_token)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.GetStatus())
	assert.EqualValues(t, "invalid rest client response when to try get access token", err.GetMessage())
}

func TestGetAccessTokenInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodGet,
		URL: "http://localhost:8081/oauth/access_token/123",
		ReqBody: ``,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message": "invalid login credentials", "status": "404", "error": "not_found"}`,
	})

	access_token,err := getAccessToken("123")
	assert.Nil(t,access_token)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.GetStatus())
	assert.EqualValues(t, "invalid error interface when trying to get access token", err.GetMessage())
}

func TestGetAccessTokenNotFounf(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodGet,
		URL: "http://localhost:8081/oauth/access_token/123",
		ReqBody: ``,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message": "not found", "status": 404, "error": "not_found"}`,
	})

	access_token,err := getAccessToken("123")
	assert.Nil(t,access_token)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.GetStatus())
	assert.EqualValues(t, "not found", err.GetMessage())
}

func TestGetAccessTokenInvalidAtResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodGet,
		URL: "http://localhost:8081/oauth/access_token/123",
		ReqBody: ``,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id": "123", "user_id": "123", "client_id": "123"}`,
	})

	access_token,err := getAccessToken("123")
	assert.Nil(t,access_token)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.GetStatus())
	assert.EqualValues(t, "error when trying to unmarshall access token response", err.GetMessage())
}

func TestGetAccessTokenNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodGet,
		URL: "http://localhost:8081/oauth/access_token/123",
		ReqBody: ``,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id": "123", "user_id": 123, "client_id": 123}`,
	})

	access_token,err := getAccessToken("123")
	assert.NotNil(t,access_token)
	assert.Nil(t, err)
	assert.EqualValues(t,access_token.Id,"123")
	assert.EqualValues(t,access_token.UserId,123)
	assert.EqualValues(t,access_token.ClientId,123)
}