package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>test DaoCloud</title>
				</head>
				<body>
					this is a my test DaoCloud
				</body>
			</html>
		`))
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}