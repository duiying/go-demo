package user

import (
	"github.com/duiying/go-demo/util"
	"log"
)

func find4Dao(id int) User {
	var u User
	row := util.Db.QueryRow("select * from user where id = ?", id)
	e := row.Scan(&u.ID, &u.Name, &u.Email, &u.Root, &u.Mtime, &u.Ctime)
	if e != nil {
		log.Printf("user id %d 没有找到", id)
	}
	return u
}
