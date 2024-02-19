package exceptions

import "net/http"

func Error404(w http.ResponseWriter) {
	http.Error(w, "404 Not Found", http.StatusNotFound)
}