package employees_test

import (
	"testing"

	"sprinteins.com/heimat/test/src/x/assert"
)

func Test_Tier5_Is_Active(t *testing.T) {
	t.Skip()
	e := employee{
		FirstName: "emp",
		LastName:  "loyee",
		Username:  "bob",
	}

	statusCode := Create(e)
	assert.Equals(t, statusCode, 200, "request was not a success")

	emp := fetchByUsername(e.Username)
	assert.Equals(t, e, emp, "fetched employee is not the same as created")

}
