package user

import (
	"github.com/duiying/go-demo/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func List(c *gin.Context) {

}

func FindBak(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultPostForm("id", "0"))
	if id == 0 {
		response.Fail(c, response.ParamsError)
		return
	}

	return
}

func Find(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	if id == 0 {
		response.Fail(c, response.ParamsError)
		return
	}

	user := find4Logic(id)

	response.Success(c, user)
}
