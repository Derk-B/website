package routes

import (
	"api/db"
	"context"
	"encoding/json"

	"github.com/uptrace/bun"
)

type BlogPageDTO struct {
	ID        int
	Title     string
	Content   string
	Timestamp int
	BannerUrl string
	BannerAlt string
}

type BlogCardDTO struct {
	ID          int
	Title       string
	Description string
	Timestamp   int
	BannerUrl   string
	BannerAlt   string
}

func ReturnBlogPosts(ctx context.Context, dbConn *bun.DB) string {
	projects := db.GetBlogPosts(ctx, dbConn)

	data, err := json.Marshal(projects)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func ReturnBlogPost(id int, ctx context.Context, dbConn *bun.DB) string {
	blog := db.GetBlogPost(id, ctx, dbConn)

	blogpostPageDTO := BlogPageDTO{
		int(blog.ID),
		blog.Title,
		blog.Content,
		blog.Timestamp,
		blog.BannerUrl,
		blog.BannerAlt,
	}

	data, err := json.Marshal(blogpostPageDTO)

	if err != nil {
		panic(err)
	}

	return string(data)
}
