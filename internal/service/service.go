package service

import (
	"net/http"

	"github.com/02amanag/setup/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ServiceStruct struct {
	usecase usecase.UsecaseStruct
}

func NewServiceStruct(object *usecase.UsecaseStruct) *ServiceStruct {
	return &ServiceStruct{
		usecase: *object,
	}
}

func (s *ServiceStruct) CustomerForm(c *gin.Context) {
	response := s.usecase.DetailsForm("param passing")
	c.JSON(http.StatusOK, gin.H{"Message": response})
}
