package projects_test

import (
	"testing"

	"sprinteins.com/heimat/test/src/x/assert"
	"sprinteins.com/heimat/test/src/x/http"
	"sprinteins.com/heimat/test/src/x/log"
)

func Test_Tier3_Clear(t *testing.T) {
	p := project{
		Name: "Heimat",
	}

	Create(p)
	statusCode := Clear()
	assert.Equals(t, statusCode, 200, "clear have failed")

	projects := fetchProjects()
	assert.Equals(t, len(projects), 0, "clear has not removed projects")

}

func Clear() int {
	res, err := http.Delete("http://localhost:5060/api/projects")
	if err != nil {
		log.Error.Println(err)
	}
	res.Body.Close()
	return res.StatusCode
}
