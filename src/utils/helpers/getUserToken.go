package helpers

import (
	"net/http"
	"strings"
)

func GetUserToken(w http.ResponseWriter) string {
	auth := w.Header().Get("Authorization")
	_, token, _ := strings.Cut(auth, " ")
	return token
}
