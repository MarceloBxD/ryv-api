package models

import (
	"time"

	"gorm.io/gorm"
)

// Article representa um artigo do blog
type Article struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Slug        string         `json:"slug" gorm:"uniqueIndex;not null"`
	Content     string         `json:"content" gorm:"type:text"`
	Excerpt     string         `json:"excerpt"`
	ImageURL    string         `json:"image_url"`
	Category    string         `json:"category"` // saúde mental, ótica, optometria
	Tags        string         `json:"tags"`     // tags separadas por vírgula
	Author      string         `json:"author"`
	AuthorID    *uint          `json:"author_id"`
	SourceURL   string         `json:"source_url"`
	PublishedAt *time.Time     `json:"published_at"`
	IsPublished bool           `json:"is_published" gorm:"default:false"`
	ViewCount   int            `json:"view_count" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// WhatsAppContact representa um contato via WhatsApp
type WhatsAppContact struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Phone     string         `json:"phone" gorm:"not null"`
	Message   string         `json:"message"`
	Source    string         `json:"source"` // página onde o contato foi feito
	ArticleID *uint          `json:"article_id,omitempty"` // se foi feito a partir de um artigo
	IPAddress string         `json:"ip_address"`
	UserAgent string         `json:"user_agent"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// Category representa uma categoria de artigos
type Category struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"uniqueIndex;not null"`
	Slug        string         `json:"slug" gorm:"uniqueIndex;not null"`
	Description string         `json:"description"`
	Color       string         `json:"color"` // cor para UI
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// User representa um usuário administrador
type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"not null"`
	Email        string         `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string         `json:"password_hash" gorm:"not null"`
	IsAdmin      bool           `json:"is_admin" gorm:"default:false"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// ScrapedArticle representa uma sugestão de post vinda do scraper
type ScrapedArticle struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	Excerpt   string         `json:"excerpt"`
	Content   string         `json:"content" gorm:"type:text"`
	ImageURL  string         `json:"image_url"`
	SourceURL string         `json:"source_url"`
	Suggested bool           `json:"suggested" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
} 