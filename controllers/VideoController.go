package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Video struct {
	Videos []string `json:"list"`
}

type VideoListController struct {
	ID int `uri:"id" binding:"required" json:"id" description:"Video Id" default:"-1"`
}

func (vcontroller *VideoListController) Handler(c *gin.Context) {
	fmt.Println(" Trying to return the Video for ", vcontroller.ID)
	vds := []string{"a", "b", "c"}
	v := Video{
		Videos: vds,
	}

	c.JSON(http.StatusOK, v)

}

/*

 {"Dil", "hall", "chall"},
 Video{
		"Videos": nil,
}
	})


*/
