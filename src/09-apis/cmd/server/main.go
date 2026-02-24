package main

import (
	"log"
	"net/http"

	"github.com/felipecaue-lb/goexpert/09-apis/configs"
	_ "github.com/felipecaue-lb/goexpert/09-apis/docs"
	"github.com/felipecaue-lb/goexpert/09-apis/internal/entity"
	"github.com/felipecaue-lb/goexpert/09-apis/internal/infra/database"
	"github.com/felipecaue-lb/goexpert/09-apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Go Expert API
// @version 1.0
// @description Product API with JWT authentication.
// @termsOfService http://swagger.io/terms/

// @contact.name Felipe Cauê
// @contact.url http://github.com/felipecaue-lb
// @contact.email felipecaue@mail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, error := configs.LoadConfig(".")
	if error != nil {
		panic(error)
	}

	db, error := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if error != nil {
		panic(error)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	println("=> Servidor iniciado!")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwt", configs.TokenAuth))
	router.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn))

	router.Use(LogRequest)

	router.Route("/products", func(router chi.Router) {
		router.Use(jwtauth.Verifier(configs.TokenAuth))
		router.Use(jwtauth.Authenticator(configs.TokenAuth))

		router.Post("/", productHandler.CreateProduct)
		router.Get("/{id}", productHandler.GetProduct)
		router.Put("/{id}", productHandler.UpdateProduct)
		router.Delete("/{id}", productHandler.DeleteProduct)
		router.Get("/", productHandler.GetAllProducts)
	})

	router.Post("/users", userHandler.CreateUser)

	router.Post("/login", userHandler.GetJWT)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	http.ListenAndServe(":8080", router)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")
		log.Printf("User => %s", user)

		log.Printf("Request => %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
