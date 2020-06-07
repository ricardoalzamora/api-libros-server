package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardoalzamora/books_server/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("pais", controllers.GetAllPaises)
		v1.POST("pais", controllers.CreateAPais)
		v1.GET("pais/:id", controllers.GetAPais)
		v1.PUT("pais/:id", controllers.UpdateAPais)
		v1.DELETE("pais/:id", controllers.DeleteAPais)
	}

	return r
}
