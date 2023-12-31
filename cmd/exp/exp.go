package main

import (
	"fmt"

	"github.com/sunnymotiani/lenslocked/models"
)

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("connected!")

	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("sun.motiani4@gmail.com", "sunny123")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	/*
		//creating table
		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS users(
				id SERIAL PRIMARY KEY,
				name TEXT,
				email TEXT UNIQUE NOT NULL
			);

			CREATE TABLE IF NOT EXISTS orders(
				id SERIAL PRIMARY KEY,
				user_id INT NOT NULL,
				amount INT,
				description TEXT
			);
		`)
		if err != nil {
			panic(err)
		}
		fmt.Println("Tables Created.")

		//INSERT SOME DATA

			name := "Jhon Calhoun"
			email := "jon3@calhoun.io"
			row := db.QueryRow(`
					INSERT INTO users (name,email)
					VALUES ($1,$2) RETURNING id;`, name, email)
			var id int
			err = row.Scan(&id)
			if err != nil {
				panic(err)
			}
			fmt.Println("User Created.id =", id)

			id = 1
			row = db.QueryRow(`
				SELECT name,email
				FROM users
				WHERE id=$1;`, id)
			var name, email string
			err = row.Scan(&name, &email)
			if err != sql.ErrNoRows {
				fmt.Printf("Error , No rows!")
			}
			if err != nil {
				panic(err)
			}
			fmt.Printf("USer information: name=%s , email =%s \n", name, email)

			userdID := 1
			for i := 1; i <= 5; i++ {
				amount := i * 100
				desc := fmt.Sprintf("fake order #%d", i)
				_, err := db.Exec(`
					INSERT INTO orders(user_id,amount,description)
					VALUES ($1,$2,$3);`, userdID, amount, desc)
				if err != nil {
					panic(err)
				}
			}
			fmt.Println("Created fake orders")


		type Order struct {
			ID          int
			UserdID     int
			Amount      int
			Description string
		}
		var orders []Order
		userID := 1
		rows, err := db.Query(`
			SELECT id,amount,description
			FROM orders
			WHERE user_id=$1;`, userID)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var order Order
			order.UserdID = userID
			err := rows.Scan(&order.ID, &order.Amount, &order.Description)
			if err != nil {
				panic(err)
			}
			orders = append(orders, order)
		}
		err = rows.Err()
		if err != nil {
			panic(err)
		}
		//CHECK FOR AN ERROR
		fmt.Println("Orders :", orders)
	*/
}
