package controllers

import (
	"fmt"
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
	paginationQuery := u.GetPaginationQueryParams(r)
	fmt.Printf("Your pagination: Page %d, PerPage %d\n", paginationQuery.Page, paginationQuery.PerPage)
	contacts := &[]models.Contact{}
	dataPagination := &models.Pagination{}
	dataPagination.Data = contacts
	queries := map[string]interface{}{"user_id": userID}
	err := DB.Model(&models.Contact{}).
		Preload("User").
		Where(queries).
		Order("is_favorited desc").
		Scopes(models.ScopePaginate(paginationQuery.Page, paginationQuery.PerPage, &models.Contact{}, dataPagination)).
		Find(&contacts).Error
	if err != nil {
		u.RespondError(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	u.Respond(w, 200, u.MessageWithData(200, "Get contact list.", dataPagination))
}
