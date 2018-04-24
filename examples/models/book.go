// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).
// Code generated by sqlbrick. DO NOT EDIT IT.

// This file is generated from: book.sql

package models

import (
	"database/sql"
	"errors"
	"strings"

	"anbillon.com/sqlbrick/typex"
	"github.com/jmoiron/sqlx"
)

// Type definition for Book which defined in sql file.
// This can be used as a model in database operation.
type Book struct {
	Id         int32            `db:"id"`
	Uid        int32            `db:"uid"`
	Name       string           `db:"name"`
	Content    typex.NullString `db:"content"`
	CreateTime typex.NullTime   `db:"create_time"`
	Price      int              `db:"price"`
}

// Type definition for BookBrick. This brick will contains all database
// operation from given sql file. Each sql file will have only one brick.
type BookBrick struct {
	db *sqlx.DB
}

// Type definition for Book transaction. This aims at sql transaction.
type BookBrickTx struct {
	tx *sqlx.Tx
}

// newBookBrick will create a Book brick. This is used
// invoke the query function generated from sql file.
func newBookBrick(db *sqlx.DB) *BookBrick {
	return &BookBrick{db: db}
}

// newBookTx will create a new transaction for Book.
func (b *BookBrick) newBookTx(tx *sqlx.Tx) *BookBrickTx {
	return &BookBrickTx{tx: tx}
}

// checkTx will check if tx is available.
func (b *BookBrickTx) checkTx() error {
	if b.tx == nil {
		return errors.New("the Begin func must be invoked first")
	}
	return nil
}

// CreateBook generated by sqlbrick, used to operate database table.
func (b *BookBrick) CreateBook() sql.Result {
	return b.db.MustExec(`CREATE TABLE IF NOT EXISTS book (
  "id"  serial NOT NULL PRIMARY KEY,
  uid int4 NOT NULL,
  name text NOT NULL,
  content varchar(255),
  create_time TIMESTAMP,
  price int NOT NULL
)`)
}

// InsertOne generated by sqlbrick, insert data into database.
func (b *BookBrick) InsertOne(args *Book) (sql.Result, error) {
	stmt, err := b.db.PrepareNamed(
		`INSERT INTO book (uid, name, content, create_time, price)
  VALUES (:uid, :name, :content, :create_time, :price)`)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(args)
}

