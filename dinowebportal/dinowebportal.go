package dinowebportal

import (
	"fmt"
	"net/http"
)
// RunWebPortal starts running the dino portal on address addr
func RunWebPortal(addr string) error {
	http.HandleFunc("/", roothandler)  // the / is the root directory
	return http.ListenAndServe(addr, nil)
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Dino web portal %s", r.RemoteAddr)
}