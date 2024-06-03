package router

import (
    "BTPN/controllers"
    "BTPN/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
    r.POST("/users/register", controllers.Register)
    r.POST("/users/login", controllers.Login)

    userRoutes := r.Group("/users")
    {
        userRoutes.PUT("/:userId", middlewares.Auth(), controllers.UpdateUser)
        userRoutes.DELETE("/:userId", middlewares.Auth(), controllers.DeleteUser)
    }

    photoRoutes := r.Group("/photos")
    {
        photoRoutes.POST("/", middlewares.Auth(), controllers.CreatePhoto)
        photoRoutes.GET("/", controllers.GetPhotos)
        photoRoutes.PUT("/:photoId", middlewares.Auth(), controllers.UpdatePhoto)
        photoRoutes.DELETE("/:photoId", middlewares.Auth(), controllers.DeletePhoto)
    }
}
