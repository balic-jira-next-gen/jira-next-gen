package store

import (
	"context"

	"github.com/concept/jira_next_gen/backend/internal/models"
)

type Store interface {
	CreateProject(ctx context.Context, project *models.Project) error
	GetProject(ctx context.Context, id int) (*models.Project, error)
	ListProjects(ctx context.Context) ([]*models.Project, error)

	CreateIssue(ctx context.Context, issue *models.Issue) error
	GetIssue(ctx context.Context, id int) (*models.Issue, error)
	UpdateIssueStatus(ctx context.Context, id int, status string) error
	ListIssues(ctx context.Context, projectID int) ([]*models.Issue, error)

	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id int) (*models.User, error)
}
