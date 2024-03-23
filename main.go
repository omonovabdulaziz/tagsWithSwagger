package main

import (
	"CODE/config"
	"CODE/controller"
	_ "CODE/docs"
	"CODE/helper"
	"CODE/model"
	"CODE/repository"
	"CODE/router"
	"CODE/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
)

// @title TagService APi
// @version 1.0
// @description  A Tag service API in Go using GIn Framework
// @BasePath /api

func main() {
	log.Info().Msg("Start Server!")
	//Database
	db := config.DatabaseConnection()
	//Validation
	validate := validator.New()
	//Hibernate
	db.Table("tags").AutoMigrate(&model.Tags{})
	//Repository
	tagRepository := repository.NewTagsRepositoryImpl(db)
	//Service
	tagService := service.NewTagsServiceImpl(tagRepository, validate)

	//Controller
	tagsController := controller.NewTagsController(tagService)

	//Router
	routers := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8088",
		Handler: routers,
	}
	helper.ErrorPanic(server.ListenAndServe())
}
