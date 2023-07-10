package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Здесь можно вызвать функции и настройки для запуска вашего веб-приложения
	// Например, setupDatabase(), setupTelegramBot(), setupWebApp() и т.д.

	// Вызов функции setupWebApp для настройки вашего веб-приложения
	setupWebApp()

	go updateProductListPeriodically()

	// Запуск веб-сервера
	log.Println("Сервер запущен на порту :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setupWebApp() {
	// Определение обработчика для маршрута "/display-products"
	http.HandleFunc("/display-products", displayProducts)
}

func displayProducts(w http.ResponseWriter, r *http.Request) {
	// Получение списка товаров из базы данных
	products, err := GetProductsFromDatabase()
	if err != nil {
		http.Error(w, "Ошибка получения списка товаров", http.StatusInternalServerError)
		return
	}

	// Вывод информации о каждом товаре
	for _, product := range products {
		fmt.Fprintf(w, "ID: %d\n", product.ID)
		fmt.Fprintf(w, "Название: %s\n", product.Name)
		fmt.Fprintf(w, "Фото: %s\n", product.Photo)
		fmt.Fprintf(w, "Описание: %s\n", product.Description)
		fmt.Fprintf(w, "Цена: %.2f\n", product.Price)
		fmt.Fprintln(w, "-------------------")
	}
}
