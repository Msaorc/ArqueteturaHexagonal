package db

import (
	"database/sql"

	"github.com/Msaorc/ArqueteturaHexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func (p *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDB) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`insert into products(id, name, price, status) values (?,?,?,?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetId(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDB) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := p.db.Exec("update products set name = ?, price = ?, status = ? where id = ?",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDB) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow("Select id from products where id=?", product.GetId()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
		return product, nil
	}

	_, err := p.update(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func NewProductDb(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}
