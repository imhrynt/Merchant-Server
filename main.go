package main

import (
	"fmt"
	"log"
	"server/database"
	"server/router"
	"time"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
)

var public_port = database.NewENV().PUBLIC_PORT

func main() {
	db := database.NewDatabase(database.NewENV())
	defer db.Close()
	go func() {
		for {
			log.Println("\033[H\033[2J")
			db.Ping()
			log.Println("Listening on port", public_port)
			stats := db.Status()
			view := fmt.Sprintf("Max_Open_Connection \t= %d\n"+
				"\t\t    Open_Connection     \t= %d\n"+
				"\t\t    In_Use	          \t= %d\n"+
				"\t\t    Idle                \t= %d\n"+
				"\t\t    Wait_Count          \t= %d\n"+
				"\t\t    Wait_Duration       \t= %s\n"+
				"\t\t    Max_Idle_Close      \t= %d\n"+
				"\t\t    Max_Idle_Time_Close \t= %d\n"+
				"\t\t    Max_Live_Time_Close \t= %d",
				stats.MaxOpenConnections,
				stats.OpenConnections,
				stats.InUse,
				stats.Idle,
				stats.WaitCount,
				stats.WaitDuration,
				stats.MaxIdleClosed,
				stats.MaxIdleTimeClosed,
				stats.MaxLifetimeClosed,
			)
			log.Println(view)
			time.Sleep(1 * time.Second)
		}
	}()
	server := router.NewAPI(fmt.Sprintf(":%s", public_port), db.DB)
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
