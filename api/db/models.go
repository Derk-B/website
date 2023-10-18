package db

import "github.com/uptrace/bun"

type Project struct {
	bun.BaseModel `bun:"table:products`

	ID          int64 `bun:",pk,autoincrement"`
	Title       string
	BannerUrl   string
	Description string
	Content     string
	Timestamp   int
}

type BlogPost struct {
	bun.BaseModel `bun:"table:blogposts`

	ID          int64 `bun:",pk,autoincrement"`
	Title       string
	BannerUrl   string
	Description string
	Content     string
	Timestamp   int
}
