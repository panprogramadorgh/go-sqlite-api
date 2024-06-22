import sqlite3
from crypt import crypt as hash, METHOD_BLOWFISH

conn = sqlite3.connect("./database.db")


def sql(conn: sqlite3.Connection, query_str: str,  *args) -> list[any]:
  cursor = conn.execute(query_str, args)
  return cursor.fetchall()


init_transaction = '''
BEGIN
'''
drop_table = '''
DROP TABLE IF EXISTS users
'''
create_table = '''
CREATE TABLE users (
  user_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  username TEXT,
  password TEXT,
  firstname TEXT,
  lastname TEXT
)
'''
insert_user = '''
INSERT INTO users (username, password, firstname, lastname) VALUES (?, ?, ?, ?)
'''
see_users = '''
SELECT * FROM users WHERE TRUE
'''
commit_transaction = '''
COMMIT
'''

sql(conn, init_transaction)
sql(conn, drop_table)
sql(conn, create_table)
usr_password = hash(word="pass1", salt=METHOD_BLOWFISH)
sql(conn, insert_user, "user1", usr_password, "paco", "sanz")
print(sql(conn, see_users))
sql(conn, commit_transaction)
