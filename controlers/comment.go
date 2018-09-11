package controlers

import (
	"github.com/goapp/models"
	"github.com/goapp/configuration"
	"github.com/goapp/commons"
	"net/http"
	"encoding/json"
	"fmt"
)

func CommentCreate(w http.ResponseWriter, r *http.Request){
	comment := models.Comment{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	db := configuration.GetConection()
	defer db.Close()

	err = db.Create(&comment).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al crear el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Comentario creado exitosamente!"
	commons.DisplayMessage(w, m)
}

func CommentGetAll(w http.ResponseWriter, r *http.Request){
	comments := []models.Comment{}
	m := models.Message{}
	user := models.User{}
	vote := models.Vote{}

	

}