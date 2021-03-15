package utils

import (

	"fmt"
	"gosample/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

func GetGinRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	// router := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	// router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/**/*")
	router.Static("htmlSupport", "htmlSupport")
	router.Use(ginsession.New())
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Home",
		})
	})

	router.GET("/assignments", model.GetAllAssignemt)

	router.POST("/assignment", model.GetQuestionForAssignment)

	router.POST("/getClasses", model.GetClassesForSchool)
	router.POST("/getSubjects", model.GetSubjectList)
	router.POST("/getChapters", model.GetChapterList)
	router.POST("/getSchool", model.GetSchool)
	router.POST("/getUsers", model.GetUserList)
	router.POST("/registerUser", model.RegisterUser)
	router.POST("/validate", model.LoginTeacher)
	router.POST("/loginUser", model.LoginUser)

	return router
}

func  GetGinGroupRouter(route *gin.Engine) *gin.RouterGroup {
	router :=route.Group("/v1")
	{
		router.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title": "Home",
			})
		})

		router.GET("/assignments", model.GetAllAssignemt)

		router.POST("/assignment", model.GetQuestionForAssignment)

		router.POST("/getClasses", model.GetClassesForSchool)
		router.POST("/getSubjects", model.GetSubjectList)
		router.POST("/getChapters", model.GetChapterList)
		router.POST("/getSchool", model.GetSchool)
		router.POST("/getUsers", model.GetUserList)
		router.POST("/registerUser", model.RegisterUser)
		router.POST("/validate", model.LoginTeacher)
		router.POST("/loginUser", model.LoginUser)
	}

return router
}
func CheckIsUserLoggedIN(context *gin.Context) bool {
	store := ginsession.FromContext(context)
	userID, ok := store.Get("userId")
	fmt.Println(" CheckIsUserLoggedIN ", userID)
	if !ok {
		context.AbortWithStatus(404)
		return false
	}
	return true
}

func SaveUser(context *gin.Context, id string) {
	store := ginsession.FromContext(context)
	store.Set("userId", id)
	err := store.Save()
	if err != nil {
		context.AbortWithError(500, err)
		return
	}
}
