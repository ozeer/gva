package designer

import (
	"github.com/ozeer/gva/global"
	"github.com/ozeer/gva/model/common/request"
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
}

func (DesignerComment) TableName() string {
	return "designer_comment"
}

func AddComment(c DesignerComment) (err error) {
	err = global.GVA_DB.Create(&c).Error
	return err
}

func DeleteComment(c DesignerComment) (err error) {
	err = global.GVA_DB.Delete(&c).Error
	return err
}

func UpdateComment(c DesignerComment) (err error) {
	err = global.GVA_DB.Save(c).Error
	return err
}

func GetCommentById(id uint) (c DesignerComment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&c).Error
	return
}

func ListComment(info request.PageInfo, condition map[string]string) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&DesignerComment{})

	var CommentList []DesignerComment
	err = db.Where("uid in ?", condition["uid"]).Count(&total).Error
	if err != nil {
		return CommentList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Where("uid in ?", condition["uid"]).Find(&CommentList).Error
	}

	return CommentList, total, err
}
