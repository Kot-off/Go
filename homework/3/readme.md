---
# Основные конструкции языка Go

## Цель задания

### Выполнив это задание, вы:

- Практически отработаете работу с различными типами данных;
- Научитесь использовать **переменные** и **константы**;
- Освоите различные способы объявления переменных;
- Ознакомитесь с организацией структуры кода в Go;
- Реализуете функции и структурируете логику программы.
---

## 🛠 Необходимые инструменты

- **Go** версии `≥ 1.18`
- Терминал (например: bash, PowerShell, CMD)
- Редактор кода (рекомендуется [VS Code](https://code.visualstudio.com/))

---

## 📋 Описание задания

Программа представляет собой простой симулятор учета товаров на складе:

- Определяет категории товаров;
- Использует глобальные и локальные переменные;
- Демонстрирует работу с разными числовыми типами;
- Реализует функции для отображения и обработки информации о товарах.

---

## Инструкция по выполнению

1. Изучите предоставленный шаблон кода.  
   Ознакомьтесь с [структурой кода и комментариями](https://go.dev/doc/tutorial/create-module) для понимания различных способов объявления переменных и констант.

2. Дополните программу своими функциями:

   - `addNewItem` — функция для добавления нового товара и возврата его ID.
   - `calculateDiscount` — функция для расчёта скидки на товар.

3. Сделайте копию шаблона для сдачи домашнего задания и сохраните в неё получившийся код.  
   Предоставьте доступ на комментирование всем, у кого есть ссылка.

---

## Рекомендации

- Обратите особое внимание на различные способы объявления переменных.
- Помните о разнице между `var` и оператором `:=`.
- Константы в Go объявляются с помощью `const`.
- Следите за областью видимости переменных и констант.
- Используйте говорящие имена для переменных и функций.
- Используйте пакет `fmt` для:
  - отображения информации о товарах в удобочитаемом формате;
  - форматирования числовых значений (цены, веса) с нужной точностью;
  - вывода предупреждений и уведомлений;
  - отображения дат в удобном для пользователя формате.

---

## 🧩 Основной функционал

### ✅ Что реализовано:

- Объявление и инициализация переменных разными способами (`var`, `:=`);
- Использование глобальных **констант** (`const`);
- Работа с разными **типами данных** (`int`, `float`, `complex`, `string`, `bool`, `time.Time`);
- Форматированный вывод с использованием пакета `fmt`;
- Отображение даты добавления товара;
- Подсчет общего количества товаров;
- Предупреждение при превышении лимита.

Вот обновлённый `README.md`, в который добавлен весь **исходный код программы**, оформленный в соответствующем разделе. Код заключён в блок Markdown с подсветкой синтаксиса Go.

---

## 🧑‍💻 Дополнительные функции (реализовать):

### ➕ `addNewItem`

> Добавляет новый товар в систему и возвращает его `ID`.

```go
func addNewItem(name string, price float64, qty int, category string) int64 {
    totalItems += qty
    newID := time.Now().UnixNano() // Простой способ генерации уникального ID
    fmt.Printf("Товар '%s' успешно добавлен. Новый ID: %d\n", name, newID)
    return newID
}
```

---

### 📉 `calculateDiscount`

> Рассчитывает скидку на товар в зависимости от условий.

```go
func calculateDiscount(price float64, percent float64) float64 {
    if percent < 0 || percent > 100 {
        fmt.Println("Ошибка: некорректное значение скидки.")
        return price
    }
    discount := price * percent / 100
    finalPrice := price - discount
    return finalPrice
}
```

---

## 🧾 Исходный код программы

```go
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
```

---

## 📎 Пример запуска

```bash
go run main.go
```

---
