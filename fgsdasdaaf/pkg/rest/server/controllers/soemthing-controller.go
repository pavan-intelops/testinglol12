package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pavan-intelops/testinglol12/fgsdasdaaf/pkg/rest/server/daos/clients/sqls"
	"github.com/pavan-intelops/testinglol12/fgsdasdaaf/pkg/rest/server/models"
	"github.com/pavan-intelops/testinglol12/fgsdasdaaf/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
	"strconv"
)

type SoemthingController struct {
	soemthingService *services.SoemthingService
}

func NewSoemthingController() (*SoemthingController, error) {
	soemthingService, err := services.NewSoemthingService()
	if err != nil {
		return nil, err
	}
	return &SoemthingController{
		soemthingService: soemthingService,
	}, nil
}

func (soemthingController *SoemthingController) UpdateSoemthing(context *gin.Context) {
	// validate input
	var input models.Soemthing
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger soemthing update
	if _, err := soemthingController.soemthingService.UpdateSoemthing(id, &input); err != nil {
		log.Error(err)
		if errors.Is(err, sqls.ErrNotExists) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

func (soemthingController *SoemthingController) DeleteSoemthing(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger soemthing deletion
	if err := soemthingController.soemthingService.DeleteSoemthing(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}
