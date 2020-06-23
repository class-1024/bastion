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

	// 页面
	page := r.Group("/")
	PageRegister(page)

	// system
	system := r.Group("/")
	SysRegister(system)

	// api
	api := r.Group("/")
	ApiRegister(api)

	// 统计
	stat := r.Group("/")
	StatRegister(stat)

	return r
}

func ApiRegister(r *gin.RouterGroup) {
	api := controller.Api{}

	r.GET("/api/piaofang", api.BoxOffice)

	r.GET("/api/users", api.GetUsers)
	r.GET("/api/user/info", middleware.MiAppAuth(), api.GetUserInfo)
	r.POST("/api/weapp/login", api.CodeLogin)
	r.POST("/api/weapp/decryptUserInfo", api.DecryptUserInfo)

	r.GET("/api/movies", api.GetMovies)
	r.GET("/api/movie/detail", api.GetMovieDetail)
	r.POST("/api/movie/new", api.CreateMovie)

	r.GET("/api/movie/logs", api.GetWatchLogs)
	r.POST("/api/movie/log", middleware.MiAppAuth(), api.CreateWatchLog)

	r.GET("/api/comments", api.GetComments)
	r.POST("/api/comment/new", middleware.MiAppAuth(), api.CreateComments)

	r.POST("/api/upload", api.Upload)
	r.POST("/api/error", api.Noop)
}

func PageRegister(r *gin.RouterGroup) {
	p := controller.Page{}
	r.GET("/", p.Docs)
	r.GET("/blog/:id", p.DocContent)
	r.GET("/movie/stat", p.Stat)
	r.GET("/movie", p.MovieForm)
	r.GET("/error", p.Errors)
}

func SysRegister(r *gin.RouterGroup) {
	s := controller.System{}
	r.GET("/sys/info", s.Info)
	r.GET("/sys/error", s.Error)
}

func StatRegister(r *gin.RouterGroup) {
	s := controller.Stat{}

	config := cors.DefaultConfig()
	r.Use(cors.New(config))

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
