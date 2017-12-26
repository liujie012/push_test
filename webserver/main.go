package webserver

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/",page)
	http.ListenAndServe("8000",nil)
}


func page (w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "page")
}


