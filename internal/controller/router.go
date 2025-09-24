package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (ctrl *Controller) RegisterEndpoints() {

	ctrl.router.GET("/ping", ctrl.Ping)
	ctrl.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ctrl.router.GET("/employees", ctrl.GetAllEmployees)
	ctrl.router.GET("/employees/:id", ctrl.GetEmployeesByID)
	ctrl.router.POST("/employees", ctrl.CreateEmployees)
	ctrl.router.PUT("/employees/:id", ctrl.UpdateEmployeesByID)
	ctrl.router.DELETE("/employees/:id", ctrl.DeleteEmployeesByID)
}

// Ping
// @Summary Health-check
// @Description Проверка сервиса
// @Tags Ping
// @Produce json
// @Success 200 {object} CommonResponse
// @Router /ping [get]
func (ctrl *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, CommonResponse{Message: "Server is up and running!"})
}

func (ctrl *Controller) RunServer(address string) error {
	// Регистрируем роуты
	ctrl.RegisterEndpoints()

	// Запускаем http-сервер
	if err := ctrl.router.Run(address); err != nil {
		return err
	}

	return nil
}
