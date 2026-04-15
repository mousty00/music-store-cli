package payment

import (
	"bufio"
	"fmt"
	"go-starter/cart"
	"os"
	"strings"
)

type CreditCard struct {
	Number     string
	HolderName string
	Expiry     string
	CVV        string
}

func (c *CreditCard) Charge(amount float64) bool {
	if amount <= 0 {
		fmt.Println("Invalid amount.")
		return false
	}
	if len(c.Number) < 12 || len(c.CVV) != 3 {
		fmt.Println("Card validation failed.")
		return false
	}
	fmt.Printf("Charging $%.2f to card %s...\n", amount, maskCard(c.Number))
	return true
}

func maskCard(number string) string {
	if len(number) < 4 {
		return "****"
	}
	return "****-****-****-" + number[len(number)-4:]
}

func Checkout(c *cart.Cart, scanner *bufio.Scanner) {
	if c.IsEmpty() {
		fmt.Println("Cannot checkout: cart is empty.")
		return
	}

	fmt.Println("\nProceeding to checkout...")
	c.Display()

	fmt.Printf("\nFinal total: $%.2f\n", c.Total())
	fmt.Println("Please enter your credit card details to complete the purchase.")

	var card CreditCard
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Card Number (16 digits): ")
	number, _ := reader.ReadString('\n')
	card.Number = strings.TrimSpace(number)

	fmt.Print("Cardholder Name: ")
	name, _ := reader.ReadString('\n')
	card.HolderName = strings.TrimSpace(name)

	fmt.Print("Expiry (MM/YY): ")
	expiry, _ := reader.ReadString('\n')
	card.Expiry = strings.TrimSpace(expiry)

	fmt.Print("CVV (3 digits): ")
	cvv, _ := reader.ReadString('\n')
	card.CVV = strings.TrimSpace(cvv)

	if card.Charge(c.Total()) {
		fmt.Println("Purchase successful! Thank you for shopping.")
		c.Clear()
	} else {
		fmt.Println("Payment failed. Please try again.")
	}
}
