package routes

import (
	"parte2/cmd/server/handlers"
	"parte2/internal/domain"
	"parte2/internal/product"
	"parte2/pkg/store"

	"github.com/gin-gonic/gin"
)

type Router struct {
	store store.ProductStorage
	db *[]domain.Product
	en *gin.Engine
}
func NewRouter(store store.ProductStorage ,en *gin.Engine, db *[]domain.Product) *Router {
	return &Router{en: en, db: db, store: store}
}

func (r *Router) SetRoutes() {
	r.Setproduct()
}
// product
func (r *Router) Setproduct() {
	// instances
	rp := product.NewRepository(r.store, r.db, 500)
	sv := product.NewService(rp)
	h := handlers.NewProduct(sv)

	ws := r.en.Group("/products")
	
	r.en.GET("/ping", h.GetPong)
	ws.GET("", h.GetProducts)
	ws.GET("/:id", h.GetProductById)
	ws.GET("/search", h.GetProductsByPrice)
	ws.POST("", h.SaveProduct)
	ws.PUT("/:id", h.UpdateProduct)
	ws.PATCH("/:id", h.PatchProduct)
	ws.DELETE("/:id", h.DeleteProduct)
}