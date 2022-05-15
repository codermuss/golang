package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/merhaba" {
		http.Error(w, "404 Sayfa Bulunamadi", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Desteklenmeyen Metod", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Merhaba, Dunya!")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
	}
	fmt.Fprintf(w, "Post istegi basarili...\n")
	ad := r.FormValue("ad")
	adres := r.FormValue("adres")
	fmt.Fprintf(w, "Ad: %s\n", ad)
	fmt.Fprintf(w, "Adres: %s\n", adres)
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/merhaba", helloHandler)

	fmt.Print("Server 8080 portunda baslatiliyor...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
