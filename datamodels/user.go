package datamodels

import(
  "time"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int64     `json:"id" form:"id"`
	Firstname      string    `json:"firstname" form:"firstname"`
  Lastname      string    `json:"lasttname" form:"lastname"`
	Username       string    `json:"username" form:"username"`
  Password       string    `json:"password" form:"password"`
  ConfirmPwd      string    `json:"confirmpwd" form:"confirmpwd"`
  Birthdate      time    `json:"birthdate" form:"birthdate"`
}
