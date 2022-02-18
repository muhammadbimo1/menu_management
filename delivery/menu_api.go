package delivery

import (
	"menu_management/delivery/appresponse"
	"menu_management/usecase"

	"github.com/gin-gonic/gin"
)

type MenuApi struct {
	usecase     usecase.MenuUseCase
	publicRoute *gin.RouterGroup
}

func NewMenuApi(publicRoute *gin.RouterGroup, useCase usecase.MenuUseCase) *MenuApi {
	api := MenuApi{
		usecase:     useCase,
		publicRoute: publicRoute,
	}
	api.initRouter()
	return &api
}

func (a *MenuApi) initRouter() {
	menuRoute := a.publicRoute.Group("/menu")
	menuRoute.GET("", a.getMenu)
}

func (a *MenuApi) getMenu(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	response := appresponse.NewJsonResponse(c)
	if id != "" {
		menu, err := a.usecase.SearchMenuById(id)
		if err != nil {
			response.SendError(*appresponse.NewInternalServerError(err, "error"))
			return
		}
		response.SendData(appresponse.NewResponseMessage("success", "menu by id", menu))
		return
	}
	if name != "" {
		menuList, err := a.usecase.SearchMenuByName(name)
		if err != nil {
			response.SendError(*appresponse.NewInternalServerError(err, "error"))
			return
		}
		response.SendData(appresponse.NewResponseMessage("success", "menu by name", menuList))
		return
	}
	menulist, err := a.usecase.GetAllMenu()
	if err != nil {
		response.SendError(*appresponse.NewInternalServerError(err, "error"))
		return
	}
	response.SendData(appresponse.NewResponseMessage("success", "all menu", menulist))
}
