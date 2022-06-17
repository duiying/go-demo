package user

// AllowedRootMap 允许的 root 值
var AllowedRootMap = map[int]string{0: "否", 1: "是"}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Root  int    `json:"root"`
	Mtime string `json:"mtime"`
	Ctime string `json:"ctime"`
}
