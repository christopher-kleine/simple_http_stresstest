package main

import (
	"net/http"

	"github.com/husobee/vestigo"
)

func main() {
	router := vestigo.NewRouter()

	router.Get("/plaintext", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":80", router)
}
