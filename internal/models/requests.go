package models

type AddRequestInt struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

type AddRequestFloat struct {
	A *float64 `json:"a"`
	B *float64 `json:"b"`
}

type SubtractRequestInt struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

type SubtractRequestFloat struct {
	A *float64 `json:"a"`
	B *float64 `json:"b"`
}

type MultiplyRequestInt struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

type MultiplyRequestFloat struct {
	A *float64 `json:"a"`
	B *float64 `json:"b"`
}

type DivideRequestInt struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

type DivideRequestFloat struct {
	A *float64 `json:"a"`
	B *float64 `json:"b"`
}
