package services

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/maslow123/library-users/pkg/pb"
	"github.com/maslow123/library-users/pkg/utils"
)

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := utils.IsValid(req)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	// check username already exists
	q := "SELECT COUNT(username) FROM users where username = $1"

	row := s.DB.QueryRowContext(ctx, q, req.Username)
	var count int32

	_ = row.Scan(&count)

	if count > 0 {
		return nil, errors.New("username-already-exists")
	}

	// hashed password
	password := req.Password
	hashedPassword := utils.HashPassword(password)

	q = `
		INSERT INTO users (name, username, hashed_password, address, phone_number, level)
		VALUES
		($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	row = s.DB.QueryRowContext(ctx, q,
		&req.Name,
		&req.Username,
		&hashedPassword,
		&req.Address,
		&req.PhoneNumber,
		&req.Level,
	)

	var lastInsertedId int64
	err = row.Scan(&lastInsertedId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp := &pb.RegisterResponse{
		Id: int32(lastInsertedId),
	}

	return resp, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	err := utils.IsValid(req)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	q := `
		SELECT id, name, username, hashed_password, address, phone_number, level, created_at
		FROM users
		WHERE username = $1
		LIMIT 1
	`

	var user pb.User
	var createdAt time.Time

	row := s.DB.QueryRowContext(ctx, q, req.Username)

	err = row.Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.HashedPassword,
		&user.Address,
		&user.PhoneNumber,
		&user.Level,
		&createdAt,
	)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return nil, errors.New("user-not-found")
		}

		return nil, err
	}

	user.CreatedAt = createdAt.Unix()

	match := utils.CheckPasswordHash(req.Password, user.HashedPassword)
	if !match {
		return nil, errors.New("wrong-password")
	}

	token, _ := s.Jwt.GenerateToken(user.Id)
	resp := &pb.LoginResponse{
		User:  &user,
		Token: token,
	}

	return resp, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	err := utils.IsValid(req)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	claims, err := s.Jwt.ValidateToken(req.Token)
	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid-token")
	}

	resp := &pb.ValidateResponse{
		UserId: claims.UserID,
	}

	return resp, nil
}
