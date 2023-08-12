package model

import (
	"gorm.io/gorm"
)

type Protocol string

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type Decode string

const (
	ProxyImageUrl Decode = "proxy_image_url"
)

type Relay struct {
	gorm.Model
	Method Method `gorm:"default:'GET'"`
	Uri    string
	Decode Decode
}
