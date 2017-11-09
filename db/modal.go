package stardictdb

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Dictionary struct {
	gorm.Model
	Code         string
	Name         string `gorm:"unique_index:idx_name_ver"`
	Version      string `gorm:"unique_index:idx_name_ver"`
	Description  string
	Date         *time.Time
	Author       string
	Email        string
	Website      string
	Type         string
	WordCount    int
	SynonymCount int
}

type Word struct {
	gorm.Model
	Dict   *Dictionary
	DictID uint   `gorm:"index:idx_dict_word"`
	Word   string `gorm:"index:idx_dict_word"`
}

type Synonym struct {
	gorm.Model
	Dict    *Dictionary
	Word    *Word
	DictID  uint
	WordID  uint
	Synonym string
}

type Content struct {
	gorm.Model
	Dict    *Dictionary
	Word    *Word
	DictID  uint
	WordID  uint
	Type    string
	Content string
}
