package models

import (
	"gorm.io/gorm"
	"time"
)

type user struct {
	gorm.Model

	FirstName string `gorm:"notnull" json:"first_name"`
	LastName  string `gorm:"notnull" json:"last_name"`
	Username string `gorm:"uniqueIndex;notnull" json:"username"`
	Email    string `gorm:"uniqueIndex;notnull" json:"email"` // check email for nith.ac.in domain, else reject
	CurrentLevel int    `gorm:"notnull" json:"current_level"` // current level of the user, starts from 1
	TotalTimeSpent int   `gorm:"notnull" json:"time_spent"` // total time spent on all levels in seconds
	PreviousLvlTime time.Time `gorm:"notnull" json:"previous_lvl_time"` // time at which the previous level was completed
}