package designer

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ozeer/gva/global"
	"github.com/ozeer/gva/model/common/request"
	"github.com/ozeer/gva/model/common/response"
	"github.com/ozeer/gva/model/designer"
	designerRes "github.com/ozeer/gva/model/designer/response"
	"github.com/ozeer/gva/utils"
	"go.uber.org/zap"
)

type CommentApi struct{}

func (d *CommentApi) Create(c *gin.Context) {
	var comment designer.DesignerComment

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(comment, utils.DesignerCommentVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = commentService.AddComment(comment)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (d *CommentApi) Delete(c *gin.Context) {
	var comment designer.DesignerComment

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(comment.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = commentService.DeleteComment(comment)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (d *CommentApi) Edit(c *gin.Context) {
	var comment designer.DesignerComment

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(comment.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(comment, utils.DesignerCommentVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = commentService.UpdateComment(&comment)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (d *CommentApi) All(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	commentList, total, err := commentService.GetCommentInfoList(c.Params, pageInfo)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     commentList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
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
