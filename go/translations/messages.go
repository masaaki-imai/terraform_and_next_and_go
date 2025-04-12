package translations

import (
	"log"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// ValidationMessages はバリデーションメッセージの翻訳を定義します。
// Laravelの翻訳ファイルと同様の役割を果たします。
var ValidationMessages = map[string]string{
	// 基本的なバリデーションルール
	"required":             "{0}は必須項目です",
	"email":                "{0}は有効なメールアドレス形式で指定してください",
	"min":                  "{0}は{1}文字以上で入力してください",
	"max":                  "{0}は{1}文字以下で入力してください",
	"numeric":              "{0}には数字を指定してください",
	"alpha":                "{0}にはアルファベットのみ使用できます",
	"alpha_num":            "{0}には英数字のみ使用できます",
	"alpha_dash":           "{0}には英数字、ハイフン、アンダースコアのみ使用できます",
	"alphanum":             "{0}には英数字のみ使用できます",
	"alphanumunicode":      "{0}には英数字のみ使用できます",
	"boolean":              "{0}は真偽値を指定してください",
	"date":                 "{0}は正しい日付形式で指定してください",
	"datetime":             "{0}は正しい日時形式で指定してください",
	"excludes":             "{0}に{1}を含めることはできません",
	"excludesall":          "{0}に{1}を含めることはできません",
	"excludesrune":         "{0}に{1}を含めることはできません",
	"format":               "{0}は{1}形式で入力してください",
	"contains":             "{0}は{1}を含む必要があります",
	"containsany":          "{0}は{1}のいずれかを含む必要があります",
	"containsrune":         "{0}は{1}を含む必要があります",
	"dive":                 "{0}の各要素にエラーがあります",
	"endswith":             "{0}は{1}で終わる必要があります",
	"eqcsfield":            "{0}は{1}と等しくなければなりません",
	"eqfield":              "{0}は{1}と等しくなければなりません",
	"eq":                   "{0}は{1}と等しくなければなりません",
	"fieldcontains":        "{0}は{1}を含む必要があります",
	"fieldexcludes":        "{0}に{1}を含めることはできません",
	"file":                 "{0}はファイルである必要があります",
	"filepath":             "{0}は有効なファイルパスである必要があります",
	"gt":                   "{0}は{1}より大きくなければなりません",
	"gte":                  "{0}は{1}以上でなければなりません",
	"gtcsfield":            "{0}は{1}より大きくなければなりません",
	"gtefield":             "{0}は{1}以上でなければなりません",
	"gtfield":              "{0}は{1}より大きくなければなりません",
	"hostname":             "{0}は有効なホスト名である必要があります",
	"ip":                   "{0}は有効なIPアドレスである必要があります",
	"ip4_addr":             "{0}は有効なIPv4アドレスである必要があります",
	"ip6_addr":             "{0}は有効なIPv6アドレスである必要があります",
	"ipv4":                 "{0}は有効なIPv4アドレスである必要があります",
	"ipv6":                 "{0}は有効なIPv6アドレスである必要があります",
	"json":                 "{0}は有効なJSON文字列である必要があります",
	"lt":                   "{0}は{1}より小さくなければなりません",
	"lte":                  "{0}は{1}以下でなければなりません",
	"ltcsfield":            "{0}は{1}より小さくなければなりません",
	"ltefield":             "{0}は{1}以下でなければなりません",
	"ltfield":              "{0}は{1}より小さくなければなりません",
	"mac":                  "{0}は有効なMACアドレスである必要があります",
	"necsfield":            "{0}は{1}と異なる必要があります",
	"nefield":              "{0}は{1}と異なる必要があります",
	"ne":                   "{0}は{1}と異なる必要があります",
	"number":               "{0}は有効な数値である必要があります",
	"oneof":                "{0}は[{1}]のいずれかである必要があります",
	"required_if":          "{1}が{2}の場合、{0}は必須です",
	"required_unless":      "{1}が{2}でない場合、{0}は必須です",
	"required_with":        "{1}が指定されている場合、{0}は必須です",
	"required_with_all":    "{1}が全て指定されている場合、{0}は必須です",
	"required_without":     "{1}が指定されていない場合、{0}は必須です",
	"required_without_all": "{1}が全て指定されていない場合、{0}は必須です",
	"startswith":           "{0}は{1}で始まる必要があります",
	"unique":               "{0}は一意である必要があります",
	"url":                  "{0}は有効なURL形式で指定してください",
	"uuid":                 "{0}は有効なUUIDである必要があります",
	"uuid3":                "{0}は有効なUUID v3である必要があります",
	"uuid4":                "{0}は有効なUUID v4である必要があります",
	"uuid5":                "{0}は有効なUUID v5である必要があります",
}

// RegisterTranslations はバリデーションメッセージの翻訳を登録します
func RegisterTranslations(v *validator.Validate, trans ut.Translator) {
	// 各バリデーションルールの翻訳を登録
	for tag, translation := range ValidationMessages {
		registerTranslation(v, trans, tag, translation)
	}
}

// registerTranslation は1つのバリデーションルールの翻訳を登録します
func registerTranslation(v *validator.Validate, trans ut.Translator, tag string, message string) {
	err := v.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, message, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		params := []string{GetJaFieldName(fe.Field())}

		// パラメータがある場合は追加
		if fe.Param() != "" {
			params = append(params, fe.Param())
		}

		t, err := ut.T(fe.Tag(), params...)
		if err != nil {
			log.Printf("警告: %sタグの翻訳に失敗しました: %v\n", fe.Tag(), err)
			return fe.Error() // フォールバック: 元のエラーメッセージを返す
		}
		return t
	})

	if err != nil {
		log.Printf("警告: %sタグの翻訳登録に失敗しました: %v\n", tag, err)
	}
}
