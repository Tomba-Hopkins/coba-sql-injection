package tambahan

import (
	"coba-sqli/database"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return 
	}

	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	if r.Method == "POST"{
		username := r.FormValue("username")
		password := r.FormValue("password")
		script := "SELECT username FROM user WHERE username = '"+ username +"' AND password = '"+ password +"' LIMIT 1"
		log.Println(script)
		row, err := db.QueryContext(ctx, script)
		if err != nil {
			panic(err)
		}
		defer row.Close()

		if row.Next(){
			var username string
			err := row.Scan(&username)
			if err != nil {
				panic(err)
			}
			log.Println("Berhasil login")
			fmt.Fprintf(w,`<script>alert("Anda berhasil Login"); window.location.href = '/dashboard';</script>`)
			return
		} else {
			fmt.Fprintf(w, `<script>alert("username password salah"); window.location.href = '/';</script>`)
		}
	}

	templ, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	templ.Execute(w, nil)
}

func HandlerRegister(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register"{
		http.NotFound(w, r)
		return 
	}
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	
	if r.Method == "POST"{
		username := r.FormValue("username")
		password := r.FormValue("password")
		script := "INSERT INTO user(username,password) VALUES(?,?)"
		_, err := db.ExecContext(ctx, script, username, password)
		if err != nil {
			panic(err)
		}
		log.Println("User baru berhasil ditambahkan")
		// http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Fprintf(w, `<script>alert("User baru berhasil ditambahkan"); window.location.href = '/';</script>`)
		return
	}
	templ, err := template.ParseFiles("register.html")
	if err != nil {
		panic(err)
	}
	templ.Execute(w, nil)
}

func HandlerDashboard(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/dashboard"{
		http.NotFound(w, r)
		return
	}

	templ, err := template.ParseFiles("dashboard.html")
	if err != nil {
		panic(err)
	}
	templ.Execute(w, nil)
}