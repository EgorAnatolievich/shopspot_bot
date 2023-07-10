package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// Product represents a product in the online store
type Product struct {
	ID          int
	Name        string
	Photo       string
	Description string
	Price       float64
}

func updateProductListPeriodically() {
	ticker := time.NewTicker(13 * time.Hour)
	for {
		select {
		case <-ticker.C:
			// Вызов функции, извлекающей товары из БД и обновляющей список
			products, err := GetProductsFromDatabase()
			if err != nil {
				log.Println("Ошибка при обновлении списка товаров:", err)
			}
			updateProductList(products)
		}
	}
}

func GetProductsFromDatabase() ([]Product, error) {
	db, err := OpenDatabaseConnection()
	if err != nil {
		return nil, err
	}

	// Запрос к базе данных для получения товаров
	rows, err := db.Query("SELECT id, name, photo, description, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Перебор результатов запроса и создание списка товаров
	products := []Product{}
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Photo, &p.Description, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Получение списка товаров из базы данных
	products, err := GetProductsFromDatabase()
	if err != nil {
		http.Error(w, "Ошибка получения списка товаров", http.StatusInternalServerError)
		return
	}

	// Загрузка шаблона HTML
	tmpl, err := template.ParseFiles("html/products.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}

	// Генерация HTML-страницы с данными о товарах
	err = tmpl.Execute(w, products)
	if err != nil {
		http.Error(w, "Ошибка генерации HTML-страницы", http.StatusInternalServerError)
		return
	}
}

func updateProductList(products []Product) {
	// Код для обновления списка товаров на странице или передачи данных на клиентскую сторону
}
