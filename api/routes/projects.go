package routes

import (
	"api/db"
	"context"
	"encoding/json"

	"github.com/uptrace/bun"
)

type ProjectPageDTO struct {
	ID        int
	Title     string
	Content   string
	Timestamp int
	BannerURL string
}

type ProjectCardDTO struct {
	ID          int
	Title       string
	Description string
	Timestamp   int
	BannerURL   string
}

func ReturnProjects(ctx context.Context, dbConn *bun.DB) string {
	projects := db.GetProjects(ctx, dbConn)

	projectCardDTOs := []ProjectCardDTO{}
	for _, project := range projects {
		projectCardDTOs = append(projectCardDTOs,
			ProjectCardDTO{
				ID:          int(project.ID),
				Description: project.Description,
				BannerURL:   project.BannerUrl,
				Timestamp:   project.Timestamp,
				Title:       project.Title,
			},
		)
	}

	data, err := json.Marshal(projectCardDTOs)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func ReturnProject(id int, ctx context.Context, dbConn *bun.DB) string {
	project := db.GetProject(id, ctx, dbConn)

	projectPageDTO := ProjectPageDTO{
		int(project.ID),
		project.Title,
		project.Content,
		project.Timestamp,
		project.BannerUrl,
	}

	data, err := json.Marshal(projectPageDTO)

	if err != nil {
		panic(err)
	}

	return string(data)
}
