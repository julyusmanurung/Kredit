package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/julyusmanurung/Kredit/api"
	"github.com/julyusmanurung/Kredit/database"
)

func main() {
	db, err := database.SetupDB()
	if err != nil {
		panic(err)
	}

	server := api.MakeServer(db)
	server.RunServer()

	// Create a ticker that will run the scheduler function every 30 minutes
	//ticker := time.NewTicker(time.Minute * 30)
	//defer ticker.Stop()
	//
	//// Run the scheduler function when the ticker fires
	//for range ticker.C {
	//	err := scheduler(db)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
}

//func scheduler(db *sql.DB) error {
//	// Query the database for some data
//	rows, err := db.Query("SELECT * FROM users")
//	if err != nil {
//		return err
//	}
//	defer rows.Close()
//
//	// Print the data to the console
//	for rows.Next() {
//		var id int
//		var name string
//		err := rows.Scan(&id, &name)
//		if err != nil {
//			return err
//		}
//		fmt.Println(id, name)
//	}
//	return rows.Err()
//}
