package locale

import (
	"encoding/json"
	"fmt"
	. "github.com/yautah/goroyale/assets"
)

type Locale struct {
	Identifier string `json:"identifier"`
	En         string `json:"en"`
	// Fr         string `json:"fr"`
	// De         string `json:"de"`
	// Es         string `json:"es"`
	// It         string `json:"it"`
	// Nl         string `json:"nl"`
	// No         string `json:"no"`
	// Tr         string `json:"tr"`
	// Jp         string `json:"jp"`
	// Kr         string `json:"kr"`
	// Ru         string `json:"ru"`
	// Ar         string `json:"ar"`
	// Pt         string `json:"pt"`
	Cn string `json:"cn"`
	// CNT        string `json:"CNT"`
	// Fa         string `json:"fa"`
	// ID         string `json:"id"`
	// Ms         string `json:"ms"`
	// Th         string `json:"th"`
	// Fi         string `json:"fi"`
}

var LocaleMap map[string]Locale = make(map[string]Locale)

func init() {
	var locales []Locale

	data, err := Asset("assets/texts.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &locales)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, locale := range locales {
		LocaleMap[locale.Identifier] = locale
	}
}

func GetTranslationById(textId string, locale string) string {
	switch locale {
	case "En":
		return LocaleMap[textId].En
	case "Cn":
		return LocaleMap[textId].Cn
	default:
		return LocaleMap[textId].Cn
	}
}

func GetLocaleById(textId string) string {
	return LocaleMap[textId].Cn
}

func GetTextById(textId string) Locale {
	return LocaleMap[textId]
}
