package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Страницей по адресу "/" автоматически выбирается файл index.html
	http.Handle("/", http.FileServer(http.Dir("static")))

	// Далее просто в зависимости от запроса сервера кидаем ему статичные html файлы
	http.HandleFunc("/travolta", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/travolta.html")
	})

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
