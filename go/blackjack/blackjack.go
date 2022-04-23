package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerHandValue := ParseCard(card1) + ParseCard(card2)
	dealerHandValue := ParseCard(dealerCard)

	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	if playerHandValue == 21 {
		if dealerHandValue >= 10 {
			return "S"
		}
		return "W"
	}

	if playerHandValue >= 17 && playerHandValue <= 20 {
		return "S"
	}

	if playerHandValue >= 12 && playerHandValue <= 16 {
		if dealerHandValue >= 7 {
			return "H"
		}
		return "S"
	}

	return "H"
}
