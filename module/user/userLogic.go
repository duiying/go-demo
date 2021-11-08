package user

import (
	response2 "github.com/duiying/go-demo/pkg/response"
)

func Find4Logic(id int) User {
	return Find4Dao(id)
}

func Search4Logic(p int, size int, id int, name string, email string) response2.List {
	total := Count4Dao(id, name, email)
	list := List4Dao(p, size, id, name, email)
	data := response2.List{
		P: p,
		Size: size,
		Total: total,
		List: list,
	}
	return data
}

func Update4Logic(user User) int {
	// 先查询记录是否存在
	exist := Find4Dao(user.ID)
	if exist.ID == 0 {
		return response2.ErrorCode
	}

	return Update4Dao(user)
}

func Create4Logic(user User) int {
	return Create4Dao(user)
}
