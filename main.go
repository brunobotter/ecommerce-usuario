package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Usuario</h1>"))
	})
	http.ListenAndServe(":5000", nil)
}
