package employees_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"sprinteins.com/heimat/test/src/x/assert"
	"sprinteins.com/heimat/test/src/x/http"
	"sprinteins.com/heimat/test/src/x/log"
)

func Test_Tier4_Employee_By_Username(t *testing.T) {

	e := employee{
		FirstName: "emp",
		LastName:  "loyee",
		Username:  "bob",
	}

	Clear()
	statusCode := Create(e)
	assert.Equals(t, statusCode, 200, "request was not a success")

	emp := fetchByUsername(e.Username)
	assert.Equals(t, e, emp, "fetched employee is not the same as created")

}

func Test_Tier4_Employee_By_Username_Not_Found(t *testing.T) {

	Clear()
	statusCode := fetchByNonExistingUsername("xXx")
	assert.Equals(t, statusCode, 404, "did not return 404 for non existing username")
}

func fetchByUsername(username string) employee {
	res, err := http.Get("http://localhost:5040/api/employees/"+username, nil)
	if err != nil {
		log.Error.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Error.Println(err)
	}

	emp := &employee{}
	err = json.Unmarshal(body, emp)
	if err != nil {
		log.Error.Printf("could not unmarshal: %s\n\n%s", err, body)
	}

	return *emp
}

func fetchByNonExistingUsername(username string) int {
	res, err := http.Get("http://localhost:5040/api/employees/"+username, nil)
	if err != nil {
		log.Error.Println(err)
	}
	res.Body.Close()

	return res.StatusCode
}
