package router

import (
	"net/http"

	"../router/handlers/api"
	"./handlers"
	"./handlers/vue"
	"./middleware"
	//go-mssqldb
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Init blablaba
func Init() {
	// Creates a default gin router
	router := gin.Default() // Grouping routes

	//自製Middleware
	//router.Use(middleware.PrintMiddleware)
	//router.Use(middleware.Print())

	//根據website的路由規則
	// router.LoadHTMLGlob("templates/*")
	// frontend
	router.Use(static.Serve("/", static.LocalFile("../", true)))
	// router.Use(static.Serve("/", static.LocalFile("../favicon.ico", true)))
	router.LoadHTMLGlob("../index.html")

	// group： v1
	v1 := router.Group("/v1")
	{
		v1.GET("/hello", handlers.Hello)

		v1.GET("/get", handlers.Get)

		//URL中的name參數
		v1.GET("/hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})

		//常規參數
		v1.GET("/welcome", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "Guest")
			lastname := c.Query("lastname")
			c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
		})
	}

	//group： v2
	v2 := router.Group("/v2")
	v2.Use(middleware.ValidateToken())
	{
		v2.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{"title": "hello Gin."})
		})

	}

	//vue
	// router.GET("/vue", handlersVue.Student)
	//group： vue
	vue := router.Group("/vue")
	{
		vue.GET("", handlersVue.Student)
		vue.GET("/mssql", handlersVue.GetData)

	}
	//group：api
	api := router.Group("/api")
	{
		api.GET("/student", handlersApi.Student)
		api.GET("/student2", handlersApi.Student2)
		api.POST("/insert", handlersApi.Insert)
		api.POST("/insert2", handlersApi.Insert2)
	}

	// 404 NotFound
	router.NoRoute(func(c *gin.Context) {
		c.HTML(200, "404.html", gin.H{})
	})

	router.Run(":8000") // listen and serve on 0.0.0.0:8000
}
