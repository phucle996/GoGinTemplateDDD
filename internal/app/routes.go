package app

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, m *Modules) {
	api := r.Group("/api")
	{
		// api version 1 using handler version 1
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			// users.Use(middleware.AuthMiddleware())
			// if you use middleware for only accept request authenticated

			{
				users.POST("", m.V1UserHandler.CreateUser)
				users.GET("", m.V1UserHandler.ListUsers)
				users.GET("/:email", m.V1UserHandler.GetUserByEmail)
				users.GET("/:username", m.V1UserHandler.GetUserByUsername)
				users.PUT("/:user_id", m.V1UserHandler.SoftDeleteUser)
				users.PUT("/:user_id", m.V1UserHandler.UpdateUser)
			}
		}
		// api version 2 using handler version 2
		v2 := api.Group("/v2")
		{
			users := v2.Group("/users")
			{
				users.POST("", m.V2UserHandler.CreateUser)
				users.GET("", m.V2UserHandler.ListUsers)
				users.GET("/:email", m.V2UserHandler.GetUserByEmail)
				users.GET("/:username", m.V2UserHandler.GetUserByUsername)
				users.PUT("/:user_id", m.V2UserHandler.SoftDeleteUser)
				users.PUT("/:user_id", m.V2UserHandler.UpdateUser)
			}
		}

	}

}
