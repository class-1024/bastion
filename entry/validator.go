package entry

import "github.com/gin-gonic/gin"

type valid interface {
	BindingValidParams(ctx *gin.Context) error
}

type Pagination struct {
	PageSize int    `form:"pagesize" json:"pagesize" validate:""`
	Page     int    `form:"page" json:"page" validate:""`
	Order    string `form:"order" json:"order" validate:""`
}

func (pagination *Pagination) BindingValidParams(ctx *gin.Context) error {
	err := CommValidator(ctx, pagination)
	if err != nil {
		return err
	}
	return nil
}

type CodeLogin struct {
	Code string `form:"code" json:"code" validate:"required"`
}

func (p *CodeLogin) BindingValidParams(ctx *gin.Context) error {
	err := CommValidator(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

type MovieDetailParams struct {
	Id int64 `form:"id" json:"id" validate:"required"`
}

func (p *MovieDetailParams) BindingValidParams(ctx *gin.Context) error {
	err := CommValidator(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

type CreateWatchLogParams struct {
	MovieId  int    `form:"movie_id" json:"movie_id" validate:"required"`
	Progress string `form:"progress" json:"progress" validate:"required"`
}

func (p *CreateWatchLogParams) BindingValidParams(ctx *gin.Context) error {
	err := CommValidator(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

type CreateCommentParams struct {
	Comment string `form:"comment" json:"comment" validate:"required"`
}

func (p *CreateCommentParams) BindingValidParams(ctx *gin.Context) error {
	err := CommValidator(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

type EncryptedUserInfo struct {
	RawData       string `form:"rawData" json:"rawData" validate:"required"`
	Signature     string `form:"signature" json:"signature" validate:"required"`
	EncryptedData string `form:"encryptedData" json:"encryptedData" validate:"required"`
	Iv            string `form:"iv" json:"iv" validate:"required"`
}

func (info *EncryptedUserInfo) BindingValidParams(ctx *gin.Context) error {
	err := CommValidator(ctx, info)
	if err != nil {
		return err
	}
	return nil
}

type AdminAccount struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

func (p *AdminAccount) BindingValidParams(ctx *gin.Context) error {
	err := CommValidator(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

type AdminAccountRegister struct {
	Username string `form:"username" json:"username" validate:"required"`
	Email    string `form:"email" json:"email" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

func (p *AdminAccountRegister) BindingValidParams(ctx *gin.Context) error {
	err := CommValidator(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

type ErrorData struct {
	Tag     string `form:"tag" json:"tag" validate:""`
	FileUrl string `form:"fileUrl" json:"fileUrl" validate:""`
	Lineno  string `form:"lineno" json:"lineno" validate:""`
	Colno   string `form:"colno" json:"colno" validate:""`
	Msg     string `form:"msg" json:"msg" validate:""`
	Desc    string `form:"desc" json:"desc" validate:""`
	Appid   string `form:"appid" json:"appid" validate:""`
	Path    string `form:"path" json:"path" validate:""`
	UId     string `form:"uid" json:"uid" validate:""`

	// 设备
	BrowserUa     string `form:"b_ua" json:"b_ua" validate:""`
	BrowserResult string `form:"b_result" json:"b_result" validate:""`
}

func (p *ErrorData) BindingValidParams(ctx *gin.Context) error {
	err := CommValidator(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

type ErrorsQuery struct {
	AppId     string `form:"appid" json:"appid" validate:""`
	Tag       string `form:"tag" json:"tag" validate:""`
	Uid       string `form:"uid" json:"uid" validate:""`
	ErrorMsg  string `form:"error_msg" json:"error_msg" validate:""`
	CreatedAt string `form:"created_at" json:"created_at" validate:""`
}

func (p *ErrorsQuery) BindingValidParams(ctx *gin.Context) error {
	err := CommValidator(ctx, p)
	if err != nil {
		return err
	}
	return nil
}
