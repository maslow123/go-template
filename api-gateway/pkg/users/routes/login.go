package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/library-api-gateway/pkg/users/pb"
	"github.com/maslow123/library-api-gateway/pkg/utils"
)

//	swagger:route POST /users/login users loginUser
//	Returns a user data and the token
//	responses:
//	 200: loginResponse
//	 502: ErrorResponse
type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.UserServiceClient) {
	req := LoginRequestBody{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	utils.SendProtoMessage(ctx, res, http.StatusOK)
}
