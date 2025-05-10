package documents

import "github.com/lai0xn/isdb/internal/repository"

type DocumentDTO struct{
  Name string `json:"name"`
  Content string `json:"content"`
  PropertyType repository.PropertyType `json:"property_type"`
  Metadata []byte `json:"metada"`
}
