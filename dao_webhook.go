package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

/*
	JSON
	{
		"email":"alex@example.com",
		"timestamp":1337197600,
		"smtp-id":"<4FB4041F.6080505@example.com>",
		"event":"processed"
	}

	MariaDB [sendgrid]> desc mail_events;
	+-----------+--------------+------+-----+---------+----------------+
	| Field     | Type         | Null | Key | Default | Extra          |
	+-----------+--------------+------+-----+---------+----------------+
	| id        | int(11)      | NO   | PRI | NULL    | auto_increment |
	| email     | varchar(200) | NO   |     | NULL    |                |
	| timestamp | varchar(50)  | NO   |     | NULL    |                |
	| smtp_id   | varchar(200) | NO   |     | NULL    |                |
	| event     | varchar(50)  | NO   |     | NULL    |                |
	+-----------+--------------+------+-----+---------+----------------+
*/


func insertWebHookEvents(data []WebHookEvent) {
	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := "INSERT INTO `sendgrid`.`mail_events` (`email`,`timestamp`,`smtp_id`,`event`) VALUES "

	dataProcessed := 0

	for _, item := range data {

		if dataProcessed > 0 {
			sql = sql + ","
		}

		sql = sql + " ('" + item.Email + "','" + strconv.Itoa(item.Timestamp) + "','" + item.SmtpId + "','" + item.Event + "')"

		dataProcessed++
	}

	if dataProcessed > 0 {
		stmt, err := db.Prepare(sql)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec()
		if err != nil {
			log.Fatal(err)
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("SQL: %v, Result: %v\n", sql, affect)
	}
}

/*
func insertSingleQueueLogTransfer(data QueueLogTransfer) {
	db, err := sql.Open("mysql", configuration.DbUsername+":"+configuration.DbPassword+"@tcp("+configuration.DbServer+":"+configuration.DbPort+")/"+configuration.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sql := "INSERT INTO `queuelog_co_transfers`.`queuelog_calls` (`unique_id`,`origin_queue`,`destination_queue`,`caller_channel`,`timestamp`) VALUES "

	sql = sql + " ('" + data.UniqueID + "','" + data.OriginQueue + "','" + data.DestinationQueue + "','" + data.CallerChannel + "','" + data.Timestamp + "')"

	stmt, err := db.Prepare(sql)
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("Error1: %v\n", err)
	}

	if err == nil {
		res, err := stmt.Exec()
		if err != nil {
			//log.Fatal(err)
			fmt.Printf("Error2: %v\n", err)
		}

		if err == nil {

			affect, err := res.RowsAffected()
			if err != nil {
				//log.Fatal(err)
				fmt.Printf("Error3: %v\n", err)
			}

			_ = affect

			if err == nil {
				fmt.Printf("SQL: %v, Result: %v\n", sql, affect)
				saveLastUniqueID(data.UniqueID)
			}
		}
	}

}
*/