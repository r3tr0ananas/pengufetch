package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Penguin = `
             %s
      _      
    ('%s')    %s
   ((   ))   %s
    ^^-^^    %s
`

var reset = "\033[0m"
var v = "\033[33mv\033[0m"

func fetch(cmd *cobra.Command, args []string) {
	fmt.Printf(
		Penguin,
		fmt.Sprintf("\033[34m%s%s", GetHostname(), reset),
		v,
		fmt.Sprintf("\033[32mOS %s%s", reset, GetOS()),
		fmt.Sprintf("\033[33mKernel %s%s", reset, GetKernel()),
		fmt.Sprintf("\033[36mUptime %s%s", reset, GetUptime()),
	)
}

func main() {
	var RootCmd = &cobra.Command{
		Use:   "pengufetch",
		Short: "Smol and cute fetch with a penguin :3",
		Run:   fetch,
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Print(err)
	}
}
