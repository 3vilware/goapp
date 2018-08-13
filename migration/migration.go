package migration 

import(
	"github.com/goapp/configuration"
	"github.com/goapp/models"
)

func Migrate(){
	db := configuration.GetConection()

	defer db.Close()

	db.CreateTable(models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id") //se unen para asegurar que no se repetiran votos
}