// UpdateSomeThing An example to show update.
func (b *BookBrick) UpdateSomeThing(args *Book) (int64, error) {
	conditionQuery := `UPDATE book SET `
	if args.Price > 0 {
		conditionQuery += `price = :price,`
	}
	if args.Content.String != "" {
		conditionQuery += `content = :content,`
	}
	conditionQuery += ` name = :name,`
	if args.CreateTime.Time.Unix() != 0 {
		conditionQuery += `create_time = :create_time`
	}
	if strings.HasSuffix(conditionQuery, ",") {
		strings.TrimSuffix(conditionQuery, ",")
	}
	conditionQuery += ` WHERE id = :id`

	stmt, err := b.db.PrepareNamed(conditionQuery)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// ComplexUpdate An example to show complex update. Second line comment.
func (b *BookBrick) ComplexUpdate(args *Book) (int64, error) {
	stmt, err := b.db.PrepareNamed(
		`UPDATE book SET price=(SELECT price FROM book, user WHERE book.uid=user.id)
  WHERE book.price <= :price AND name = :name`)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// SelectAll generated by sqlbrick, select data from database.
func (b *BookBrick) SelectAll(dest interface{}) error {
	return b.db.Select(dest, `SELECT * FROM book`)
}

// CountBooks generated by sqlbrick, select data from database.
func (b *BookBrick) CountBooks(dest interface{}, uid interface{}) error {
	stmt, err := b.db.PrepareNamed(
		`SELECT COUNT(*) FROM book WHERE uid = :uid`)
	if err != nil {
		return err
	}

	// create map arguments for sqlx
	args := map[string]interface{}{
		"uid": uid,
	}

	row := stmt.QueryRowx(args)
	if row.Err() != nil {
		return row.Err()
	}

	return row.Scan(dest)
}

// SelectById An example to show SelectById.
func (b *BookBrick) SelectById(dest interface{}, id interface{}) error {
	stmt, err := b.db.PrepareNamed(
		`SELECT * FROM book WHERE id = :id ORDER BY name ASC`)
	if err != nil {
		return err
	}

	// create map arguments for sqlx
	args := map[string]interface{}{
		"id": id,
	}

	row := stmt.QueryRowx(args)
	if row.Err() != nil {
		return row.Err()
	}

	return row.StructScan(dest)
}

// SelectByUid generated by sqlbrick, select data from database.
func (b *BookBrick) SelectByUid(dest interface{}, uid interface{}) error {
	stmt, err := b.db.PrepareNamed(
		`SELECT * FROM book WHERE uid = :uid ORDER BY name ASC`)
	if err != nil {
		return err
	}

	// create map arguments for sqlx
	args := map[string]interface{}{
		"uid": uid,
	}

	rows, err := stmt.Queryx(args)
	if err != nil {
		return err
	}

	return sqlx.StructScan(rows, dest)
}

// DeleteById An example to show DeleteById.
func (b *BookBrick) DeleteById(id interface{}) (int64, error) {
	stmt, err := b.db.PrepareNamed(`DELETE FROM book WHERE id = :id`)
	if err != nil {
		return 0, err
	}

	// create map arguments for sqlx
	args := map[string]interface{}{
		"id": id,
	}

	result, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// TxInsert generated by sqlbrick, insert data into database.
func (b *BookBrickTx) TxInsert(args *Book) (sql.Result, error) {
	if err := b.checkTx(); err != nil {
		return nil, err
	}

	stmt, err := b.tx.PrepareNamed(
		`INSERT INTO book (uid, name, content, create_time, price)
  VALUES (:uid, :name, :content, :create_time, :price)`)
	if err != nil {
		return nil, err
	}

	if result, err := stmt.Exec(args); err != nil {
		if rbe := b.tx.Rollback(); rbe != nil {
			return nil, rbe
		}
		return nil, err
	} else {
		return result, nil
	}
}

// TxSelect generated by sqlbrick, select data from database.
func (b *BookBrickTx) TxSelect(dest interface{}, uid interface{}) error {
	if err := b.checkTx(); err != nil {
		return err
	}

	stmt, err := b.tx.PrepareNamed(
		`SELECT COUNT(*) FROM book WHERE uid = :uid`)
	if err != nil {
		return err
	}

	// create map arguments for sqlx
	args := map[string]interface{}{
		"uid": uid,
	}

	row := stmt.QueryRowx(args)
	if row.Err() != nil {
		return row.Err()
	}

	return row.Scan(dest)
}

// TxDeleteById generated by sqlbrick, delete data from database.
// Affected rows will return if there's no error.
func (b *BookBrickTx) TxDeleteById(id interface{}) (int64, error) {
	if err := b.checkTx(); err != nil {
		return 0, err
	}

	stmt, err := b.tx.PrepareNamed(`DELETE FROM book WHERE id = :id`)
	if err != nil {
		return 0, err
	}

	// create map arguments for sqlx
	args := map[string]interface{}{
		"id": id,
	}

	result, err := stmt.Exec(args)
	if err != nil {
		if rbe := b.tx.Rollback(); rbe != nil {
			return 0, rbe
		}
		return 0, err
	}

	return result.RowsAffected()
}

// TxUpdate generated by sqlbrick, update data in database.
func (b *BookBrickTx) TxUpdate(args *Book) (int64, error) {
	conditionQuery := `UPDATE book SET `
	if args.Price > 0 {
		conditionQuery += `price = :price,`
	}
	if args.Content.String != "" {
		conditionQuery += `content = :content,`
	}
	conditionQuery += ` name = :name WHERE id = :id`

	if err := b.checkTx(); err != nil {
		return 0, err
	}

	stmt, err := b.tx.PrepareNamed(conditionQuery)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(args)
	if err != nil {
		if rbe := b.tx.Rollback(); rbe != nil {
			return 0, rbe
		}
		return 0, err
	}

	return result.RowsAffected()
}
