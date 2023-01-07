package application

import (
	"net/http"

	logs "github.com/mt1976/ebEstimates/logs"
)

func Resources_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc("/favicon.ico", Resources_HandlerFavicon)
	mux.HandleFunc("/favicon-32x32.png", Resources_HandlerFavicon32)
	mux.HandleFunc("/site.webmanifest", Resources_HandlerManifest)
	mux.HandleFunc("/favicon-16x16.png", Resources_HandlerFavicon16)
	mux.HandleFunc("/browserconfig.xml", Resources_HandlerBrowserConfiguration)
	//mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	logs.Publish("Core", "Resources"+" Impl")
}

func Resources_HandlerFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon.ico")
}

func Resources_HandlerFavicon16(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon-16x16.png")
}

func Resources_HandlerFavicon32(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon-32x32.png")
}

func Resources_HandlerManifest(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "site.webmanifest")
}

func Resources_HandlerBrowserConfiguration(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "browserconfig.xml")
}
