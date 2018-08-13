package models

import (
	"github.com/dgrijalva/jwt-go" //Paquete para los tokens
)

//no se utiliza gorm porque no se interactua con la base de datos
type Claim struct{
	User `json: "user"`
	jwt.StandardClaims
}