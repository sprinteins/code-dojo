package statement

import (
	"fmt"
	"math"
)

// Statement _
func Statement(invoice Invoice, plays PlayMap) string {
	var totalAmount = 0
	var volumeCredits = 0.0
	var result = "Statement for " + invoice.Customer + "\n"
	var format = func(number float64) string {
		return fmt.Sprintf("$%.2f", number)
	}

	for _, perf := range invoice.Performances {
		var play, _ = plays[perf.PlayID]
		var thisAmount = 0

		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (perf.Audience - 30)
			}
			break
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*(perf.Audience-20)
			}
			thisAmount += 300 * perf.Audience
			break
		default:
			panic("unknown type " + play.Type)
		}

		volumeCredits += math.Max(float64(perf.Audience)-30.0, 0.0)
		if "comedy" == play.Type {
			volumeCredits += math.Floor(float64(perf.Audience) / 5.0)
		}

		result += "    " + play.Name + ": " + format(float64(thisAmount)/100.0) + " (" + fmt.Sprintf("%d", perf.Audience) + " seats)\n"
		totalAmount += thisAmount
	}
	result += "  Amount owed is " + format(float64(totalAmount)/100.0) + "\n"
	result += "  You earned " + fmt.Sprintf("%d", int(volumeCredits)) + " credits"

	return result
}

// Invoice _
type Invoice struct {
	Customer     string
	Performances []Performance
}

// Performance _
type Performance struct {
	PlayID   string
	Audience int
}

// Plays defines different plays
func Plays() PlayMap {
	plays := make(PlayMap)

	plays["hamlet"] = Play{
		Name: "Hamlet",
		Type: "tragedy",
	}

	plays["as-like"] = Play{
		Name: "As You Like It",
		Type: "comedy",
	}

	plays["othello"] = Play{
		Name: "Othello",
		Type: "tragedy",
	}

	return plays
}

// PlayMap _
type PlayMap = map[string]Play

// Play _
type Play struct {
	Name string
	Type string
}
