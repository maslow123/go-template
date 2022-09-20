package utils

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/runtime/protoiface"

	"github.com/golang/protobuf/jsonpb"
)

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func SendProtoMessage(ctx *gin.Context, res protoiface.MessageV1, status int) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Status(status)

	m := jsonpb.Marshaler{EmitDefaults: true, OrigName: true}
	m.Marshal(ctx.Writer, res)
}
