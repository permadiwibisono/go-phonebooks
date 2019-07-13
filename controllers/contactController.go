package controllers

import (
	"fmt"
	"go-phonebooks/middlewares"
	"go-phonebooks/models"
	u "go-phonebooks/utils"
	"net/http"

	"github.com/jinzhu/gorm"
)

type ContactControllerType struct {
	Controller
}

func (i *ContactControllerType) GetPrefixUrl() string {
	return i.PrefixURL
}

func (i *ContactControllerType) GetRoutes() map[string]Route {
	return i.Routes
}
func (i *ContactControllerType) GetMiddlewares() map[string][]string {
	return i.Middlewares
}

var ContactController = &ContactControllerType{}

func init() {
	ContactController.PrefixURL = "/contacts"

	routes := map[string]Route{
		"Index": Route{
			Method:  http.MethodGet,
			Name:    "Contacts.Get.List",
			Handler: ContactController.Index,
		},
	}
	middlewares := map[string][]string{
		"Index": []string{"jwt"},
	}
	ContactController.Routes = routes
	ContactController.Middlewares = middlewares
}

func (self *ContactControllerType) Index(w http.ResponseWriter, r *http.Request, DB *gorm.DB) {
	userID := r.Context().Value("user").(uint)
	getPaginationQuery := r.Context().Value("pagination")
	pagination := middlewares.PaginationQuery{
		Page:    1,
		PerPage: 16,
	}
	if getPaginationQuery != nil {
		pagination = getPaginationQuery.(middlewares.PaginationQuery)
	}
	fmt.Printf("Your pagination: Page %d, PerPage %d\n", pagination.Page, pagination.PerPage)
	contacts := &[]models.Contact{}
	dataPagination := &models.Pagination{}
	queries := map[string]interface{}{"user_id": userID}
	err := DB.Model(&models.Contact{}).
		Preload("User").
		Where(queries).
		Order("is_favorited desc").
		Scopes(models.ScopePaginate(int(pagination.Page), int(pagination.PerPage), &models.Contact{}, dataPagination)).
		Find(&contacts).Error
	if err != nil {
		u.RespondError(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	dataPagination.Data = contacts
	u.Respond(w, 200, u.MessageWithData(200, "Get contact list.", dataPagination))
}
