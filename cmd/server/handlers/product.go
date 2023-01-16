package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"parte2/internal/domain"
	"parte2/internal/product"
	"parte2/pkg/web/request"
	"parte2/pkg/web/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	service product.Service

}

func NewProduct(service product.Service) *Product {
	return &Product{service}
}

func (p *Product) GetPong(ctx *gin.Context) {
	//Response
	ctx.String(http.StatusOK, "Pong")
}

func (p *Product) GetProducts(ctx *gin.Context) {
	//Response
	products, err := p.service.GetAllProducts()

	if err != nil {
		response.Failure(ctx, http.StatusNotFound, err)
		return
	}

	response.Success(ctx,  http.StatusOK, products)
}

func (p *Product) GetProductById(ctx *gin.Context) {
	//Request
	id, err := strconv.Atoi(ctx.Param("id"))
	
	
	//Process
	var productSearched domain.Product
	if err == nil {
		productSearched, err = p.service.GetProductById(id)
	}
	
	//Response
	switch err {
	case product.ErrorProductNotFound:
		response.Failure(ctx, http.StatusNotFound, err)
	case nil:
		response.Success(ctx, http.StatusOK, productSearched)
	default:
		response.Failure(ctx, http.StatusBadRequest, product.ErrorInvalidRequest)	
	}
}

func (p *Product)  GetProductsByPrice(ctx *gin.Context) {
	//Request
	fmt.Println(ctx.Query("price"))
	priceGt , err := strconv.ParseFloat(ctx.Query("price"), 8)
	
	//Process
	var productsSearched []domain.Product
	if err == nil {
		productsSearched, err = p.service.GetProductsByPrice(priceGt)
	}

	//Response 
	switch err {
	case product.ErrorProductNotFound:
		response.Failure(ctx, http.StatusNotFound, err)
	case nil:
		response.Success(ctx, http.StatusOK, productsSearched)
	default:
		response.Failure(ctx, http.StatusBadRequest, product.ErrorInvalidRequest)	
	
	}

}

func (p *Product) SaveProduct(ctx *gin.Context) {
	//Request 
	var req domain.Product
	err := ctx.ShouldBind(&req)
	
	validate := validator.New()
	err = validate.Struct(&req)

	//Process
	var productSaved domain.Product

	if err == nil {
		productSaved, err = p.service.SaveProduct(req)
	}
	
	//Response 
	
	switch err {
	case product.ErrorInvalidData :
		response.Failure(ctx, http.StatusUnprocessableEntity, err)
	case product.ErrorCodeValueExist :
		response.Failure(ctx, http.StatusConflict, err)
	case nil :
		response.Success(ctx, http.StatusCreated, productSaved)
	default :
		response.Failure(ctx, http.StatusBadRequest,product.ErrorInvalidRequest)

	}

}

func (p *Product) UpdateProduct(ctx *gin.Context) {
	//Request
	token := ctx.GetHeader("token")
	
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, "No está autorizado")
		return
	}

	var productToUpdate domain.Product
	id, err := strconv.Atoi(ctx.Param("id"))
	
	err = ctx.ShouldBindJSON(&productToUpdate)
	productToUpdate.Id = id

	//Process
	var productUpdated domain.Product

	if err == nil {
		productUpdated, err = p.service.UpdateProduct(id, productToUpdate)
	}

	//Response 
	switch err {
	case product.ErrorProductNotFound :
		response.Failure(ctx, http.StatusNotFound, err)
	case nil :
		response.Success(ctx, http.StatusOK, productUpdated)
	default :
		response.Failure(ctx, http.StatusBadRequest, product.ErrorInvalidRequest)
	}
}

func (p *Product) PatchProduct(ctx *gin.Context) {
	//Request
	token := ctx.GetHeader("token")
	
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, "No está autorizado")
		return
	}

	var productToPatch request.ProductRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	
	err = json.NewDecoder(ctx.Request.Body).Decode(&productToPatch)
	
	//Process
	var productPatched domain.Product

	if err == nil {
		productPatched, err = p.service.PatchProduct(id, productToPatch)
	}

	//Response 
	switch err {
	case product.ErrorProductNotFound :
		response.Failure(ctx, http.StatusNotFound, err)
	case nil :
		response.Success(ctx, http.StatusOK, productPatched)
	default :
		response.Failure(ctx, http.StatusBadRequest, product.ErrorInvalidRequest)
	}
}

func (p *Product) DeleteProduct(ctx *gin.Context) {
	//Request
	token := ctx.GetHeader("token")
	
	if token != os.Getenv("TOKEN") {
		response.Failure(ctx, http.StatusUnauthorized, product.ErrorUnauthorized)
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	
	//Process
	if err == nil {
		err = p.service.DeleteProduct(id)
	}
	
	//Response
	switch err {
	case product.ErrorProductNotFound:
		response.Failure(ctx, http.StatusNotFound, err)
	case nil:
		response.Success(ctx, http.StatusNoContent, "")
	default:
		response.Failure(ctx, http.StatusBadRequest, product.ErrorInvalidRequest)	
	}

}
 


