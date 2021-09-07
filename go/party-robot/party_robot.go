package partyrobot

import (
	"fmt"
	"strings"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and stands out his age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour string, direction string, distance float64) string {
	welcome := Welcome(name)
	tableAssignment := fmt.Sprintf("You have been assigned to table %X. Your table is %v, exactly %.1f meters from here.", table, direction, distance)
	nextToNeighbour := fmt.Sprintf("You will be sitting next to %v.", neighbour)

	str := []string{welcome, tableAssignment, nextToNeighbour}
	return strings.Join(str, "\n")
}
