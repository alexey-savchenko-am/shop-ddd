package product

type Repository interface {
	Save(p *Product) error
	ByID(id ID) (*Product, error)
}
