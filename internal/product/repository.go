package product

import (
	"parte2/internal/domain"
	"parte2/pkg/store"
)

type Repository interface {
	GetAllProducts()                   []domain.Product
	GetProductById(int)                (domain.Product, error)
	GetProductsByPrice(float64)        []domain.Product
	SaveProduct(domain.Product)        domain.Product
	UpdateProduct(int,domain.Product)  (domain.Product, error)
	PatchProduct( int, domain.Product) (domain.Product, error)
	DeleteProduct(int)  				error
}

type repository struct {
	Store store.ProductStorage
	Products *[]domain.Product
	LastID int
}

func NewRepository(store store.ProductStorage, products *[]domain.Product, lastID int)  Repository {
	return &repository{store, products, lastID}
}

func (r *repository) GetAllProducts() []domain.Product {
	return *r.Products 
}

func (r *repository) GetProductById(id int) (productSearched domain.Product, err error) {
	productSearched, err = r.Store.FindProduct(id)
	return
}

func (r *repository) GetProductsByPrice(priceGt float64) (productsSearched []domain.Product) {
	for _, product := range *r.Products {
		if product.Price > priceGt  {
			productsSearched = append(productsSearched, product)
		}
	}
	return 
}

func (r *repository) SaveProduct(product domain.Product) (productStoraged domain.Product) {
	r.LastID++
	product.Id = r.LastID
	productStoraged = product

	*r.Products = append(*r.Products, productStoraged)
	
	return 

}

func (r *repository) UpdateProduct(id int, product domain.Product) (productUpdated domain.Product, err error) {
	
	for i, productInStorage := range *r.Products {
		if productInStorage.Id != id {
			continue
		}
		(*r.Products)[i] = product
		productUpdated = (*r.Products)[i]
		err = r.Store.UpdateProduct(*r.Products)
	}
	return
}

func (r *repository) PatchProduct(id int, product domain.Product) (productUpdated domain.Product, err error) {
	for i, productInStorage := range *r.Products {
		if productInStorage.Id != id {
			continue
		}
		(*r.Products)[i] = product
		productUpdated = (*r.Products)[i]
		err = r.Store.UpdateProduct(*r.Products)
	}
	return
}

func (r *repository) DeleteProduct(id int) (err error) {
	for i, product := range *r.Products {
		if product.Id != id {
			continue
		}
		*r.Products = append((*r.Products)[:i],(*r.Products)[i+1:]...)
		err = r.Store.DeleteProduct(*r.Products)
	}
	return
}