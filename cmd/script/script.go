package script

import (
	"errors"
	"fmt"
	"github.com/duiying/go-demo/cmd/script/handler"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
)

var (
	mode   string
	Script = &cobra.Command{
		Use: "script",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New(color.RedString("script requires at least one arg"))
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}
)

func init() {
	// 初始化 env
	Script.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode; eg:dev,test,prod")
	err := godotenv.Load(fmt.Sprintf(".env.%s", mode))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 注册子命令
	Script.AddCommand(handler.Hello)
	Script.AddCommand(handler.World)
}
