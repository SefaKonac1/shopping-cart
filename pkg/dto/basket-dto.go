package dto

import "github.com/SefaKonac1/shopping-cart.git/pkg/entity"

type BasketDTO struct {
	Products   []entity.Product
	TotalPrice float32
}

func (basket *BasketDTO) ShowAllBasketItems() ([]entity.Product, float32) {
	return basket.Products, basket.TotalPrice
}

func (basket *BasketDTO) AddItemIntoCart(p *entity.Product) ([]entity.Product, float32) {
	basket.Products = append(basket.Products, *p)
	basket.TotalPrice += p.Price
	return basket.Products, basket.TotalPrice
}

func (basket *BasketDTO) DeleteItemFromCart(Id uint) {
	for i := range basket.Products {
		if basket.Products[i].ID == Id {
			basket.Products = append(basket.Products[:i], basket.Products[i+1:]...)
			return
		}
	}

	panic("Item Id is not found in cart !")
}

func (basket *BasketDTO) CompleteOrder() ([]entity.Product, float32) {
	return basket.ShowAllBasketItems()
}
