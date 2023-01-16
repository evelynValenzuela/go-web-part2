package request

type ProductRequest struct {
	Name 		string	`json:"name"       validate:"required"`
	Quantity 	int		`json:"quantity"   validate:"required"`
	IsPublised 	bool	`json:"is_published"`
	Price 		float64	`json:"price"       validate:"required"`
}