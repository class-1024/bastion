package controller

import (
	"bastion/dao"
	"bastion/database"
	"bastion/entry"
	"bastion/pkg/constant"
	"bastion/pkg/errno"
	"bastion/pkg/response"
	"bastion/service"
	"bastion/service/movie"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/teris-io/shortid"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type Api struct {
}

func (a *Api) Noop(c *gin.Context) {
	response.Fail(c, errno.InternalServerError, errors.New("开发中"))
	return
}

// @Summary 上传文件
// @Produce  json
// @Param file formData file true "上传文件"
// @Success 200 {object} response.Response
// @Router /api/upload [post]
func (a *Api) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	id, _ := shortid.Generate()
	ext := filepath.Ext(file.Filename)
	dst := "web/public/uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + id + ext

	e := c.SaveUploadedFile(file, dst)
	if e != nil {
		response.Error(c, http.StatusInternalServerError, e)
		return
	}

	// 成功
	response.Success(c, dst)
	return
}

// @Summary code登录
// @Produce  json
// @Param code query string true "wx.login 获取的 code"
// @Success 200 {object} response.Response
// @Router /api/weapp/login [post]
func (a *Api) CodeLogin(c *gin.Context) {
	session := sessions.Default(c)

	p := entry.CodeLogin{}
	err := p.BindingValidParams(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	miniApp := service.WeApp{
		AppName:   viper.GetString("weapp.name"),
		AppId:     viper.GetString("weapp.appId"),
		AppSecret: viper.GetString("weapp.appSecret"),
	}
	// 通过code获取 openid session_key
	res, err := miniApp.GetSessionKey(p.Code)
	if err != nil {
		response.Fail(c, errno.InternalServerError, err)
		return
	}

	// 添加或者查找
	user, err := service.FindOrCreateUserByOpenid(res.Openid)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	// 登录
	session.Set(constant.SessionKeyWeApp, user)
	err = session.Save()
	if err != nil {
		response.Fail(c, errno.ErrorSession, err)
		return
	}

	// 成功
	response.Success(c, user)
	return
}

// @Summary 用户信息
// @Produce  json
// @Success 200 {object} response.Response
// @Router /api/user/info [get]
func (a *Api) GetUserInfo(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(constant.SessionKeyWeApp)
	u := user.(database.MUser)

	userInfo, err := dao.FindUserByOpenid(u.Openid)
	if gorm.IsRecordNotFoundError(err) {
		response.Fail(c, errno.ErrorNotFound, err)
		return
	}
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	// 更新
	session.Set(constant.SessionKeyWeApp, userInfo)
	err = session.Save()
	if err != nil {
		response.Fail(c, errno.ErrorSession, err)
		return
	}

	response.Success(c, userInfo)
	return
}

// @Summary 小程序用户信息解密
// @Produce  json
// @Success 200 {object} response.Response
// @Router /api/weapp/decryptUserInfo [post]
func (a *Api) DecryptUserInfo(c *gin.Context) {
	p := entry.EncryptedUserInfo{}
	err := p.BindingValidParams(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	session := sessions.Default(c)
	user := session.Get(constant.SessionKeyWeApp)
	u := user.(database.MUser)

	miniApp := service.WeApp{
		AppName:   viper.GetString("weapp.name"),
		AppId:     viper.GetString("weapp.appId"),
		AppSecret: viper.GetString("weapp.appSecret"),
	}

	// 解密
	info, e := miniApp.DecryptUserInfo(u.Openid, p.RawData, p.EncryptedData, p.Signature, p.Iv)
	if e != nil {
		response.Fail(c, errno.ErrorDecryptUserData, e)
		return
	}

	// 更新
	updateErr := dao.UpdateUser(info)
	if updateErr != nil {
		response.Fail(c, errno.ErrorUpdateData, updateErr)
		return
	}

	response.Success(c, u, "解密成功")
	return
}

// @Summary 用户列表
// @Produce json
// @Param pagesize query integer false "一页数量"
// @Param page query integer false "页码"
// @Param order query string false "排序"
// @Success 200 {object} response.Response
// @Router /api/users [get]
func (a *Api) GetUsers(c *gin.Context) {
	p, err := CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	rows, total, e := dao.FindAllUsers(p.PageSize, p.Page, p.Order)

	if e != nil {
		response.Fail(c, errno.ErrorQueryData, e)
		return
	}

	response.Success(c, response.PageData{
		Page:     p.Page,
		Pagesize: p.PageSize,
		Total:    total,
		Rows:     rows,
	})
	return
}

// @Summary 评论列表
// @Produce  json
// @Param pagesize query integer false "一页数量"
// @Param page query integer false "页码"
// @Param order query string false "排序"
// @Success 200 {object} response.Response
// @Router /api/comments [get]
func (a *Api) GetComments(c *gin.Context) {
	p, err := CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	rows, total, err := dao.FindAllComments(p.PageSize, p.Page, p.Order)

	if err != nil {
		response.Fail(c, errno.ErrorQueryData, err)
		return
	}

	response.Success(c, response.PageData{
		Page:     p.Page,
		Pagesize: p.PageSize,
		Total:    total,
		Rows:     rows,
	})
	return
}

// @Summary 电影列表
// @Produce  json
// @Param pagesize query integer false "一页数量"
// @Param page query integer false "页码"
// @Param order query string false "排序"
// @Success 200 {object} response.Response
// @Router /api/movies [get]
func (a *Api) GetMovies(c *gin.Context) {
	p, err := CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	rows, total, e := dao.FindAllMovies(p.PageSize, p.Page, p.Order)

	if e != nil {
		response.Fail(c, errno.ErrorQueryData, e)
		return
	}

	response.Success(c, response.PageData{
		Page:     p.Page,
		Pagesize: p.PageSize,
		Total:    total,
		Rows:     rows,
	})
	return
}

// @Summary 添加电影
// @Produce  json
// @Success 200 {object} response.Response
// @Router /api/movie/new [post]
func (a *Api) CreateMovie(c *gin.Context) {
	response.Fail(c, errno.InternalServerError, errors.New("开发中"))
	return
}

// @Summary 电影详情
// @Produce  json
// @Param id query int true "电影id"
// @Success 200 {object} response.Response
// @Router /api/movie/detail [get]
func (a *Api) GetMovieDetail(c *gin.Context) {
	p := entry.MovieDetailParams{}
	err := p.BindingValidParams(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	desc, e := dao.FindMovieById(p.Id)
	if gorm.IsRecordNotFoundError(e) {
		response.Fail(c, errno.ErrorNotFound, e)
		return
	}
	if e != nil {
		response.Fail(c, errno.ErrorQueryData, e)
		return
	}

	response.Success(c, desc)
	return
}

// @Summary 观看记录列表
// @Produce  json
// @Param pagesize query integer false "一页数量"
// @Param page query integer false "页码"
// @Param order query string false "排序"
// @Success 200 {object} response.Response
// @Router /api/movie/logs [get]
func (a *Api) GetWatchLogs(c *gin.Context) {
	p, err := CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	rows, total, err := dao.FindAllMovieLogs(p.PageSize, p.Page, p.Order)
	if err != nil {
		response.Fail(c, errno.ErrorQueryData, err)
		return
	}

	response.Success(c, response.PageData{
		Page:     p.Page,
		Pagesize: p.PageSize,
		Total:    total,
		Rows:     rows,
	})
	return
}

// @Summary 添加观看记录
// @Produce  json
// @Param movie_id body integer true "电影id"
// @Param progress body string true "进度 00:02:32"
// @Success 200 {object} response.Response
// @Router /api/movie/log [post]
func (a *Api) CreateWatchLog(c *gin.Context) {
	p := entry.CreateWatchLogParams{}
	err := p.BindingValidParams(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	session := sessions.Default(c)
	user := session.Get(constant.SessionKeyWeApp)
	u := user.(database.MUser)

	fmt.Printf("%v \n", user)
	fmt.Printf("%v \n", u)

	// 兼容客户端传参
	isExist, log, err := service.LogIsExist(int(u.ID), p.MovieId)
	if err != nil {
		response.Fail(c, errno.InternalServerError, err)
		return
	}

	// 更新记录
	if isExist {
		dao.UpdateMovieLog(int(log.ID), p.Progress)
		response.Success(c, nil, "更新")
		return
	}

	// 添加记录
	err = dao.CrateMovieLog(int(u.ID), p.MovieId, p.Progress)
	if err != nil {
		response.Fail(c, errno.InternalServerError, err)
		return
	}
	response.Success(c, nil, "添加")
	return
}

// @Summary 添加评论
// @Produce  json
// @Param comment body string true "评论"
// @Success 200 {object} response.Response
// @Router /api/comment/new [post]
func (a *Api) CreateComments(c *gin.Context) {
	p := entry.CreateCommentParams{}
	err := p.BindingValidParams(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	session := sessions.Default(c)
	user := session.Get(constant.SessionKeyWeApp)
	u := user.(database.MUser)

	err = dao.CrateComment(int(u.ID), p.Comment)
	if err != nil {
		response.Fail(c, errno.InternalServerError, err)
		return
	}

	response.Success(c, nil, "已提交")
	return
}

// @Summary 票房数据
// @Produce  json
// @Success 200 {object} response.Response
// @Router /api/piaofang [get]
func (a *Api) BoxOffice(c *gin.Context) {
	res, err := movie.GetBoxOffice()
	if err != nil {
		response.Fail(c, errno.InternalServerError, err)
		return
	}
	response.Success(c, res.Data)
	return
}

func CheckPagination(c *gin.Context) (*entry.Pagination, error) {
	p := entry.Pagination{
		PageSize: 50,
		Page:     1,
	}
	err := p.BindingValidParams(c)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
