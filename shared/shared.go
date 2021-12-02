package shared

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func ReadData() ([]string, error) {
	fileName := fmt.Sprintf("data/day%02d.txt", viper.GetInt("day"))
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
