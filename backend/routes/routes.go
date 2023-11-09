package routes

import (
	"github.com/A-Cer23/adminbot-backend/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(router *gin.Engine, db *pgxpool.Pool) {
	api := router.Group("/api")
	{
		guild := api.Group("/guild")
		{
			guild.POST("/", func(c *gin.Context) {
				handlers.CreateGuild(c, db)
			})
		}
	}
}
