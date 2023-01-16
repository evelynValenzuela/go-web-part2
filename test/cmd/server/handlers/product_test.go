package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"parte2/cmd/server/handlers"
	"parte2/internal/domain"
	"parte2/internal/product"
	"parte2/pkg/store"
	"parte2/pkg/web/response"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "12345")
	store := store.NewStore("/Users/evvalenzuela/Documents/Bootcamp/web/clase-web01/parte2/test/cmd/server/handlers/products.json")
	products, _ := store.ReadFile()
	repo := product.NewRepository(store, &products, 500)
	service := product.NewService(repo)
	p := handlers.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.GET("/", p.GetProducts)
	pr.GET("/:id", p.GetProductById)
	pr.POST("/", p.SaveProduct)
	pr.PUT("/:id", p.UpdateProduct)
	pr.DELETE("/:id", p.DeleteProduct)
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url,  bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "12345")

	return req, httptest.NewRecorder()

}

func TestGetProductsSucces(t *testing.T) {
	//arrange
	var responseExpected response.Response
	
	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/products/", "" )
	
	//act
	r.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(),&responseExpected)
	
	//assert
	assert.Equal(t, 200, rr.Code)
	assert.Nil(t, err)
}

func TestGetProductByIdSucces(t *testing.T) {
	var responseExpected response.Response
	
	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/products/2","")

	//act
	r.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &responseExpected)
	fmt.Printf("%v",rr.Body )

	//assert
	assert.Equal(t, 200, rr.Code)
	assert.Nil(t, err)
}

func TestSaveProduct(t *testing.T) {
	//arrange
	var responseExpected response.Response
	
	r := createServer()
	product := domain.Product{Id: 501,Name: "Té Hatsu", Quantity: 123, CodeValue: "12345", IsPublised: true, Expiration: "01/12/2022", Price: 45.67}
	bodyReq, _ := json.Marshal(product)
	req, rr := createRequestTest(http.MethodPost, "/products/", string(bodyReq))

	//act
	r.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &responseExpected)

	//assert
	assert.Equal(t, rr.Code, 201)
	assert.Nil(t, err)
}

func TestDeleteProductSucces(t *testing.T) {
	//arrange
	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/products/2","")

	//act
	r.ServeHTTP(rr, req)

	//assert
	assert.Equal(t, 204, rr.Code)
}

func TestGetProductByIdFailure(t *testing.T) {
	//arrange 
	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/products/2a", "")
	fmt.Printf("%v \n", rr.Body)

	//act
	r.ServeHTTP(rr, req)

	//assert
	assert.Equal(t, 400, rr.Code)
}

func TestUpdateProductFailed(t *testing.T) {
	//arrange 

	product := domain.Product{Id: 501,Name: "Té Hatsu", Quantity: 123, CodeValue: "12345", IsPublised: true, Expiration: "01/12/2022", Price: 45.67}
	bodyReq, _ := json.Marshal(product)

	r := createServer()
	req, rr := createRequestTest(http.MethodPut, "/products/501", string(bodyReq))

	//act
	r.ServeHTTP(rr, req)

	//assert
	assert.Equal(t, 404, rr.Code)
}

func TestDeleteProductFailed(t *testing.T) {

	//arrange
	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/products/1", "")
	req.Header.Del("token")

	//act
	r.ServeHTTP(rr, req)

	//assert
	assert.Equal(t, 401, rr.Code)

}

