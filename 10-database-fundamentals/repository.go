package main

import (
	"errors"
	"fmt"
)

// Product, veritabanı modelidir.
type Product struct {
	ID    int
	Name  string
	Price float64
}

// ProductRepository, veri erişim katmanı arayüzüdür.
type ProductRepository interface {
	Save(p *Product) error
	GetByID(id int) (*Product, error)
	GetAll() []Product
}

// MemoryDB, Repository arayüzünü uygulayan bir in-memory veritabanıdır.
type MemoryDB struct {
	data map[int]Product
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{data: make(map[int]Product)}
}

func (db *MemoryDB) Save(p *Product) error {
	db.data[p.ID] = *p
	return nil
}

func (db *MemoryDB) GetByID(id int) (*Product, error) {
	p, ok := db.data[id]
	if !ok {
		return nil, errors.New("ürün bulunamadı")
	}
	return &p, nil
}

func (db *MemoryDB) GetAll() []Product {
	products := make([]Product, 0, len(db.data))
	for _, p := range db.data {
		products = append(products, p)
	}
	return products
}

func main() {
	repo := NewMemoryDB()

	// Veri kaydetme
	repo.Save(&Product{ID: 1, Name: "MacBook Pro", Price: 50000.0})
	repo.Save(&Product{ID: 2, Name: "iPhone 15", Price: 70000.0})

	// Veri çekme
	products := repo.GetAll()
	fmt.Println("Kayıtlı Ürünler:")
	for _, p := range products {
		fmt.Printf("- %s: %.2f TL\n", p.Name, p.Price)
	}
}
