// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).
// Code generated by sqlbrick. DO NOT EDIT IT.

package models

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

// Type definition for SqlBrick. It contains all bricks depends on the number of
// sql files. It also wraps some sqlx func for for convenience.
type SqlBrick struct {
	Db   *sqlx.DB
	Book *BookBrick
	User *UserBrick
}

// Type definition for brick transaction. This aims at sql transaction.
// If you want a transaction, then invoke 'Begin' to get this struct.
type BrickTx struct {
	tx   *sqlx.Tx
	Book *BookBrickTx
	User *UserBrickTx
}

// NewSqlBrick create a new SqlBrick to operate all bricks.
func NewSqlBrick(db *sqlx.DB) *SqlBrick {
	return &SqlBrick{
		Db:   db,
		Book: newBookBrick(db),
		User: newUserBrick(db),
	}
}

// Begin will start a new transaction for bricks. If any query
// is defined as tx sql, this must be invoked.
func (b *SqlBrick) Begin() (*BrickTx, error) {
	tx, err := b.Db.Beginx()
	if err != nil {
		return nil, err
	}

	return &BrickTx{
		tx:   tx,
		Book: b.Book.newBookTx(tx),
		User: b.User.newUserTx(tx),
	}, nil
}

// Commit will end a transaction for brick. Begin must be invoked
// before Commit. Otherwise there will be an error.
func (b *BrickTx) Commit() error {
	if b.tx == nil {
		return errors.New("the Begin func must be invoked first")
	}

	return b.tx.Commit()
}