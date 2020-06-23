package router

import (
	"bastion/controller"
	"bastion/database"
	"bastion/middleware"
	"bastion/utils"
	"bastion/utils/pprof"
	"encoding/gob"
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"html/template"
)

func init() {
	// sessions encode
	gob.Register(database.MUser{})
	gob.Register(database.StatAdmin{})
}

func Init() *gin.Engine {
	r := gin.New()

	store := cookie.NewStore([]byte("secret"))

	// 全局中间件
	r.Use(
		middleware.Recovery(),
		middleware.RequestLog(),
		middleware.Translation(),
		sessions.Sessions("bastion", store))

	// 模版
	r.SetFuncMap(template.FuncMap{
		"fmtDate":   utils.FmtDate,
		"parseHtml": utils.ParseHtml,
	})
	r.LoadHTMLGlob("web/views/*")

	// 静态资源
	r.Static("/web/public", "web/public")
	r.StaticFile("favicon.ico", "web/public/favicon.ico")

	// swagger doc
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// prometheus 监控
	r.Use(ginprom.PromMiddleware(nil))
	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))

	// pprof
	pprof.Register(r)

	// system
	system := r.Group("/")
	SysRegister(system)

	// 统计
	stat := r.Group("/")
	StatRegister(stat)

	return r
}


func SysRegister(r *gin.RouterGroup) {
	s := controller.System{}
	r.GET("/sys/info", s.Info)
	r.GET("/sys/error", s.Error)
}

func StatRegister(r *gin.RouterGroup) {
	s := controller.Stat{}

	r.Use(cors.Default())

	r.POST("/api/stat/admin/login", s.AdminLogin)
	r.POST("/api/stat/admin/create", s.AdminCreate)
	r.POST("/api/stat/admin/update", middleware.StatAdminAuth(), s.AdminUpdate)
	r.GET("/api/stat/admin/info", middleware.StatAdminAuth(), s.AdminInfo)
	r.GET("/api/stat/admin/list", middleware.StatAdminAuth(), s.AdminList)

	r.POST("/api/stat/project", middleware.StatAdminAuth(), s.CreateProject)
	r.GET("/api/stat/projects", middleware.StatAdminAuth(), s.FindAllProjects)
	r.GET("/api/stat/project/:id", middleware.StatAdminAuth(), s.FindProjectById)

	r.POST("/api/stat/error", s.CreateError)
	r.GET("/img/stat/error", s.ImgCreateError)

	r.GET("/api/stat/device", s.FindDeviceByUid)
	r.GET("/api/stat/devices", s.FindAllDevice)
	r.GET("/api/stat/errors", s.FindErrorsWithParams)

	r.POST("/api/stat/behavior", s.CreateBehavior)
	r.GET("/api/stat/behaviors", middleware.StatAdminAuth(), s.FindAllBehaviors)
	r.GET("/api/stat/behavior/:id", middleware.StatAdminAuth(), s.FindBehaviorById)

	r.GET("/api/test/fail", s.TestFail)
	r.GET("/api/test/error", s.TestError)
	r.GET("/api/test/timeout", s.TestTimeOut)
}
