package projects_test

import (
	"encoding/json"
	"testing"
	
	"sprinteins.com/heimat/test/src/x/assert"
	"sprinteins.com/heimat/test/src/x/http"
	"sprinteins.com/heimat/test/src/x/log"
)

func Test_Tier2_Add(t *testing.T) {
	p := project{
		Name: "Heimat",
	}

	projectsBefore := fetchProjects()

	statusCode := Create(p)
	assert.Equals(t, statusCode, 200, "request was not a success")

	projectsAfter := fetchProjects()
	assert.Equals(t, len(projectsAfter), len(projectsBefore)+1, "project has not been added")

}

func Create(p project) int {
	payload, err := json.Marshal(p)
	if err != nil {
		log.Error.Printf("could not create employee: %s", err)
	}
	res, err := http.Post("http://localhost:5060/api/projects", nil, payload)
	if err != nil {
		log.Error.Println(err)
	}
	res.Body.Close()
	return res.StatusCode
}

