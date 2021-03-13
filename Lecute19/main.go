package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)
func init() {

	   // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/hosting_db")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

	 // defer the close till after the main function has finished
    // executing
    defer db.Close()

    // perform a db.Query insert
    insert, err := db.Query("INSERT INTO `request` (`ID`, `name`, `company`, `email`, `status`) VALUES (NULL, 'Mamun Khan', 'abc', 'mamun@gmail.com', '1');")

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

	  // if there is an error inserting, handle it
	  if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()

	fmt.Println("db connection successful")
}
func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/request", request)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer((http.Dir("assets")))))
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
	 
}

func features(w http.ResponseWriter, r *http.Request) {
	
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}
func docs(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

// method 01
/*func request(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	name := r.FormValue("name")
	company := r.FormValue("company")
	email := r.FormValue("email")

	fmt.Println(name, company, email)

	fmt.Fprint(w, `received`) //response
	// fmt.Fprint(w. `received %s %s %s`, name, company, email) page a dektachile
}*/

// method_02
func request(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for key, val := range r.Form {

		fmt.Println(key, val)
	}

	fmt.Fprint(w, `Done`)
}