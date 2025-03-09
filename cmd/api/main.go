package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/xhos/xhos.dev.backend/internal/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("starting...")
	fmt.Println(`
          __                        __                __                  __                     __
   _  __ / /_   ____   _____   ____/ /___  _   __    / /_   ____ _ _____ / /__ ___   ____   ____/ /
  | |/_// __ \ / __ \ / ___/  / __  // _ \| | / /   / __ \ / __  // ___// //_// _ \ / __ \ / __  / 
 _>  < / / / // /_/ /(__  )_ / /_/ //  __/| |/ /_  / /_/ // /_/ // /__ / ,<  /  __// / / // /_/ /  
/_/|_|/_/ /_/ \____//____/(_)\__,_/ \___/ |___/(_)/_.___/ \__,_/ \___//_/|_| \___//_/ /_/ \__,_/   
                                                                                                   
`)
	fmt.Println("API server is running on port 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
