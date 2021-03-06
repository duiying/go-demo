package user

import (
	"fmt"
	"github.com/duiying/go-demo/pkg/logger"
	"github.com/duiying/go-demo/pkg/mysql"
	response2 "github.com/duiying/go-demo/pkg/response"
	"strings"
)

func Find4Dao(id int) User {
	var u User
	sql := "SELECT * FROM `user` WHERE id = ?"
	row := mysql.Db.QueryRow(sql, id)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Root, &u.Mtime, &u.Ctime)
	if err != nil {
		logger.Error("user id 没有找到", "id", id)
	}
	return u
}

func List4Dao(p int, size int, id int, name string, email string) []User {
	where := " WHERE 1 = 1"
	if id != 0 {
		where += fmt.Sprintf(" AND `id` = %d", id)
	}
	if name != "" {
		where += fmt.Sprintf(" AND `name` = '%s'", name)
	}
	if email != "" {
		where += fmt.Sprintf(" AND `email` = '%s'", email)
	}
	offset := (p - 1) * size
	order := " ORDER BY `id` ASC"
	limit := fmt.Sprintf(" LIMIT %d, %d", offset, size)
	sql := "SELECT * FROM `user`" + where + order + limit
	rows, err := mysql.Db.Query(sql)
	// 创建一个切片来存储数据
	list := make([]User, 0)
	if err != nil {
		return list
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		_ = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Root, &user.Ctime, &user.Mtime)
		// 追加到切片中
		list = append(list, user)
	}
	return list
}

func Count4Dao(id int, name string, email string) int {
	var count int
	where := " WHERE 1 = 1"
	if id != 0 {
		where += fmt.Sprintf(" AND `id` = %d", id)
	}
	if name != "" {
		where += fmt.Sprintf(" AND `name` = '%s'", name)
	}
	if email != "" {
		where += fmt.Sprintf(" AND `email` = '%s'", email)
	}

	sql := "SELECT count(*) FROM `user`" + where

	err := mysql.Db.QueryRow(sql).Scan(&count)

	if err == nil {
		return count
	}
	return 0
}

func Update4Dao(user User) int {
	sql := "UPDATE `user` SET"
	if user.Email != "" {
		sql += fmt.Sprintf(" `email` = '%s',", user.Email)
	}
	if user.Name != "" {
		sql += fmt.Sprintf(" `name` = '%s',", user.Name)
	}
	if user.Root != -1 {
		sql += fmt.Sprintf(" `name` = %d,", user.Root)
	}
	sql = strings.TrimRight(sql, ",")
	where := fmt.Sprintf(" WHERE `id` = %d", user.ID)
	sql += where
	res, err := mysql.Db.Exec(sql)
	if err != nil {
		logger.Error("SQL 错误了", "sql", sql)
		return response2.ErrorCode
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0
	}
	return int(affected)
}

func Create4Dao(user User) int {
	sql := "INSERT INTO `user` (name, email, root) values(?, ?, ?)"

	res, err := mysql.Db.Exec(sql, user.Name, user.Email, user.Root)
	if err != nil {
		logger.Error("创建用户失败", "name", user.Name, "email", user.Email, "root", user.Root)
		return response2.ErrorCode
	}

	id, _ := res.LastInsertId()
	return int(id)
}
