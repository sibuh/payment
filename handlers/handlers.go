package handlers

import (
	"log"
	"net/http"
	"os"

	"payment/db"
	"payment/model"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
)

func GetProducts(c *gin.Context) {

	products := db.GetAllProducts()

	c.JSON(http.StatusOK, gin.H{"products": products})

}

func CreateProduct(c *gin.Context) {
	var product model.Product
	c.BindJSON(&product)
	createdProduct := db.CreateProduct(product)

	c.JSON(http.StatusOK, createdProduct)

}

func PublishableKey(c *gin.Context) {

	err := godotenv.Load("secrets.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c.JSON(http.StatusOK, gin.H{
		"publishableKey": os.Getenv("STRIPE_PUBLISHABLE_KEY"),
	})
}

func HandleCreatePaymentIntent(c *gin.Context) {

	var product model.Product
	err := godotenv.Load("secrets.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("ShouldBindJSON: %v", err)
		return
	}

	data := db.GetAProduct(product.Id)

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(2000 + int64(data.Price)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	log.Printf("pi.New: %v", pi.ClientSecret)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("pi.New: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"clientSecret": pi.ClientSecret,
	})
}
