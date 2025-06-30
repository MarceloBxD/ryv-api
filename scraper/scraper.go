package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ScrapedArticle struct {
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Excerpt     string    `json:"excerpt"`
	URL         string    `json:"url"`
	Source      string    `json:"source"`
	PublishedAt time.Time `json:"published_at"`
	Category    string    `json:"category"`
	Tags        []string  `json:"tags"`
}

type ScraperService struct {
	client *http.Client
}

func NewScraperService() *ScraperService {
	return &ScraperService{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Sites para scraping (focados em saúde mental, óptica e bem-estar)
var targetSites = []struct {
	URL      string
	Selector string
	Category string
}{
	{
		URL:      "https://www.psychologytoday.com/us/blog",
		Selector: ".blog-post",
		Category: "Saúde Mental",
	},
	{
		URL:      "https://www.verywellmind.com",
		Selector: ".article-card",
		Category: "Saúde Mental",
	},
	{
		URL:      "https://www.allaboutvision.com",
		Selector: ".article",
		Category: "Ótica",
	},
	{
		URL:      "https://www.aoa.org/news",
		Selector: ".news-item",
		Category: "Optometria",
	},
	{
		URL:      "https://www.healthline.com/health/mental-health",
		Selector: ".article-card",
		Category: "Saúde Mental",
	},
}

func (s *ScraperService) ScrapeArticles() ([]ScrapedArticle, error) {
	var articles []ScrapedArticle

	for _, site := range targetSites {
		siteArticles, err := s.scrapeSite(site.URL, site.Selector, site.Category)
		if err != nil {
			fmt.Printf("Erro ao fazer scraping de %s: %v\n", site.URL, err)
			continue
		}
		articles = append(articles, siteArticles...)
	}

	return articles, nil
}

func (s *ScraperService) scrapeSite(url, selector, category string) ([]ScrapedArticle, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Headers para simular um navegador real
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var articles []ScrapedArticle

	doc.Find(selector).Each(func(i int, selection *goquery.Selection) {
		article := s.extractArticle(selection, category, url)
		if article.Title != "" && article.Content != "" {
			articles = append(articles, article)
		}
	})

	return articles, nil
}

func (s *ScraperService) extractArticle(selection *goquery.Selection, category, baseURL string) ScrapedArticle {
	article := ScrapedArticle{
		Category: category,
		Source:   baseURL,
	}

	// Extrair título
	title := selection.Find("h1, h2, h3, .title, .headline").First().Text()
	article.Title = strings.TrimSpace(title)

	// Extrair conteúdo
	content := selection.Find("p, .content, .body").Text()
	article.Content = strings.TrimSpace(content)

	// Extrair excerpt (primeiros 200 caracteres)
	if len(article.Content) > 200 {
		article.Excerpt = article.Content[:200] + "..."
	} else {
		article.Excerpt = article.Content
	}

	// Extrair URL
	if href, exists := selection.Find("a").First().Attr("href"); exists {
		if strings.HasPrefix(href, "http") {
			article.URL = href
		} else {
			article.URL = baseURL + href
		}
	}

	// Extrair tags
	tags := selection.Find(".tags, .categories, .keywords").Text()
	if tags != "" {
		article.Tags = strings.Split(tags, ",")
		for i, tag := range article.Tags {
			article.Tags[i] = strings.TrimSpace(tag)
		}
	}

	// Definir data de publicação (padrão para agora)
	article.PublishedAt = time.Now()

	return article
}

// Função para buscar artigos relacionados usando APIs externas
func (s *ScraperService) SearchRelatedContent(query string) ([]ScrapedArticle, error) {
	// Exemplo usando NewsAPI (você precisará de uma API key)
	apiKey := "YOUR_NEWS_API_KEY" // Substitua pela sua chave
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&language=pt&sortBy=publishedAt&apiKey=%s", 
		query, apiKey)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var newsResponse struct {
		Articles []struct {
			Title       string    `json:"title"`
			Description string    `json:"description"`
			URL         string    `json:"url"`
			PublishedAt time.Time `json:"publishedAt"`
			Source      struct {
				Name string `json:"name"`
			} `json:"source"`
		} `json:"articles"`
	}

	if err := json.Unmarshal(body, &newsResponse); err != nil {
		return nil, err
	}

	var articles []ScrapedArticle
	for _, newsArticle := range newsResponse.Articles {
		article := ScrapedArticle{
			Title:       newsArticle.Title,
			Content:     newsArticle.Description,
			Excerpt:     newsArticle.Description,
			URL:         newsArticle.URL,
			Source:      newsArticle.Source.Name,
			PublishedAt: newsArticle.PublishedAt,
			Category:    "Notícias Relacionadas",
		}
		articles = append(articles, article)
	}

	return articles, nil
}

// Função para limpar e processar conteúdo
func (s *ScraperService) ProcessContent(content string) string {
	// Remover HTML tags
	re := regexp.MustCompile(`<[^>]*>`)
	content = re.ReplaceAllString(content, "")

	// Remover caracteres especiais
	content = strings.ReplaceAll(content, "&nbsp;", " ")
	content = strings.ReplaceAll(content, "&amp;", "&")
	content = strings.ReplaceAll(content, "&lt;", "<")
	content = strings.ReplaceAll(content, "&gt;", ">")

	// Limpar espaços extras
	content = strings.TrimSpace(content)
	re = regexp.MustCompile(`\s+`)
	content = re.ReplaceAllString(content, " ")

	return content
}

// Função para categorizar automaticamente o conteúdo
func (s *ScraperService) CategorizeContent(title, content string) string {
	text := strings.ToLower(title + " " + content)

	keywords := map[string][]string{
		"Saúde Mental": {
			"ansiedade", "depressão", "estresse", "bem-estar", "psicologia",
			"terapia", "meditação", "mindfulness", "saúde mental", "emocional",
		},
		"Ótica": {
			"óculos", "lentes", "visão", "olhos", "óptica", "armação",
			"proteção uv", "óculos de sol", "lentes progressivas",
		},
		"Optometria": {
			"optometria", "exame ocular", "oftalmologia", "presbiopia",
			"miopia", "astigmatismo", "catarata", "glaucoma",
		},
		"Dicas de Saúde": {
			"saúde", "bem-estar", "qualidade de vida", "hábitos saudáveis",
			"prevenção", "cuidados", "dicas", "conselhos",
		},
	}

	maxScore := 0
	bestCategory := "Dicas de Saúde"

	for category, words := range keywords {
		score := 0
		for _, word := range words {
			if strings.Contains(text, word) {
				score++
			}
		}
		if score > maxScore {
			maxScore = score
			bestCategory = category
		}
	}

	return bestCategory
} 