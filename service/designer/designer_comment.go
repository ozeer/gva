package designer

import (
	"strings"

	"github.com/ozeer/gva/model/common/request"
	"github.com/ozeer/gva/model/designer"
	designerRes "github.com/ozeer/gva/model/designer/response"
	"github.com/ozeer/gva/utils/oss"
)

type CommentService struct{}

func (t *CommentService) GetCommentDetail(id uint) (c designerRes.DesignerCommentDetail, err error) {
	comment, err := designer.GetCommentById(id)

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

func (t *CommentService) GetCommentInfoList(info request.PageInfo, condition map[string]string) (list interface{}, total int64, err error) {

	return nil, 0, nil
}
