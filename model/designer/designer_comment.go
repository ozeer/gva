package designer

import "github.com/ozeer/gva/global"

type DesignerComment struct {
	global.GVA_MODEL
	Point          string `json:"point" gorm:"comment:总体评价打分"`                   // 评价打分
	Content        string `json:"content" gorm:"comment:评价内容"`                   // 评价内容
	Pics           string `json:"pics" gorm:"comment:评价图片"`                      // 装修图片
	HouseType      uint   `json:"house_type" gorm:"default:0;comment:房屋类型"`      // 房屋类型:0(默认)|毛胚1|精装2|老房3
	DecorationMode uint   `json:"decoration_mode" gorm:"default:0;comment:装修方式"` // 装修方式:0(默认)|全包1|半包2|清包3
	HouseArea      uint   `json:"house_area" gorm:"default:0;comment:房屋面积"`      // 房屋面积
	TotalCost      uint   `json:"total_cost" gorm:"default:0;comment:总花费"`       // 总花费
	City           string `json:"city" gorm:"comment:房屋所在城市"`                    // 房屋所在城市
	Contract       string `json:"contract" gorm:"comment:合同截图"`                  // 合同截图
}

func (DesignerComment) TableName() string {
	return "designer_comment"
}
