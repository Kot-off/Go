// Программа для управления небольшим складом товаров
package main

import (
	"fmt"
	"time"
)

// Глобальные константы для категорий товаров
const (
	CategoryElectronics = "Электроника"
	CategoryFood        = "Продукты"
	CategoryClothes     = "Одежда"
	MaxItems            = 100 // Максимальное количество товаров на складе
)

// Глобальная переменная для подсчета товаров
var totalItems int

func main() {
	// Объявление и инициализация переменных разными способами

	var itemName string
	itemName = "Смартфон"

	var itemPrice float64 = 999.99

	var minQuantity, maxQuantity int = 5, 20

	var isAvailable = true

	quantity := 15

	var (
		itemID     int64  = 12345
		itemColor  string = "Черный"
		itemWeight float32
		dateAdded  time.Time = time.Now()
	)

	itemWeight = 0.3

	percentInStock := float64(quantity) / float64(MaxItems) * 100

	category := CategoryElectronics

	displayItemInfo(itemID, itemName, quantity, itemPrice, isAvailable, category)

	fmt.Println("\nДополнительная информация:")
	fmt.Printf("Цвет товара: %s\n", itemColor)
	fmt.Printf("Вес товара: %.2f кг\n", itemWeight)
	fmt.Printf("Минимальное количество: %d, Максимальное количество: %d\n", minQuantity, maxQuantity)
	fmt.Printf("Процент от максимального количества на складе: %.1f%%\n", percentInStock)
	fmt.Printf("Дата добавления: %s\n", dateAdded.Format("02-01-2006"))

	const (
		Small  = 1
		Medium = 5
		Large  = 10
	)

	var (
		shortValue   int8      = 127
		intValue     int       = 1000000
		uintValue    uint      = 10000
		floatValue   float32   = 123.456
		complexValue complex64 = 1 + 2i
	)

	fmt.Println("\nРазные числовые типы:")
	fmt.Printf("int8: %d\n", shortValue)
	fmt.Printf("int: %d\n", intValue)
	fmt.Printf("uint: %d\n", uintValue)
	fmt.Printf("float32: %f\n", floatValue)
	fmt.Printf("complex64: %v\n", complexValue)

	productCode := "ЯЩИК-12345"
	fmt.Printf("\nКод товара: %s, длина: %d символов, длина: %d байт\n", productCode, len([]rune(productCode)), len(productCode))

	updateTotalItems(quantity)
	fmt.Printf("\nОбщее количество товаров на складе: %d\n", totalItems)
}

// Функция для отображения информации о товаре
func displayItemInfo(id int64, name string, qty int, price float64, isAvailable bool, category string) {
	fmt.Println("=== Информация о товаре ===")
	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Название: %s\n", name)
	fmt.Printf("Категория: %s\n", category)
	fmt.Printf("Количество: %d\n", qty)
	fmt.Printf("Цена: %.2f руб.\n", price)

	if isAvailable {
		fmt.Println("Статус: В наличии")
	} else {
		fmt.Println("Статус: Нет в наличии")
	}
}

// Функция для обновления общего количества товаров
func updateTotalItems(qty int) {
	totalItems += qty

	if totalItems > MaxItems {
		fmt.Println("Предупреждение: Превышено максимальное количество товаров на складе!")
		totalItems = MaxItems
	}
}

// Добавляет новый товар в систему и возвращает его ID.
func addNewItem(name string, price float64, qty int, category string) int64 {
	totalItems += qty
	newID := time.Now().UnixNano() // Генерация уникального ID
	fmt.Printf("Товар '%s' успешно добавлен. Новый ID: %d\n", name, newID)
	return newID
}

// Рассчитывает скидку на товар в зависимости от условий.
func calculateDiscount(price float64, percent float64) float64 {
	if percent < 0 || percent > 100 {
		fmt.Println("Ошибка: некорректное значение скидки.")
		return price
	}
	discount := price * percent / 100
	finalPrice := price - discount
	return finalPrice
}
