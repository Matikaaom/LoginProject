// package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	r := gin.Default()
// 	r.GET("/aom", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "aom",
// 		})
// 	})
// 	r.Run("localhost:8080")
// }

// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"net/http"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	_ "github.com/go-sql-driver/mysql"
// 	"golang.org/x/crypto/bcrypt"
// )

// var jwtKey = []byte("your_secret_key")

// type Credentials struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	var creds Credentials
// 	json.NewDecoder(r.Body).Decode(&creds)

// 	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
// 	if err != nil {
// 		http.Error(w, "Database connection error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()

// 	var storedPassword string
// 	err = db.QueryRow("SELECT password FROM users WHERE username = ?", creds.Username).Scan(&storedPassword)
// 	if err != nil {
// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	if err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(creds.Password)); err != nil {
// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	expirationTime := time.Now().Add(24 * time.Hour)
// 	claims := &Claims{
// 		Username: creds.Username,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		http.Error(w, "Error generating token", http.StatusInternalServerError)
// 		return
// 	}

// 	http.SetCookie(w, &http.Cookie{
// 		Name:     "token",
// 		Value:    tokenString,
// 		Expires:  expirationTime,
// 		HttpOnly: true,
// 	})
// }

// func main() {
// 	http.HandleFunc("/login", loginHandler)
// 	http.ListenAndServe(":8080", nil)
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"main/controller"
// 	"main/data"
// 	"main/docs"
// 	"main/middleware/login"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// 	httpSwagger "github.com/swaggo/http-swagger"
// )

// const (
// 	port = 8082
// )

// @title JWT Login API
// @version 1.0.0
// @contact.name Junio Cesar Ferreira
// @license.name MIT
// @BasePath /
// func main() {
// 	fmt.Println("initializing go server")

// 	loginMethod := login.NewJwt()
// 	respository := data.NewDummy()
// 	controller := controller.New(loginMethod, respository)

// 	// Inicialização do roteador do Gorilla Mux
// 	router := mux.NewRouter()

// 	router.HandleFunc("/login", controller.Login).Methods("POST", "OPTIONS")
// 	router.HandleFunc("/protected", loginMethod.Authenticate(controller.ProtectedEndpoint)).Methods("GET", "OPTIONS")

// 	//Documentação da API
// 	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

// 	fmt.Printf("listening port %d\n", port)

// 	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
// }
