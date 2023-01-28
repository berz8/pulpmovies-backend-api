package models

import "time"

type User struct {
  ID uint `json:"id" gorm:"primary_key"`
  Username string `json:"username"`
  Email string `json:"email"`
  Birthday *time.Time `json:"birthday"`
  Biography *string `json:"biography"`
  ProfilePath *string `json:"profilePath"`
  AccountStatus *string `json:"accountStatus"`
  FullName *string `json:"fullName"`
  Password string `json:"-"`
  CreatedAt time.Time
  UpdatedAt time.Time
}
