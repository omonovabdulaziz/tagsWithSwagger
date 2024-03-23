package service

import (
	"CODE/request"
	"CODE/response"
)

type TagsService interface {
	Crate(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
