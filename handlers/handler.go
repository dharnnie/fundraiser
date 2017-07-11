package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
)

// Payload works as payload badboy
type Payload struct {
	Name    string `bson: "name"`
	Email   string `bson: "email"`
	Message string `bson: "message"`
	Time    string `bson: "time"`
}

// Info is a struct used to share thank u msg on hmpg
type Info struct {
	Mess  string
	Mess1 string
	Mess2 string
}

// ServeResource serves static files(css and js) from server
func ServeResource(w http.ResponseWriter, r *http.Request) {
	path := "templates/" + r.URL.Path
	http.ServeFile(w, r, path)
}

// Home Serves the Landing page
func Home(w http.ResponseWriter, r *http.Request) {
	data := Info{
		Mess:  "Hello, thanks again for stopping by here, I'm glad you did. I created this website to raise funds for my tuition at African Leadership University. Kindly take some time to read through, that will be great!",
		Mess1: "Thank you for visiting!",
		Mess2: "Please help me study at ALU!",
	}
	t, err := template.ParseFiles("templates/index-image.html")
	if err != nil {
		log.Println("Could not parse home")
	}
	t.Execute(w, data)
}

// Message handles the messages from the site
func Message(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	t := time.Now().Format("2006.01.02")
	data := Payload{
		r.FormValue("name"),
		r.FormValue("email"),
		r.FormValue("message"),
		t,
	}
	log.Println(data)
	data.SendMessage()
	tmp, err := template.ParseFiles("templates/index-image.html")
	if err != nil {
		log.Println("Could not parse at SendMsg", err)
	}
	msg := Info{
		Mess:  "Thank you for your interest in my cause",
		Mess1: "Thank you for your message!",
		Mess2: "I'll revert shortly",
	}
	tmp.Execute(w, msg)
}

// SendMessage saves the message in my mongo database
func (p *Payload) SendMessage() {
	newSession := connect()
	collection := newSession.DB("funds").C("mails")
	err := collection.Insert(p)
	if err != nil {
		log.Println("Could not add to email")
	}
}

func connect() *mgo.Session {
	session, err := mgo.Dial("mongodb://dharnnie:dharnniefunds@ds153412.mlab.com:53412/funds")
	if err != nil {
		log.Println("Could not Dial DB", err)
	}
	return session
}
