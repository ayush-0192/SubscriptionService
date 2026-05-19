package main

const webPort = "80"

func main() {
	// connect to database

	db := initDB()
	db.Ping()
	// create session

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
