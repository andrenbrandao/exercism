package twelve

import (
	"fmt"
	"strings"
)

var gifts = []string{
	"twelve Drummers Drumming",
	"eleven Pipers Piping",
	"ten Lords-a-Leaping",
	"nine Ladies Dancing",
	"eight Maids-a-Milking",
	"seven Swans-a-Swimming",
	"six Geese-a-Laying",
	"five Gold Rings",
	"four Calling Birds",
	"three French Hens",
	"two Turtle Doves",
	"a Partridge in a Pear Tree",
}

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

func Verse(day int) string {
	var chosenGifts string
	if day == 1 {
		chosenGifts = gifts[len(days)-1]
	} else {
		chosenGifts = strings.Join(gifts[len(days)-day:len(days)-1], ", ")
		chosenGifts = fmt.Sprintf("%s, and %s", chosenGifts, gifts[len(days)-1])
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", days[day-1], chosenGifts)

}

func Song() string {
	numVerses := len(days)
	verses := make([]string, 0, numVerses)

	for i := 1; i <= numVerses; i++ {
		verses = append(verses, Verse(i))
	}

	return strings.Join(verses, "\n")
}
