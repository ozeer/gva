package designer

import (
	"github.com/gin-gonic/gin"
	"github.com/ozeer/gva/global"
	"github.com/ozeer/gva/model/common/response"
	"github.com/ozeer/gva/model/designer"
	"go.uber.org/zap"
)

type CourseApi struct{}

func (e *CourseApi) List(c *gin.Context) {

}

func (e *CourseApi) Detail(c *gin.Context) {

}

func (e *CourseApi) Create(c *gin.Context) {
	var course designer.DesignerCourse
	err := c.ShouldBindJSON(&course)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = courseService.CreateDesignerCourse(course)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (e *CourseApi) Delete(c *gin.Context) {

}

func (e *CourseApi) Edit(c *gin.Context) {

}
