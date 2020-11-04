package controllers

import (
	"fmt"
	"gotello/config"
	"html/template"
	"net/http"
)

func viewIndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("app/views/index.html")
	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartWebServer() error {
	http.HandleFunc("/", viewIndexHandler)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", config.Config.Address, config.Config.Port), nil)
}