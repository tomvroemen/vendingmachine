package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	cache "github.com/tomvroemen/buzz-cache"
)

var DB = cache.New(cache.NoExpiration, time.Duration(300*time.Second)) // databases
var AC = cache.New(cache.NoExpiration, time.Duration(300*time.Second)) // auth credentials

var ACCEPTEDAMOUNTS [5]int = [5]int{5, 10, 20, 50, 100}

type Product struct { // amountAvailable, cost, productName and sellerId
	AmountAvailable int    `json:"amountAvailable"`
	Cost            string `json:"cost"`
	ProductName     string `json:"productName"`
	SellerId        string `json:"sellerId"`
}

type User struct { // username, password, deposit and role
	Username int    `json:"username"`
	Password string `json:"password"`
	Deposit  int    `json:"deposit"`
	Role     string `json:"role"`
}

func main() {
	server()
}

func server() {
	fmt.Println("starting HTTP server")
	mux := http.NewServeMux()
	//	mux.Handle("/checkout/", CheckoutHandler())
	mux.Handle("/user/", http.StripPrefix("/user/", UserHandler()))
	mux.Handle("/product/", http.StripPrefix("/product/", ProductHandler()))
	mux.Handle("/deposit/", http.StripPrefix("/deposit/", DepositHandler())) // TODO
	mux.Handle("/buy/", http.StripPrefix("/buy/", BuyHandler()))             // TODO
	mux.Handle("/reset/", http.StripPrefix("/reset/", ResetHandler()))       // TODO
	//mux.Handle("/", IndexHandler())
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		panic(err)
	}
}

func Auth(r http.Request, username, password string) bool {
	user, pass, ok := r.BasicAuth()
	if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
		return false
	}
	return true
}

/*
CRUD	HTTP
Create	POST
Read	GET
Update	PUT
Delete	DELETE
*/

func UserHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		//fmt.Println(r.Method, r.URL.Path)

		result := map[string]interface{}{"status": "no data to return"}
		cnt := 0

		if r.Method == "POST" {
			//result =
			result["status"] = "created " + strconv.Itoa(cnt) + " entries"
		} else { // other methods are Authenticated
			if Auth(*r, username, password) { // TODO: source credentials from DB
				w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}
			if r.Method == "GET" {
				//result =
				result["status"] = "read " + strconv.Itoa(cnt) + " entries"
			} else if r.Method == "PUT" {
				//result =
				result["status"] = "updated " + strconv.Itoa(cnt) + " entries"
			} else if r.Method == "DELETE" {
				//result =
				result["status"] = "deleted " + strconv.Itoa(cnt) + " entries"
			}
		}
		js, _ := json.MarshalIndent(result, "", "    ")
		w.Write([]byte(js))
		return
	}
	return http.HandlerFunc(fn)
}

func ProductHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		//fmt.Println(r.Method, r.URL.Path)

		result := map[string]interface{}{"status": "no data to return"}
		cnt := 0

		if r.Method == "GET" {
			//result =
			result["status"] = "read " + strconv.Itoa(cnt) + " entries"
		} else { // other methods are Authenticated
			if Auth(*r, username, password) { // TODO: source credentials from DB
				w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}
			if r.Method == "POST" {
				//result =
				result["status"] = "created " + strconv.Itoa(cnt) + " entries"
			} else if r.Method == "PUT" {
				//result =
				result["status"] = "updated " + strconv.Itoa(cnt) + " entries"
			} else if r.Method == "DELETE" {
				//result =
				result["status"] = "deleted " + strconv.Itoa(cnt) + " entries"
			}
		}
		js, _ := json.MarshalIndent(result, "", "    ")
		w.Write([]byte(js))
		return
	}
	return http.HandlerFunc(fn)
}
