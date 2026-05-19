package main

import (
	"database/sql"
	"log"
	"os"
	"time"
	_"github.com/jackc/pgconn"
	_"github.com/jackc/pgx/v4"
	_"github.com/jackc/pgx/v4/stdlib"
	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/redisstore"
	"github.com/gomodule/redigo/redis"
	"sync"
	"net/http"
	"fmt"
)

const webPort = ":8080"

func main() {
	// connect to database

	db := initDB()
	db.Ping()
	// create session
	session := initSession()

	// create channels

	// create waitgroups

	// set up application config

	// set up mail

	//listen for web connection
}

func initDB() *sql.DB {
	conn := connectToDB()

	if conn == nil {
		log.Panic("cant`t connect to database")
	}
	return conn
}

func connectToDB() *sql.DB {
	count := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)

		if err != nil {
			log.Println("postgres not ready yet")

		} else {
			log.Println("connected to database")
			return connection
		}

		if count > 10 {
			return nil
		}
		log.Println("Backing for 1 second")
		time.Sleep(1*time.Second)
		count++
		continue
	}
}

func openDB(dsn string) (*sql.DB,error) {
	db,err := sql.Open("pgx",dsn)

	if err != nil {
		return nil,err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initSession() *scs.SessionManager {
	session := scs.New()
	session.Store = redisstore.New(intitRedis())
	session.Lifetime = 24*time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func intitRedis() *redis.Pool {
	redisPool := &redis.Pool {
		MaxIdle: 10,
		Dial: func() (redis.Conn,error) {
			return redis.Dial("tcp",os.Getenv("REDIS"))
		},
	}
	return redisPool
}
