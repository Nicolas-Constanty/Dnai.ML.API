package api

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type ApiHandler struct {}

// Home handler (/)
func (t *ApiHandler) HomeHandler(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Welcome to ML Studio API"))
}

// Token handler (/token)
func (t *ApiHandler) AuthHandler(w http.ResponseWriter, r *http.Request) {
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
