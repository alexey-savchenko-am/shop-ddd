package domain

import (
	"fmt"

	"github.com/alexey-savchenko-am/shop-ddd/internal/common"
)

var (
	ErrProductSKURequired     = common.Error{Code: "product_sku_required", Message: "product sku is required"}
	ErrProductNameRequired    = common.Error{Code: "product_name_required", Message: "product name is required"}
	ErrProductInvalidPrice    = common.Error{Code: "product_invalid_price", Message: "product price must be greater than zero"}
	ErrProductInvalidCurrency = common.Error{Code: "product_invalid_currency", Message: "product currency is invalid"}
	ErrProductNotFound        = common.Error{Code: "product_not_found", Message: "product not found"}
)

func ErrProductInvalidID(id string) common.Error {
	return common.Error{
		Code:    "product_invalid_id",
		Message: fmt.Sprintf("product ID %s is not valid", id),
	}
}
