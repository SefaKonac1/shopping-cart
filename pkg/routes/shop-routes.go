package routes

import (
	"github.com/SefaKonac1/shopping-cart.git/pkg/controller"
	"github.com/gorilla/mux"
)

var ShoppingCartRoutes = func(router *mux.Router) {
	router.HandleFunc("/products/", controller.ListProducts).Methods("GET")
	router.HandleFunc("/basketItem/{itemId}", controller.AddItemToCart).Methods("POST")
	router.HandleFunc("/basketItem/", controller.ShowCartItems).Methods("GET")
	router.HandleFunc("/basketItem/{itemId}", controller.DeleteItemFromCart).Methods("DELETE")
	router.HandleFunc("/completeOrder/", controller.CompleteShopping).Methods("HEAD")
}
