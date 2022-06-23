package cmd

import (
	"errors"
	"github.com/duiying/go-demo/cmd/script"
	"github.com/duiying/go-demo/cmd/server"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "entry",
	SilenceUsage: false,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(color.RedString("requires at least one arg"))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(server.Server) // 服务
	rootCmd.AddCommand(script.Script) // 脚本
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
