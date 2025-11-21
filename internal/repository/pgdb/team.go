package pgdb

import (
	"PR-service/internal/entity"
	"PR-service/internal/repo/repoerrs"
	"PR-service/pkg/postgres"
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	log "github.com/sirupsen/logrus"
)


type TeamRepo struct {
	*postgres.Postgres
}

func NewTeamRepo(pg *postgres.Postgres) *AccountRepo {
	return &AccountRepo{pg}
}


func (r *TeamRepo) CreateTeam(ctx context.Context, teamName string) error {
    sql, args, _ := r.Builder.
        Insert("teams").
        Columns("team_name").          
        Values(teamName).              
        ToSql()                        

    _, err := r.Pool.Exec(ctx, sql, args...)
    if err != nil {
        log.Debugf("err: %v", err)
        var pgErr *pgconn.PgError
        if ok := errors.As(err, &pgErr); ok {
            if pgErr.Code == "23505" {  
                return repoerrs.ErrAlreadyExists  
            }
        }
        return fmt.Errorf("TeamRepo.CreateTeam - r.Pool.Exec: %v", err)
    }

    return nil  
}