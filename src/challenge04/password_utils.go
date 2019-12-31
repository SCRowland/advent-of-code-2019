package main

import "strconv"

func isSixDigitNumber(num int) bool {
	strNum := strconv.Itoa(num)
	if len(strNum) == 6 {
		return true
	}
	return false
}

func hasRepeatedDigits(num int) bool {
	strNum := strconv.Itoa(num)

	hundredThousands := strNum[0]
	tenThousands := strNum[1]
	thousands := strNum[2]
	hundreds := strNum[3]
	tens := strNum[4]
	units := strNum[5]

	if hundredThousands == tenThousands ||
		tenThousands == thousands ||
		thousands == hundreds ||
		hundreds == tens ||
		tens == units {
		return true
	}
	return false
}

func hasAPureDoubleRepeatedDigit(num int) bool {
	strNum := strconv.Itoa(num)

	// XXnnnn
	if strNum[0] == strNum[1] && strNum[1] != strNum[2] {
		return true
	}

	// nXXnnn
	if strNum[1] == strNum[2] && strNum[0] != strNum[1] && strNum[3] != strNum[2] {
		return true
	}

	// nnXXnn
	if strNum[2] == strNum[3] && strNum[1] != strNum[2] && strNum[4] != strNum[3] {
		return true
	}

	// nnnXXn
	if strNum[3] == strNum[4] && strNum[2] != strNum[3] && strNum[5] != strNum[4] {
		return true
	}

	// nnnnXX
	if strNum[4] == strNum[5] && strNum[3] != strNum[4] {
		return true
	}

	return false
}

func hasIncreasingDigits(num int) bool {
	strNum := strconv.Itoa(num)

	hundredThousands, _ := strconv.Atoi(string(strNum[0]))
	tenThousands, _ := strconv.Atoi(string(strNum[1]))
	thousands, _ := strconv.Atoi(string(strNum[2]))
	hundreds, _ := strconv.Atoi(string(strNum[3]))
	tens, _ := strconv.Atoi(string(strNum[4]))
	units, _ := strconv.Atoi(string(strNum[5]))

	if units >= tens &&
		tens >= hundreds &&
		hundreds >= thousands &&
		thousands >= tenThousands &&
		tenThousands >= hundredThousands {
		return true
	}

	return false
}

// NumberPotentialPasswords counts potential problems between start and end
func NumberPotentialPasswords(start, end int) int {
	legitimatePasswordCount := 0

	for i := start; i <= end; i++ {
		if isSixDigitNumber(i) &&
			hasIncreasingDigits(i) &&
			hasRepeatedDigits(i) &&
			hasAPureDoubleRepeatedDigit(i) {
			legitimatePasswordCount++
		}
	}

	return legitimatePasswordCount
}
