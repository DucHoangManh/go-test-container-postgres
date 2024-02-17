package repository

type Product struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Type  string `db:"type"`
	Code  string `db:"code"`
	Price int64  `db:"price"`
}
