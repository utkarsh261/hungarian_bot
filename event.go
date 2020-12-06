package main

import (
	"database/sql"
	"log"
	"strings"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/mattn/go-sqlite3"
)

func addevent(ID int64, msg string) {
	db, err := sql.Open("sqlite3", "./testdb.db")
	if err != nil {
		log.Panic("Failed to connect to database")
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS event (data TEXT)")
	if err != nil {
		log.Panic(err)
	}
	args := strings.Fields(msg)
	space := " "
	var to_insert strings.Builder
	for i := 1; i < len(args); i++ {
		to_insert.WriteString(args[i])
		to_insert.WriteString(space)
	}
	actual := to_insert.String()
	if len(actual) < 1 {
		log.Panic("Cant Concat slice to the string")
	}
	if len(args) > 1 {
		stmt, err := db.Prepare("INSERT INTO event (data) VALUES (?)")
		if err != nil {
			log.Panic(err)
		}
		res, err := stmt.Exec(actual)
		if err != nil {
			log.Panic(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Panic(err)
		}
		log.Println(id)
		db.Close()
		log.Print("VALUE ADDED INTO DATABASE")
		bot.Send(tbot.NewMessage(ID, "Event added succesfully!"))
	} else {
		log.Print("NOT ENOUGH ARGUMENTS")
		bot.Send(tbot.NewMessage(ID, "Please provide details of the event, i.e. Title and expected Date/Time."))
	}
}

func listevents(ID int64) {
	db, err := sql.Open("sqlite3", "./testdb.db")
	if err != nil {
		log.Panic("Failed to connect to database")
	}
	rows, err := db.Query("SELECT * FROM event")
	if err != nil {
		log.Panic(err)
	}
	var sb strings.Builder
	newline := "\n\n"
	for rows.Next() {
		var info string
		err = rows.Scan(&info)
		if err != nil {
			log.Panic(err)
		}
		sb.WriteString(info)
		sb.WriteString(newline)
	}
	show := sb.String()
	bot.Send(tbot.NewMessage(ID, show))
	// Use to delete after testing.
	//_, err = db.Exec("DELETE FROM event")
	//if err != nil {
	//	log.Panic(err)
	//}
	db.Close()
}

func nextevent(ID int64) {
	// Would be logical to do this after events support date/time.
}
