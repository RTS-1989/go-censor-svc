package textParser

import (
	"github.com/RTS-1989/go-censor-svc/pkg/models"
	"strings"
)

func CheckCensored(text string, censored []*models.CensoredWords) bool {
	words := strings.Split(text, " ")
	for _, word := range words {
		for _, cw := range censored {
			if word == cw.Word {
				return true
			}
		}
	}

	return false
}
