package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	env string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "order-service-go",
	Short: "Command Utility Tools for Bitwyre Projects",
	Long: `Command Utility Tools for Bitwyre Projects. You can execute any command 
related to this project. For example: Start Service with specific environment, or execute data seeder, run the test.etc`,
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:       "start",
	Short:     "Start the server with specific environment: [ dev | prod ]",
	Long:      `Start the server with specific environment: [ dev | prod ]. By default it will use dev environment`,
	ValidArgs: []string{"dev", "prod"},
	Run: func(cmd *cobra.Command, args []string) {
		startApp(cmd)
	},
}

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup Project for the first time",
	Long:  `Setup Project for the first time. If you new with this project, you can execute this command. It will prepare all required dependency`,
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("/bin/sh", "run.sh", "setup").Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(out))
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
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(setupCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test.yaml)")
	startCmd.Flags().StringVarP(&env, "env", "e", "dev", "Environment: dev | prod")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startApp(cmd *cobra.Command) bool {
	for _, v := range cmd.ValidArgs {
		if v == env {
			err := os.Setenv("ENV", env)
			if err != nil {
				return false
			}
			Server()
			return true
		}
	}

	log.Fatal("Env value only valid with: ", cmd.ValidArgs, ". given value:", env)
	return false
}
