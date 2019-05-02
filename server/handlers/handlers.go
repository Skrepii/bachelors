package handlers

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	adminUser, adminPass := "mike", "wasowski"
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error in parsing form: ", err.Error())
	}

	fmt.Printf("Received a request => username: %s, password: %s\n", r.FormValue("username"), r.FormValue("password"))

	if r.FormValue("username") != adminUser || r.FormValue("password") != adminPass {
		_, _ = fmt.Fprint(w, "Access denied")
		return
	}

	http.Redirect(w, r, "/user.html", http.StatusSeeOther)
}
