package repository

import (
	"RGT/konis/lib"
	"RGT/konis/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func FindAllCarts(id int) ([]models.CartsJoin, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	// sql := `
	// 	select c.id, c.quantity, pv.id, p.id, u.id from carts c
	// 	join product_variants pv on pv.id = c.variant_id
	// 	join products p on p.id = c.product_id
	// 	join users u on u.id = c.user_id
	// 	where c.id = 1;`

	sql := `SELECT carts.id, carts.transaction_detail_id, carts.quantity, product_variants.name as variant, product_sizes.name as size, products.title, products.price ,image  FROM carts
			INNER JOIN product_variants ON carts.variant_id = product_variants.id
			INNER JOIN product_sizes ON carts.sizes_id = product_sizes.id
			INNER JOIN products on carts.product_id = products.id
			INNER JOIN product_images on product_images.product_id  = products.id
			WHERE carts.user_id = $1`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return []models.CartsJoin{}, err
	}

	rows, err := pgx.CollectRows(query, pgx.RowToStructByPos[models.CartsJoin])

	if err != nil {
		return []models.CartsJoin{}, err
	}
	fmt.Println(rows)

	return rows, err
}
func CreateCarts(data models.Carts) (models.Carts, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	ctx := context.Background()

	sqlTx := `
		INSERT INTO transaction_details ("quantity", "product_id", "variant_id", "product_size_id")
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var transactionDetailId int
	err := db.QueryRow(ctx, sqlTx, data.Quantity, data.ProductId, data.VariantId, data.ProductSizeId).Scan(&transactionDetailId)
	if err != nil {
		fmt.Println("Failed insert transaction_details:", err)
		return models.Carts{}, err
	}

	sqlCart := `
		INSERT INTO carts ("transaction_detail_id", "quantity", "variant_id", "sizes_id", "product_id", "user_id")
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING "id", "transaction_detail_id", "quantity", "variant_id", "sizes_id", "product_id", "user_id"
	`

	row, err := db.Query(ctx, sqlCart,
		transactionDetailId,
		data.Quantity,
		data.VariantId,
		data.ProductSizeId,
		data.ProductId,
		data.UserId,
	)
	if err != nil {
		fmt.Println("Gagal insert carts:", err)
		return models.Carts{}, err
	}

	cart, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Carts])
	if err != nil {
		fmt.Println("Gagal collect carts:", err)
		return models.Carts{}, err
	}

	fmt.Println("Inserted cart:", cart)
	return cart, nil
}

func GetCartsByUserId(id int) ([]models.Carts, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * from carts WHERE user_id = $1`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return []models.Carts{}, err
	}

	selectedRow, err := pgx.CollectRows(query, pgx.RowToStructByName[models.Carts])

	if err != nil {
		return []models.Carts{}, err
	}

	return selectedRow, err
}

func DeleteCarts(data models.Carts, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())
	fmt.Println(id)
	sql := `DELETE FROM carts WHERE user_id=$1`

	_, err := db.Exec(context.Background(), sql, id)

	if err != nil {
		return err
	}

	return nil
}
