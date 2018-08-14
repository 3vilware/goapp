package controlers

import(
	"net/http"
	"github.com/goapp/models"
	"fmt"
	"github.com/goapp/configuration"
	"github.com/goapp/commons"
	"encoding/json"
	"crypto/sha256"
	"encoding/base64"
	"crypto/md5"
	"log"
)

//Login es el controlador del login
func Login(w http.ResponseWriter, r *http.Request){
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	db := configuration.GetConection()
	defer db.Close()

	c := sha256.Sum256([]byte(user.Password))
	pwd := base64.URLEncoding.EncodeToString(c[:32]) // leer del inicio a la posicion 32

	db.Where("email = ? and password = ?", user.Email, pwd).First(&user) // User ya tiene pass & email pero con el puntero la rellena
	if user.ID > 0 {
		user.Password = "" //si se encontro algo pass va en blanco para que no lo devuelva en el json
		token := commons.GenerateJWT(user)
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatalf("Error al convertir token a json: %s", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}else{
		m := models.Message{
			Message: "Usuario o contraseña no válidos",
			Code: http.StatusUnauthorized,
		}
		commons.DisplayMessage(w, m)
	}

}

// UserCreate permite registrar un usuario
func UserCreate(w http.ResponseWriter, r *http.Request){
	user := models.User{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el usuario: %s",err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return //paramos con esto siempre
	}

	if user.Password != user.ConfirmPassword{
		m.Message = "Las contraseñas no coinciden"
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	c := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", c)

	user.Password = pwd

	picmd5 := md5.Sum([]byte(user.Email))
	picstr := fmt.Sprintf("%x", picmd5) // almacena el string del hash que de por el usuario
	user.Picture = "https://gravatar.com/avatar" + picstr + "?s=100"

	db := configuration.GetConection()
	defer db.Close()

	err = db.Create(&user).Error

	if err != nil {
		m.Message = fmt.Sprintf("Error al crear registro: %s",err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return 
	}

	m.Message = "Usuario creado con éxito!"
	m.Code = http.StatusCreated
	commons.DisplayMessage(w, m)
}