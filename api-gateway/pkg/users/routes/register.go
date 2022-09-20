package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/library-api-gateway/pkg/users/pb"
	"github.com/maslow123/library-api-gateway/pkg/utils"
)

//	swagger:route POST /users/register users registerUser
//	Returns a user ID when user created
//	responses:
//	 201: registerResponse
//	 502: ErrorResponse
type RegisterRequestBody struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Level       int32  `json:"level"`
}

func Register(ctx *gin.Context, c pb.UserServiceClient) {
	req := RegisterRequestBody{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Name:        req.Name,
		Username:    req.Username,
		Password:    req.Password,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Level:       req.Level,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
