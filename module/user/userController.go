package user

import (
	"github.com/duiying/go-demo/pkg/constant"
	redis2 "github.com/duiying/go-demo/pkg/redis"
	"github.com/duiying/go-demo/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

// Search 多条
func Search(c *gin.Context) {
	p, _ := strconv.Atoi(c.DefaultQuery("p", constant.DefaultP))
	size, _ := strconv.Atoi(c.DefaultQuery("size", constant.DefaultSize))
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	name := c.DefaultQuery("name", "")
	email := c.DefaultQuery("email", "")

	data := Search4Logic(p, size, id, name, email)
	response.Success(c, data)
}

// Find 单条
func Find(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	if id == 0 {
		response.Fail(c, response.ParamsError)
		return
	}

	user := Find4Logic(id)

	if user.ID == 0 {
		response.Fail(c, response.ExistError)
		return
	}

	response.Success(c, user)
}

// Update 更新
func Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))

	// 更新 user，id 必传
	if id == 0 {
		response.Fail(c, response.ParamsError)
		return
	}
	name := c.DefaultQuery("name", "")
	email := c.DefaultQuery("email", "")
	root, _ := strconv.Atoi(c.DefaultQuery("root", "-1"))

	var user User
	user.ID = id
	user.Root = root
	user.Name = name
	user.Email = email

	// 如果传了 root 参数，验证 root 参数的合法性
	if user.Root != -1 {
		_, ok := AllowedRootMap[user.Root]
		if !ok {
			response.Fail(c, response.ParamsError)
			return
		}
	}

	// 需要更新的字段，至少传 1 个
	if user.Root == -1 && user.Name == "" && user.Email == "" {
		response.Fail(c, response.ParamsError)
		return
	}

	affected := Update4Logic(user)
	if affected == response.ErrorCode {
		response.Fail(c, response.ExistError)
		return
	}

	response.Success(c, affected)
}

// Create 创建
func Create(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	email := c.DefaultQuery("email", "")
	root, _ := strconv.Atoi(c.DefaultQuery("root", "0"))

	var user User
	user.Root = root
	user.Name = name
	user.Email = email

	// 如果传了 root 参数，验证 root 参数的合法性
	_, ok := AllowedRootMap[user.Root]
	if !ok {
		response.Fail(c, response.ParamsError)
		return
	}

	// 需要更新的字段，至少传 1 个
	if user.Name == "" || user.Email == "" {
		response.Fail(c, response.ParamsError)
		return
	}

	lastInsertId := Create4Logic(user)
	if lastInsertId == response.ErrorCode {
		response.Fail(c, response.CreateError)
		return
	}

	response.Success(c, lastInsertId)
}

// Redis 测试 Redis
func Redis(c *gin.Context) {
	key := "key1"
	val := "val1"
	_, _ = redis2.GetInstance().Do("SET", key, val)
	res, _ := redis.String(redis2.GetInstance().Do("GET", key))
	response.Success(c, res)
}
