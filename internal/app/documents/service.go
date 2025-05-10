package documents

import (
	"context"
	"io"
	"log"
	"mime/multipart"

	"github.com/lai0xn/isdb/internal/repository"
)

type Service struct {
  r *repository.Queries
}

func NewDocService(r *repository.Queries) *Service{
  return &Service{
    r: r,
  }
}

func (s *Service) CreateDocument(d DocumentDTO) (repository.Document,error){
  document,err := s.r.CreateDocument(context.Background(),repository.CreateDocumentParams{
    Content: d.Content,
    PropertyType: d.PropertyType,
    Name: d.Name,
    Metadata: d.Metadata,
  })
  if err != nil {
    log.Println(err)
    return repository.Document{},err
  }
  return document,nil
}

func (s *Service) ParseDocument(file multipart.File) (string,error){
  content,err := io.ReadAll(file)
  if err != nil {
    return "",err
  }
  return string(content),nil
}
