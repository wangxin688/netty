package middleware

import "github.com/gin-gonic/gin"

type I18nLanuage string

const LocaleCtx string = "i18nLanguage"

const (
	I18nLanuageEn I18nLanuage = "en"
	I18nLanuageZh I18nLanuage = "zh"
)

func I18nMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("Accept-Language")
		if locale == "" {
			locale = string(I18nLanuageEn)
		}
		c.Header("Accept-Language", locale)
		c.Set(LocaleCtx, locale)
		c.Next()
	}
}
