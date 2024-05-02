package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
}

type sqlStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &sqlStore{
		db:      db,
		Queries: New(db),
	}

}

func (s *sqlStore) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err : %v, rb err : %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
