package designer

import (
	"github.com/ozeer/gva/global"
	"github.com/ozeer/gva/model/designer"
)

type CourseService struct{}

func (c *CourseService) CreateDesignerCourse(e designer.DesignerCourse) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}
