package entity

import "time"

type PullRequest struct {
    PullRequestID   string     `db:"pull_request_id"`
    PullRequestName string     `db:"pull_request_name"`
    AuthorID        string     `db:"author_id"`
    Status          string     `db:"status"` // OPEN | MERGED
    CreatedAt       time.Time  `db:"created_at"`
    MergedAt        *time.Time `db:"merged_at"` // null
}