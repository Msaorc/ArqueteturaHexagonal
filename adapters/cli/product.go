package cli

import (
	"fmt"

	"github.com/Msaorc/ArqueteturaHexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, producId string, productName string, price float32) (string, error) {
	var result = ""
	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f",
			product.GetId(), product.GetName(), product.GetPrice())
	case "enable":
		product, err := service.Get(producId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s hab been enabled.", res.GetName())
	case "disable":
		product, err := service.Get(producId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s hab been disable.", res.GetName())
	default:
		res, err := service.Get(producId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			res.GetId(), res.GetName(), res.GetPrice(), res.GetStatus())
	}
	return result, nil
}
