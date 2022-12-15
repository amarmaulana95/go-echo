package schemas

type SchemaProduct struct {
	ID          string `json:"id" validate:"uuid"`
	Name        string `json:"name"`
	Price       string `json:"price" validate:"required,price"`
	Description string `json:"descrription"`
	Tags        string `json:"tags"`
	Category_id string `json:"category_id"`
	// Role      string `json:"role" validate:"required,lowercase"`
}

type ResponStatusDataView struct {
	Page        uint32      `json:"page"`
	Per_page    uint32      `json:"per_page"`
	Total_pages uint32      `json:"total_pages"`
	Total       uint32      `json:"total"`
	Data        interface{} `json:"data"`
}
