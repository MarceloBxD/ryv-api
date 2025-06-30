package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Claims estrutura para claims do JWT
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	IsAdmin bool  `json:"is_admin"`
	jwt.RegisteredClaims
}

// AuthMiddleware middleware para autenticação JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token de autorização não fornecido",
			})
			c.Abort()
			return
		}

		// Verificar se o header começa com "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Formato de token inválido. Use 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		// Validar o token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token inválido ou expirado",
			})
			c.Abort()
			return
		}

		// Adicionar claims ao contexto
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("is_admin", claims.IsAdmin)

		c.Next()
	}
}

// AdminMiddleware middleware para verificar se o usuário é admin
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get("is_admin")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Usuário não autenticado",
			})
			c.Abort()
			return
		}

		if !isAdmin.(bool) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Acesso negado. Apenas administradores podem acessar este recurso",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RateLimitMiddleware middleware básico de rate limiting
func RateLimitMiddleware() gin.HandlerFunc {
	// Implementação básica - em produção usar Redis ou similar
	return func(c *gin.Context) {
		// Por enquanto, apenas passa adiante
		// TODO: Implementar rate limiting real
		c.Next()
	}
} 