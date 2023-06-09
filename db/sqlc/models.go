// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"time"
)

type Category struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID         int32     `json:"id"`
	Products   []int32   `json:"products"`
	Quantities []int32   `json:"quantities"`
	User       int32     `json:"user"`
	Sum        int32     `json:"sum"`
	Paid       bool      `json:"paid"`
	Status     string    `json:"status"`
	Created    time.Time `json:"created"`
}

type Product struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Category int32  `json:"category"`
	Photo    []byte `json:"photo"`
	Price    int32  `json:"price"`
	Shop     int32  `json:"shop"`
}

type Shop struct {
	ID       int32   `json:"id"`
	Name     string  `json:"name"`
	Products []int32 `json:"products"`
}

type User struct {
	ID             int32   `json:"id"`
	Username       string  `json:"username"`
	Email          string  `json:"email"`
	HashedPassword string  `json:"hashed_password"`
	Favourites     []int32 `json:"favourites"`
}
