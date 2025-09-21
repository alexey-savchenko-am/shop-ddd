package postgres

type ProductModel struct {
	ID       string `gorm:"type:uuid;primaryKey;column:id"`
	SKU      string `gorm:"column:sku"`
	Name     string `gorm:"column:name"`
	Price    int64  `gorm:"column:price"`
	Currency string `gorm:"column:currency"`
}

func (ProductModel) TableName() string {
	return "products"
}
