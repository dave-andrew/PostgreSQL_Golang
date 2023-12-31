// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewPost struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	UserID      string `json:"UserID"`
}

type NewUser struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

type Post struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	UserID      string `json:"UserID" gorm:"foreignKey:ID"`
	User *User `json:"User"`
}

type User struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}
