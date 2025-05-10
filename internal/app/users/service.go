package users

import (
	"context"

	"github.com/lai0xn/isdb/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
  repo *repository.Queries 
}

func NewService(r *repository.Queries) *Service{
  return &Service{
    repo: r,
  }
}


func (s *Service) GetUserByID(id pgtype.UUID) (*repository.User,error){
  return nil,nil
}

func (s *Service) GetUserByEmail(email string) (*repository.User,error) {
  user ,err :=s.repo.GetUserByEmail(context.Background(),email)
  if err != nil {
    return nil,err
  }

  return &user,nil

}
