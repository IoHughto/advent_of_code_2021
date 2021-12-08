package day07

import (
	"advent_of_code_2021/shared"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		data        []string
		expectedErr error
	}{
		{[]string{"1,2,3"}, nil},
		{[]string{}, errUnexpectedData},
		{[]string{"a", "b"}, errUnexpectedData},
		{[]string{"a"}, errCantParseInt},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test-%d", i), func(t *testing.T) {
			err := Run(test.data)
			//if err != test.expectedErr {
			if !errors.Is(err, test.expectedErr) {
				t.Errorf("got %s, expected %s", err, test.expectedErr)
			}
		})
	}
}

func TestConvertToInts(t *testing.T) {
	tests := []struct {
		value       string
		expected    []int
		expectedErr error
	}{
		{"1,2,3", []int{1, 2, 3}, nil},
		{"16,1,2,0,4,2,7,1,2,14", []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, nil},
		{"asdf", nil, errCantParseInt},
	}

	for _, test := range tests {
		t.Run(test.value, func(t *testing.T) {
			ans, err := convertToInts(test.value)
			//if err != test.expectedErr {
			if !errors.Is(err, test.expectedErr) {
				t.Errorf("got %s, expected %s", err, test.expectedErr)
			}
			if !reflect.DeepEqual(ans, test.expected) {
				t.Errorf("got %v, expected %v", ans, test.expected)
			}
		})
	}
}

func TestFindMinFuel(t *testing.T) {
	rawData, err := shared.ReadTestData(7)
	if err != nil {
		t.Errorf("Error parsing test data")
	}
	data, err := convertToInts(rawData[0])
	if err != nil {
		t.Errorf("Couldn't convert test data")
	}
	tests := []struct {
		fuelFunction func(int) int
		expected     int
	}{
		{simple, 37},
		{cumulative, 168},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("Test-%d", i), func(t *testing.T) {
			ans := findMinFuel(data, test.fuelFunction)
			if ans != test.expected {
				t.Errorf("got %d, expected %d", ans, test.expected)
			}
		})
	}
}
