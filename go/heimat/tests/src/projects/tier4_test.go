package projects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"sprinteins.com/heimat/test/src/x/assert"
	"sprinteins.com/heimat/test/src/x/http"
	"sprinteins.com/heimat/test/src/x/log"
)

func Test_Tier4_Project_By_Name(t *testing.T) {

	projToCreate := project{
		Name: "Heimat",
	}

	Clear()
	statusCode := Create(projToCreate)
	assert.Equals(t, statusCode, 200, "request was not a success")

	projFetched := fetchByName(projToCreate.Name)
	assert.Equals(t, projFetched, projToCreate, "fetched project is not the same as created")

}

func Test_Tier4_Employee_By_Username_Not_Found(t *testing.T) {

	Clear()
	statusCode := fetchByNonExistingName("xXx")
	assert.Equals(t, statusCode, 404, "did not return 404 for non existing username")
}

func fetchByName(name string) project {
	res, err := http.Get("http://localhost:5060/api/projects/"+name, nil)
	if err != nil {
		log.Error.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Error.Println(err)
	}

	project := &project{}
	err = json.Unmarshal(body, project)
	if err != nil {
		log.Error.Printf("could not unmarshal: %s\n\n%s", err, body)
	}

	return *project
}

func fetchByNonExistingName(name string) int {
	res, err := http.Get("http://localhost:5060/api/projects/"+name, nil)
	if err != nil {
		log.Error.Println(err)
	}
	res.Body.Close()

	return res.StatusCode
}
