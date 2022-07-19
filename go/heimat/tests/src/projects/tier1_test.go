package projects_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"sprinteins.com/heimat/test/src/x/http"
	"sprinteins.com/heimat/test/src/x/log"
)

func Test_Tier1_Projects(t *testing.T) {
	fetchProjects()
}

func fetchProjects() []project {
	res, err := http.Get("http://localhost:5060/api/projects", nil)
	if err != nil {
		log.Error.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Error.Println(err)
	}

	projects := make([]project, 0)
	if len(body) == 0 {
		return projects
	}
	err = json.Unmarshal(body, &projects)
	if err != nil {
		log.Error.Printf("could not unmarshal: %s\n\n%s", err, body)
	}
	return projects
}
