package designer

import "github.com/ozeer/gva/service"

type ApiGroup struct {
	CommentApi
	CourseApi
}

var (
	commentService = service.ServiceGroupApp.DesignerServiceGroup.CommentService
	courseService  = service.ServiceGroupApp.DesignerServiceGroup.CourseService
)
