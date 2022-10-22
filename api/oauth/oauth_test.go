package oauth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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