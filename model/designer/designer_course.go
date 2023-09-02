package designer

import "github.com/ozeer/gva/global"

type DesignerCourse struct {
	global.GVA_MODEL
	Name         string `json:"name" gorm:"comment:课程名称"`         // 课程名称
	Title        string `json:"title" gorm:"comment:课程标题"`        // 课程标题
	SubTitle     string `json:"sub_title" gorm:"comment:课程副标题"`   // 课程副标题
	Introduction string `json:"introduction" gorm:"comment:课程介绍"` // 课程介绍
	CoverImg     string `json:"cover_img" gorm:"comment:封面图"`     // 课程封面图
	Video        string `json:"video" gorm:"comment:视频"`          // 课程视频
}

func (DesignerCourse) TableName() string {
	return "designer_course"
}
