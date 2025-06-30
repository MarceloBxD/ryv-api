package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"ryv-api/database"
	"ryv-api/models"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("🔧 Criando primeiro administrador do sistema")
	fmt.Println("=============================================")

	// Inicializar banco de dados
	database.InitDatabase()

	// Verificar se já existe um admin
	var adminCount int64
	database.DB.Model(&models.User{}).Where("is_admin = ?", true).Count(&adminCount)
	
	if adminCount > 0 {
		fmt.Println("⚠️  Já existe pelo menos um administrador no sistema.")
		fmt.Print("Deseja continuar mesmo assim? (s/N): ")
		
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))
		
		if response != "s" && response != "sim" && response != "y" && response != "yes" {
			fmt.Println("Operação cancelada.")
			return
		}
	}

	// Coletar dados do admin
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nome completo: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Senha (mínimo 6 caracteres): ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Validações básicas
	if len(name) < 2 {
		log.Fatal("Nome deve ter pelo menos 2 caracteres")
	}

	if len(email) < 5 || !strings.Contains(email, "@") {
		log.Fatal("Email inválido")
	}

	if len(password) < 6 {
		log.Fatal("Senha deve ter pelo menos 6 caracteres")
	}

	// Verificar se o email já existe
	var existingUser models.User
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		log.Fatal("Email já cadastrado no sistema")
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Erro ao processar senha:", err)
	}

	// Criar usuário admin
	admin := models.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hashedPassword),
		IsAdmin:      true,
	}

	if err := database.DB.Create(&admin).Error; err != nil {
		log.Fatal("Erro ao criar administrador:", err)
	}

	fmt.Println("✅ Administrador criado com sucesso!")
	fmt.Printf("Nome: %s\n", admin.Name)
	fmt.Printf("Email: %s\n", admin.Email)
	fmt.Printf("ID: %d\n", admin.ID)
	fmt.Println("\n🔐 Use estas credenciais para fazer login no painel administrativo")
	fmt.Println("🌐 Acesse: http://localhost:3001/api/auth/login")
} 