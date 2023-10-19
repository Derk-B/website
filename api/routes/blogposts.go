package routes

import (
	"api/db"
	"context"
	"encoding/json"

	"github.com/uptrace/bun"
)

func ReturnBlogPosts(ctx context.Context, dbConn *bun.DB) string {
	projects := db.GetBlogPosts(ctx, dbConn)

	data, err := json.Marshal(projects)

	if err != nil {
		panic(err)
	}

	return string(data)
}
