package main

import (
	"fmt"
	"log"
	"os"

	"github.com/magnusfarje/gosoc"
	"github.com/magnusfarje/gosoc/providers/facebook"
	"github.com/magnusfarje/gosoc/providers/google"
)

var (
	facebookToken = os.Getenv("FACEBOOK_ACCESS_TOKEN") // Facebook access token
	googleToken   = os.Getenv("GOOGLE_ID_TOKEN")       // Google id token
)

func main() {
	gosoc.AddProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), true),
		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), true))

	fmt.Println("Validate google id_token")
	googleProvider, err := gosoc.GetProvider("google")
	if err != nil {
		log.Fatal(err)
	}
	googleUser, err := googleProvider.ValidateToken(googleToken)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", googleUser)

	fmt.Println("\n\nValidate facebook access_token")
	facebookProvider, err := gosoc.GetProvider("facebook")
	if err != nil {
		log.Fatal(err)
	}
	facebookUser, err := facebookProvider.ValidateToken(facebookToken)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", facebookUser)

}
