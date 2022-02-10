package models

import (
	"time"

	"gorm.io/gorm"
)

// struct user
type Products struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(255)" json:"name" form:"name"`
	UsersID       uint      `json:"user_id" form:"user_id"`
	SubcategoryID int       `json:"subcategory_id" form:"subcategory_id"`
	CityID        int       `json:"city_id" form:"city_id"`
	Caption       string    `gorm:"type:longtext;not null" json:"caption" form:"caption"`
	Longitude     float64   `gorm:"type:varchar(30);not null" json:"lon" form:"lon"`
	Latitude      float64   `gorm:"type:varchar(30);not null" json:"lat" form:"lat"`
	Rating        float64   `gorm:"type:decimal(2,1)" json:"rating" form:"rating"`
	Photos        []Photos  `gorm:"foreignKey:ProductsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comment       []Comment `gorm:"foreignKey:ProductsID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// struct body create product
type BodyCreateProducts struct {
	Name          string `gorm:"type:varchar(255)" json:"name" form:"name"`
	SubcategoryID int    `json:"subcategory_id" form:"subcategory_id"`
	CityID        int    `json:"city_id" form:"city_id"`
	Caption       string `gorm:"type:longtext;not null" json:"caption" form:"caption"`
}

// struct get product
type GetAllProduct struct {
	ID               uint
	UsersID          uint
	Name             string
	Subcategory_Name string
	SubcategoryID    int
	CityID           int
	City_Name        string
	Description      string
	Url              string
}

// struct get product
type GetProduct struct {
	ID               uint
	UsersID          uint
	CreatedAt        time.Time
	Nama             string
	Phone_Number     string
	Name             string
	SubcategoryID    int
	Subcategory_Name string
	CityID           int
	City_Name        string
	Description      string
	Longitude        float64
	Latitude         float64
	Url              []string
}
