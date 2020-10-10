package data

// User data model
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// User payload
type UserPayload struct {
	*User
	Role string `json:"role"`
}
