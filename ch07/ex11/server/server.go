// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"strconv"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/set", db.set)
	http.HandleFunc("/select", db.sel)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.del)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

type ItemResult struct {
	TotalCount int
	Items      []*Item
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var result ItemResult
	result.TotalCount = len(db)
	for item, price := range db {
		entry := &Item{item, price}
		result.Items = append(result.Items, entry)
	}
	if err := itemList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
func (db database) sel(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) set(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if _, ok := db[item]; !ok {
		val, err := strconv.ParseFloat(price, 32)
		if err == nil {
			db[item] = dollars(val)
			fmt.Fprintf(w, "%s: %s\n", item, price)
		} else {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "invalid price: %q\n", item)
		}
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "It has already existed the item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	} else {
		val, err := strconv.ParseFloat(price, 32)
		if err == nil {
			db[item] = dollars(val)
			fmt.Fprintf(w, "%s: %s\n", item, price)
		} else {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "invalid price: %q\n", item)
		}
	}
}

func (db database) del(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

type Item struct{
	Name string
	Price dollars
}

var itemList = template.Must(template.New("itemList").Parse(`<!DOCTYPE html>
<html>
<head>
	<title>Item List</title>
</head>
<body>
	<style>table {border-collapse: collapse;} th,td {padding: 0.5em; border: black 1px solid;}</style>
	<h1>Item List</h1>
	<table>
		<tr>
			<th>Name</th>
			<th>Price</th>
		</tr>
		{{range .Items}}
		<tr>
		  <td>{{.Name}}</td>
		  <td>{{.Price}}</td>
		</tr>
		{{end}}
	</table>
	</body>
</html>`))
