package challenge10

import "testing"

var angleTestData = []struct {
	from  Point
	to    Point
	angle Angle
}{
	{
		Point{1, 1},
		Point{1, 2},
		90.0,
	},
	{
		Point{1, 1},
		Point{2, 1},
		360.0,
	},
	{
		Point{1, 1},
		Point{1, 0},
		270.0,
	},
	{
		Point{1, 1},
		Point{0, 1},
		180.0,
	},
}

func TestAngle(t *testing.T) {
	for _, testData := range angleTestData {
		angle := testData.from.Angle(&testData.to)
		if !AngleEquals(angle, testData.angle) {
			t.Errorf(
				"Angle from %v to %v is %f not %f",
				testData.from, testData.to, angle, testData.angle,
			)
		}
	}
}

var angleEqualityTestData = []struct {
	a1       Angle
	a2       Angle
	expected bool
}{
	{
		Angle(1.5707965),
		Angle(1.5707964),
		true,
	},
	{
		Angle(1.570796),
		Angle(1.570796),
		true,
	},
	{
		Angle(1.570795),
		Angle(1.570796),
		false,
	},
	{
		Angle(1.570786),
		Angle(1.570796),
		false,
	},
}

func TestAngleEquals(t *testing.T) {
	for _, td := range angleEqualityTestData {
		got := AngleEquals(td.a1, td.a2)
		if got != td.expected {
			t.Errorf("AngleEquals(%v, %v) = %t not %t", td.a1, td.a2, got, td.expected)
		}
	}
}
