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
					<title>����DaoCloud</title>
				</head>
				<body>
					�����ҵ�һ�����Գ���
				</body>
			</html>
		`))
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}