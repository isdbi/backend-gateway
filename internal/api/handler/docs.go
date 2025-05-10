package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lai0xn/isdb/internal/api"
	"github.com/lai0xn/isdb/internal/app/documents"
	"github.com/lai0xn/isdb/internal/repository"
	"github.com/lai0xn/isdb/pkg/utils"
)

func (a *API) ParseDocument(w http.ResponseWriter,r *http.Request){
  err := r.ParseMultipartForm(10 << 20)

  if err != nil {
    utils.WriteJSONError(w,http.StatusInternalServerError,err)
    return
  }

  file,h,err := r.FormFile("doc")
  if err != nil {
    utils.WriteJSONError(w,http.StatusInternalServerError,err)
    return
  }
  defer file.Close()

  content,err := a.docSrv.ParseDocument(file)

  if err != nil {
    utils.WriteJSONError(w,http.StatusInternalServerError,err)
    return
  }

  metadata,err := json.Marshal(r.FormValue("metadata"))

  if err != nil {
    utils.WriteJSONError(w,http.StatusInternalServerError,err)
    return
  }

  document,err := a.docSrv.CreateDocument(documents.DocumentDTO{
    Name: h.Filename,
    Content: content,
    PropertyType: repository.PropertyType(r.FormValue("property_type")),
    Metadata: metadata,
  })


   if err != nil {
    utils.WriteJSONError(w,http.StatusInternalServerError,err)
    return
  }
  
  utils.WriteJSONResponse(w,http.StatusOK,api.Map{
    "message":"document parsed",
    "data":document,
  })
  

}
