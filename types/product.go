package types

type Product struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Timestamp string `json:"timestamp"`
}
