package handler

import (
	"github.com/lai0xn/isdb/internal/app/auth"
	"github.com/lai0xn/isdb/internal/app/documents"
	"github.com/lai0xn/isdb/internal/app/users"
)

type API struct {
	authSrv *auth.Service
  userSrv *users.Service
  docSrv *documents.Service
}

func NewApi(authSrv *auth.Service,userSrv *users.Service,docSrv *documents.Service) *API {
	return &API{
		authSrv: authSrv,
    userSrv: userSrv,
    docSrv: docSrv,
	}
}
