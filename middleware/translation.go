package middleware

import (
	"bastion/utils"
	"github.com/gin-gonic/gin"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	zh_tw_translations "gopkg.in/go-playground/validator.v9/translations/zh_tw"
)

//设置Translation
func Translation() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := utils.Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(utils.Validate, trans)
			break
		case "en":
			_ = en_translations.RegisterDefaultTranslations(utils.Validate, trans)
			break
		case "zh_tw":
			_ = zh_tw_translations.RegisterDefaultTranslations(utils.Validate, trans)
			break
		default:
			_ = zh_translations.RegisterDefaultTranslations(utils.Validate, trans)
			break
		}

		//自定义错误内容
		//app.Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		//	return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
		//}, func(ut ut.Translator, fe validator.FieldError) string {
		//	t, _ := ut.T("required", fe.Field())
		//	return t
		//})

		//设置trans到context
		c.Set("trans", trans)
		c.Next()
	}
}
