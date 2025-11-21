CREATE TABLE IF NOT EXISTS pull_request_reviewers (
    pull_request_id TEXT NOT NULL REFERENCES pull_requests(pull_request_id) ON DELETE CASCADE,
    reviewer_id TEXT NOT NULL REFERENCES users(user_id),
    PRIMARY KEY (pull_request_id, reviewer_id)
);
