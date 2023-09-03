package designer

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ozeer/gva/global"
	"github.com/ozeer/gva/model/common/response"
	"github.com/ozeer/gva/model/designer"
	designerRes "github.com/ozeer/gva/model/designer/response"
	"github.com/ozeer/gva/utils"
	"go.uber.org/zap"
)

type CommentApi struct{}

func (d *CommentApi) All(c *gin.Context) {

}

func (d *CommentApi) Detail(c *gin.Context) {
	var comment designer.DesignerComment
	err := c.ShouldBindQuery(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(comment.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 1、从service类中获取信息
	data, err := commentService.GetCommentDetail(comment.ID)

	log.Println(comment, err)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	// 2、处理数据结构

	// 3、返回数据
	response.OkWithDetailed(designerRes.DesignerCommentDetailResponse{Info: data}, "获取成功", c)
}

func (d *CommentApi) Add(c *gin.Context) {

}

func (d *CommentApi) Delete(c *gin.Context) {

}

func (d *CommentApi) Edit(c *gin.Context) {

}

func (d *CommentApi) Get(c *gin.Context) {

}
