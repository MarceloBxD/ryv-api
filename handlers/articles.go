package handlers

import (
	"net/http"
	"ryv-api/database"
	"ryv-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetArticles retorna todos os artigos publicados
func GetArticles(c *gin.Context) {
	var articles []models.Article
	
	query := database.DB.Where("is_published = ?", true).Order("published_at DESC")
	
	// Filtro por categoria
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}
	
	// Filtro por tag
	if tag := c.Query("tag"); tag != "" {
		query = query.Where("tags LIKE ?", "%"+tag+"%")
	}
	
	// Paginação
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit
	
	var total int64
	query.Model(&models.Article{}).Count(&total)
	
	if err := query.Offset(offset).Limit(limit).Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar artigos"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": (int(total) + limit - 1) / limit,
		},
	})
}

// GetArticle retorna um artigo específico por slug
func GetArticle(c *gin.Context) {
	slug := c.Param("slug")
	
	var article models.Article
	if err := database.DB.Where("slug = ? AND is_published = ?", slug, true).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artigo não encontrado"})
		return
	}
	
	// Incrementar contador de visualizações
	database.DB.Model(&article).Update("view_count", article.ViewCount+1)
	
	c.JSON(http.StatusOK, article)
}

// CreateArticle cria um novo artigo
func CreateArticle(c *gin.Context) {
	var article models.Article
	
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	
	if err := database.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar artigo"})
		return
	}
	
	c.JSON(http.StatusCreated, article)
}

// UpdateArticle atualiza um artigo existente
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	
	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artigo não encontrado"})
		return
	}
	
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	
	if err := database.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar artigo"})
		return
	}
	
	c.JSON(http.StatusOK, article)
}

// DeleteArticle remove um artigo
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	
	if err := database.DB.Delete(&models.Article{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar artigo"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Artigo deletado com sucesso"})
}

// GetCategories retorna todas as categorias
func GetCategories(c *gin.Context) {
	var categories []models.Category
	
	if err := database.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar categorias"})
		return
	}
	
	c.JSON(http.StatusOK, categories)
} 