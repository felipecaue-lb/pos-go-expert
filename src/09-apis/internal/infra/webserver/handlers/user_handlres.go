package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/felipecaue-lb/goexpert/09-apis/internal/dto"
	"github.com/felipecaue-lb/goexpert/09-apis/internal/entity"
	"github.com/felipecaue-lb/goexpert/09-apis/internal/infra/database"
	"github.com/go-chi/jwtauth/v5"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

// GetJWT godoc
// @Summary Get JWT token
// @Description Authenticate user and return a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.GetJWTInput true "User credentials"
// @Success 200 {object} dto.GetJWTOutput
// @Failure 400 {object} dto.ErrorOutput "Bad Request"
// @Failure 401 {object} dto.ErrorOutput "Unauthorized"
// @Router /login [post]
func (handler *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var input dto.GetJWTInput

	error := json.NewDecoder(r.Body).Decode(&input)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	user, error := handler.UserDB.FindByEmail(input.Email)
	if error != nil {
		w.WriteHeader(http.StatusUnauthorized)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !user.ValidatePassword(input.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		error := dto.ErrorOutput{Message: "Invalid credentials"}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Duration(jwtExpiresIn) * time.Second),
	})

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserInput true "User information"
// @Success 201 {object} entity.User
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [post]
func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	error := json.NewDecoder(r.Body).Decode(&user)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	newUser, error := entity.NewUser(user.Name, user.Email, user.Password)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	error = handler.UserDB.Create(newUser)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := dto.ErrorOutput{Message: error.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
