package handler

import "github.com/gin-gonic/gin"

// объект для работы с endpoints
type Handler struct {
}

// функция для инициализации endpoints
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.CreateList)         // создание задачи
			lists.GET("/", h.getAllLists)         // получение списка задач
			lists.GET("/:id", h.getListByID)      // получение информации о определенной задаче
			lists.PATCH("/:id", h.updateListByID) // обновление определенной задачи
			lists.DELETE(":id", h.deleteListByID) // удаление определенной задачи
		}
	}

	return router
}
