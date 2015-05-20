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
					<title>测试DaoCloud</title>
				</head>
				<body>
					这是我的一个测试程序
				</body>
			</html>
		`))
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}