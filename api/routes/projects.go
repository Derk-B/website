package routes

import (
	"api/db"
	"context"
	"encoding/json"

	"github.com/uptrace/bun"
)

func ReturnProjects(ctx context.Context, dbConn *bun.DB) string {
	projects := db.GetProjects(ctx, dbConn)

	data, err := json.Marshal(projects)

	if err != nil {
		panic(err)
	}

	return string(data)
}
