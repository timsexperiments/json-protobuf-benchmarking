package json

type Person struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	Email   string   `json:"email,omitempty"`
	Hobbies []string `json:"hobbies,omitempty"`
}
