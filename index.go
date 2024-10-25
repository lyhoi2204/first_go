package main

import (
	"fmt"
	"net/http"
)

// Hàm xử lý route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!") // Gửi phản hồi về client
}

func main() {
	// Đăng ký route "/hello" với hàm xử lý helloHandler
	http.HandleFunc("/hello", helloHandler)

	// Lắng nghe và phục vụ trên cổng 8080
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil) // nil là bộ xử lý mặc định
}
