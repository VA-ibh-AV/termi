/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"time"

	"github.com/VA-ibh-AV/termi/internal/ai"
	"github.com/VA-ibh-AV/termi/internal/ui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "termi [prompt]",
	Short: "AI-powered terminal command helper",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return cmd.Help()
		}

		prompt := args[0]

		stop := ui.StartSpinner(cmd.Context(), "thinking...")

		start := time.Now()
		resp, err := ai.Generate(cmd.Context(), prompt)

		stop()

		if err != nil {
			return err
		}

		ui.PrintAIResponse(resp.Command, time.Since(start))

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.termi.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
