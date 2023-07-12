package model

import "github.com/ArtisanCloud/PowerLibs/v3/database"

type Lora struct {
	database.PowerModel

	Preview string `gorm:"comment:预览图" json:"preview"`
	Hash    string `gorm:"comment:Hash" json:"hash"`
	Alias   string `gorm:"comment:别名" json:"alias"`
	Name    string `gorm:"comment:名称" json:"name"`
}
