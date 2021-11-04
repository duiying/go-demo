package user

type UserFindValidate struct {
	id int `validate:"min=1"`
}