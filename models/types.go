package models

import "gorm.io/gorm"

// Generated by https://quicktype.io

type ServerInfo struct {
	ID           string
	Region       string
	Availability string
	Url          string
}
type ServerDB struct {
	gorm.Model
	ID           string
	Region       string
	Availability string
	Url          string
}
