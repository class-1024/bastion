package controller

import (
	"bastion/dao"
	"bastion/database"
	"bastion/entry"
	"bastion/pkg/constant"
	"bastion/pkg/errno"
	"bastion/pkg/response"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Stat struct {
}

// @Summary 管理员登录
// @Produce json
// @Router /api/stat/admin/login [post]
func (s *Stat) AdminLogin(c *gin.Context) {
	session := sessions.Default(c)

	p := entry.AdminAccount{}
	err := p.BindingValidParams(c)
	if err != nil {
		response.Fail(c, errno.ErrorNotFound, err)
		return
	}

	// 查找
	user, err := dao.FindAdminUserByUserName(p.Username)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.Fail(c, errno.ErrorUserNotFound, err)
			return
		}
		response.Fail(c, errno.ErrorNotFound, err)
		return
	}

	// 校验
	if p.Password != user.Password {
		response.Fail(c, errno.ErrorUsePassword, errors.New("密码错误"))
		return
	}

	session.Set(constant.SessionKeyStatAdmin, user)
	err = session.Save()
	if err != nil {
		response.Fail(c, errno.ErrorSession, err)
		return
	}

	response.Success(c, user, "登录成功")
	return
}

// @Summary 创建管理员
// @Produce json
// @Router /api/stat/admin/create [post]
func (s *Stat) AdminCreate(c *gin.Context) {
	p := entry.AdminAccountRegister{}
	err := p.BindingValidParams(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	err = dao.CreateAdminUser(database.StatAdmin{
		Username: p.Username,
		Password: p.Password,
		Email:    p.Email,
	})
	if err != nil {
		response.Fail(c, errno.ErrorCreateData, err)
		return
	}

	response.Success(c, nil, "创建成功")
}

// @Summary 管理员信息
// @Produce json
// @Router /api/stat/admin/info [get]
func (s *Stat) AdminInfo(c *gin.Context) {
	session := sessions.Default(c)
	get := session.Get(constant.SessionKeyStatAdmin)
	response.Success(c, get, )
	return

}

// @Summary 所有管理员
// @Produce json
// @Router /api/stat/admin/list [get]
func (s *Stat) AdminList(c *gin.Context) {
	p, err := entry.CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	rows, total, e := dao.FindAllAdminUsers(p.PageSize, p.Page, p.Order)
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

// @Summary 管理员信息更新
// @Produce json
// @Router /api/stat/admin/update [post]
func (s *Stat) AdminUpdate(c *gin.Context) {
}

// @Summary 创建项目
// @Produce json
// @Router /api/stat/project [post]
func (s *Stat) CreateProject(c *gin.Context) {
}

// @Summary 查找所有项目
// @Produce json
// @Router /api/stat/projects [get]
func (s *Stat) FindAllProjects(c *gin.Context) {
	p, err := entry.CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	rows, total, e := dao.FindAllProjects(p.PageSize, p.Page, p.Order)
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

// @Summary 查找项目
// @Produce json
// @Router /api/stat/project/:id [get]
func (s *Stat) FindProjectById(c *gin.Context) {
}

// @Summary 上报错误
// @Produce json
// @Router /api/stat/errors [post]
func (s *Stat) CreateError(c *gin.Context) {
	p := entry.ErrorData{}
	_ = p.BindingValidParams(c)

	data := database.StatError{
		Appid:       p.Appid,
		Path:        p.Path,
		ErrorMsg:    p.Msg,
		ErrorString: p.Desc,
		UId:         p.UId,
	}

	e := dao.CreateErrors(data)
	if e != nil {
		c.String(http.StatusOK, e.Error())
		return
	}
	c.String(http.StatusOK, "ok")
}

// @Summary errlock上报
// @Produce json
// @Router /api/stat/errors [post]
func (s *Stat) ImgCreateError(c *gin.Context) {
	p := entry.ErrorData{}
	_ = p.BindingValidParams(c)

	// 设备信息
	if p.BrowserUa != "" {
		data := database.StatDevice{
			UId:           p.UId,
			BrowserUa:     p.BrowserUa,
			BrowserResult: p.BrowserResult,
		}
		if e := dao.CreateIfNotExistDevice(data); e != nil {
			c.String(http.StatusOK, e.Error())
			return
		}
	}

	// 创建错误
	if p.Msg != "" {
		data := database.StatError{
			Appid:       p.Appid,
			Tag:         p.Tag,
			Path:        p.Path,
			ErrorMsg:    p.Msg,
			ErrorString: p.Desc,
			FileURL:     p.FileUrl,
			Lineno:      p.Lineno,
			Colno:       p.Colno,
			UId:         p.UId,
		}
		if e := dao.CreateErrors(data); e != nil {
			c.String(http.StatusOK, e.Error())
			return
		}
	}
	c.String(http.StatusOK, "ok")
}

// @Summary 查找所有错误
// @Produce json
// @Router /api/stat/errors [get]
func (s *Stat) FindErrorsWithParams(c *gin.Context) {
	p, err := entry.CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	query := entry.ErrorsQuery{}
	err = query.BindingValidParams(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	findErrors, total, e := dao.FindErrors(query, p.PageSize, p.Page, p.Order)
	if e != nil {
		response.Fail(c, errno.ErrorQueryData, e)
		return
	}

	response.Success(c, response.PageData{
		Page:     p.Page,
		Pagesize: p.PageSize,
		Total:    total,
		Rows:     findErrors,
	})
	return
}

// @Summary 根据id查找错误
// @Produce json
// @Router /api/stat/error/:id [get]
func (s *Stat) FindErrorById(c *gin.Context) {
}

// @Summary 根据id查找错误
// @Produce json
// @Router /api/stat/device/:uid [get]
func (s *Stat) FindDeviceByUid(c *gin.Context) {
	val := c.Query("uid")
	if val != "" {
		fmt.Println("", val)
	}

	res, err := dao.FindByUid(val)
	if err != nil {
		response.Fail(c, errno.ErrorNotFound, err)
		return
	}
	response.Success(c, res)
}

func (s *Stat) FindAllDevice(c *gin.Context) {
	p, err := entry.CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	device, total, e := dao.FindAllDevice(p.PageSize, p.Page, p.Order)
	if e != nil {
		response.Fail(c, errno.ErrorQueryData, e)
		return
	}

	response.Success(c, response.PageData{
		Page:     p.Page,
		Pagesize: p.PageSize,
		Total:    total,
		Rows:     device,
	})
}

// @Summary 上报打点
// @Produce json
// @Router /api/stat/behavior [post]
func (s *Stat) CreateBehavior(c *gin.Context) {
}

// @Summary 查找所有打点
// @Produce json
// @Router /api/stat/behaviors [get]
func (s *Stat) FindAllBehaviors(c *gin.Context) {
}

// @Summary 根据id查找打点
// @Produce json
// @Router /api/stat/behavior/:id [get]
func (s *Stat) FindBehaviorById(c *gin.Context) {
}

func (s *Stat) TestFail(c *gin.Context) {
	response.Fail(c, errno.ErrorQueryData, errors.New("query failed"))
}

func (s *Stat) TestError(c *gin.Context) {
	response.Error(c, http.StatusInternalServerError, errors.New("server error"))
}

func (s *Stat) TestTimeOut(c *gin.Context) {
	time.Sleep(time.Second * 11)
	response.Success(c, nil, "ok")
}
