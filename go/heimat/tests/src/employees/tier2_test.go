package employees_test

import (
	"encoding/json"
	"testing"

	"sprinteins.com/heimat/test/src/x/assert"
	"sprinteins.com/heimat/test/src/x/http"
	"sprinteins.com/heimat/test/src/x/log"
)

func Test_Tier2_Add(t *testing.T) {
	e := employee{
		FirstName: "emp",
		LastName:  "loyee",
		Username:  "bob",
	}

	beforeEmps := fetchEmployees()

	statusCode := Create(e)
	assert.Equals(t, statusCode, 200, "request was not a success")

	afterEmps := fetchEmployees()
	assert.Equals(t, len(afterEmps), len(beforeEmps)+1, "employee has not been added")

}

func Create(e employee) int {
	payload, err := json.Marshal(e)
	if err != nil {
		log.Error.Printf("could not create employee: %s", err)
	}
	res, err := http.Post("http://localhost:5040/api/employees", nil, payload)
	if err != nil {
		log.Error.Println(err)
	}
	res.Body.Close()
	return res.StatusCode
}
