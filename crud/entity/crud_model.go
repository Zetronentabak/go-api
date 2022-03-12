package entity

type GetDataResponse struct {
	Data []DataProduct `json:"product"`
}
type DataProduct struct {
	ProductID   string `json:"product_id,omitempty" bson:"product_id"`
	ProductName string `json:"product_name" bson:"product_name"`
}
type DataProductRequest struct {
	ProductID   string `json:"product_id,omitempty" bson:"product_id"`
	ProductName string `json:"product_name,omitempty" bson:"product_name"`
}
