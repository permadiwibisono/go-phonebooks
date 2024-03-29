package bootstrap

import (
	"fmt"
	"go-phonebooks/config"
	"go-phonebooks/controllers"
	"go-phonebooks/middlewares"
	"go-phonebooks/models"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (app *App) dbConnect(config *config.DBConfig) *gorm.DB {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset,
		config.Loc,
	)
	fmt.Println(dbURI)

	conn, err := gorm.Open(config.Dialect, dbURI)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected with databases...")
	}
	return conn
}

func (app *App) EnvLoad() {
	fmt.Println("Load environment variables...")
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
}

func (app *App) Initialize() {
	app.EnvLoad()
	config := config.GetConfig()
	db := app.dbConnect(config.DB)
	app.DB = models.AutoMigrate(db)
	app.Router = mux.NewRouter()
	app.setRouters()
}

func (app *App) setRouters() {
	app.RegisterControllerRouters(controllers.HomeController)
	app.RegisterControllerRouters(controllers.AuthController)
	app.RegisterControllerRouters(controllers.ContactController)
	app.Router.Use(middlewares.Logger)
	app.Router.Use(middlewares.PaginationQueryParams)
}

func (app *App) RegisterControllerRouters(ctrl controllers.IController) {
	entryPoint := "/api"
	apiRouter := app.Router.PathPrefix(entryPoint).Subrouter()
	routeList := ctrl.GetRoutes()
	for key, r := range routeList {
		regex, _ := regexp.Compile("/$")
		prefixURL := ctrl.GetPrefixUrl()
		prefixURL = regex.ReplaceAllString(prefixURL, "")
		path := prefixURL + r.URL
		fmt.Printf("PATH: %s%s\n", entryPoint, path)
		middlewaresArray := ctrl.GetMiddlewares()[key]
		app.SetRoute(apiRouter, &r, path, middlewaresArray)
	}
}

func (app *App) SetRoute(r *mux.Router, routeSetting *controllers.Route, path string, middlewaresArray []string) {
	if len(middlewaresArray) > 0 || middlewaresArray != nil {
		handlerFunc := http.HandlerFunc(app.wrapRequestHandler(routeSetting.Handler))
		enabledMiddlewares := []string{}
		var midFuncs http.Handler
		for _, value := range middlewaresArray {
			if value == "jwt" || value == "hello" {
				enabledMiddlewares = append(enabledMiddlewares, value)
			}
		}
		if len(enabledMiddlewares) > 0 {
			// fmt.Println(enabledMiddlewares)
			midFuncs = app.ApplyMiddlewares(midFuncs, enabledMiddlewares, 0, handlerFunc)
		}
		if midFuncs != nil {
			r.Handle(path, midFuncs).
				Methods(routeSetting.Method).
				Name(routeSetting.Name)
		} else {
			r.HandleFunc(path, app.wrapRequestHandler(routeSetting.Handler)).
				Methods(routeSetting.Method).
				Name(routeSetting.Name)
		}
	} else {
		r.HandleFunc(path, app.wrapRequestHandler(routeSetting.Handler)).
			Methods(routeSetting.Method).
			Name(routeSetting.Name)
	}
}

func setMiddleware(middlewareName string, handlerFunc http.HandlerFunc) http.Handler {
	if middlewareName == "jwt" {
		return middlewares.JwtAuthMiddleware(handlerFunc)
	} else if middlewareName == "hello" {
		return middlewares.HelloMiddleware(handlerFunc)
	}
	return nil
}

func appendMiddleware(middlewareName string, otherMiddlewares http.Handler) http.Handler {
	if middlewareName == "jwt" {
		return middlewares.JwtAuthMiddleware(otherMiddlewares)
	} else if middlewareName == "hello" {
		return middlewares.HelloMiddleware(otherMiddlewares)
	}
	return otherMiddlewares
}

func (app *App) ApplyMiddlewares(myMiddleware http.Handler, middlewareArr []string, start int, handler http.HandlerFunc) http.Handler {
	if start < len(middlewareArr) {
		if myMiddleware == nil {
			myMiddleware = setMiddleware(middlewareArr[start], handler)
		} else {
			myMiddleware = appendMiddleware(middlewareArr[start], myMiddleware)
		}
		return app.ApplyMiddlewares(myMiddleware, middlewareArr, start+1, handler)
	}
	return myMiddleware
}

func (app *App) wrapRequestHandler(handler controllers.RequestHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, app.DB)
	}
}
