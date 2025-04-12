package validation

import (
	"net/http"
	"errors"
	"L-cart/translations"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(c *gin.Context, err error) bool {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		// バリデーションエラーを翻訳
		translatedErrors := []string{}
		for _, e := range validationErrors {
			translatedErrors = append(translatedErrors, e.Translate(translations.GetTranslator()))
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"isValid": false,
			"result":  translatedErrors,
		})
		return true
	}
	return false
}