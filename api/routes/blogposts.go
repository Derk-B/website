package routes

import (
	"api/db"
	"encoding/json"
)

func ReturnBlogPosts() string {
	projects := db.GetBlogPosts()

	data, err := json.Marshal(projects)

	if err != nil {
		panic(err)
	}

	return string(data)
}
