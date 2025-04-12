package translations

import (
	"strings"
)

// AttributeJaNames は、フィールド名と日本語名のマッピングを定義します。
// Laravelの'attributes'配列と同様の役割を果たします。
var AttributeJaNames = map[string]string{
	// ユーザー関連
	"email":    "メールアドレス",
	"password": "パスワード",
	"name":     "名前",
	"username": "ユーザー名",

	// 住所関連
	"address":  "住所",
	"address1": "市区町村",
	"address2": "町名・番地",
	"address3": "ビル・マンション名",
	"zipcode":  "郵便番号",
	"city":     "市区町村",
	"state":    "都道府県",
	"country":  "国",

	// 連絡先関連
	"phone":  "電話番号",
	"mobile": "携帯電話番号",
	"fax":    "FAX番号",

	// 個人情報関連
	"birthday": "生年月日",
	"age":      "年齢",
	"gender":   "性別",

	// 支払い関連
	"card_number":    "カード番号",
	"card_name":      "カード名義",
	"card_expiry":    "有効期限",
	"card_cvv":       "セキュリティコード",
	"payment_method": "支払い方法",

	// 商品関連
	"product_name": "商品名",
	"product_code": "商品コード",
	"price":        "価格",
	"quantity":     "数量",
	"description":  "説明",

	// 日付関連
	"date":       "日付",
	"start_date": "開始日",
	"end_date":   "終了日",
	"created_at": "作成日時",
	"updated_at": "更新日時",

	// その他
	"title":    "タイトル",
	"content":  "内容",
	"message":  "メッセージ",
	"comment":  "コメント",
	"status":   "ステータス",
	"type":     "種類",
	"category": "カテゴリー",
	"tag":      "タグ",
	"url":      "URL",
	"file":     "ファイル",
	"image":    "画像",
	"document": "書類",
}

// GetJaFieldName はフィールド名を日本語に変換します
func GetJaFieldName(field string) string {
	// 小文字に変換（JSONタグは通常小文字のため）
	fieldLower := field
	if len(fieldLower) > 0 {
		fieldLower = strings.ToLower(fieldLower)
	}

	// マッピングから日本語名を取得
	if jaName, exists := AttributeJaNames[fieldLower]; exists {
		return jaName
	}

	// マッピングにない場合は元のフィールド名を返す
	return field
}
