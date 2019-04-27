package main

import (
	"time"

	completeRepository "onlineApplicationAPI/src/application/complete/repository"
	completeService "onlineApplicationAPI/src/application/complete/services"
	completeUsecase "onlineApplicationAPI/src/application/complete/usecase"
	"onlineApplicationAPI/src/application/submit/delibery/http"
	"onlineApplicationAPI/src/application/submit/repository"
	"onlineApplicationAPI/src/application/submit/service"
	"onlineApplicationAPI/src/application/submit/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	ginEngine := createGinHttpHandler()

	authenticationService := service.NewDefaultAuhtenticationServices()
	emailService := service.NewDefaultEmailServices()
	bullySystem := completeService.NewDefaultBullyServices()

	applicationRepository := completeRepository.NewMysqlApplicationRepository()
	completeQueueUseCase := completeUsecase.NewCompleteUseCase(emailService, bullySystem, applicationRepository)

	fileRepository := repository.NewDefaultFileRepository()
	submitUseCase := usecase.NewSubmitUseCase(emailService, completeQueueUseCase, fileRepository)

	http.NewSubmitGinHTTPHandler(ginEngine, submitUseCase, authenticationService)

	ginEngine.Run(":8080")

}

func createGinHttpHandler() *gin.Engine {
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	config.AllowAllOrigins = true
	server.Use(cors.New(config))
	return server
}
