package main

import (
	"log"
	"net/http"
)

func main() {
	// Servir arquivos estáticos
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	log.Println("🌐 Servidor do painel administrativo iniciado na porta 3000")
	log.Println("📱 Acesse: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
} 