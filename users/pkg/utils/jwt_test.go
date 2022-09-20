package utils

import (
	"testing"

	"github.com/maslow123/library-users/pkg/config"
	"github.com/stretchr/testify/require"
)

type Server struct {
	Jwt JwtWrapper
}

func TestJwt(t *testing.T) {
	c, err := config.LoadConfig("../config/envs", "test")
	require.NoError(t, err)

	jwt := JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "test-jwt",
		ExpirationHours: 1,
	}

	userID := int32(RandomInt(0, 1))

	s := Server{
		Jwt: jwt,
	}

	token, err := s.Jwt.GenerateToken(userID)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// check valid token
	valid, err := s.Jwt.ValidateToken(token)
	require.NoError(t, err)
	require.NotNil(t, valid)

	// invalid token
	invalidToken := RandomString("token-", 50)
	valid, err = s.Jwt.ValidateToken(invalidToken)
	require.Error(t, err)
	require.Nil(t, valid)

}
