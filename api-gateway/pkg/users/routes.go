//	Package users Library API.
//
//	Documentation for Library API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	swagger:meta
package users

import (
	"github.com/gin-gonic/gin"
	"github.com/maslow123/library-api-gateway/pkg/config"
	"github.com/maslow123/library-api-gateway/pkg/users/pb"
	"github.com/maslow123/library-api-gateway/pkg/users/routes"
)

//	ErrorResponse returns error string from for each service
//	swagger: response errorResponse
type ErrorResponseWrapper struct {
	//	Error Response
	//	in: body
	Body *pb.ErrorResponse
}

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
		Router: r,
	}

	a := InitAuthMiddleware(svc)

	routes := r.Group("/users")
	routes.Use(a.CORSMiddleware)
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

//	User ID returns in the response
//	swagger:response registerResponse
type RegisterResponseWrapper struct {
	//	User ID
	//	in: body
	Body *pb.RegisterResponse
}

//	swagger:parameters registerUser
type RegisterUserParams struct {
	//	The body to create new user
	//	in: body
	//	required: true
	Body pb.RegisterRequest
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

//	User Data returns in the response
//	swagger:response loginResponse
type LoginResponse struct {
	//	User data
	//	in: body
	Body *pb.LoginResponse
}

//	swagger:parameters loginUser
type LoginUserParams struct {
	//	The body to login user form
	//	in: body
	//	required: true
	Body pb.LoginRequest
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
