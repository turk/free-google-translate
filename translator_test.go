package translator

import (
	"net/http"
	"testing"
)

func TestTranslate(t *testing.T) {
	client := http.Client{}
	translator := NewTranslator(&client)
	text := "Hello"
	mdText := "Buna ziua"

	translated, _ := translator.Translate(text, "en", "ro")
	if translated != mdText {
		t.Logf("expected: %s", mdText)
		t.Logf("given: %s", translated)
		t.Fail()
	}
}
