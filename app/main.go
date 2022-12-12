package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html>\n<head>\n<title>Skaffold Test PAge</title>\n</head>\n<body>")
		ext_ip := "localhost"

		for name, values := range r.Header {
			if name == "X-Forwarded-For" {
				ext_ip = string(r.Header["X-Forwarded-For"][0])
			}
			for _, value := range values {
				fmt.Fprintf(w, "%s = %s<br>\n", name, value)
				fmt.Println(name, value)
			}
		}

		fmt.Fprintf(w, "<h3>Hello, %s you've requested the path: %s</h3>\n", ext_ip, r.URL.Path)
		// fmt.Fprintf(w, "<h3>Hello, you've requested the path: %s</h3>\n", r.URL.Path)
		fmt.Fprintf(w, "</body>\n</html>")
	})
	http.ListenAndServe(":8080", nil)
}
