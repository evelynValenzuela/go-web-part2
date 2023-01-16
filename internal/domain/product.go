package domain

type Product struct {
	Id 			int		`json:"id"`
	Name 		string	`json:"name"       validate:"required"`
	Quantity 	int		`json:"quantity"   validate:"required"`
	CodeValue 	string	`json:"code_value" validate:"required"`
	IsPublised 	bool	`json:"is_published"`
	Expiration 	string	`json:"expiration" validate:"required"`
	Price 		float64	`json:"price"       validate:"required"`
}