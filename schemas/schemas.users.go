package schemas

type SchemaUser struct {
	ID string `json:"id" validate:"uuid"`
	// FirstName string `json:"first_name" validate:"required,lowercase"`
	// LastName  string `json:"last_name" validate:"required,lowercase"`
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	// Role      string `json:"role" validate:"required,lowercase"`
}

type ResponStatusDataViewUser struct {
	Page        uint32      `json:"page"`
	Per_page    uint32      `json:"per_page"`
	Total_pages uint32      `json:"total_pages"`
	Total       uint32      `json:"total"`
	Data        interface{} `json:"data"`
}

type SchemaTotal struct {
	Total int32 `json:"total"`
}
