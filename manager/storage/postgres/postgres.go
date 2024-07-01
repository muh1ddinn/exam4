package postgres

import (
	"context"
	"crmservice/config"
	"crmservice/storage"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	Pool *pgxpool.Pool
	cfg  config.Config
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	pgxPoolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	pgxPoolConfig.MaxConns = 50
	pgxPoolConfig.MaxConnLifetime = time.Hour

	newPool, err := pgxpool.NewWithConfig(ctx, pgxPoolConfig)
	if err != nil {
		return nil, err
	}

	return Store{
		Pool: newPool,
		cfg:  cfg,
	}, nil
}
func (s Store) CloseDB() {
	s.Pool.Close()
}

func (s Store) Manager_student() storage.Manager_crmRepoRepoI {
	newstudnet := Newmanager_crmRepo(s.Pool)

	return newstudnet
}

func (s Store) Managerteacher() storage.Managerteacher_RepoRepoI {
	newteacher := Newmanagerteacher(s.Pool)

	return newteacher
}

func (s Store) Managersupportteacher() storage.Managersupportteacher_RepoRepoI {
	newsupportteacher := Newmanagersupportteacher(s.Pool)

	return newsupportteacher
}

func (s Store) Manageradministratio() storage.ManageradministrationRepo_RepoRepoI {
	newadministration := Newmanageradministration(s.Pool)

	return newadministration
}

func (s Store) Managergroup() storage.ManagerGroupService_RepoRepoI {
	newgroup := Newmanagergroup(s.Pool)

	return newgroup
}
