package auth

import (
	"context"
	"errors"

	"github.com/lai0xn/isdb/internal/repository"
	"github.com/lai0xn/isdb/pkg/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repository *repository.Queries
  verifier Verifier
}

func NewService(r *repository.Queries,v Verifier) *Service {
	return &Service{
		repository: r,
    verifier: v,
	}
}



func (s *Service) SignupUser(ctx context.Context, req SignupUserRequest) error {
	// Validate the request
	if err := req.Validate(); err != nil {
		return err
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	id := uuid.New()
	// Convert to pgtype.UUID
	pgUUID := pgtype.UUID{}
	copy(pgUUID.Bytes[:], id[:]) // Copy the 16 bytes
	pgUUID.Valid = true

	// Convert the request to a repository model
	params := repository.CreateUserParams{
		ID:         pgUUID,
		Email:      req.Email,
		Password:   hashedPassword, // Use the hashed password
		Name:       req.Name,
		FamilyName: req.FamilyName,
		Age:        pgtype.Int4{Int32: int32(req.Age), Valid: true},
	}

	// Call the repository to create the user
	err = s.repository.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SendVerification(user repository.User) error {
  err := s.verifier.Send(user)
  if err != nil {
    return err
  }
  return nil
}

func (s *Service) Verify(user repository.User, otp string) error {
  err := s.verifier.Verify(user,otp)
  if err != nil {
    return err
  }
  err = s.repository.ActivateUser(context.Background(),user.ID)
  return nil
}


func (s *Service) LoginUser(ctx context.Context, req LoginUserRequest) (map[string]string, error) {
	// Validate the request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Fetch the user by email
	user, err := s.repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	// Verify the password
	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}
	tokenPair, err := utils.GenerateTokenPair(user.Email, user.ID.String())

	if err != nil {
		return nil, errors.New("Failed to generate token")
	}

	return tokenPair, nil
}
