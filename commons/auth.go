package commons

import (
	"time"
	"log"
	"io/ioutil"
	"crypto/rsa"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/goapp/models"
)

//Variables globales
var (
	privateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
)

func init(){
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")

	if err != nil {
		log.Fatal("No se pudo leer el archivo privado, Error:  ", err)
	}

	publicteBytes, err := ioutil.ReadFile("./keys/public.rsa")

	if err != nil {
		log.Fatal("No se pudo leer el archivo publico.")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a Private Key")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicteBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a Public Key")
	}
}

//Generar token para el cliente
func GenerateJWT(user models.User) string{
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*2).Unix(), 
			Issuer: "Ricardo Amador",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)

	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}

	return result
}