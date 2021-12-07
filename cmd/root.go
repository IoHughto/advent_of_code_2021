package cmd

import (
	day01 "advent_of_code_2021/days/01"
	day02 "advent_of_code_2021/days/02"
	day03 "advent_of_code_2021/days/03"
	day04 "advent_of_code_2021/days/04"
	day05 "advent_of_code_2021/days/05"
	day06 "advent_of_code_2021/days/06"
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

	fmt.Printf("Day %02d\n", viper.GetInt("day"))
	switch viper.GetInt("day") {
	case 1:
		err = day01.Run()
	case 2:
		err = day02.Run()
	case 3:
		err = day03.Run()
	case 4:
		err = day04.Run()
	case 5:
		err = day05.Run()
	case 6:
		err = day06.Run()
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
