package routes

import (
	"api/db"
	"encoding/json"
)

func ReturnProjects() string {
	projects := db.GetProjects()

	data, err := json.Marshal(projects)

	if err != nil {
		panic(err)
	}

	return string(data)
}
