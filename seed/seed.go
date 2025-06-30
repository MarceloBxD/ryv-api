package main

import (
	"log"
	"ryv-api/database"
	"ryv-api/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Inicializar banco de dados
	database.InitDatabase()
	
	// Artigos de exemplo
	articles := []models.Article{
		{
			Title:       "Como a Saúde Ocular Afeta o Bem-Estar Mental",
			Slug:        "saude-ocular-bem-estar-mental",
			Content:     `<p>A conexão entre saúde ocular e bem-estar mental é mais profunda do que muitos imaginam. Nossos olhos são responsáveis por capturar cerca de 80% das informações que processamos diariamente, e quando há problemas de visão, isso pode impactar significativamente nossa qualidade de vida e saúde mental.</p>

<h2>Impacto dos Problemas Visuais na Saúde Mental</h2>
<p>Estudos mostram que pessoas com problemas de visão não corrigidos têm maior probabilidade de desenvolver:</p>
<ul>
<li>Ansiedade e estresse</li>
<li>Depressão</li>
<li>Isolamento social</li>
<li>Dificuldades de concentração</li>
</ul>

<h2>Sinais de que sua Visão Pode Estar Afetando seu Bem-Estar</h2>
<p>Fique atento a estes sinais:</p>
<ul>
<li>Dores de cabeça frequentes</li>
<li>Fadiga visual</li>
<li>Dificuldade para focar em tarefas</li>
<li>Irritabilidade ao ler ou usar telas</li>
</ul>

<h2>Como Cuidar da Saúde Ocular</h2>
<p>Algumas dicas importantes:</p>
<ul>
<li>Faça exames oftalmológicos regulares</li>
<li>Use óculos com proteção UV</li>
<li>Descanse os olhos a cada 20 minutos</li>
<li>Mantenha uma boa iluminação ao ler</li>
</ul>

<p>Lembre-se: cuidar da sua visão é cuidar da sua mente. Agende uma consulta conosco para avaliar sua saúde ocular!</p>`,
			Excerpt:     "Descubra como problemas de visão podem afetar sua saúde mental e bem-estar geral.",
			ImageURL:    "https://images.unsplash.com/photo-1571019613454-1cb2f99b2d8b?w=800&h=400&fit=crop",
			Category:    "Saúde Mental",
			Tags:        "saúde mental, visão, bem-estar, óculos",
			Author:      "Equipe RYV",
			IsPublished: true,
			PublishedAt: &[]time.Time{time.Now().AddDate(0, 0, -5)}[0],
		},
		{
			Title:       "Óculos de Sol: Mais que Moda, Proteção para seus Olhos",
			Slug:        "oculos-sol-protecao-olhos",
			Content:     `<p>Os óculos de sol não são apenas um acessório de moda - eles são essenciais para proteger seus olhos dos raios ultravioleta (UV) que podem causar danos irreversíveis à visão.</p>

<h2>Por que Proteger seus Olhos do Sol?</h2>
<p>A exposição excessiva aos raios UV pode causar:</p>
<ul>
<li>Catarata precoce</li>
<li>Degeneração macular</li>
<li>Pterígio (crescimento anormal da conjuntiva)</li>
<li>Queimaduras na córnea</li>
</ul>

<h2>Como Escolher Óculos de Sol Adequados</h2>
<p>Na hora de escolher seus óculos de sol, considere:</p>
<ul>
<li>Proteção UV 100%</li>
<li>Lentes polarizadas para reduzir reflexos</li>
<li>Tamanho adequado para cobrir toda a área dos olhos</li>
<li>Qualidade das lentes e armação</li>
</ul>

<h2>Diferentes Tipos de Lentes</h2>
<p>Cada tipo de lente tem sua função:</p>
<ul>
<li><strong>Lentes Cinzas:</strong> Reduzem a intensidade da luz sem alterar as cores</li>
<li><strong>Lentes Marrons:</strong> Melhoram o contraste e são ideais para dirigir</li>
<li><strong>Lentes Verdes:</strong> Oferecem boa percepção de cores</li>
<li><strong>Lentes Azuis:</strong> Reduzem o brilho da neve e água</li>
</ul>

<p>Visite nossa loja para encontrar os óculos de sol perfeitos para você!</p>`,
			Excerpt:     "Descubra por que os óculos de sol são essenciais para proteger seus olhos e como escolher o par ideal.",
			ImageURL:    "https://images.unsplash.com/photo-1511499767150-a48a237f0083?w=800&h=400&fit=crop",
			Category:    "Ótica",
			Tags:        "óculos de sol, proteção UV, saúde ocular, moda",
			Author:      "Equipe RYV",
			IsPublished: true,
			PublishedAt: &[]time.Time{time.Now().AddDate(0, 0, -3)}[0],
		},
		{
			Title:       "Ansiedade e Problemas Visuais: Uma Relação Bidirecional",
			Slug:        "ansiedade-problemas-visuais-relacao",
			Content:     `<p>A ansiedade e os problemas visuais frequentemente andam de mãos dadas, criando um ciclo que pode ser difícil de quebrar. Entender essa relação é fundamental para buscar o tratamento adequado.</p>

<h2>Como a Ansiedade Afeta a Visão</h2>
<p>A ansiedade pode causar diversos sintomas visuais:</p>
<ul>
<li>Visão embaçada temporária</li>
<li>Sensibilidade à luz</li>
<li>Olhos secos</li>
<li>Tremores visuais</li>
<li>Dificuldade de foco</li>
</ul>

<h2>Como Problemas Visuais Causam Ansiedade</h2>
<p>Por outro lado, problemas de visão podem gerar:</p>
<ul>
<li>Medo de perder a visão</li>
<li>Ansiedade social</li>
<li>Estresse ao dirigir ou trabalhar</li>
<li>Preocupação constante com a saúde</li>
</ul>

<h2>Quebrando o Ciclo</h2>
<p>Para interromper esse ciclo negativo:</p>
<ul>
<li>Procure um oftalmologista regularmente</li>
<li>Use óculos adequados quando necessário</li>
<li>Pratique técnicas de relaxamento</li>
<li>Mantenha uma rotina de exercícios</li>
<li>Considere terapia para ansiedade</li>
</ul>

<h2>Quando Buscar Ajuda</h2>
<p>Procure ajuda profissional se você:</p>
<ul>
<li>Está evitando atividades por medo de problemas visuais</li>
<li>Sentindo ansiedade constante relacionada à visão</li>
<li>Notando mudanças súbitas na visão</li>
<li>Preocupado excessivamente com a saúde dos olhos</li>
</ul>

<p>Lembre-se: cuidar da saúde mental e ocular é fundamental para uma vida plena e saudável.</p>`,
			Excerpt:     "Entenda a relação complexa entre ansiedade e problemas visuais e como quebrar esse ciclo.",
			ImageURL:    "https://images.unsplash.com/photo-1559757148-5c350d0d3c56?w=800&h=400&fit=crop",
			Category:    "Saúde Mental",
			Tags:        "ansiedade, visão, saúde mental, tratamento",
			Author:      "Equipe RYV",
			IsPublished: true,
			PublishedAt: &[]time.Time{time.Now().AddDate(0, 0, -1)}[0],
		},
		{
			Title:       "Lentes Progressivas: A Solução Moderna para Presbiopia",
			Slug:        "lentes-progressivas-solucao-presbiopia",
			Content:     `<p>A presbiopia, popularmente conhecida como "vista cansada", é uma condição natural que afeta a maioria das pessoas após os 40 anos. As lentes progressivas oferecem uma solução elegante e moderna para esse problema.</p>

<h2>O que é Presbiopia?</h2>
<p>A presbiopia é a perda gradual da capacidade de focar objetos próximos, causada pelo envelhecimento natural do cristalino. É um processo inevitável que afeta:</p>
<ul>
<li>Leitura de textos pequenos</li>
<li>Uso de smartphones</li>
<li>Trabalho no computador</li>
<li>Atividades que requerem visão de perto</li>
</ul>

<h2>Vantagens das Lentes Progressivas</h2>
<p>As lentes progressivas oferecem:</p>
<ul>
<li>Visão nítida em todas as distâncias</li>
<li>Transição suave entre as zonas de foco</li>
<li>Aparência natural (sem linha divisória)</li>
<li>Conforto visual superior</li>
<li>Adaptação rápida</li>
</ul>

<h2>Tipos de Lentes Progressivas</h2>
<p>Existem diferentes opções disponíveis:</p>
<ul>
<li><strong>Lentes Progressivas Standard:</strong> Solução econômica e eficaz</li>
<li><strong>Lentes Progressivas Premium:</strong> Campo de visão mais amplo</li>
<li><strong>Lentes Progressivas Digitais:</strong> Personalizadas para seu estilo de vida</li>
</ul>

<h2>Dicas para Adaptação</h2>
<p>Para uma adaptação mais rápida:</p>
<ul>
<li>Use os óculos constantemente nos primeiros dias</li>
<li>Mova a cabeça, não apenas os olhos</li>
<li>Pratique em ambientes familiares</li>
<li>Tenha paciência - a adaptação leva tempo</li>
</ul>

<p>Agende uma consulta conosco para descobrir se as lentes progressivas são a solução ideal para você!</p>`,
			Excerpt:     "Conheça as lentes progressivas, a solução moderna e eficaz para a presbiopia.",
			ImageURL:    "https://images.unsplash.com/photo-1582750433449-648ed127bb54?w=800&h=400&fit=crop",
			Category:    "Optometria",
			Tags:        "lentes progressivas, presbiopia, vista cansada, óculos",
			Author:      "Equipe RYV",
			IsPublished: true,
			PublishedAt: &[]time.Time{time.Now().AddDate(0, 0, -2)}[0],
		},
		{
			Title:       "Depressão e Isolamento: Como Problemas Visuais Podem Piorar",
			Slug:        "depressao-isolamento-problemas-visuais",
			Content:     `<p>Problemas visuais não tratados podem levar ao isolamento social e, em casos mais graves, à depressão. Entender essa conexão é crucial para buscar ajuda adequada.</p>

<h2>O Ciclo do Isolamento</h2>
<p>Quando há problemas de visão, muitas pessoas:</p>
<ul>
<li>Evitam sair de casa por medo de acidentes</li>
<li>Deixam de participar de atividades sociais</li>
<li>Sentem vergonha de usar óculos</li>
<li>Perdem a independência</li>
<li>Desenvolvem baixa autoestima</li>
</ul>

<h2>Sinais de Alerta</h2>
<p>Fique atento a estes sinais em você ou em alguém próximo:</p>
<ul>
<li>Mudança no comportamento social</li>
<li>Dificuldade para reconhecer pessoas</li>
<li>Evita atividades que antes gostava</li>
<li>Sentimentos de inutilidade</li>
<li>Alterações no sono ou apetite</li>
</ul>

<h2>Como Quebrar o Ciclo</h2>
<p>Algumas estratégias que podem ajudar:</p>
<ul>
<li>Buscar tratamento oftalmológico adequado</li>
<li>Usar óculos ou lentes de contato quando necessário</li>
<li>Participar de grupos de apoio</li>
<li>Manter contato com amigos e família</li>
<li>Considerar terapia psicológica</li>
</ul>

<h2>Prevenção é Fundamental</h2>
<p>Para evitar que problemas visuais afetem sua saúde mental:</p>
<ul>
<li>Faça exames oftalmológicos regulares</li>
<li>Use correção visual adequada</li>
<li>Mantenha uma rede de apoio social</li>
<li>Pratique atividades físicas</li>
<li>Busque ajuda profissional quando necessário</li>
</ul>

<p>Lembre-se: sua visão é preciosa e merece cuidados adequados. Não hesite em buscar ajuda!</p>`,
			Excerpt:     "Entenda como problemas visuais podem levar ao isolamento social e depressão, e como prevenir isso.",
			ImageURL:    "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=800&h=400&fit=crop",
			Category:    "Saúde Mental",
			Tags:        "depressão, isolamento, problemas visuais, saúde mental",
			Author:      "Equipe RYV",
			IsPublished: true,
			PublishedAt: &[]time.Time{time.Now().AddDate(0, 0, -4)}[0],
		},
	}
	
	// Inserir artigos
	for _, article := range articles {
		var existingArticle models.Article
		if err := database.DB.Where("slug = ?", article.Slug).First(&existingArticle).Error; err != nil {
			if err.Error() == "record not found" {
				if err := database.DB.Create(&article).Error; err != nil {
					log.Printf("Erro ao criar artigo %s: %v", article.Title, err)
				} else {
					log.Printf("✅ Artigo criado: %s", article.Title)
				}
			}
		} else {
			log.Printf("⚠️ Artigo já existe: %s", article.Title)
		}
	}
	
	// Criar usuário admin inicial
	adminEmail := "admin@ryv.com.br"
	adminPassword := "admin123" // Troque após o primeiro login!
	var existingAdmin models.User
	err := database.DB.Where("email = ?", adminEmail).First(&existingAdmin).Error
	if err != nil {
		hash, _ := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		admin := models.User{
			Name:         "Administrador",
			Email:        adminEmail,
			PasswordHash: string(hash),
			IsAdmin:      true,
		}
		if err := database.DB.Create(&admin).Error; err != nil {
			log.Printf("Erro ao criar admin: %v", err)
		} else {
			log.Printf("✅ Usuário admin criado: %s (senha: %s)", adminEmail, adminPassword)
		}
	} else {
		log.Printf("⚠️ Usuário admin já existe: %s", adminEmail)
	}
	
	log.Println("🎉 Seed concluído com sucesso!")
} 