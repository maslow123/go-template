package services

import (
	"database/sql"

	"github.com/maslow123/library-users/pkg/utils"
)

type Server struct {
	DB  *sql.DB
	Jwt utils.JwtWrapper
}
