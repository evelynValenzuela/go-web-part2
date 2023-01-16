package product

import (
	"errors"
	"fmt"
	"parte2/internal/domain"
	"parte2/pkg/web/request"
	"regexp"
	"time"
)

var (
	ErrorInvalidDate = errors.New("Error: El formato de fecha es inválido")
	ErrorProductNotFound = errors.New("Error: El producto NO existe")
	ErrorProductsNotFound = errors.New("No se encontraron productos con el criterio establecido")
	ErrorCodeValueExist = errors.New("Error: El code value del producto ya existe")
	ErrorInvalidData= errors.New("Error: No se permiten valores vacíos")
	ErrorInvalidRequest= errors.New("Error: Invalid Request")
	ErrorUnauthorized= errors.New("Error: No está autorizado")
)

type Service interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductById(id int) (domain.Product, error)
	GetProductsByPrice(priceGt float64) (productsSearched []domain.Product, err error)
	SaveProduct(product domain.Product) (productStoraged domain.Product, err error)
	UpdateProduct(id int, product domain.Product) (productUpdated domain.Product, err error)
	PatchProduct(id int, product request.ProductRequest) (productUpdated domain.Product, err error)
	DeleteProduct(id int) (err error)
}


type service struct {
	//Repo
	repo Repository

	//extrernal Apis
}


func NewService(r Repository) *service {
	return &service{repo: r}

}

func (s *service) GetAllProducts() (products []domain.Product, err error ){
	products = s.repo.GetAllProducts()

	if len(products) == 0 {
		err = ErrorProductsNotFound
	}
	
	return 
}

func (s *service) GetProductById(id int) (productSearched domain.Product, err error) {

	productSearched, err = s.repo.GetProductById(id)

	if productSearched.Id == 0 {
		err = ErrorProductNotFound
	}

	return
}

func (s *service) GetProductsByPrice(priceGt float64) (productsSearched []domain.Product, err error) {
	productsSearched = s.repo.GetProductsByPrice(priceGt)

	if len(productsSearched) == 0 {
		err = ErrorProductsNotFound
	}
	return
}

func (s *service) SaveProduct(product domain.Product) (productStoraged domain.Product, err error) {
	
	for _, productInStorage := range s.repo.GetAllProducts() {
		if(productInStorage.CodeValue == product.CodeValue) {
			err = ErrorCodeValueExist
			return 
		}
	}

	date := product.Expiration
	expresion := "[0-3][0-9]/[0-1][0-9]/[0-2][0-9][0-9][0-9]"
	match, _ := regexp.MatchString(expresion, date)
	
	if !match {
		err = ErrorInvalidDate
		return
	} else {
		_, err = time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", date[6:10], date[3:5], date[0:2]))
				
		if err != nil {
			err = ErrorInvalidDate
			return
		}
	
	}

	productStoraged =  s.repo.SaveProduct(product)
	return
}

func (s *service) UpdateProduct(id int, product domain.Product) (productUpdated domain.Product, err error) {
	_, err = s.GetProductById(id)

	if err != nil {
		return
	}
	productUpdated, err =  s.repo.UpdateProduct(id, product)

	return

}

func (s *service) PatchProduct(id int, product request.ProductRequest) (productUpdated domain.Product, err error) {
	productInStorage, err := s.GetProductById(id)

	if err != nil {
		return
	}

	productInStorage.Name = product.Name
	productInStorage.Quantity = product.Quantity
	productInStorage.IsPublised = product.IsPublised
	productInStorage.Price = product.Price
	productUpdated, err =  s.repo.PatchProduct(id, productInStorage)

	return

}

func (s *service) DeleteProduct(id int) (err error) {
	_, err = s.GetProductById(id)

	if err != nil {
		return
	}

	err = s.repo.DeleteProduct(id)
	return

}
