package main

import (
	"io"
	"net/http"
	"os"
)

func hello(res http.ResponseWriter, req *http.Request) {
	
	errorPage := 
	`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Go Eray Go</title>
	</head>
	<body>
		<p>Could not find the requested page.</p>
	</body>
	</html>
	`
	
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	// Open html file
	htmlFile, err := os.Open("index.html")
	if err != nil {
		io.WriteString(
			res, 
			errorPage,
		)
		return
	}
	defer htmlFile.Close()

	// Search file statistics
	stats, err := htmlFile.Stat()
	if err != nil {
		io.WriteString(res, errorPage)
		return
	}

	// Read the byte buffer of file
	bs := make([]byte, stats.Size())
	_, err = htmlFile.Read(bs)

	if err != nil {
		io.WriteString(res, errorPage)
		return
	}

	io.WriteString(
		res,
		string(bs),
	)
}

func main() {
	http.HandleFunc("/hello", hello)

	// Static file handling
	http.Handle("/assets/", 
	http.StripPrefix(
		"/assets/",
		http.FileServer(http.Dir("assets")),
		),
	)

	http.ListenAndServe(":9000", nil)
}