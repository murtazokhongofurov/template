package models

import "time"

type User struct {
	UserID      int        `json:"user_id" db:"user_id" redis:"user_id" validate:"omitempty"`
	FirstName   string     `json:"first_name" db:"first_name" redis:"first_name" validate:"required,lte=30"`
	LastName    string     `json:"last_name" db:"last_name" redis:"last_name" validate:"required,lte=30"`
	Email       string     `json:"email" db:"email" redis:"email" validate:"omitempty,lte=60,email"`
	Password    string     `json:"password,omitempty" db:"password" redis:"password" validate:"omitempty,required,gte=6"`
	Role        *string    `json:"role,omitempty" db:"role" redis:"role" validate:"omitempty"`
	About       *string    `json:"about,omitempty" db:"about" redis:"about" validate:"omitempty"`
	Avatar      *string    `json:"avatar,omitempty" db:"avatar" redis:"avatar" validate:"omitempty,lte=512,url"`
	PhoneNumber *string    `json:"phone_number,omitempty" db:"phone_number" redis:"phone_number" validate:"omitempty,lte=20"`
	Address     *string    `json:"address,omitempty" db:"address" redis:"address" validate:"omitempty,lte=250"`
	City        *string    `json:"city,omitempty" db:"city" redis:"city" validate:"omitempty,lte=24"`
	Country     *string    `json:"country,omitempty" db:"country" redis:"country" validate:"omitempty,lte=24"`
	Gender      *string    `json:"gender,omitempty" db:"gender" redis:"gender" validate:"omitempty,lte=10"`
	Postcode    *int       `json:"postcode,omitempty" db:"postcode" redis:"postcode" validate:"omitempty"`
	Birthday    *time.Time `json:"birthday,omitempty" db:"birthday" redis:"birthday" validate:"omitempty,lte=10"`
	CreatedAt   time.Time  `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
	LoginDate   time.Time  `json:"login_date" db:"login_date" redis:"login_date"`
}

// All Users response
type UsersList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Users      []*User `json:"users"`
}

// Find user query
type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
