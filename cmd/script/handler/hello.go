package handler

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Hello = &cobra.Command{
		Use: "hello",
		PreRun: func(cmd *cobra.Command, args []string) {
			prepend()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return helloHandler(args)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			release()
		},
	}
)

// 业务逻辑写在这里
func helloHandler(args []string) error {
	fmt.Println(args)
	fmt.Println("hello handler")
	return nil
}
