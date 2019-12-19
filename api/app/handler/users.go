package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	uuid "github.com/google/uuid"

	"cafekalaa/api/app/model"
	"cafekalaa/api/utils"
)

func GetAllUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	user := model.User{}
	reqToken := r.Header.Get("Authorization")
	userID, okExtractUUID := utils.ExtractClaims(reqToken)

	if !okExtractUUID {
		fmt.Println("Can't get UUID for Get All Users section")
		respondError(w, http.StatusUnauthorized, "Token Invalid")
		return
	}

	userUUID, errUUID := uuid.Parse(userID["id"].(string))
	if errUUID != nil {
		panic(errUUID)
	}

	err := db.QueryRowContext(ctx, "SELECT id FROM users WHERE id=$1", userUUID).Scan(&user.Id)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusNotFound, "No UserId exist")
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:

		rows, err := db.QueryContext(ctx, "select * from users")
		if err != nil {
			panic(err)
		}
		users := make([]model.User, 0)
		for rows.Next() {
			var user model.User
			if err := rows.Scan(&user.Name, &user.Avatar, &user.Mobile, &user.BirthDay, &user.Identify, &user.Cart, &user.Credit, &user.Password, &user.Type, &user.Email, &user.Id); err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
		respondJSON(w, http.StatusOK, users)
	}
}

func RegisterUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	stmt, err := db.Prepare("INSERT INTO users(id,name,avatar,mobile, email,birth_day, identify, cart, credit, password, type) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);")

	if err != nil {
		panic(err)
	}

	uuid, err := uuid.NewUUID()

	res, err := stmt.Exec(uuid, user.Name, user.Avatar, user.Mobile, user.Email, user.BirthDay, user.Identify, user.Cart, user.Credit, user.Password, user.Type)

	if err != nil && res != nil {
		panic(err)
	}

	respondJSON(w, http.StatusCreated, user)
}

func FindUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	ctx := context.Background()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := db.QueryRowContext(ctx, "SELECT * FROM users WHERE name=$1", user.Name).Scan(
		&user.Name, &user.Avatar, &user.Mobile, &user.Email, &user.BirthDay, &user.Identify, &user.Cart, &user.Credit, &user.Password, &user.Type, &user.Id)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusOK, nil)
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created on %s\n", user.Name, user.Password)
		respondJSON(w, http.StatusOK, user)
	}
}

func FindUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	rows, err := db.QueryContext(ctx, "select * from users where name=$1", user.Name)

	if err != nil {
		panic(err)
	}

	users := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Name, &user.Avatar, &user.Mobile, &user.BirthDay, &user.Identify, &user.Cart, &user.Credit, &user.Password, &user.Type, &user.Email, &user.Id); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	respondJSON(w, http.StatusOK, users)
}

func UpdateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := db.QueryRowContext(ctx, "UPDATE users SET name=$1,avatar=$2,mobile=$3, email=$4,birth_day=$5, identify=$6, cart=$7, credit=$8, password=$9, type=$10 WHERE name=$11;",
		user.Name, user.Avatar, user.Mobile, user.Email, user.BirthDay, user.Identify, user.Cart, user.Credit, user.Password, user.Type, user.Name).Scan(
		&user.Name, &user.Avatar, &user.Mobile, &user.Email, &user.BirthDay, &user.Identify, &user.Cart, &user.Credit, &user.Password, &user.Type, &user.Id)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusOK, user)
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created on %s\n", user.Name, user.Password)
		respondJSON(w, http.StatusOK, user)
	}
}

func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := db.QueryRowContext(ctx, "DELETE FROM users WHERE name=$1;", user.Name).Scan(&user.Name)

	switch {
	case err == sql.ErrNoRows:
		respondJSON(w, http.StatusOK, nil)
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created on %s\n", user.Name, user.Password)
		respondJSON(w, http.StatusOK, nil)
	}

}

//LoginUser
func LoginUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	type Credentials struct {
		Id       uuid.UUID `json:"-"`
		Mobile   string    `json:"mobile", db:"mobile"`
		Password string    `json:"password", db:"password"`
	}

	type UserToken struct {
		Token string `json:"token"`
	}

	user := Credentials{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	userPassword := user.Password

	err := db.QueryRowContext(ctx, "SELECT id, mobile, password FROM users WHERE mobile=$1", user.Mobile).Scan(
		&user.Id, &user.Mobile, &user.Password)

	switch {
	case err == sql.ErrNoRows:

		respondError(w, http.StatusNotFound, err.Error())
		return
	case err != nil:

		log.Fatalf("query error: %v\n", err)
	default:

		match := utils.CheckPasswordHash(userPassword, user.Password)

		if !match {

			respondError(w, http.StatusBadRequest, "wrong password")
			return
		}

		userToken := &UserToken{}
		token := utils.MakeTokenFromUUID(user.Id)
		userToken.Token = token
		respondJSON(w, http.StatusOK, userToken)

	}

}

func SendSmsVerfication(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	mobile := user.Mobile
	password := user.Password

	if len(mobile) == 11 && strings.HasPrefix(mobile, "09") {
		if len(password) >= 6 {
			respondJSON(w, http.StatusOK, "کد فعال سازی به شماره وارد شده ارسال شد")
			userCode := SendVerficationSms(mobile)

			stmt, err := db.Prepare("INSERT INTO user_auth(mobile, password, code) VALUES($1,$2,$3);")

			if err != nil {
				panic(err)
			}

			userPass, _ := utils.HashPassword(user.Password)

			res, err := stmt.Exec(user.Mobile, userPass, userCode)

			if err != nil && res != nil {
				panic(err)
			}

			respondJSON(w, http.StatusCreated, user)

		} else {
			respondError(w, http.StatusNotFound, "رمز عبور حداقل باید ۶ کاراکتر باشد")
		}
	} else {
		respondError(w, http.StatusNotFound, "لطفا شماره موبایل را به درستی وارد نمایید")
	}

}

func GetOtpFromUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	type UserData struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		Code     string `json:"code"`
	}

	userData := UserData{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userData); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	err := db.QueryRowContext(ctx, "SELECT * FROM user_auth WHERE mobile=$1 And code=$2",
		userData.Mobile, userData.Code).Scan(
		&userData.Mobile, &userData.Password, &userData.Code)

	switch {
	case err == sql.ErrNoRows:
		respondError(w, http.StatusNotFound, err.Error())
		return
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:

		uuid, _ := uuid.NewUUID()

		stmt, err := db.Prepare("INSERT INTO users(name,avatar,mobile,birth_day,identify,cart,credit,password,type,email,id) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);")

		if err != nil {
			panic(err)
		}

		res, err := stmt.Exec("", "", userData.Mobile, "", "", "", "", userData.Password, "", "", uuid)

		if err != nil && res != nil {
			panic(err)
		}
		respondJSON(w, http.StatusOK, userData)
	}
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {

	type UserToken struct {
		Token string `json:"token"`
	}

	reqToken := r.Header.Get("Authorization")

	userID, okExtractUUID := utils.ExtractClaimsForRefresh(reqToken)

	if !okExtractUUID {
		fmt.Println("Can't get UUID from token in refresh section")
		return
	}

	userUUID, errUUID := uuid.Parse(userID["id"].(string))
	if errUUID != nil {
		panic(errUUID)
	}

	token := utils.MakeTokenFromUUID(userUUID)

	userToken := &UserToken{}
	userToken.Token = token
	respondJSON(w, http.StatusOK, userToken)
}
