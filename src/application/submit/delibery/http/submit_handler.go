package http

import (
	"context"
	"net/http"
	"time"

	commonEntity "onlineApplicationAPI/src/application/common/entity"
	"onlineApplicationAPI/src/application/submit"
	"onlineApplicationAPI/src/application/submit/entity"

	"github.com/gin-gonic/gin"
)

type submitHandler struct {
	authService   submit.AuthenticationService
	submitUseCase submit.UseCase
}

func NewSubmitGinHTTPHandler(gin *gin.Engine, submitUserCase submit.UseCase, auth submit.AuthenticationService) *gin.Engine {
	handler := &submitHandler{
		authService:   auth,
		submitUseCase: submitUserCase,
	}
	gin.POST("/v1/application/submit", handler.Submit)
	return gin
}

type SubmitApplicationRequest struct {
	Email         string    `form:"email" binding:"required" json:"email"`
	Name          string    `form:"name" binding:"required" json:"name"`
	PersonalID    string    `form:"personalID" binding:"required" json:"personal_id"`
	BirthLocation string    `form:"birthLocation" binding:"required" json:"birth_location"`
	BirthDate     time.Time `form:"birthDate" binding:"required" json:"birth_date"`
}

func (handler *submitHandler) Submit(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	userUUID, tErr := handler.authService.VerifyToken(ctx, c.Request)
	if tErr != nil {
		c.JSON(http.StatusUnauthorized, tErr)
		return
	}
	var requestData SubmitApplicationRequest
	err := c.Bind(&requestData)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	fileHeader, err := c.FormFile("cv")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	newApplication := entity.NewApplication{
		Email:      requestData.Email,
		Name:       requestData.Name,
		PersonalID: requestData.PersonalID,
		Birth: struct {
			Date     time.Time
			Location commonEntity.Location
		}{Date: requestData.BirthDate, Location: struct {
			Name      string
			Latitude  *float64
			Longitude *float64
		}{Name: requestData.BirthLocation, Latitude: nil, Longitude: nil}},
		CreatedBy: userUUID,
	}

	ok, err := handler.submitUseCase.SubmitApplication(ctx, newApplication, file)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	if ok {
		c.Status(http.StatusCreated)
	} else {
		c.Status(http.StatusInternalServerError)
	}

}
