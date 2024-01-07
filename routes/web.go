package routes

import (
	"net/http"

	"github.com/dybaay/go-server/controllers"
)

func Routes() {
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("./views"))))

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/form/view", controllers.ShowForm)
	http.HandleFunc("/form/submit", controllers.SubmitForm)
}
