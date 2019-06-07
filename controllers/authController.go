package controllers

import (
	"encoding/json"
	"go-phonebooks/models"
	u "go-phonebooks/utils"
	h "go-phonebooks/utils/hash"
	"net/http"
	"strings"
)

var AuthController = &Controller{PrefixURL: "/auth"}

func init() {
	routes := map[string]Route{
		"Profile":  Route{Method: http.MethodGet, Name: "Auth.Get.Profile"},
		"Login":    Route{URL: "/login", Method: http.MethodPost, Name: "Auth.Post.Login"},
		"Register": Route{URL: "/register", Method: http.MethodPost, Name: "Auth.Post.Register"},
	}
	middlewares := map[string][]string{
		"Profile": []string{"jwt"},
	}
	AuthController.Routes = routes
	AuthController.Middlewares = middlewares
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	LoginRequest
	ConfirmPassword string `json:"confirm_password"`
}

func (self *Controller) Profile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	user := &models.User{}
	err := models.GetDB().Model(&models.User{}).Find(&user, userID).Error
	if err != nil {
		u.RespondError(w, http.StatusBadRequest, "Bad Request!", nil)
		return
	}
	u.Respond(w, 200, u.MessageWithData(200, "Succeeded!", user))
}

func validateLoginRequest(login *LoginRequest) (map[string]interface{}, bool) {
	errors := make(map[string][]string)
	if !strings.Contains(login.Email, "@") {
		msg := "Email address is required."
		x, y := errors["email"]
		if !y {
			errors["email"] = []string{msg}
		} else {
			errors["email"] = append(x, msg)
		}
	}
	if len(login.Password) < 6 {
		msg := "Password is required."
		x, y := errors["password"]
		if !y {
			errors["password"] = []string{msg}
		} else {
			errors["password"] = append(x, msg)
		}
	}
	if len(errors) > 0 {
		return map[string]interface{}{"errors": errors}, false
	}
	return nil, true
}

func (self *Controller) Login(w http.ResponseWriter, r *http.Request) {
	login := &LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(login)
	if err != nil {
		u.RespondError(w, 400, "Invalid request", nil)
		return
	}
	errValidation, isValid := validateLoginRequest(login)
	if !isValid {
		u.RespondError(w, http.StatusUnprocessableEntity, "Please fullfilled your form!", errValidation)
		return
	}
	user := &models.User{}
	errFindModel := models.GetDB().
		Where("email = ?", login.Email).
		First(&user).Error
	if errFindModel != nil {
		u.RespondError(w, http.StatusBadRequest, "Invalid credentials!", nil)
		return
	}

	isValidPassword, _ := h.Compare(user.Password, login.Password)
	if !isValidPassword {
		u.RespondError(w, http.StatusBadRequest, "Invalid credentials!", nil)
		return
	}

	user.GenerateToken()
	w.Header().Add("X_API_TOKEN", user.Token)
	u.Respond(w, 200, u.MessageWithData(200, "Succeeded!", user))
}

func validateRegisterRequest(register *RegisterRequest) (map[string]interface{}, bool) {
	errors := make(map[string][]string)
	if !strings.Contains(register.Email, "@") {
		msg := "Email address is required."
		x, y := errors["email"]
		if !y {
			errors["email"] = []string{msg}
		} else {
			errors["email"] = append(x, msg)
		}
	}
	if len(register.Password) < 6 {
		msg := "Password is required."
		x, y := errors["password"]
		if !y {
			errors["password"] = []string{msg}
		} else {
			errors["password"] = append(x, msg)
		}
	}
	if register.Password != register.ConfirmPassword {
		msg := "Confirm password not same with the password field."
		x, y := errors["confirm_password"]
		if !y {
			errors["confirm_password"] = []string{msg}
		} else {
			errors["confirm_password"] = append(x, msg)
		}
	}
	if len(errors) > 0 {
		return map[string]interface{}{"errors": errors}, false
	}
	return nil, true
}

func (self *Controller) Register(w http.ResponseWriter, r *http.Request) {
	register := &RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(register)
	if err != nil {
		u.RespondError(w, 400, "Invalid request", nil)
		return
	}
	errValidation, isValid := validateRegisterRequest(register)
	if !isValid {
		u.RespondError(w, http.StatusUnprocessableEntity, "Please fullfilled your form!", errValidation)
		return
	}
	user := &models.User{
		Email:    register.Email,
		Password: register.Password,
	}
	modelValidation, isValid := user.Validate()
	if !isValid {
		u.RespondError(w, http.StatusUnprocessableEntity, "Please fullfilled your form!", modelValidation)
		return
	}
	user.Save()
	user.GenerateToken()
	w.Header().Add("X_API_TOKEN", user.Token)
	u.Respond(w, 200, u.MessageWithData(200, "Succeeded!", user))
}
