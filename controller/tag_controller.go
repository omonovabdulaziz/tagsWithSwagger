package controller

import (
	"CODE/helper"
	"CODE/request"
	"CODE/response"
	"CODE/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TagsController struct {
	tagService service.TagsService
}

func NewTagsController(tagService service.TagsService) *TagsController {
	return &TagsController{
		tagService: tagService,
	}
}

// Create		godoc
// @Summary			Create tags
// @Description		Save tags data in Db.
// @Param			tags body request.CreateTagsRequest true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
func (controller *TagsController) Create(ctx *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	controller.tagService.Crate(createTagsRequest)
	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, webResponse)
}

// Update godoc
// @Summary  UpdateTags
// @Description UpdateTags data.
// @Param tagId path string true "update tags by id"
// @Param tags body request.CreateTagsRequest true "Update tags"
// @Produce application/json
// @Tags tags
// @Success 200 {object} response.Response{}
// @Router  /tags [put]
func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagRequest.Id = id
	controller.tagService.Update(updateTagRequest)
	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, webResponse)
}

// DeleteTags godoc
// @Summary  DeleteTags
// @Description RemoveTags By Id from data.
// @Produce application/json
// @Tags tags
// @Success 200 {object} response.Response{}
// @Router  /tags/{tagID} [delete]
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, webResponse)
}

// FindAll godoc
// @Summary Get All tags
// @Description Return list tags
// @Tags tags
// @Success 200 {object} response.Response{}
// @Router  /tags [get]
func (controller *TagsController) FindAll(ctx *gin.Context) {
	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   controller.tagService.FindAll(),
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, webResponse)
}

// FindById godoc
// @Summary Get Single tags By Id
// @Description Return tags by id
// @Produce application/json
// @Tags tags
// @Success 200 {object} response.Response{}
// @Router  /tags/{tagId} [get]
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   controller.tagService.FindById(id),
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, webResponse)
}
