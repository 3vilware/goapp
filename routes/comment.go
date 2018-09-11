package routes

import(
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/goapp/controlers"
)

func SetCommentRouter(router *mux.Router){
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controlers.CommentCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controlers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}