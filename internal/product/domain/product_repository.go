package domain

type ProductRepository interface {
	Save(p *Product) error
	ByID(id ProductID) (*Product, error)
}
