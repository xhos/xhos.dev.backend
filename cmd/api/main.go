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
                                              __    __               __                  __
   ____ ___  __  __________  _________  _____/ /_  / /_  ____ ______/ /_____  ____  ____/ /
  / __  __ \/ / / / ___/ _ \/ ___/ __ \/ ___/ __/ / __ \/ __  / ___/ //_/ _ \/ __ \/ __  / 
 / / / / / / /_/ (__  )  __(__  ) /_/ / /  / /__ / /_/ / /_/ / /__/ ,< /  __/ / / / /_/ /  
/_/ /_/ /_/\__,_/____/\___/____/\____/_/   \__(_)_.___/\__,_/\___/_/|_|\___/_/ /_/\__,_/   
                                                                                                                                                                            
`)

	// applied to all routes
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.CORS,
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
