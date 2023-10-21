package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

// Use this to connect to the database for every query.
//
// Don't forget to call ctx.Done() when done querying.
//
// Returns a context and *bun.DB that you can use for queries on the database
func ConnectToDB() (context.Context, *bun.DB) {
	ctx := context.Background()

	godotenv.Load("../database/.env")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	address := os.Getenv("DB_IPADDRESS")
	port := os.Getenv("DB_PORT")

	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(address+":"+port),
		pgdriver.WithUser(username),
		pgdriver.WithDatabase(database),
		pgdriver.WithInsecure(true),
		pgdriver.WithPassword(password),
	)

	sqldb := sql.OpenDB(pgconn)

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())

	return ctx, db
}

// Use this to insert a new blogpost into the database
func InsertBlogPost(post BlogPost, ctx context.Context, db *bun.DB) {
	db.NewInsert().Model(&post).Exec(ctx)
}

// Use this to insert a new project into the database
func InsertProject(project Project, ctx context.Context, db *bun.DB) {
	_, err := db.NewInsert().Model(&project).Exec(ctx)

	if err != nil {
		fmt.Println(err)
	}
}

// Call this to get blogposts from the database.
//
// Returns a list of blogposts
func GetBlogPosts(ctx context.Context, db *bun.DB) []BlogPost {
	posts := []BlogPost{}

	db.NewSelect().Model(&BlogPost{}).Scan(ctx, &posts)

	return posts
}

// Call this to get projects from the database.
//
// Returns a list of projects
func GetProjects(ctx context.Context, db *bun.DB) []Project {
	projects := []Project{}

	db.NewSelect().Model(&Project{}).Scan(ctx, &projects)

	return projects
}

// Call this to get a project by ID from the databse
//
// Returns a project that has the specified ID
func GetBlogPost(id int, ctx context.Context, db *bun.DB) BlogPost {
	blogs := []BlogPost{}

	db.NewSelect().Model(&BlogPost{}).Where("ID = ?", id).Scan(ctx, &blogs)

	return blogs[0]
}
