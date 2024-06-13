/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/muesli/termenv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)


var p = termenv.ColorProfile()


var (
	cyan  = termenv.String().Foreground(p.Color("6")).Styled
	green = termenv.String().Foreground(p.Color("10")).Styled
	yellow = termenv.String().Foreground(p.Color("3")).Styled
	red = termenv.String().Foreground(p.Color("1")).Styled
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goport",
	Short: cyan("A simple and Fast port scanning tool written in GoLang"),


	Run: func(cmd *cobra.Command, args []string) { 

		cmd.Help()
	},
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func customHelpFunc(cmd *cobra.Command, args []string) {
	tile := figure.NewFigure("GoPort", "", true)

	fmt.Println(green(tile.String()))
	fmt.Println(green("A simple and Fast port scanning tool written in GoLang"))
	
	fmt.Println("Usage:")
	fmt.Println(green("  goport [flags]"))
	fmt.Println(green("  goport [command]"))
	fmt.Println()
	
	fmt.Println("Available Commands:")
	fmt.Println(yellow("  help        Help about any command"))
	fmt.Println(yellow("  ping        Used for ping scans"))
	fmt.Println()
	fmt.Println("Flags:")
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		fmt.Println(termenv.String(fmt.Sprintf(green("  -%s, --%s\t%s"), flag.Shorthand, flag.Name, flag.Usage)))
	})
	fmt.Println()
	fmt.Println(termenv.String(green(`Use "goport [command] --help" for more information about a command.`)))
}

func init() {

	rootCmd.SetHelpFunc(customHelpFunc)


}




