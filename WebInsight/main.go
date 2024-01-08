package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func analyzeWebsite(url string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("falha ao criar a requisição: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro ao acessar o site: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("erro ao acessar o site. Status: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", fmt.Errorf("falha ao analisar o conteúdo HTML: %w", err)
	}

	title := doc.Find("title").Text()
	description := doc.Find("meta[name=description]").AttrOr("content", "Descrição não encontrada.")

	summary := fmt.Sprintf("Título: %s\nDescrição: %s\n", title, description)

	return summary, nil
}

func main() {
	url := ""
	summary, err := analyzeWebsite(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(summary)
}
