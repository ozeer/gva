package response

import (
	"time"

	"github.com/ozeer/gva/model/designer"
)

type DesignerCommentDetailResponse struct {
	Info DesignerCommentDetail `json:"info"`
}

type DesignerCommentDetail struct {
	ID             uint      `json:"id"`              // 评论ID
	Point          string    `json:"point"`           // 评价打分
	Content        string    `json:"content"`         // 评价内容
	Pics           []string  `json:"pics"`            // 装修图片
	HouseType      uint      `json:"house_type"`      // 房屋类型:0(默认)|毛胚1|精装2|老房3
	DecorationMode uint      `json:"decoration_mode"` // 装修方式:0(默认)|全包1|半包2|清包3
	HouseArea      uint      `json:"house_area"`      // 房屋面积
	TotalCost      uint      `json:"total_cost"`      // 总花费
	City           string    `json:"city"`            // 房屋所在城市
	Contract       string    `json:"contract"`        // 合同截图
	CreatedAt      time.Time `json:"created_at"`      // 创建时间
	UpdatedAt      time.Time `json:"updated_at"`      // 更新时间
}

type DesignerCommentList struct {
	List  []designer.DesignerComment `json:"list"`  // 评论列表
	Total uint                       `json:"total"` // 评论总条数
}
