package utils

import (
	"log"
	"strings"

	tu "github.com/mymmrac/telego/telegoutil"
)

func TrimStrings(items []string) []string {
	result := make([]string, len(items))
	for i, s := range items {
		result[i] = strings.TrimSpace(s)
	}
	return result
}

func SplitTranslations(text string) []string {
	return TrimStrings(strings.Split(text, ","))
}

func EqualStrings(a, b string) bool {
	return strings.EqualFold(strings.TrimSpace(a), strings.TrimSpace(b))
}

func JoinStrings(items []string) string {
	return strings.Join(items, ", ")
}

func StringsToEntities(translations []string) []tu.MessageEntityCollection {
	log.Printf("translations: %v", translations)
	if len(translations) == 0 {
		return []tu.MessageEntityCollection{}
	}
	if len(translations) == 1 {
		return []tu.MessageEntityCollection{tu.Entity(translations[0]).Code()}
	}
	result := make([]tu.MessageEntityCollection, len(translations)*2-1)
	for i, t := range translations {
		index := i * 2
		if index == 0 {
			result[index] = tu.Entity(t).Code()
			continue
		}
		result[index-1] = tu.Entity(", ")
		result[index] = tu.Entity(t).Code()
	}
	return result
}

func Prepend[A interface{}](a []A, b ...A) []A {
	return append(b, a...)
}
