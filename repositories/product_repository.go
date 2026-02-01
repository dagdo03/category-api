package repositories

import (
	"categories-api/model"
	"database/sql"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAllProducts() ([]model.Product, error) {
	query := "SELECT id, name, price, stock, category_id FROM products"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := make([]model.Product, 0)
	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (repo *ProductRepository) CreateProduct(product *model.Product) error {
	query := `
        INSERT INTO products (name, price, stock, category_id) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id`

	err := repo.db.QueryRow(query,
		product.Name,
		product.Price,
		product.Stock,
		product.CategoryID,
	).Scan(&product.ID)

	if err != nil {
		return err
	}
	categoryQuery := "SELECT id, name, description FROM categories WHERE id = $1"
	var cat model.Category
	err = repo.db.QueryRow(categoryQuery, product.CategoryID).Scan(
		&cat.ID,
		&cat.Name,
		&cat.Description,
	)

	if err != nil {
		return err
	}

	product.Category = &cat
	return nil
}

func (repo *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	query := `
        SELECT 
            p.id, p.name, p.price, p.stock, p.category_id,
            c.id, c.name, c.description
        FROM products p
        LEFT JOIN categories c ON p.category_id = c.id
        WHERE p.id = $1`

	var p model.Product
	var c model.Category

	err := repo.db.QueryRow(query, id).Scan(
		&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID,
		&c.ID, &c.Name, &c.Description,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	if p.CategoryID != 0 {
		p.Category = &c
	}

	return &p, nil
}

func (repo *ProductRepository) UpdateProductById(product *model.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"

	result, err := repo.db.Exec(query,
		product.Name,
		product.Price,
		product.Stock,
		product.CategoryID,
		product.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (repo *ProductRepository) DeleteProductById(id int) error {
	query := "DELETE FROM products WHERE id = $1"

	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}
