package aipaper

import (
	"github.com/newdee/aipaper-util/config"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func GetProductList() ([]Product, error) {
	cfg, err := config.Get(config.AIPaper)
	if err != nil {
		return nil, err
	}
	var productList []Product
	err = cfg.GetWithUnmarshal("product", &productList, &config.JSONUnmarshaler{})
	if err != nil {
		return nil, err
	}
	return productList, nil
}
