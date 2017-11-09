package cmd

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/wenerme/stardict/genproto/v1/stardictdata"
)

var data *stardictdata.StardictData

func open(file string) error {
	return nil
}
