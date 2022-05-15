package clock

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

/*
Let's try to think on how to solve by using the examples
below.

**Example 1**
totalMinutes = -100

-60 minutes = -1h
-40 minutes = -40min

00:00 = 1440
23:00 = 23*60 = 1380
22:20 = 1340 min


**Example 2**
totalMinutes = 70

00:00 = 1440
01:10 = 1510

00:00 -> 00:00
0min     1440min

How to convert from minutes to hours and minutes?
Take 1510 % 1440 = 70
Then, divide 70 / 60 to get hours
and the remainder are the minutes.

When the totalMinutes is negative?
Then, we fix the negative value by adding 1440
and modding it again.

*/
func New(h, m int) Clock {
	totalMinutes := h*60 + m
	clockBased := (totalMinutes%1440 + 1440) % 1440

	hours := clockBased / 60
	minutes := clockBased - hours*60

	return Clock{hours, minutes}
}

func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
