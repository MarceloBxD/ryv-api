package main

import (
	"log"
	"ryv-api/database"
	"ryv-api/handlers"
	"ryv-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar banco de dados
	database.InitDatabase()
	db := database.DB

	// Configurar Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Configurar CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Inicializar handlers
	recommendationHandler := handlers.NewRecommendationHandler(db)
	authHandler := handlers.NewAuthHandler(db)

	// Rotas da API
	api := r.Group("/api")
	{
		// Rotas públicas de artigos
		articles := api.Group("/articles")
		{
			articles.GET("", handlers.GetArticles)
			articles.GET("/categories", handlers.GetCategories)
			articles.GET("/:slug", handlers.GetArticle)
		}

		// Rota de recomendação diária
		api.GET("/articles/daily-recommendation", recommendationHandler.DailyRecommendation)

		// Rotas do WhatsApp
		whatsapp := api.Group("/whatsapp")
		{
			whatsapp.POST("/contact", handlers.CreateWhatsAppContact)
		}

		// Rotas de autenticação
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
			auth.POST("/create-admin", authHandler.CreateAdmin) // Apenas para setup inicial
		}

		// Rotas protegidas (requerem autenticação)
		protected := api.Group("/admin")
		protected.Use(middleware.AuthMiddleware())
		{
			// Perfil do usuário
			protected.GET("/profile", authHandler.GetProfile)

			// Rotas de artigos (admin)
			adminArticles := protected.Group("/articles")
			{
				adminArticles.POST("", handlers.CreateArticle)
				adminArticles.PUT("/:id", handlers.UpdateArticle)
				adminArticles.DELETE("/:id", handlers.DeleteArticle)
			}

			// Rotas de contatos WhatsApp (admin)
			adminWhatsApp := protected.Group("/whatsapp")
			{
				adminWhatsApp.GET("/contacts", handlers.GetWhatsAppContacts)
				adminWhatsApp.GET("/stats", handlers.GetWhatsAppContactStats)
			}
		}
	}

	// Rota de health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "RYV API está funcionando!",
		})
	})

	log.Println("🚀 Servidor RYV API iniciado na porta 3001")
	log.Fatal(r.Run(":3001"))
} 