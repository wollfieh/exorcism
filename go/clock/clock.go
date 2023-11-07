package clock

import "fmt"

const (
	dayminutes = 60 * 24
)

type Clock int

// func shake shakes the clock to rearrange hours and minutes
// each day (dayminutes long) is the same on the timeline
// if we're on the negative side we can slide one to the right
func shake(c int) Clock {
	if c < 0 {
		// slide one day to the right
		c = (c % dayminutes) + dayminutes
	} else {
		c %= dayminutes
	}

	return Clock(c)
}
func New(h, m int) Clock {
	return shake(h*60 + m)

}

func (c Clock) Add(m int) Clock {
	return shake(int(c) + m)
}

func (c Clock) Subtract(m int) Clock {
	return shake(int(c) - m)
}

func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", (c / 60), c%60)

}
