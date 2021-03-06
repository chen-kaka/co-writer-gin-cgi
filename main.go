package main

import (
	"fmt"
	"net/http"
	"runtime"

	"./controllers"
	"./db"
	"./services/middlewares"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	// use redis as session restore
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("co-writer-session", store))

	//init mongodb connection
	db.Connect()

	r.Use(CORSMiddleware())
	//use mongodb as database store
	r.Use(middlewares.Connect)
	//use postgresdb as database store
	//db.Init()

	v1 := r.Group("/app")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

		v1.POST("/user/signin", user.Signin)
		v1.POST("/user/signup", user.Signup)
		v1.GET("/user/signout", user.Signout)

		/*** START Article ***/
		article := new(controllers.ArticleController)

		v1.POST("/article", article.Create)
		v1.GET("/articles", article.All)
		v1.GET("/article/:id", article.One)
		v1.PUT("/article/:id", article.Update)
		v1.DELETE("/article/:id", article.Delete)

		repository := new(controllers.RepositoryController)
		v1.GET("/repository/list", repository.QueryMyRepos)
		v1.POST("/repository/create", repository.CreateRepo)

		userInfo := new(controllers.UserInfoController)
		v1.GET("/user_info/info", userInfo.QueryInfo)
		v1.POST("/user_info/create", userInfo.CreateUser)

		comment := new(controllers.CommentController)
		v1.GET("/comment/query_repo_comment",comment.QueryRepoComment)
		v1.POST("/comment/create_repo_comment",comment.CreateRepoComment)
	}

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.01",
			"goVersion":             runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	r.Run(":9000")
}
