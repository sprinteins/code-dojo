package employees_test

import (
	"testing"

	"sprinteins.com/heimat/test/src/x/assert"
	"sprinteins.com/heimat/test/src/x/http"
	"sprinteins.com/heimat/test/src/x/log"
)

func Test_Tier3_Clear(t *testing.T) {
	e := employee{
		FirstName: "emp",
		LastName:  "loyee",
		Username:  "bob",
	}

	Create(e)
	statusCode := Clear()
	assert.Equals(t, statusCode, 200, "clear have failed")

	emps := fetchEmployees()
	assert.Equals(t, len(emps), 0, "clear has not removed employees")

}

func Clear() int {
	res, err := http.Delete("http://localhost:5040/api/employees")
	if err != nil {
		log.Error.Println(err)
	}
	res.Body.Close()
	return res.StatusCode
}
