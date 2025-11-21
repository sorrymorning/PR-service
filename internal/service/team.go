package service

import (
	"PR-service/internal/entity"
	"PR-service/internal/repository"
	"PR-service/internal/repository/repoerrs"
	"context"
)

type TeamService struct {
	teamRepo repo.Team
}

func NewTeamService(teamRepo repo.Team) *TeamService {
	return &TeamService{teamRepo: teamRepo}
}

func (s *TeamService) CreateTeam(ctx context.Context) (int, error) {
	id, err := s.teamRepo.CreateTeam(ctx)
	if err != nil {
		if err == repoerrs.ErrAlreadyExists {
			return 0, ErrAccountAlreadyExists
		}
		return 0, ErrCannotCreateAccount
	}

	return id, nil
}