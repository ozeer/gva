package designer

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ozeer/gva/api/v1"
)

type CommentRouter struct{}

func (t *CommentRouter) InitCommentRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	commentRouter := Router.Group("DesignerComment")
	commentApi := v1.ApiGroupApp.DesignerApiGroup.CommentApi
	{
		commentRouter.POST("Add", commentApi.Add)
		commentRouter.POST("Delete", commentApi.Delete)
		commentRouter.POST("Edit", commentApi.Edit)
		commentRouter.GET("All", commentApi.All)
		commentRouter.GET("Detail", commentApi.Detail)
		commentRouter.GET("Get", commentApi.Get)
	}
	return commentRouter
}
