package services

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/maslow123/library-users/pkg/pb"
	"github.com/maslow123/library-users/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	username := utils.RandomString("username-", 10)
	testCases := []struct {
		name      string
		isSuccess bool
		req       *pb.RegisterRequest
		err       string
	}{
		{
			"OK",
			true,
			&pb.RegisterRequest{
				Name:        utils.RandomString("name-", 10),
				Username:    username,
				Password:    utils.RandomString("", 10),
				Address:     utils.RandomString("address-", 50),
				PhoneNumber: fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			"",
		},
		{
			"Username already exists",
			false,
			&pb.RegisterRequest{
				Name:        utils.RandomString("name-", 10),
				Username:    username,
				Password:    utils.RandomString("", 10),
				Address:     utils.RandomString("address-", 50),
				PhoneNumber: fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			"username-already-exists",
		},
		{
			"Required Name",
			false,
			&pb.RegisterRequest{
				Name:        "",
				Username:    username,
				Password:    utils.RandomString("", 10),
				Address:     utils.RandomString("address-", 50),
				PhoneNumber: fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			"Key: 'RegisterRequest.Name'",
		},
		{
			"Required Username",
			false,
			&pb.RegisterRequest{
				Name:        utils.RandomString("name-", 10),
				Username:    "",
				Password:    utils.RandomString("", 10),
				Address:     utils.RandomString("address-", 50),
				PhoneNumber: fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			"Key: 'RegisterRequest.Username'",
		},
		{
			"Required Password",
			false,
			&pb.RegisterRequest{
				Name:        utils.RandomString("name-", 10),
				Username:    username,
				Password:    "",
				Address:     utils.RandomString("address-", 50),
				PhoneNumber: fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			"Key: 'RegisterRequest.Password'",
		},
		{
			"Required Address",
			false,
			&pb.RegisterRequest{
				Name:        utils.RandomString("name-", 10),
				Username:    username,
				Password:    utils.RandomString("password-", 10),
				Address:     "",
				PhoneNumber: fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
			},
			"Key: 'RegisterRequest.Address'",
		},
		{
			"Required Phone Number",
			false,
			&pb.RegisterRequest{
				Name:        utils.RandomString("name-", 10),
				Username:    username,
				Password:    utils.RandomString("password-", 10),
				Address:     utils.RandomString("address-", 50),
				PhoneNumber: "",
			},
			"Key: 'RegisterRequest.PhoneNumber'",
		},
	}

	ctx := context.Background()
	conn := checkConnection(ctx, t)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.Register(ctx, tc.req)

			if tc.err != "" {
				require.True(t, strings.Contains(err.Error(), tc.err))
			}
			if tc.isSuccess {
				require.NotZero(t, resp.Id)
				require.NoError(t, err)

				return
			}

			require.Nil(t, resp)
			require.Error(t, err)
		})
	}
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		name         string
		errorMessage string
		isSuccess    bool
		req          *pb.LoginRequest
	}{
		{
			"OK",
			"",
			true,
			&pb.LoginRequest{},
		},
		{
			"Username is required",
			"Key: 'LoginRequest.Username",
			false,
			&pb.LoginRequest{
				Username: "",
				Password: utils.RandomString("password-", 5),
			},
		},
		{
			"Password is required",
			"Key: 'LoginRequest.Password",
			false,
			&pb.LoginRequest{
				Username: utils.RandomString("username-", 5),
				Password: "",
			},
		},
		{
			"User not found",
			"user-not-found",
			false,
			&pb.LoginRequest{
				Username: utils.RandomString("username-", 5),
				Password: utils.RandomString("password-", 5),
			},
		},
		{
			"Wrong password",
			"wrong-password",
			true,
			&pb.LoginRequest{},
		},
	}

	ctx := context.Background()
	conn := checkConnection(ctx, t)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			var req pb.LoginRequest
			var user *pb.RegisterRequest

			if tc.isSuccess {
				user = createRandomUser(t, client)

				req.Username = user.Username
				req.Password = user.Password

				if tc.errorMessage == "wrong-password" {
					req.Password = utils.RandomString("password-", 5)
					loginError(t, ctx, client, &req, tc.errorMessage)

					return
				}

				resp, err := client.Login(ctx, &req)

				require.NoError(t, err)
				require.NotEmpty(t, resp.Token)

				require.Equal(t, user.Name, resp.User.Name)
				require.Equal(t, user.Address, resp.User.Address)
				require.Equal(t, user.PhoneNumber, resp.User.PhoneNumber)

				return
			}

			loginError(t, ctx, client, tc.req, tc.errorMessage)
		})
	}
}

func TestValidate(t *testing.T) {
	testCases := []struct {
		name         string
		errorMessage string
		isSuccess    bool
		req          *pb.ValidateRequest
	}{
		{
			"OK",
			"",
			true,
			nil,
		},
		{
			"Token is required",
			"Key: 'ValidateRequest.Token",
			false,
			&pb.ValidateRequest{
				Token: "",
			},
		},
		{
			"Invalid token",
			"invalid-token",
			false,
			&pb.ValidateRequest{
				Token: "invalid-token",
			},
		},
	}

	ctx := context.Background()
	conn := checkConnection(ctx, t)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			// create user
			newUser := createRandomUser(t, client)

			// login user
			reqLogin := &pb.LoginRequest{
				Username: newUser.Username,
				Password: newUser.Password,
			}

			user, err := client.Login(ctx, reqLogin)
			require.NoError(t, err)
			require.NotNil(t, user)
			require.NotEmpty(t, user.Token)

			// validate user token
			if tc.isSuccess {
				reqValidate := &pb.ValidateRequest{
					Token: user.Token,
				}

				resp, err := client.Validate(ctx, reqValidate)
				require.NoError(t, err)
				require.NotZero(t, resp.UserId)
				require.Equal(t, resp.UserId, user.User.Id)

				return
			}

			_, err = client.Validate(ctx, tc.req)
			require.Error(t, err)
			require.ErrorContains(t, err, tc.errorMessage)
		})
	}
}
func createRandomUser(
	t *testing.T,
	client pb.UserServiceClient,
) *pb.RegisterRequest {

	ctx := context.Background()
	req := &pb.RegisterRequest{
		Name:        utils.RandomString("name-", 10),
		Username:    utils.RandomString("username-", 5),
		Password:    utils.RandomString("", 10),
		Address:     utils.RandomString("address-", 50),
		PhoneNumber: fmt.Sprintf("%d", utils.RandomInt(100000, 1000000)),
	}

	resp, err := client.Register(ctx, req)
	require.NotZero(t, resp.Id)
	require.NoError(t, err)

	return req
}

func loginError(t *testing.T, ctx context.Context, client pb.UserServiceClient, req *pb.LoginRequest, errorMessage string) {
	_, err := client.Login(ctx, req)
	require.Error(t, err)
	require.ErrorContains(t, err, errorMessage)
}
