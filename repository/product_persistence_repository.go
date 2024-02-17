package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductPersistenceRepository struct {
	db *sqlx.DB
}

func NewProductPersistenceRepository(db *sqlx.DB) *ProductPersistenceRepository {
	return &ProductPersistenceRepository{db: db}
}

func (r *ProductPersistenceRepository) Create(product Product) (Product, error) {
	created := Product{}
	rows, err := r.db.NamedQuery(
		"INSERT INTO product (name, type, code, price) VALUES (:name, :type, :code, :price) RETURNING *", product,
	)
	if err != nil {
		return created, fmt.Errorf("ProductPersistenceRepository Create: %w", err)
	}
	for rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return created, fmt.Errorf("ProductPersistenceRepository Create: %w", err)
		}
	}
	return created, nil
}

func (r *ProductPersistenceRepository) Update(product Product) (Product, error) {
	updated := Product{}
	rows, err := r.db.NamedQuery(
		"UPDATE product SET name=:name, type=:type, code=:code, price=:price WHERE id=:id RETURNING *", product,
	)
	if err != nil {
		return updated, fmt.Errorf("ProductPersistenceRepository Update: %w", err)
	}
	for rows.Next() {
		if err := rows.StructScan(&updated); err != nil {
			return updated, fmt.Errorf("ProductPersistenceRepository Update: %w", err)
		}
	}
	return updated, nil
}

func (r *ProductPersistenceRepository) Delete(id int64) (bool, error) {
	_, err := r.db.Exec("DELETE FROM product WHERE id=$1", id)
	if err != nil {
		return false, fmt.Errorf("ProductPersistenceRepository Delete: %w", err)
	}
	return true, nil
}

func (r *ProductPersistenceRepository) GetAll() ([]Product, error) {
	var products []Product
	rows, err := r.db.Queryx("SELECT * FROM product")
	if err != nil {
		return products, fmt.Errorf("ProductPersistenceRepository GetAll: %w", err)
	}
	for rows.Next() {
		product := Product{}
		if err := rows.StructScan(&product); err != nil {
			return products, fmt.Errorf("ProductPersistenceRepository GetAll: %w", err)
		}
		products = append(products, product)
	}
	return products, nil
}
