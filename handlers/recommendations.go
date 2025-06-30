package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"ryv-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RecommendationHandler struct {
	db *gorm.DB
}

func NewRecommendationHandler(db *gorm.DB) *RecommendationHandler {
	return &RecommendationHandler{db: db}
}

// DailyRecommendation retorna uma recomendação inteligente baseada em psicologia e marketing
func (h *RecommendationHandler) DailyRecommendation(c *gin.Context) {
	var articles []models.Article
	
	// Buscar artigos publicados
	if err := h.db.Where("is_published = ?", true).Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar artigos"})
		return
	}

	if len(articles) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nenhum artigo encontrado"})
		return
	}

	// Algoritmo inteligente de recomendação baseado em:
	// 1. Psicologia: Escolher artigos que criam conexão emocional
	// 2. Marketing: Priorizar conteúdo que gera engajamento
	// 3. Timing: Considerar o momento do dia e contexto

	recommendation := h.selectIntelligentRecommendation(articles)
	
	// Calcular tempo de leitura estimado
	readingTime := h.calculateReadingTime(recommendation.Content)
	
	response := gin.H{
		"id":          recommendation.ID,
		"title":       recommendation.Title,
		"excerpt":     recommendation.Excerpt,
		"category":    recommendation.Category,
		"slug":        recommendation.Slug,
		"imageURL":    recommendation.ImageURL,
		"readingTime": readingTime,
		"motivation":  h.generateMotivation(recommendation.Category),
	}

	c.JSON(http.StatusOK, response)
}

func (h *RecommendationHandler) selectIntelligentRecommendation(articles []models.Article) models.Article {
	// Usar seed baseado na data para garantir consistência diária
	today := time.Now().Format("2006-01-02")
	seed := int64(0)
	for _, char := range today {
		seed += int64(char)
	}
	rand.Seed(seed)

	// Pontuação baseada em psicologia e marketing
	type scoredArticle struct {
		article models.Article
		score   float64
	}

	var scoredArticles []scoredArticle

	for _, article := range articles {
		score := h.calculateArticleScore(article)
		scoredArticles = append(scoredArticles, scoredArticle{
			article: article,
			score:   score,
		})
	}

	// Ordenar por pontuação
	for i := 0; i < len(scoredArticles)-1; i++ {
		for j := i + 1; j < len(scoredArticles); j++ {
			if scoredArticles[i].score < scoredArticles[j].score {
				scoredArticles[i], scoredArticles[j] = scoredArticles[j], scoredArticles[i]
			}
		}
	}

	// Retornar o artigo com maior pontuação
	return scoredArticles[0].article
}

func (h *RecommendationHandler) calculateArticleScore(article models.Article) float64 {
	score := 0.0

	// Fatores psicológicos
	psychologicalFactors := map[string]float64{
		"Saúde Mental": 1.5,  // Alto engajamento emocional
		"Ótica":        1.2,  // Interesse prático
		"Optometria":   1.3,  // Conhecimento técnico
		"Dicas de Saúde": 1.4, // Aplicação prática
	}

	if factor, exists := psychologicalFactors[article.Category]; exists {
		score += factor
	}

	// Fatores de marketing
	// Títulos que criam curiosidade
	curiosityWords := []string{"como", "por que", "quando", "descubra", "revele", "secreto"}
	for _, word := range curiosityWords {
		if contains(article.Title, word) {
			score += 0.3
		}
	}

	// Palavras emocionais
	emotionalWords := []string{"transformar", "conectar", "bem-estar", "saúde", "vida", "felicidade"}
	for _, word := range emotionalWords {
		if contains(article.Title, word) || contains(article.Excerpt, word) {
			score += 0.2
		}
	}

	// Fatores temporais
	now := time.Now()
	articleDate := *article.PublishedAt
	
	// Artigos mais recentes têm pontuação adicional
	daysSincePublished := now.Sub(articleDate).Hours() / 24
	if daysSincePublished < 7 {
		score += 0.5 // Artigos da última semana
	} else if daysSincePublished < 30 {
		score += 0.3 // Artigos do último mês
	}

	// Fatores de engajamento
	if article.ViewCount > 0 {
		score += float64(article.ViewCount) * 0.01 // Mais visualizações = mais popular
	}

	// Variação aleatória para evitar sempre o mesmo artigo
	score += rand.Float64() * 0.5

	return score
}

func (h *RecommendationHandler) calculateReadingTime(content string) string {
	// Média de 200 palavras por minuto
	wordCount := len(content) / 5 // Estimativa aproximada
	minutes := wordCount / 200
	
	if minutes < 1 {
		return "1 min"
	} else if minutes < 5 {
		return "2-3 min"
	} else if minutes < 10 {
		return "5-7 min"
	} else {
		return "10+ min"
	}
}

func (h *RecommendationHandler) generateMotivation(category string) string {
	motivations := map[string][]string{
		"Saúde Mental": {
			"💡 Desbloqueie insights poderosos sobre sua mente",
			"🧠 Conecte-se com seu bem-estar emocional",
			"🌟 Transforme sua perspectiva sobre saúde mental",
			"❤️ Cuide da sua mente como cuida do seu corpo",
		},
		"Ótica": {
			"👁️ Descubra como cuidar da sua visão",
			"🔍 Veja o mundo com novos olhos",
			"✨ Tecnologia que transforma sua experiência visual",
			"🌍 Enxergue a vida com mais clareza",
		},
		"Optometria": {
			"🔬 Ciência avançada para sua saúde ocular",
			"📊 Dados que revelam a verdade sobre sua visão",
			"🎯 Soluções precisas para problemas visuais",
			"⚡ Conhecimento que ilumina seu caminho",
		},
		"Dicas de Saúde": {
			"💪 Pequenas mudanças, grandes resultados",
			"🌱 Cultive hábitos que transformam sua vida",
			"⚡ Energia e vitalidade ao seu alcance",
			"🚀 Acelere seu potencial de bem-estar",
		},
	}

	if categoryMotivations, exists := motivations[category]; exists {
		return categoryMotivations[rand.Intn(len(categoryMotivations))]
	}
	
	return "🌟 Descubra insights valiosos para sua vida"
}

func contains(text, word string) bool {
	return len(text) >= len(word) && 
		   (text == word || 
		    (len(text) > len(word) && 
		     (text[:len(word)] == word || 
		      text[len(text)-len(word):] == word ||
		      containsSubstring(text, word))))
}

func containsSubstring(text, word string) bool {
	for i := 0; i <= len(text)-len(word); i++ {
		if text[i:i+len(word)] == word {
			return true
		}
	}
	return false
} 