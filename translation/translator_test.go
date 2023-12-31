package translation_test

import (
	"testing"

	"github.com/mtolley/hello-api/translation"
)

func TestTranslate(t *testing.T) {
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{Word: "hello", Language: "english", Translation: "hello"},
		{Word: "hello", Language: "german", Translation: "hallo"},
		{Word: "hello", Language: "finnish", Translation: "hei"},
		{Word: "hello", Language: "dutch", Translation: ""},
		{Word: "bye", Language: "dutch", Translation: ""},
		{Word: "bye", Language: "german", Translation: ""},
		{Word: "hello", Language: "German", Translation: "hallo"},
		{Word: "Hello", Language: "german", Translation: "hallo"},
		{Word: "hello ", Language: "german", Translation: "hallo"},
		{Word: "hello ", Language: "french", Translation: "bonjour"},
	}

	for _, test := range tt {
		res := translation.Translate(test.Word, test.Language)
		if res != test.Translation {
			t.Errorf(`"expected %s to be %s from %s but received %s`, test.Word, test.Translation, test.Language, res)
		}
	}
}
