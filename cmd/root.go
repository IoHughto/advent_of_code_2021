package cmd

import (
	day01 "advent_of_code_2021/days/01"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var rootCmd = &cobra.Command{
	Use:     "aoc21",
	Short:   "Runs days for AOC2021",
	Long:    "CLI for running individual days in Advent of Code 2021",
	PreRunE: Prepare,
	RunE:    Run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	viper.AutomaticEnv()

	rootCmd.Flags().IntP("day", "d", 0, "Day of the challenge")
}

func Run(_ *cobra.Command, _ []string) error {
	var err error

	switch viper.GetInt("day") {
	case 1:
		err = day01.Run()
	default:
		err = fmt.Errorf("%w: %d", errUnexpectedDay, viper.GetInt("day"))
	}

	return err
}

func Prepare(cmd *cobra.Command, _ []string) error {
	return viper.BindPFlags(cmd.Flags())
}

// Sentinel errors
var errUnexpectedDay = fmt.Errorf("unexpected day")
