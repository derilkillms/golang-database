package golang_database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id,name) VALUES('m','Muhammad')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)

	}
	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID: ", id)
		fmt.Println("Name: ", name)
	}

	defer rows.Close()
}

func TestQueryComplex(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, name, email, balance, rating, birth_date, merried, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id, name, email      string
			balance              int32
			rating               float64
			birthDate, createdAt time.Time
			married              bool
		)
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("=====================")
		fmt.Println("Id :", id, "Name :", name, "Email :", email, "Balance :", balance, "Rating :", rating, "Birth Date:", birthDate, "Married :", married, "Created At :", createdAt)
	}

}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	id := "md"
	name := "Muhammad Deril"

	query := "INSERT INTO customer(id, name) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, query, id, name)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "deril@gmail.com"
	comment := "Test komen"

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}

// Saat kita menggunakan Function Query atau EXEC yang menggunakan parameter, sebenarnya implementasi dibawahnya menggunakan Prepare Statement. Jadi tahapan pertama statement nya disiapkan terlebih dahulu, seteal itu baru di isi dengan parameter.

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	//prepare
	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "deril" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke" + strconv.Itoa(i)

		//execute query
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Commnet Id", id)
	}
}

// Secara default, semua perintah SQL yang kita kirim menggunakan golang akan otomatis di commit, atau istilahnya auto commit. Namun kita bisa menggunakan fitur transaksi sehingga SQL yang kita kirim tidak akan otomatis di commit ke database.

// Untuk memulai transaksi, kita bisa menggunakan function (DB) Begin(), dimana akan menghasilkan struct Tx yang merupakan representasi Transaction. Struct Tx ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana hampir semua function di DB ada di Tx, seperti Exec, Query, atau prepare. Setelah selesai melakukan transaksi, kita bisa menggunakan function (Tx) Commit() untuk melakukan commit atau Rollback().
func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	// emulai transaksi
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	// do transaction
	for i := 0; i < 10; i++ {
		email := "deril" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Commnet Id", id)
	}
	// save dengan melakukan commit
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
