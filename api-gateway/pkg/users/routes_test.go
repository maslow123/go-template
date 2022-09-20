package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/library-api-gateway/pkg/users/pb"
	"github.com/maslow123/library-api-gateway/pkg/utils"
	"github.com/stretchr/testify/require"
)

type ServiceError struct {
	Error string
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		name          string
		body          gin.H
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": "user1",
				"password": "111111",
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Required username",
			body: gin.H{
				"username": "",
				"password": "111111",
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				loginError(t, recorder.Body, "Unknown desc = Key: 'LoginRequest.Username'")
			},
		},
		{
			name: "Required password",
			body: gin.H{
				"username": "user1",
				"password": "",
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				loginError(t, recorder.Body, "Unknown desc = Key: 'LoginRequest.Password'")
			},
		},
		{
			name: "User not found",
			body: gin.H{
				"username": "invalid user",
				"password": "wrong password",
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				loginError(t, recorder.Body, "user-not-found")
			},
		},
		{
			name: "Wrong password",
			body: gin.H{
				"username": "user1",
				"password": "wrong password",
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				loginError(t, recorder.Body, "wrong-password")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := NewServer(t)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestRegister(t *testing.T) {
	testCases := []struct {
		name          string
		body          gin.H
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":         utils.RandomString("name-", 10),
				"username":     utils.RandomString("username-", 10),
				"password":     utils.RandomString("password-", 10),
				"address":      utils.RandomString("address-", 30),
				"phone_number": fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				checkResponseRegister(t, recorder.Body, true, "")
			},
		},
		{
			name: "Username already exists",
			body: gin.H{
				"name":         utils.RandomString("name-", 10),
				"username":     "user1",
				"password":     utils.RandomString("password-", 10),
				"address":      utils.RandomString("address-", 30),
				"phone_number": fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				checkResponseRegister(t, recorder.Body, false, "username-already-exists")
			},
		},
		{
			name: "Required Name",
			body: gin.H{
				"name":         "",
				"username":     utils.RandomString("username-", 10),
				"password":     utils.RandomString("password-", 10),
				"address":      utils.RandomString("address-", 30),
				"phone_number": fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				checkResponseRegister(t, recorder.Body, false, "Unknown desc = Key: 'RegisterRequest.Name'")
			},
		},
		{
			name: "Required Username",
			body: gin.H{
				"name":         utils.RandomString("name-", 10),
				"username":     "",
				"password":     utils.RandomString("password-", 10),
				"address":      utils.RandomString("address-", 30),
				"phone_number": fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				checkResponseRegister(t, recorder.Body, false, "Unknown desc = Key: 'RegisterRequest.Username'")
			},
		},
		{
			name: "Required Password",
			body: gin.H{
				"name":         utils.RandomString("name-", 10),
				"username":     utils.RandomString("username-", 10),
				"password":     "",
				"address":      utils.RandomString("address-", 30),
				"phone_number": fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				checkResponseRegister(t, recorder.Body, false, "Unknown desc = Key: 'RegisterRequest.Password'")
			},
		},
		{
			name: "Required Address",
			body: gin.H{
				"name":         utils.RandomString("name-", 10),
				"username":     utils.RandomString("username-", 10),
				"password":     utils.RandomString("password-", 10),
				"address":      "",
				"phone_number": fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				checkResponseRegister(t, recorder.Body, false, "Unknown desc = Key: 'RegisterRequest.Address'")
			},
		},

		{
			name: "Required Phone Number",
			body: gin.H{
				"name":         utils.RandomString("name-", 10),
				"username":     utils.RandomString("username-", 10),
				"password":     utils.RandomString("password-", 10),
				"address":      utils.RandomString("address-", 30),
				"phone_number": "",
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				checkResponseRegister(t, recorder.Body, false, "Unknown desc = Key: 'RegisterRequest.PhoneNumber'")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := NewServer(t)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users/register"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}

func checkResponseRegister(
	t *testing.T,
	body *bytes.Buffer,
	isSuccess bool,
	errorMessage string,
) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	if isSuccess {
		var resp pb.RegisterResponse

		err = json.Unmarshal(data, &resp)
		require.NoError(t, err)
		require.NotZero(t, resp.Id)
		require.Equal(t, "", errorMessage)
		return
	}

	var resp ServiceError
	err = json.Unmarshal(data, &resp)
	require.NoError(t, err)

	log.Println("Error: ", resp.Error)
	require.Contains(t, resp.Error, errorMessage)

}

func loginError(t *testing.T, body *bytes.Buffer, errorMessage string) {

	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var resp ServiceError
	err = json.Unmarshal(data, &resp)
	require.NoError(t, err)

	require.Contains(t, resp.Error, errorMessage)
}
