package main 

import (
	"fmt"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after request phase!")
	})
}
	
func mainLogic( w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing middleware .. . .")
	w.Write([]byte("OK"))
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	http.ListenAndServe(":8000", nil)
}