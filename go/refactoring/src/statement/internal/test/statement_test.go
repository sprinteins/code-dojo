package statement_test

import (
	"testing"

	"bitbucket.sprinteins.com/users/trusz/repos/code-dojo/go/refactoring/src/statement"
	th "bitbucket.sprinteins.com/users/trusz/repos/code-dojo/go/refactoring/src/statement/internal/test/testhelper"
)

func TestStatement(t *testing.T) {

	var tests = []statementTest{
		{
			desc: "simple test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) { testStatement(t, tt) })
	}

}

func testStatement(t *testing.T, tt statementTest) {
	invoices := th.InvoicesWithAudiences(55, 35, 40)
	plays := statement.Plays()
	actualOutput := statement.Statement(invoices[0], plays)

	expectedOutput := th.GenerateExpectedOutput(th.TemplateInput{
		HamletPrice:    "$650.00",
		HamletSeats:    "55",
		AsYouLikePrice: "$580.00",
		AsYouLikeSeats: "35",
		OthelloPrice:   "$500.00",
		OthelloSeats:   "40",
		AmountOwned:    "$1730.00",
		CreditsEarned:  "47",
	})

	th.Equals(t, actualOutput, expectedOutput)

}

type statementTest struct {
	desc string
}
