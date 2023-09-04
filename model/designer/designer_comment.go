package designer

import (
	"github.com/ozeer/gva/global"
)

type DesignerComment struct {
	global.GVA_MODEL
	Point          string `json:"point" gorm:"comment:总体评价打分"`                                   // 评价打分
	Content        string `json:"content" gorm:"type:text;comment:评价内容"`                         // 评价内容
	Pics           string `json:"pics" gorm:"type:varchar(300);comment:评价图片"`                    // 装修图片
	HouseType      uint   `json:"house_type" gorm:"type:tinyint(1);default:0;comment:房屋类型"`      // 房屋类型:0(默认)|毛胚1|精装2|老房3
	DecorationMode uint   `json:"decoration_mode" gorm:"type:tinyint(1);default:0;comment:装修方式"` // 装修方式:0(默认)|全包1|半包2|清包3
	HouseArea      uint   `json:"house_area" gorm:"type:tinyint(1);default:0;comment:房屋面积"`      // 房屋面积
	TotalCost      uint   `json:"total_cost" gorm:"type:tinyint(1);default:0;comment:总花费"`       // 总花费
	City           string `json:"city" gorm:"type:varchar(100);comment:房屋所在城市"`                  // 房屋所在城市
	Contract       string `json:"contract" gorm:"type:varchar(120);comment:合同截图"`                // 合同截图
	Status         uint   `json:"status" gorm:"type:tinyint(1);default:1;comment:数据状态"`          // 数据状态 0删除 1正常
}

func (DesignerComment) TableName() string {
	return "designer_comment"
}
