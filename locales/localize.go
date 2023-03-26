package locales

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var i18nBundle *i18n.Bundle

func BootLocalize() *i18n.Bundle {
	languages := []string{"en", "th"}
	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	for _, lang := range languages {
		path, _ := filepath.Abs(fmt.Sprintf("locales/%s.json", lang))
		bundle.MustLoadMessageFile(path)
	}

	return bundle
}

func GetLocalize(lang string, message *string) {
	if i18nBundle == nil {
		i18nBundle = BootLocalize()
	}

	localizer := i18n.NewLocalizer(i18nBundle, lang)
	value, err := localizer.LocalizeMessage(&i18n.Message{
		ID: *message,
	})

	if err == nil && value != "" {
		*message = value
	}
}
