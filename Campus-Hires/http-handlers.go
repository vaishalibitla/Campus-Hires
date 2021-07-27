package main 
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"      
	"golang.org/x/crypto/bcrypt"
	_ "github.com/lib/pq"
)
func handleLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./www/login.html")
}
func handleRegister(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./www/register.html")
}
func handleSignup(w http.ResponseWriter, r *http.Request){
                var err error
                r.ParseForm()
	params := r.PostForm
	username := params.Get("username")
                password := params.Get("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	_, err = db.Query("insert into users values ($1, $2)", username, string(hashedPassword))
                if err != nil {
		renderError(w, err)
		return
	}
    http.Redirect(w, r, "/login.html", 302)
}
func handleSignin(w http.ResponseWriter, r *http.Request){
	var err error
                r.ParseForm()
	params := r.PostForm
	username := params.Get("username")
    	password := params.Get("password")
	result := db.QueryRow("select password from users where username=$1", username)
	storedCreds := &Credentials{}
	err = result.Scan(&storedCreds.password)
	if err != nil {
		renderError(w, err)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.password), []byte(password)); err != nil {
                     renderError(w, err)
	}
	http.Redirect(w, r, "/list", 302)
}
func handleSend(w http.ResponseWriter, r *http.Request) {
	var err error
	r.ParseForm()
	params := r.PostForm
                email := params.Get("email")
	query := params.Get("query")
	_, err = insertQuery(email, query)
	if err != nil {
		renderErrorPage(w, err)
		return
	}
	http.Redirect(w, r, "/contact.html", 302)
}
func handleSave(w http.ResponseWriter, r *http.Request) {
                var id=0
	var err error
	r.ParseForm()
	params := r.PostForm
	idStr := params.Get("id")
	if len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}
	name := params.Get("name")
	email := params.Get("email")
	phoneStr := params.Get("phone")
	phone := 0
	if len(phoneStr) > 0 {
		phone, err = strconv.Atoi(phoneStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}
	joiningDateStr := params.Get("joiningDate")
	var joiningDate time.Time
	if len(joiningDateStr) > 0 {
		joiningDate, err = time.Parse("2006-01-02", joiningDateStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}
	if id == 0 {
		_, err = insertMember(name, email, phone, joiningDate)
	} else {
		_, err = updateMember(id, name, email, phone, joiningDate)
	}
	if err != nil {
		renderErrorPage(w, err)
		return
	}
	http.Redirect(w, r, "/list", 302)
}
func handleList(w http.ResponseWriter, r *http.Request) {
	members, err := allMembers()
	if err != nil {
		renderErrorPage(w, err)
		return
	}
	buf, err := ioutil.ReadFile("www/index.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}
	var page = AllMember{AllMembers: members}
	allMember := string(buf)
	t := template.Must(template.New("allMember").Parse(allMember))
	t.Execute(w, page)
}
func handleView(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")
	var currentMember = Member{}
	currentMember.JoiningDate = time.Now()
	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
		currentMember, err = getMember(id)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}
	buf, err := ioutil.ReadFile("www/member.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}
	var page = MemberPage{TargetMem: currentMember}
	memberPage := string(buf)
	t := template.Must(template.New("memberPage").Parse(memberPage))
	err = t.Execute(w, page)
	if err != nil {
		renderErrorPage(w, err)
		return
	}
}
func handleDelete(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")
	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
		n, err := removeMember(id)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
		fmt.Printf("Rows removed: %v\n", n)
	}
	http.Redirect(w, r, "/list", 302)
}
func renderErrorPage(w http.ResponseWriter, errorMsg error) {
	buf, err := ioutil.ReadFile("www/errorpage.html")
	if err != nil {
		log.Printf("%v\n", err)
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	var page = ErrorPage{ErrorMsg: errorMsg.Error()}
	errorPage := string(buf)
	t := template.Must(template.New("errorPage").Parse(errorPage))
	t.Execute(w, page)
}
func renderError(w http.ResponseWriter, errorMsg error) {
	buf, err := ioutil.ReadFile("www/error.html")
	if err != nil {
		log.Printf("%v\n", err)
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	var page = ErrorPage{ErrorMsg: errorMsg.Error()}
	errorPage := string(buf)
	t := template.Must(template.New("errorPage").Parse(errorPage))
	t.Execute(w, page)
}
func handleAbout(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./www/about.html")
}
func handleTechnologies(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./www/technologies.html")
}
func handleContact(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./www/contact.html")
}