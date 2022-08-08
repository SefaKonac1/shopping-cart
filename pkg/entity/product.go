package entity

import (
	config "github.com/SefaKonac1/shopping-cart.git/pkg/configuration"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	ID    uint    `gorm:"primaryKey"`
	Name  string  `gorm:""json:"name"`
	Price float32 `json:"price"`
	VAT   float32 `json:"vat"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func (p *Product) CreateProduct() *Product {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllProducts() []Product {
	var Products []Product
	db.Find(&Products)
	return Products
}

func GetProductById(Id uint) (*Product, *gorm.DB) {
	var getProduct Product
	db := db.Where("ID=?", Id).Find(&getProduct)
	return &getProduct, db
}
