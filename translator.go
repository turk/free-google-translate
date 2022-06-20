package translator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Translator struct {
	client *http.Client
}

func NewTranslator(client *http.Client) Translator {
	return Translator{client: client}
}

func (t Translator) Translate(text, sourceLang, targetLang string) (string, error) {
	var result []interface{}
	var translated []string

	urlStr := fmt.Sprintf(
		"https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s",
		sourceLang,
		targetLang,
		url.QueryEscape(text),
	)

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	res, err := t.client.Do(req)

	if err != nil {
		return "err", errors.New("error getting translate.googleapis.com")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "err", errors.New("error reading response body")
	}

	if res.StatusCode != 200 {
		return "err", errors.New("translation failed")
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "err", errors.New("error unmarshaling body")
	}

	if len(result) > 0 {
		data := result[0]
		for _, slice := range data.([]interface{}) {
			for _, translatedText := range slice.([]interface{}) {
				translated = append(translated, fmt.Sprintf("%v", translatedText))
				break
			}
		}
		return strings.Join(translated, ""), nil
	}
	return "err", errors.New("translation not found")
}
