package server

import (
	"database/sql"
	"log"
	"net/http"

	//handlers
	handlerAuth "forum/auth/delivery/http"
	handlerLike "forum/like/delivery/http"
	handlerPost "forum/post/delivery/http"

	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	db := InitDB()
	mux := http.NewServeMux()
	//Register handlers
	handlerAuth.RegisterAuth(db, mux)
	handlerPost.RegisterPost(db, mux)
	handlerLike.RegisterLike(db, mux)

	err := http.ListenAndServe("localhost:8080", mux)
	log.Println(err)
}

func InitDB() *sql.DB {
	DB, err := sql.Open("sqlite3", "../../forumDB.db")
	if err != nil {
		panic(err)
	}
	migrate(DB, Users)
	migrate(DB, Posts)
	migrate(DB, LikePost)
	migrate(DB, LikeComment)
	return DB
}

func migrate(db *sql.DB, query string) {
	statement, err := db.Prepare(query)
	if err == nil {
		_, creationError := statement.Exec()
		if creationError == nil {
			log.Println("Table created successfully")
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}

}

const Users string = `
CREATE TABLE IF NOT EXISTS "user" (
	"ID"	INTEGER NOT NULL UNIQUE,
	"NicName"	TEXT NOT NULL UNIQUE,
	"Email"	TEXT NOT NULL UNIQUE,
	"Password"	TEXT NOT NULL,
	"Token"	TEXT,
	PRIMARY KEY("ID" AUTOINCREMENT)
);
`
const Posts string = `
CREATE TABLE IF NOT EXISTS "posts" (
	"id"	INTEGER NOT NULL UNIQUE,
	"title"	TEXT NOT NULL UNIQUE,
	"content"	TEXT NOT NULL UNIQUE,
	"create_at"	TEXT NOT NULL,
	"update_at"	TEXT NOT NULL,
	"id_user"	INTEGER NOT NULL,
	FOREIGN KEY("id_user") REFERENCES "user"("ID"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
`

const LikePost string = `
CREATE TABLE IF NOT EXISTS "likes_posts" (
	"id_post"	INTEGER NOT NULL,
	"id_user"	INTEGER NOT NULL,
	"liked"	INTEGER NOT NULL DEFAULT 0,
	FOREIGN KEY("id_user") REFERENCES "user"("ID"),
	FOREIGN KEY("id_post") REFERENCES "posts"("id")
);
`

const LikeComment string = `
CREATE TABLE IF NOT EXISTS "likes_comment" (
	"id_comment"	INTEGER NOT NULL,
	"id_user"	INTEGER NOT NULL,
	"liked"	INTEGER NOT NULL DEFAULT 0,
	FOREIGN KEY("id_user") REFERENCES "user"("ID"),
	FOREIGN KEY("id_comment") REFERENCES "comments"("id")
);
`
