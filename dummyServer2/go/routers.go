package swagger

import (

    //"net/http"
    "github.com/gorilla/mux"

)

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
  
    router.HandleFunc("/addleave", AddLeave).Methods("POST", "GET")

    return router

}