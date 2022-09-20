package users

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/library-api-gateway/pkg/config"
	"github.com/stretchr/testify/require"
)

func NewServer(t *testing.T) *ServiceClient {
	r := gin.Default()

	config, err := config.LoadConfig("../config/envs", "test")
	require.NoError(t, err)

	svc := RegisterRoutes(r, &config)
	return svc
}
