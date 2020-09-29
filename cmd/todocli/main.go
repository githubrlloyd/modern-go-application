package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/sagikazarmark/modern-go-application/internal/app/todocli"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "todocli",
		Short: "TODO CLI manages TODOs.",
	}
	//Hard-coded pwd example
	pwd := "68af404b513073584c4b6f22b6c63e6b"
	
	todocli.Configure(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
