package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LogInTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogOut      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList(db *gorm.DB) []*UserBasic {
	userdatas := make([]*UserBasic, 100)
	db.Find(&userdatas)
	for v, _ := range userdatas {
		log.Println(v)
	}
	return userdatas
}
