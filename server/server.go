package server

import (
	"database/sql"
	http2 "forum/services/auth/delivery/http"
	handlerCategory "forum/services/category/delivery/http"
	handlerComment "forum/services/comment/delivery/http"
	handlerLike "forum/services/like/delivery/http"
	handlerPost "forum/services/post/delivery/http"
	handlerTag "forum/services/tag/delivery/http"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	db := InitDB()

	// load templates
	t := template.Must(template.ParseGlob("./static/*.gohtml"))
	t = template.Must(t.ParseGlob("./static/components/*.gohtml"))

	mux := http.NewServeMux()
	//middleware
	mid := http2.NewAuthentication(db)
	//Register handlers
	http2.RegisterAuth(db, mux, t)
	handlerPost.RegisterPost(db, mux, *mid)
	handlerLike.RegisterLike(db, mux, *mid)
	handlerComment.RegisterPost(db, mux, *mid)
	handlerTag.RegisterTag(db, mux, *mid)
	handlerCategory.RegisterCategory(db, mux, mid)

	//Register static handlers
	mux.Handle("/static/", http.FileServer(http.Dir("../../static")))

	handler := Logging(mux)

	err := http.ListenAndServe("localhost:8080", handler)
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
	migrate(DB, Comments)
	migrate(DB, Tags)
	migrate(DB, TagPosts)
	migrate(DB, Categories)
	migrate(DB, CategoryPost)
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

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		rec := statusRecorder{w, 200}
		start := time.Now()
		next.ServeHTTP(&rec, req)
		log.Printf("%s %d %s %s", req.Method, rec.status, req.RequestURI, time.Since(start))
	})
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

const Comments string = `
CREATE TABLE IF NOT EXISTS "comments" (
	"id"	INTEGER NOT NULL UNIQUE,
	"text"	TEXT NOT NULL,
	"id_user"	INTEGER NOT NULL,
	"id_post"	INTEGER NOT NULL,
	"create_at"	TEXT NOT NULL,
	FOREIGN KEY("id_post") REFERENCES "posts"("id"),
	FOREIGN KEY("id_user") REFERENCES "user"("ID"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
`

const Tags string = `
CREATE TABLE IF NOT EXISTS "tags" (
	"id"	INTEGER NOT NULL UNIQUE,
	"title"	TEXT NOT NULL UNIQUE,
	PRIMARY KEY("id" AUTOINCREMENT)
);
`

const TagPosts string = `
CREATE TABLE  IF NOT EXISTS "tag_posts" (
	"id"	INTEGER NOT NULL UNIQUE,
	"id_tag"	INTEGER NOT NULL,
	"id_post"	INTEGER NOT NULL,
	FOREIGN KEY("id_post") REFERENCES "posts"("id"),
	FOREIGN KEY("id_tag") REFERENCES "tags"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
`
const Categories string = `
CREATE TABLE IF NOT EXISTS "categories" (
	"id"	INTEGER NOT NULL UNIQUE,
	"title"	TEXT NOT NULL UNIQUE,
	"description"	TEXT NOT NULL UNIQUE,
	PRIMARY KEY("id" AUTOINCREMENT)
);
`

const CategoryPost string = `
CREATE TABLE IF NOT EXISTS "category_posts" (
	"id"	INTEGER NOT NULL UNIQUE,
	"id_category"	INTEGER NOT NULL,
	"id_post"	INTEGER NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("id_post") REFERENCES "posts"("id"),
	FOREIGN KEY("id_category") REFERENCES "categories"("id")
);
`
