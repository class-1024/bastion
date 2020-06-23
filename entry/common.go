package entry

import (
	"bastion/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

func CommValidator(ctx *gin.Context, i interface{}) error {
	if err := ctx.ShouldBind(i); err != nil {
		return fmt.Errorf("validator error: %w", err)
	}

	v := ctx.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = utils.Uni.GetTranslator("zh")
	}

	err := utils.Validate.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var sliceErrs []string
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}
