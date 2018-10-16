package api

import (
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type MlApi struct {
	Router *mux.Router
}

type Credentials struct {
	Password string
}

var Pass []byte

var mySigningKey = []byte("secret")

func NewMlApi() MlApi {
	mlapi := MlApi{}
	mlapi.Router = mux.NewRouter()
	return mlapi
}
func (api * MlApi) Start(routeConf string, password string,port string) {
	GenerateRoute(api.Router, routeConf)
	Pass, _ = bcrypt.GenerateFromPassword([]byte(password), 8)
	http.Handle("/", api.Router)
	log.Fatal(http.ListenAndServe(":" + port, api.Router))
}