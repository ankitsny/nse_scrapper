package parser

import (
	"errors"
	"fmt"
	"net/http"
	"nse_scrapper/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// HTMLBodyToTextList :
func HTMLBodyToTextList(resp *http.Response) ([]string, error) {

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		fmt.Println("ERROR WHILE PARSING", err.Error())
		return nil, errors.New("Parse Error")
	}

	data := strings.Split(doc.Text(), "\n")

	data = utils.FilterStr(data, func(s string) bool {
		s = strings.TrimSpace(s)
		if len(s) == 0 || s == "\n" || s == ":" {
			return false
		}
		return true
	})

	data = utils.MapStr(data, func(s string) string {
		return strings.TrimSpace(s)
	})
	return data, nil
}
