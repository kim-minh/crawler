package utils

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jackc/pgx/v5/pgtype"
)

func ExtractText(htmlString pgtype.Text) pgtype.Text {
	newHtml := strings.ReplaceAll(htmlString.String, "&lt;", "<")
	newHtml = strings.ReplaceAll(newHtml, "&gt;", ">")

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(newHtml))
	if err != nil {
		LogError(err)
		return pgtype.Text{}
	}

	var textBuilder strings.Builder
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		textBuilder.WriteString(s.Text())
	})

	regex := regexp.MustCompile(`\\(.)`)
	text := regex.ReplaceAllString(textBuilder.String(), "")
	text = strings.TrimSpace(text)
	return pgtype.Text{String: text, Valid: true}
}
