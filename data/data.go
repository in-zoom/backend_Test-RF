package data

import (
    jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json: "phoneNumber"`
}
 

type Auth struct{
	Username string    `json:"username"`
	Password string  `json:"password"`
 }
 
 type Token struct {
	UserId int
	jwt.StandardClaims
}