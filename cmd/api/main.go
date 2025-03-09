package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/xhos/xhos.dev.backend/internal/handlers"
	"github.com/xhos/xhos.dev.backend/internal/middleware"
)

func main() {
	fmt.Println("starting...")
	fmt.Println(`
          __                        __                __                  __                     __
   _  __ / /_   ____   _____   ____/ /___  _   __    / /_   ____ _ _____ / /__ ___   ____   ____/ /
  | |/_// __ \ / __ \ / ___/  / __  // _ \| | / /   / __ \ / __  // ___// //_// _ \ / __ \ / __  / 
 _>  < / / / // /_/ /(__  )_ / /_/ //  __/| |/ /_  / /_/ // /_/ // /__ / ,<  /  __// / / // /_/ /  
/_/|_|/_/ /_/ \____//____/(_)\__,_/ \___/ |___/(_)/_.___/ \__,_/ \___//_/|_| \___//_/ /_/ \__,_/   
                                                                                                   
`)

	// applied to all routes
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Auth,
	)

	router := handlers.SetupRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("api server is running on port 8080")
	
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
