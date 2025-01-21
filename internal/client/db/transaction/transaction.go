package transaction

import (
	"github.com/biryanim/auth/internal/client/db"
	"github.com/biryanim/auth/internal/client/db/pg"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type manager struct {
	db db.Transactor
}

func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{db: db}
}

func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Hadler) (err error) {
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	ctx = pg.MakeContextTx(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("recovered from panic: %v", r)
		}

		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrap(errRollback, "failed to rollback transaction")
			}
			return
		}

		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "failed to commit transaction")
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed to executing code inside transaction")
	}

	return err
}

func (m *manager) ReadCommited(ctx context.Context, f db.Hadler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
