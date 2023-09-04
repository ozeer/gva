package designer

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ozeer/gva/api/v1"
)

type CourseRouter struct{}

func (e *CourseRouter) InitCourseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	courseRouter := Router.Group("DesignerCourse")
	courseApi := v1.ApiGroupApp.DesignerApiGroup.CourseApi
	{
		courseRouter.POST("Create", courseApi.Create)
		courseRouter.POST("Delete", courseApi.Delete)
		courseRouter.POST("Edit", courseApi.Edit)
		courseRouter.GET("All", courseApi.All)
		courseRouter.GET("Detail", courseApi.Detail)
	}
	return courseRouter
}
