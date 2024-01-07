package miniblog

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "miniblog",
		Short:        "A go practical project",
		Long:         "",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any aruments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	return cmd

}

// 实际业务代码入口
func run() error {
	fmt.Println("Hello MiniBlog!")
	return nil
}
