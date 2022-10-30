package route

import (
	"github.com/gin-gonic/gin"
	"1536509937/ku-bbs/internal/app/frontend"
)

func RegisterFrontedRoute(engine *gin.Engine) {
	group := engine.Group("/")

	group.GET("/register", frontend.Auth.RegisterPage)

	group.POST("/register", frontend.Auth.RegisterSubmit)

	group.GET("/login", frontend.Auth.LoginPage)

	group.POST("/login", frontend.Auth.LoginSubmit)

	group.Use(visitor)

	group.GET("/", frontend.Home.HomePage)

	group.GET("/logout", frontend.Auth.LogoutSubmit)

	group.GET("/publish", frontend.Topic.PublishPage)

	group.POST("/publish", frontend.Topic.PublishSubmit)

	group.GET("/topics/:id", frontend.Topic.DetailPage)

	group.POST("/topics/:id/delete", frontend.Topic.DeleteSubmit)

	group.GET("/topics/:id/edit", frontend.Topic.EditPage)

	group.POST("/topics/:id/edit", frontend.Topic.EditSubmit)

	group.POST("/topics/:id/comment-state", frontend.Topic.SettingCommentStateSubmit)

	group.POST("/comments", frontend.Comment.PublishSubmit)

	group.POST("comments/delete", frontend.Comment.DeleteSubmit)

	group.GET("/user", frontend.User.HomePage)

	group.GET("/user/edit", frontend.User.EditPage)

	group.POST("/user/edit", frontend.User.EditSubmit)

	group.POST("/md-upload", frontend.File.MDUploadSubmit)

	group.GET("/notice", frontend.Notice.HomePage)

	group.POST("/likes", frontend.Like.LikeSubmit)

	group.POST("/follows", frontend.Follow.FollowSubmit)

	group.POST("/checkins", frontend.Checkin.StoreSubmit)

	group.POST("/reports", frontend.Report.ReportSubmit)
}
