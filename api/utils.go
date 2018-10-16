package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

// Route struct definition
type Route struct {
	Method string
	Path  string
	Handler string
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