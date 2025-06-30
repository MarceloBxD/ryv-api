package database

import (
	"log"
	"ryv-api/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDatabase inicializa a conexão com o banco de dados
func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("ryv_blog.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate das tabelas
	err = DB.AutoMigrate(&models.Article{}, &models.WhatsAppContact{}, &models.Category{}, &models.User{}, &models.ScrapedArticle{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Criar categorias padrão se não existirem
	createDefaultCategories()
	
	log.Println("Database connected and migrated successfully")
}

// createDefaultCategories cria as categorias padrão do blog
func createDefaultCategories() {
	categories := []models.Category{
		{
			Name:        "Saúde Mental",
			Slug:        "saude-mental",
			Description: "Artigos sobre saúde mental e bem-estar",
			Color:       "#10B981", // verde
		},
		{
			Name:        "Ótica",
			Slug:        "otica",
			Description: "Dicas e informações sobre óculos e lentes",
			Color:       "#3B82F6", // azul
		},
		{
			Name:        "Optometria",
			Slug:        "optometria",
			Description: "Informações técnicas sobre saúde ocular",
			Color:       "#8B5CF6", // roxo
		},
		{
			Name:        "Dicas de Saúde",
			Slug:        "dicas-saude",
			Description: "Dicas gerais de saúde e bem-estar",
			Color:       "#F59E0B", // amarelo
		},
	}

	for _, category := range categories {
		var existingCategory models.Category
		if err := DB.Where("slug = ?", category.Slug).First(&existingCategory).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				DB.Create(&category)
			}
		}
	}
} 