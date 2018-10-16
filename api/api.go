package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
	"github.com/gorilla/mux"
)

type Route struct {
	Method string
	Path  string
	Handler string
}

type ApiHandler struct {}

func FuncPointer(m map[string]interface{}, name string) (result reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	result = f
	return
}

func ParseRoute(path string) (result []Route) {
	var datas []Route
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		args := strings.Split(line, " ")
		data := Route{ Method: args[0], Path: args[1], Handler: args[2]}
		datas = append(datas, data)
		fmt.Println(data)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return datas
}

func GenerateRoute(r *mux.Router, path string) {
	routes := ParseRoute(path)
	var t ApiHandler
	for _, route := range routes {
		route := route
		method := reflect.ValueOf(&t).MethodByName(route.Handler)
		if method.IsValid() {
			handler := func(w http.ResponseWriter, r *http.Request){
				size := 2
				inputs := make([]reflect.Value, size)
				inputs[0] = reflect.ValueOf(w)
				inputs[1] = reflect.ValueOf(r)
				method.Call(inputs)
			}
			r.HandleFunc(route.Path, handler).Methods(route.Method)
		}
	}
}

var Pass []byte

type Credentials struct {
	Password string
}

var mySigningKey = []byte("secret")

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

func (t *ApiHandler) HomeHandler(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Welcome to ML Studio API"))
}

func (t *ApiHandler) TokenHandler(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = bcrypt.CompareHashAndPassword(Pass, []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		RespondError(w, http.StatusUnauthorized,"Wrong password")
		return
	}

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "Ado Kukic"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	data := map[string]string{"auth" : tokenString}
	/* Finally, write the token to the browser window */
	RespondJSON(w, http.StatusOK, data)
}