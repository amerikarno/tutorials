package controller

import (
	"net/http"
	"wnexam/service"

	"github.com/gin-gonic/gin"
)

//Groups struct add
type Groups struct{
	AgeGroups service.Ages `json:"AgeGroups"`
}

//Province struct added
type Province struct{
	Province []service.Province `json:"Province"`
}
//WnController interface added
type WnController interface{
	GetCovidSum(gc *gin.Context)
}

type wnController struct{
	svc service.WnService
}

//NewWnController added
func NewWnController(svc service.WnService) WnController{
	return wnController{svc}
}


func (c wnController) GetCovidSum(gc *gin.Context){
	provinces, _ := c.svc.GetProvince()
	ages, _ := c.svc.GetAges()
	age := Groups{
		AgeGroups: *ages,
	}
	gc.JSON(http.StatusOK,provinces)
	gc.JSON(http.StatusOK,age)
}