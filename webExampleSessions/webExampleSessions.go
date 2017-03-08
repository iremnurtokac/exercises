package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// ---- my substitue starts

type Loging struct {
	Writing    http.ResponseWriter // w
	Requesting *http.Request       // r
}

func (lg *Loging) logInfo(session *sessions.Session) sessions.Store {

	session, _ = store.Get(lg.Requesting, "cookie-name")
	//check if the user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(lg.Writing, "Forbidden", http.StatusForbidden)
		return session.Store()
	}
	//Print secret message
	fmt.Fprintln(lg.Writing, "The cake is a lie!")
	return session.Store()

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		session.Values["authenticated"] = true
		session.Save(lg.Requesting, lg.Writing)
	}
	return session.Store()

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		session.Values["authenticated"] = false
		session.Save(lg.Requesting, lg.Writing)
	}
	return session.Store()
}

// ----- my substitute ends

// ---- original code starts

/*
func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	//check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	//Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	//Authentication goes here
	//...

	//Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}
func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	//Revoke users authentication

	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}


*/

// ---- original code ends
