package main

import "testing"

func TestNumberPotentialPasswords(t *testing.T) {
	start := 356261
	end := 846303
	expected := 334

	got := NumberPotentialPasswords(start, end)
	if got != expected {
		t.Errorf("NumberPotentialPasswords(%d, %d) = %d, not %d", start, end, got, expected)
	}
}

var numbers = []struct {
	number     int
	isSixDigit bool
}{
	{
		1,
		false,
	},
	{
		12,
		false,
	},
	{
		123,
		false,
	},
	{
		1234,
		false,
	},
	{
		12345,
		false,
	},
	{
		123456,
		true,
	},
	{
		1234567,
		false,
	},
	{
		12345678,
		false,
	},
	{
		1234567890,
		false,
	},
}

func TestIsSixDigitNumber(t *testing.T) {
	for _, tt := range numbers {
		got := isSixDigitNumber(tt.number)
		if got != tt.isSixDigit {
			t.Errorf("isSixDigitNumber(%d) = %t, should be %t", tt.number, got, tt.isSixDigit)
		}
	}
}

var RepeatedDigitNumbers = []struct {
	number            int
	hasRepeatedDigits bool
}{
	{
		112345,
		true,
	},
	{
		123456,
		false,
	},
	{
		123455,
		true,
	},
}

func TestHasRepeatedDigits(t *testing.T) {
	for _, tt := range RepeatedDigitNumbers {
		got := hasRepeatedDigits(tt.number)
		if got != tt.hasRepeatedDigits {
			t.Errorf("hasRepeatedDigits(%d) = %t, should be %t", tt.number, got, tt.hasRepeatedDigits)
		}
	}
}

var increasingDigitNumbers = []struct {
	number            int
	hasRepeatedDigits bool
}{
	{
		123456,
		true,
	},
	{
		123455,
		true,
	},
	{
		111111,
		true,
	},
	{
		112233,
		true,
	},
	{
		123454,
		false,
	},
	{
		121456,
		false,
	},
	{
		103456,
		false,
	},
}

func TestHasIncreasingDigits(t *testing.T) {
	for _, tt := range increasingDigitNumbers {
		got := hasIncreasingDigits(tt.number)
		if got != tt.hasRepeatedDigits {
			t.Errorf("hasIncreasingDigits(%d) = %t, should be %t", tt.number, got, tt.hasRepeatedDigits)
		}
	}
}

var hasOnlyDoulbeRepeatedDigitNumbers = []struct {
	number                      int
	hasAPureDoubleRepeatedDigit bool
}{
	{
		112233,
		true,
	},
	{
		123444,
		false,
	},
	{
		111122,
		true,
	},
	{
		222345,
		false,
	},
}

func TestHasAPureDoubleRepeatedDigit(t *testing.T) {
	for _, tt := range hasOnlyDoulbeRepeatedDigitNumbers {
		got := hasAPureDoubleRepeatedDigit(tt.number)
		if got != tt.hasAPureDoubleRepeatedDigit {
			t.Errorf("hasAPureDoubleRepeatedDigit(%d) = %t, should be %t", tt.number, got, tt.hasAPureDoubleRepeatedDigit)
		}
	}
}
