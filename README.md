gosoc - Social token handler
===
Go package for managing Social sign-in access tokens.   
Used to validate, and get user profile data from token.



---

### Currently supported providers
* Facebook
* Google

---

### Install
```sh
go get github.com/Magnusfarje/gosoc
```
---

### Example
```go
func main() {
    // Add provider(s) with key, secret and bool to validate token origin is your "app"
	gosoc.AddProviders(google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), true))
    
    // Get provider
	googleProvider, _ := gosoc.GetProvider("google")
    
    // Validate token and get user profile
	googleUser, _ := googleProvider.ValidateToken(os.Getenv("GOOGLE_ID_TOKEN"))

	fmt.Printf("%+v", googleUser)
}
```
```
{
	ID:104556323466866496426 
	Mail:anders@andersson.se 
	Provider:google 
	Picture:https://lh3.googleusercontent.com/-5FJS2A2GDl4/AAAAAAAAAAI/AAAAAAAAAAA/APaXHhRansdfsdYnHGiMaKY-7zkP5gWH8Q/s12-c/photo.jpg 
	FirstName:Anders 
	LastName:Andersson 
	Expire:2016-09-29 22:19:52 +0200 CEST
}
```
See [example](https://gitlab.com/magnusfarje/gosoc/tree/master/gosoc_example)

### Custom providers
Implement the `Provider` interface to add providers. 



