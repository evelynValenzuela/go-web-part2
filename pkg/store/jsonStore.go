package store

import (
	"encoding/json"
	"os"
	"parte2/internal/domain"
)

type ProductStorage interface {
	ReadFile()  ([]domain.Product, error)
	FindProduct(int) (domain.Product,error)
	UpdateProduct([]domain.Product) error
	DeleteProduct([]domain.Product) error
}

type store struct {
	Filepath string
}

func NewStore(filepath string) ProductStorage {
	return &store{filepath}
}

func (s *store) ReadFile() (products []domain.Product, err error) {
	file, err := os.ReadFile(s.Filepath)
	err = json.Unmarshal(file, &products)

	return 

}

func (s *store) FindProduct(id int) (product domain.Product, err error) {
	var products []domain.Product
	file, err := os.ReadFile(s.Filepath)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &products)

	for _, productInStorage := range products {
		if productInStorage.Id != id {
			continue
		}
		product = productInStorage
	}

	return

}

func (s *store) UpdateProduct(products []domain.Product) (err error) {
	file, err := os.OpenFile(s.Filepath, os.O_TRUNC | os.O_RDWR, 0644)

	productsToJson, err := json.Marshal(products)
	_, err = file.Write(productsToJson)
	
	return

}

func (s *store) DeleteProduct(products []domain.Product) (err error) {
	file, err := os.OpenFile(s.Filepath, os.O_TRUNC | os.O_RDWR, 0644)

	productsToJson, err := json.Marshal(products)
	_, err = file.Write(productsToJson)
	
	return

}


