package postgres

type ProductModel struct {
	ID string `gorm:"primaryKey"`
	SKU string
	Name string
	Price int64
	Currency string
}

func (ProductModel) TableName() string {
	return "products"
}