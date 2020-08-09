type User struct {
	FirstName, LastName string
	Location            *UserLocation
}

type UserLocation struct {
	City    string
	Country string
}

// List of functions
func NewUser(firstName, lastName string) *User {
	return &User{FirstName: firstName,
		LastName: lastName,
		Location: &UserLocation{
			City:    "Santa Monica",
			Country: "USA",
		},
	}
}

// List of methods
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
  us:=User {FirstName: "Matt",
		LastName: "Damon",
		Location: &UserLocation{
			City:    "Santa Monica",
			Country: "USA",}}
  fmt.Println(us.Greeting())
}
