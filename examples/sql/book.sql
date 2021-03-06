-- Create table if not exsited
{define name CreateBook}
CREATE TABLE IF NOT EXISTS book (
  "id"  serial NOT NULL PRIMARY KEY,
  uid int4 NOT NULL,
  name text NOT NULL,
  content varchar(255),
  create_time TIMESTAMP,
  price int NOT NULL
);
{end define}

{define name AddOne}
INSERT INTO book (uid, name, content, create_time, price)
  VALUES (${uid}, ${name}, ${content}, ${create_time}, ${price});
{end define}

-- An example to show update.
{define name UpdateSomeThing}
UPDATE book SET
{if price > 0} price = ${price}, {end if}
{if content != ""} content = ${content}, {end if}
name = ${name},
{if create_time.Unix() != 0} create_time = ${create_time} {end if}
WHERE id = ${id};
{end define}

-- An example to show complex update.
-- Second line comment.
{define name ComplexUpdate}
UPDATE book SET price=(SELECT price FROM book, user WHERE book.uid=user.id)
  WHERE book.price <= ${price} AND name = ${name};
{end define}

{define name SelectAll}
SELECT * FROM book;
{end define}

{define name CountBooks, mapper basicType}
SELECT COUNT(*) FROM book WHERE uid = ${uid};
{end define}

-- An example to show SelectById.
{define name SelectById, mapper single}
SELECT * FROM book WHERE id = ${id} ORDER BY name ASC;
{end define}

{define name SelectByUid, mapper array}
SELECT * FROM book WHERE uid = ${uid} ORDER BY name ASC;
{end define}

-- An example to show DeleteById.
{define name DeleteById}
DELETE FROM book WHERE id = ${id};
{end define}

{define name DeleteByIdAndUid}
DELETE FROM book WHERE id = ${id} and uid = ${uid};
{end define}

{define name TxInsert, tx true}
INSERT INTO book (uid, name, content, create_time, price)
  VALUES (${uid}, ${name}, ${content}, ${create_time}, ${price});
{end define}

{define name TxSelect, mapper basicType, tx true}
SELECT COUNT(*) FROM book WHERE uid = ${uid};
{end define}

{define name TxDeleteById, tx true}
DELETE FROM book WHERE id = ${id};
{end define}

{define name TxUpdate, tx true}
UPDATE book SET
{if price > 0} price = ${price}, {end if}
{if content != ""} content = ${content}, {end if}
name = ${name} WHERE id = ${id};
{end define}