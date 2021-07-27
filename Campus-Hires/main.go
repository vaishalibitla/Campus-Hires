package main 
import (
	"database/sql"
	"log"
	"net/http"
             _ "github.com/lib/pq"
)
const hashCost = 8
var db *sql.DB
func init() {
	tmpDB, err := sql.Open("postgres", "dbname=campus_hires user=postgres password=vaish215 host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
}
func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))
   	http.HandleFunc("/",handleRegister)
	http.HandleFunc("/signin", handleSignin)
	http.HandleFunc("/signup", handleSignup)
    	http.HandleFunc("/login.html", handleLogin)
	http.HandleFunc("/register.html", handleRegister)
    	http.HandleFunc("/list", handleList)
	http.HandleFunc("/member.html", handleView)
	http.HandleFunc("/save", handleSave)
	http.HandleFunc("/delete", handleDelete)
    	http.HandleFunc("/about.html",handleAbout)
    	http.HandleFunc("/technologies.html",handleTechnologies)
    	http.HandleFunc("/contact.html",handleContact)
    	http.HandleFunc("/send", handleSend) 
	log.Fatal(http.ListenAndServe(":8085", nil))
}
