package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/wenerme/letsgo/cobra"
	"github.com/wenerme/stardict/util"
)

var convConf = &_convConf{}

type _convConf struct {
	OutFile string
}

func (self *_convConf) Install(fs *pflag.FlagSet, cmd *wcobra.Command) error {
	return nil
}

// convCmd represents the conv command
var convCmd = &cobra.Command{
	Use:     "conv",
	Aliases: []string{"c"},
	Args:    cobra.MinimumNArgs(2),
	Short:   "Conversion tool for stardict",
	Long:    `convert pb/json/stardict -> pb/json/sqlite`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if data, err = stardictutil.Read(args[0]); err != nil {
			logrus.WithError(err).Fatal("failed to open file")
		}

		if err = stardictutil.Write(data, args[1]); err != nil {
			logrus.WithError(err).Fatal("failed to write file")
		}
	},
}

func init() {
	RootCmd.AddCommand(convCmd)
	wcobra.Wrap(convCmd).Install(convConf)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
