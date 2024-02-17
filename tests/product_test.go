package tests

import (
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/stretchr/testify/require"
	"go-test-container-postgres/repository"
)

func Test_Product(t *testing.T) {
	t.Parallel()
	t.Run(
		"test create product", func(t *testing.T) {
			t.Parallel()
			db := NewDatabase(t)
			repo := repository.NewProductPersistenceRepository(db)
			created, err := repo.Create(
				repository.Product{
					Name:  "cake",
					Type:  "food",
					Code:  "c1",
					Price: 100,
				},
			)
			require.Nil(t, err)
			require.Equal(t, "cake", created.Name)
		},
	)

	t.Run(
		"test get all products", func(t *testing.T) {
			t.Parallel()
			db := NewDatabase(t)
			repo := repository.NewProductPersistenceRepository(db)
			_, err := repo.Create(
				repository.Product{
					Name:  "cake",
					Type:  "food",
					Code:  "c1",
					Price: 100,
				},
			)
			require.Nil(t, err)
			products, err := repo.GetAll()
			require.Nil(t, err)
			require.Len(t, products, 1)
		},
	)

	t.Run(
		"test update product", func(t *testing.T) {
			t.Parallel()
			db := NewDatabase(t)
			repo := repository.NewProductPersistenceRepository(db)
			created, err := repo.Create(
				repository.Product{
					Name:  "cake",
					Type:  "food",
					Code:  "c1",
					Price: 100,
				},
			)
			require.Nil(t, err)
			created.Name = "new cake"
			updated, err := repo.Update(created)
			require.Nil(t, err)
			require.Equal(t, "new cake", updated.Name)
		},
	)
}
