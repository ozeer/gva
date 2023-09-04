package designer

import (
	"strings"

	"github.com/ozeer/gva/global"
	"github.com/ozeer/gva/model/common/request"
	"github.com/ozeer/gva/model/designer"
	designerRes "github.com/ozeer/gva/model/designer/response"
	"github.com/ozeer/gva/utils"
	"github.com/ozeer/gva/utils/oss"
)

type CommentService struct{}

func (t *CommentService) GetCommentDetail(id uint) (c designerRes.DesignerCommentDetail, err error) {
	comment, err := t.GetCommentById(id)

	// 处理图片
	picsArr := strings.Split(comment.Pics, ",")
	var pics []string
	for _, v := range picsArr {
		picUrl := oss.GetImgUrl(v)
		pics = append(pics, picUrl)
	}

	info := designerRes.DesignerCommentDetail{
		ID:             comment.ID,
		CreatedAt:      comment.CreatedAt,
		UpdatedAt:      comment.UpdatedAt,
		Point:          comment.Point,
		Content:        comment.Content,
		Pics:           pics,
		HouseType:      comment.HouseType,
		DecorationMode: comment.DecorationMode,
		HouseArea:      comment.HouseArea,
		TotalCost:      comment.TotalCost,
		City:           comment.City,
		Contract:       oss.GetImgUrl(comment.Contract),
	}

	return info, err
}

func (t *CommentService) GetCommentInfoList(condition any, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&designer.DesignerComment{})

	var CommentList []designer.DesignerComment
	err = db.Where("status = ?", global.STATUS_NORMAL).Count(&total).Error
	if err != nil {
		return CommentList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Where("status = ?", global.STATUS_NORMAL).Find(&CommentList).Error
	}
	return CommentList, total, err
}

func (t *CommentService) AddComment(c designer.DesignerComment) (err error) {
	err = global.GVA_DB.Create(&c).Error
	return err
}

func (t *CommentService) DeleteComment(c designer.DesignerComment) (err error) {
	err = global.GVA_DB.Model(&c).Where("id = ?", c.ID).Updates(map[string]interface{}{"status": global.STATUS_DELETE, "deleted_at": utils.Date()}).Error
	return err
}

func (t *CommentService) UpdateComment(c *designer.DesignerComment) (err error) {
	err = global.GVA_DB.Save(c).Error
	return err
}

func (t *CommentService) GetCommentById(id uint) (c designer.DesignerComment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&c).Error
	return
}

func (t *CommentService) ListComment(info request.PageInfo, condition map[string]string) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&designer.DesignerComment{})

	var CommentList []designer.DesignerComment
	err = db.Where("uid in ?", condition["uid"]).Count(&total).Error
	if err != nil {
		return CommentList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Where("uid in ?", condition["uid"]).Find(&CommentList).Error
	}

	return CommentList, total, err
}
