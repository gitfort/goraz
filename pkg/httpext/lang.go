package httpext

import (
	"fmt"
	"golang.org/x/text/language"
	"net/http"
	"regexp"
)

var (
	languages = map[string]language.Tag{
		"en": language.English,
		"us": language.English,
		"uk": language.English,
		"de": language.German,
		"fa": language.Persian,
		"ir": language.Persian,
	}
)

func GetLang(req *http.Request) language.Tag {
	url := fmt.Sprintf("%v%v", req.Host, req.URL)
	for key, val := range languages {
		if ok, _ := regexp.MatchString(fmt.Sprintf("\\b%s\\b", key), url); ok {
			return val
		}
	}
	return language.Und
}
