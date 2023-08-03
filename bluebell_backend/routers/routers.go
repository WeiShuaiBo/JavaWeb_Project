package routers

import (
	"bluebell_backend/controller"
	"bluebell_backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r := gin.Default()
	//if err := endless.ListenAndServe(, r); err != nil {
	//	log.Fatalf("listen: %s\n", err)
	//}
	//srv := &http.Server{
	//	Addr:    ":8080",
	//	Handler: r,
	//}
	//
	//go func() {
	//	//开启一个goroutine 启动服务
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//}()

	r.Use(Cors())
	v1 := r.Group("/api/v1")
	// 生成验证码
	v1.GET("/captcha", func(c *gin.Context) {
		utils.Captcha(c, 4)
	})
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/signup", controller.SignUpHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)

	v1.Use(controller.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.PostDetailHandler)
		v1.GET("/post", controller.PostListHandler)

		v1.GET("/posts2", controller.PostList2Handler)

		v1.POST("/vote", controller.VoteHandler)

		v1.POST("/comment", controller.CommentHandler)
		v1.GET("/comment", controller.CommentListHandler)

		v1.POST("/createProject", controller.CreateProject)
		v1.POST("/createProject1", controller.CreateProject2)
		//用户的个人信息
		v1.GET("/getInf", controller.ListUserInformation)
		//用户修改个人信息
		v1.POST("/updateInf", controller.UpdateUserInformation)

	}

	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//zap.L().Info("shutdown server ...")
	////创建一个5s超市的context
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := srv.Shutdown(ctx); err != nil {
	//	zap.L().Fatal("server shutdownL: ")
	//}
	//log.Println("server exiting")
	//
	//r.NoRoute(func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "404",
	//	})
	//})
	return r
}

// 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		}
		if method == "OPTIONS" {
			/* 添加 */
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "OPTIONS")
			c.Header("Access-Control-Allow-Headers", "*")
			/* 添加 */
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
