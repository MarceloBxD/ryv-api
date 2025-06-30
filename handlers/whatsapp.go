package handlers

import (
	"net/http"
	"ryv-api/database"
	"ryv-api/models"

	"github.com/gin-gonic/gin"
)

// CreateWhatsAppContact registra um novo contato via WhatsApp
func CreateWhatsAppContact(c *gin.Context) {
	var contact models.WhatsAppContact
	
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	
	// Capturar informações do cliente
	contact.IPAddress = c.ClientIP()
	contact.UserAgent = c.GetHeader("User-Agent")
	
	// Se não foi especificada uma fonte, usar a página atual
	if contact.Source == "" {
		contact.Source = c.GetHeader("Referer")
	}
	
	if err := database.DB.Create(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar contato"})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "Contato registrado com sucesso",
		"contact": contact,
	})
}

// GetWhatsAppContacts retorna todos os contatos (para admin)
func GetWhatsAppContacts(c *gin.Context) {
	var contacts []models.WhatsAppContact
	
	if err := database.DB.Order("created_at DESC").Find(&contacts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar contatos"})
		return
	}
	
	c.JSON(http.StatusOK, contacts)
}

// GetWhatsAppContactStats retorna estatísticas dos contatos
func GetWhatsAppContactStats(c *gin.Context) {
	var totalContacts int64
	var todayContacts int64
	var thisWeekContacts int64
	var thisMonthContacts int64
	
	// Total de contatos
	database.DB.Model(&models.WhatsAppContact{}).Count(&totalContacts)
	
	// Contatos de hoje
	database.DB.Model(&models.WhatsAppContact{}).
		Where("DATE(created_at) = DATE('now')").
		Count(&todayContacts)
	
	// Contatos desta semana
	database.DB.Model(&models.WhatsAppContact{}).
		Where("created_at >= DATE('now', '-7 days')").
		Count(&thisWeekContacts)
	
	// Contatos deste mês
	database.DB.Model(&models.WhatsAppContact{}).
		Where("created_at >= DATE('now', '-30 days')").
		Count(&thisMonthContacts)
	
	c.JSON(http.StatusOK, gin.H{
		"total":      totalContacts,
		"today":      todayContacts,
		"this_week":  thisWeekContacts,
		"this_month": thisMonthContacts,
	})
} 