package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/KeiranScript/e-z-rec" // Replace with your actual module path
)

func main() {
	var mode string

	// Define root command
	var rootCmd = &cobra.Command{
		Use:   "screen-recorder",
		Short: "A pretty tool to record your screen with wf-recorder",
		Run: func(cmd *cobra.Command, args []string) {
			err := modules.RecordScreen(mode)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		},
	}

	// Add flag for selecting recording mode
	rootCmd.Flags().StringVarP(&mode, "mode", "m", "fullscreen", "Recording mode: partial, window, fullscreen")

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
