package handler

import (
	"cafekalaa/api/app/model"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	uuid "github.com/google/uuid"
)

func AddToCart(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	cart := model.Cart{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cart); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	stmt, err := db.Prepare("INSERT INTO carts(id,user_id,product_id,quantity,price) VALUES($1,$2,$3,$4,$5);")

	if err != nil {
		panic(err)
	}

	uuid, err := uuid.NewUUID()

	res, err := stmt.Exec(uuid, cart.UserID, cart.ProductID, cart.Quantity, cart.Price)

	if err != nil && res != nil {
		panic(err)
	}

	respondJSON(w, http.StatusCreated, cart)

}

func GetCartPrice(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	//cart := model.Cart{}

	type Cart struct {
		UserID uuid.UUID `json:"userid"`
	}

	type Price struct {
		Price string
	}

	cart := Cart{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cart); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	rows, err := db.QueryContext(ctx, "select price from carts where user_id=$1", cart.UserID)

	if err != nil {
		panic(err)
	}

	cartPrice := make([]Price, 0)
	for rows.Next() {
		var price Price
		if err := rows.Scan(&price.Price); err != nil {
			log.Fatal(err)
		}

		cartPrice = append(cartPrice, price)
	}

	var allPrice int64
	for _, element := range cartPrice {

		i2, err := strconv.ParseInt(element.Price, 10, 64)

		if err != nil {
			panic(err)
		}

		allPrice += i2
	}

	respondJSON(w, http.StatusOK, allPrice)
}

func DeleteFromCart(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	type CartID struct {
		CartID uuid.UUID `json:"cartid"`
	}

	cartID := CartID{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cartID); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	err := db.QueryRowContext(ctx, "DELETE FROM carts WHERE id=$1;", cartID.CartID).Scan(&cartID.CartID)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusOK, nil)
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		respondJSON(w, http.StatusOK, nil)
	}

}
