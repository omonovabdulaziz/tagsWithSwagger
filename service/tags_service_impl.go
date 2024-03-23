package service

import (
	"CODE/helper"
	"CODE/model"
	"CODE/repository"
	"CODE/request"
	"CODE/response"
	"github.com/go-playground/validator/v10"
)

type TagServiceImpl struct {
	TagsRepository repository.TagsRepository
	validate       *validator.Validate
}

func (t TagServiceImpl) Crate(tags request.CreateTagsRequest) {
	err := t.validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t TagServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}

func (t TagServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

func (t TagServiceImpl) FindById(tagsId int) response.TagsResponse {
	tags, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)
	tagsResponse := response.TagsResponse{
		Id:   tags.ID,
		Name: tags.Name,
	}
	return tagsResponse
}

func (t TagServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, m := range result {
		tag := response.TagsResponse{
			Id:   m.ID,
			Name: m.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}

func NewTagsServiceImpl(tagsRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagServiceImpl{
		TagsRepository: tagsRepository,
		validate:       validate,
	}
}
