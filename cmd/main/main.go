package main

import (
	"github.com/SefaKonac1/shopping-cart.git/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.ShoppingCartRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
