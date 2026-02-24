package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/felipecaue-lb/goexpert/09-apis/internal/dto"
	"github.com/felipecaue-lb/goexpert/09-apis/internal/entity"
	"github.com/felipecaue-lb/goexpert/09-apis/internal/infra/database"
	entityPkg "github.com/felipecaue-lb/goexpert/09-apis/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(productDB database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: productDB}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param request body dto.CreateProductInput true "Product data"
// @Success 201 {object} dto.ErrorOutput
// @Failure 400 {object} dto.ErrorOutput "Bad Request"
// @Failure 500 {object} dto.ErrorOutput "Internal Server Error"
// @Router /products [post]
// @Security ApiKeyAuth
func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	error := json.NewDecoder(r.Body).Decode(&product)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	p, error := entity.NewProduct(product.Name, product.Price)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	error = handler.ProductDB.Create(p)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 400 {object} dto.ErrorOutput "Bad Request"
// @Failure 404 {object} dto.ErrorOutput "Not Found"
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (handler *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.ErrorOutput{Message: "ID is required"}
		json.NewEncoder(w).Encode(error)
		return
	}

	product, error := handler.ProductDB.FindByID(id)
	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary Update a product by ID
// @Description Update a product by its ID with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body entity.Product true "Updated product data"
// @Success 200 {object} dto.ErrorOutput
// @Failure 400 {object} dto.ErrorOutput "Bad Request"
// @Failure 404 {object} dto.ErrorOutput "Not Found"
// @Failure 500 {object} dto.ErrorOutput "Internal Server Error"
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (handler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.ErrorOutput{Message: "ID is required"}
		json.NewEncoder(w).Encode(error)
		return
	}

	var product entity.Product
	error := json.NewDecoder(r.Body).Decode(&product)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	product.ID, error = entityPkg.ParseID(id)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, error = handler.ProductDB.FindByID(id)
	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	error = handler.ProductDB.Update(&product)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} dto.ErrorOutput
// @Failure 400 {object} dto.ErrorOutput "Bad Request"
// @Failure 404 {object} dto.ErrorOutput "Not Found"
// @Failure 500 {object} dto.ErrorOutput "Internal Server Error"
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (handler *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.ErrorOutput{Message: "ID is required"}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, error := handler.ProductDB.FindByID(id)
	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	error = handler.ProductDB.Delete(id)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get a list of all products with pagination and sorting
// @Tags products
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param sort query string false "Sort order (e.g., asc, desc)"
// @Success 200 {array} entity.Product
// @Failure 500 {object} dto.ErrorOutput "Internal Server Error"
// @Router /products [get]
// @Security ApiKeyAuth
func (handler *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, error := strconv.Atoi(page)
	if error != nil {
		pageInt = 0
	}

	limitInt, error := strconv.Atoi(limit)
	if error != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	products, error := handler.ProductDB.FindAll(pageInt, limitInt, sort)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
