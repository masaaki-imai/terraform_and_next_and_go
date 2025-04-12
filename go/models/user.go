package models

import "time"

type User struct {
	ID             uint      `gorm:"primaryKey;column:id" json:"id"`
	Email          string    `gorm:"column:email" json:"email"`
	Password       string    `gorm:"column:password" json:"password"`
	Name           string    `gorm:"column:name" json:"name"`
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	PhoneNumber int    `gorm:"column:phone_number" json:"phone_number"`
	Address1    string `gorm:"column:address_1" json:"address_1"`
	Address2    string `gorm:"column:address_2" json:"address_2"`
	Address3    string `gorm:"column:address_3" json:"address_3"`
	PostCode1   int    `gorm:"column:post_code_1" json:"post_code_1"`
	PostCode2   int    `gorm:"column:post_code_2" json:"post_code_2"`
	LastTimeLogin  *time.Time `gorm:"column:last_time_login" json:"last_time_login"`
	CommonFields
}
