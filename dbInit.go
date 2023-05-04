package forum

import (
	"database/sql"
	"log"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("mysql", user + ":"+ passwd +"@tcp(" + host +":3306)/" + dbname)
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("CREATE TABLE IF NOT EXISTS LikeComment (Comment_id int, User_id int, IsLike int)")
	db.Exec("CREATE TABLE IF NOT EXISTS LikePost (Post_id int, User_id int, IsLike int)")
	db.Exec("CREATE TABLE IF NOT EXISTS ReportPost (Post_id int UNIQUE, User_id int)")
	db.Exec("CREATE TABLE IF NOT EXISTS ReportComment (Post_id int, Comment_id int UNIQUE, User_id int)")
	db.Exec("CREATE TABLE IF NOT EXISTS Comment (Id int UNIQUE PRIMARY KEY NOT NULL AUTO_INCREMENT, Post_id int, User_id int, Text CHAR(255), CreatedAt DATETIME)")
	db.Exec("CREATE TABLE IF NOT EXISTS Post (Id int UNIQUE PRIMARY KEY NOT NULL AUTO_INCREMENT, User_id int, Title CHAR(255) UNIQUE,Text CHAR(255), CreatedAt DATETIME, UpdatedAT DATETIME)")
	db.Exec("CREATE TABLE IF NOT EXISTS User (Id int UNIQUE PRIMARY KEY NOT NULL AUTO_INCREMENT,UUID CHAR(255) , Username CHAR(255) UNIQUE, Email CHAR(255) UNIQUE, Password CHAR(255), Role CHAR(255), CreatedAt DATETIME)")
	db.Exec("CREATE TABLE IF NOT EXISTS Categories (Id int UNIQUE PRIMARY KEY NOT NULL AUTO_INCREMENT, Name CHAR(255))")
	db.Exec("CREATE TABLE IF NOT EXISTS Post_Categories (Post_id int, Categorie_id int)")
	db.Exec("CREATE TABLE IF NOT EXISTS RequestMod (User_id int, Reason CHAR(255))")
}

