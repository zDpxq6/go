// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 115.

// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}
type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type memorizingSort struct {
	t        []*Track
	memories []func(x, y *Track) bool
}

func (x memorizingSort) Len() int      { return len(x.t) }
func (x memorizingSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x memorizingSort) Less(i, j int) bool {
	result := false
	for _, less := range x.memories {
		if less != nil {
			result = less(x.t[i], x.t[j])
		}
	}
	return result
}

// Web Server

var tracksList = template.Must(template.New("tracksList").Parse(`<!DOCTYPE html>
<html>
<head>
	<title>Track List</title>
</head>
<body>
	<style>table {border-collapse: collapse;} th,td {padding: 0.5em; border: black 1px solid;}</style>
	<h1>Track List</h1>
	<table>
		<tr>
			<th><a href="?column=Title">Title</a></th>
			<th><a href="?column=Artist">Artist</a></th>
			<th><a href="?column=Album">Album</a></th>
			<th><a href="?column=Year">Year</a></th>
			<th><a href="?column=Length">Length</a></th>
		</tr>
		{{range .}}
		<tr>
		  <td>{{.Title}}</td>
		  <td>{{.Artist}}</td>
		  <td>{{.Album}}</td>
		  <td>{{.Year}}</td>
		  <td>{{.Length}}</td>
		</tr>
		{{end}}
	</table>
	</body>
</html>`))

func choose(key string) func(x, y *Track) bool {
	switch key {
	case "Title":
		return func(x, y *Track) bool {
			return x.Title < y.Title
		}
	case "Artist":
		return func(x, y *Track) bool {
			return x.Artist < y.Artist
		}
	case "Album":
		return func(x, y *Track) bool {
			return x.Album < y.Album
		}
	case "Year":
		return func(x, y *Track) bool {
			return x.Year < y.Year
		}
	case "Length":
		return func(x, y *Track) bool {
			return x.Length < y.Length
		}
	}
	return nil
}

var memories = make([]func(x, y *Track) bool, max*2)
var tmp int

const max int = 8

func main() {
	sockAddress := "localhost:8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		column := r.FormValue("column")
		memories[0], memories[1], memories[2] = memories[1], memories[2], memories[3]
		memories[3] = choose(column)
		sort.Sort(memorizingSort{tracks, memories})
		if err := tracksList.Execute(w, tracks); err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(sockAddress, nil))
	return
}
