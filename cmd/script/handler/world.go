package handler

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	World = &cobra.Command{
		Use: "world",
		PreRun: func(cmd *cobra.Command, args []string) {
			prepend()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return worldHandler(args)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			release()
		},
	}
)

// 业务逻辑写在这里
func worldHandler(args []string) error {
	fmt.Println(args)
	fmt.Println("world handler")
	return nil
}
