package employees_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"sprinteins.com/heimat/test/src/x/http"
	"sprinteins.com/heimat/test/src/x/log"
)

func Test_Tier1_Employees(t *testing.T) {
	fetchEmployees()
}

func fetchEmployees() []employee {
	res, err := http.Get("http://localhost:5040/api/employees", nil)
	if err != nil {
		log.Error.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Error.Println(err)
	}

	emps := make([]employee, 0)
	if len(body) == 0 {
		return emps
	}
	err = json.Unmarshal(body, &emps)
	if err != nil {
		log.Error.Printf("could not unmarshal: %s\n\n%s", err, body)
	}
	return emps
}
