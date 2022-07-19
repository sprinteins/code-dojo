package testhelper

import (
	"bytes"
	"html/template"
	"math/rand"

	"bitbucket.sprinteins.com/users/trusz/repos/code-dojo/go/refactoring/src/statement"
)

// InvoicesWithRandomNoOfAudience _
func InvoicesWithRandomNoOfAudience() []statement.Invoice {

	return InvoicesWithRandomNoOfAudienceMax(100)

}

// InvoicesWithRandomNoOfAudienceMax _
func InvoicesWithRandomNoOfAudienceMax(maxAudience int) []statement.Invoice {

	hamletAudience := rand.Intn(maxAudience)
	asLikeAudience := rand.Intn(maxAudience)
	othelloAudience := rand.Intn(maxAudience)

	return InvoicesWithAudiences(hamletAudience, asLikeAudience, othelloAudience)

}

// InvoicesWithAudiences _
func InvoicesWithAudiences(hamletAudience int, asLikeAudience int, othelloAudience int) []statement.Invoice {
	invoices := make([]statement.Invoice, 1)

	performances := make([]statement.Performance, 3)
	performances[0] = statement.Performance{
		PlayID:   "hamlet",
		Audience: hamletAudience,
	}
	performances[1] = statement.Performance{
		PlayID:   "as-like",
		Audience: asLikeAudience,
	}
	performances[2] = statement.Performance{
		PlayID:   "othello",
		Audience: othelloAudience,
	}

	invoices[0] = statement.Invoice{
		Customer:     "BigCo",
		Performances: performances,
	}

	return invoices
}

// GenerateExpectedOutput _
func GenerateExpectedOutput(tempIn TemplateInput) string {
	tmpl, err := template.New("test").Parse(outputTemplate)
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, tempIn)
	if err != nil {
		panic(err)
	}
	return tpl.String()
}

// OutputTemplate _
var outputTemplate = `Statement for BigCo
    Hamlet: {{.HamletPrice}} ({{.HamletSeats}} seats)
    As You Like It: {{.AsYouLikePrice}} ({{.AsYouLikeSeats}} seats)
    Othello: {{.OthelloPrice}} ({{.OthelloSeats}} seats)
  Amount owed is {{.AmountOwned}}
  You earned {{.CreditsEarned}} credits`

// TemplateInput _
type TemplateInput struct {
	HamletPrice    string
	HamletSeats    string
	AsYouLikePrice string
	AsYouLikeSeats string
	OthelloPrice   string
	OthelloSeats   string
	AmountOwned    string
	CreditsEarned  string
}
