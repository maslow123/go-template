package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/library-api-gateway/pkg/config"
	"github.com/maslow123/library-api-gateway/pkg/users"
)

func main() {
	c, err := config.LoadConfig("./pkg/config/envs", "dev")

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	_ = *users.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
