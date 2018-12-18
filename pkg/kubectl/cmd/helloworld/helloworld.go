package helloworld

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/util/i18n"
)

// NewCmdHelloWorld creates a command object for print a hello world
// from "helloworld" action.
func NewCmdHelloWorld() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello-world",
		Short: i18n.T("Print 'hello world'"),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello World")
		},
	}
	return cmd
}
