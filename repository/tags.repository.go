package repository

import "CODE/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tags int)
	FindById(tagsId int) (tags model.Tags, err error)
	FindAll() []model.Tags
}
