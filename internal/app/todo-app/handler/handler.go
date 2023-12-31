package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyabukanov123/todo-app/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	// инициалиазаци роута
	router := gin.New()

	// объявляем методы сгупировав их по маршрутам
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-id", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}

		items := lists.Group(":id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:items_id", h.getItemById)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteList)
		}
	}
	return router
}
