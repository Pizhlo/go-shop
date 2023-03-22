package main

import (
	"log"
	"net/http"
)

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// текст "Привет из Snippetbox" как тело ответа.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // TODO: improve
		http.NotFound(w, r)
		return
	}
    w.Write([]byte("My new shop!"))
}
 
// Обработчик для отображения всех товаров.
func showProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bce товары"))
}

// Обработчик для отображения информации о пользователе.
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Пользователь"))
}
 
// Обработчик для регистрации.
func registerUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма для регистрации"))
}

// Обработчик для авторизации.
func authorizeUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма для авторизации"))
}

func main() {    
	mux := router()
    log.Println("Запуск веб-сервера на http://127.0.0.1:4000")

    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}

func router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", showProducts)
	mux.HandleFunc("/user", getUser) // TODO: user/:id
	mux.HandleFunc("/register", registerUser)
	mux.HandleFunc("/auth", authorizeUser)
    mux.HandleFunc("/", home)

	return mux
}
