package translations

import (
	"log"
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
)

var (
	Uni        *ut.UniversalTranslator
	Translator ut.Translator
)

func InitValidator() {
	// 日本語ロケールの設定
	japanese := ja.New()
	Uni = ut.New(japanese, japanese)

	// 日本語トランスレーターの取得
	var found bool
	Translator, found = Uni.GetTranslator("ja")
	if !found {
		log.Println("日本語トランスレーターが見つかりませんでした")
	}

	// Ginのバインディングバリデーターを取得
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// フィールド名のカスタマイズ（JSONタグを使用）
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("json")
			if name == "" {
				return fld.Name
			}
			return name
		})

		// 日本語翻訳の登録
		if err := ja_translations.RegisterDefaultTranslations(v, Translator); err != nil {
			log.Println("日本語翻訳の登録に失敗しました:", err)
		}

		// 翻訳関数の登録
		RegisterTranslations(v, Translator)
	}
}

// GetTranslator は現在のトランスレーターを返します
func GetTranslator() ut.Translator {
	return Translator
}
