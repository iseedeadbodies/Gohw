package main

import "fmt"

type Product struct {
	Name     string
	price    float64
	Quantity int
}

func newProduct(name string, price float64, quantity int) Product {
	if price < 0 {
		price = 0
	}

	if quantity < 0 {
		quantity = 0
	}

	return Product{
		Name:     name,
		price:    price,
		Quantity: quantity,
	}
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p *Product) SetPrice(newPrice float64) {
	if newPrice < 0 {
		fmt.Println("Error: price cannot be negative")
		return
	}

	p.price = newPrice
}

func (p *Product) Buy(amount int) {
	if amount <= 0 {
		fmt.Println("Error: amount must be positive")
		return
	}

	if amount > p.Quantity {
		fmt.Println("Not enough product")
		return
	}

	p.Quantity -= amount
	fmt.Println("Bought:", amount)
}

func (p *Product) Restock(amount int) {
	if amount <= 0 {
		fmt.Println("Error: amount must be positive")
		return
	}

	p.Quantity += amount
}

func (p Product) Info() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Price:", p.price)
	fmt.Println("Quantity:", p.Quantity)
}

func main() {
	product := newProduct("Phone", 300000, 5)

	product.Info()

	fmt.Println()

	product.Buy(2)
	product.Buy(10)

	fmt.Println()

	product.SetPrice(280000)
	product.Restock(7)

	product.Info()
}
