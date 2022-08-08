package controller

import (
	"encoding/json"
	"fmt"
	"github.com/SefaKonac1/shopping-cart.git/pkg/dto"
	"github.com/SefaKonac1/shopping-cart.git/pkg/entity"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var BasketDto dto.BasketDTO

func ListProducts(w http.ResponseWriter, r *http.Request) {
	productLists := entity.GetAllProducts()
	res, _ := json.Marshal(productLists)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddItemToCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["itemId"]
	ID, err := strconv.ParseInt(productId, 0, 0)

	if err != nil {
		fmt.Errorf("error while parsing ")
	}

	productDetails, _ := entity.GetProductById(uint(ID))

	if productDetails != nil {
		BasketDto.AddItemIntoCart(productDetails)
	} else {
		panic("wanted item to add into basket is not exist in database !")
	}
}

func ShowCartItems(w http.ResponseWriter, r *http.Request) {
	cartItems, totalPrices := BasketDto.ShowAllBasketItems()
	w.Header().Set("Content-Type", "pkglication/json")

	res, _ := json.Marshal(cartItems)
	res2, _ := json.Marshal(totalPrices)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
	w.Write(res2)
}

func DeleteItemFromCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productId := vars["ItemId"]

	ID, err := strconv.ParseInt(productId, 0, 0)

	if err != nil {
		fmt.Errorf("error while parsing ")
	}

	productDetails, _ := entity.GetProductById(uint(ID))

	if productDetails != nil {
		BasketDto.DeleteItemFromCart(uint(ID))
	} else {
		panic("wanted item to add into basket is not exist in database !")
	}
}

func CompleteShopping(w http.ResponseWriter, r *http.Request) {
	BasketDto.CompleteOrder()
}
