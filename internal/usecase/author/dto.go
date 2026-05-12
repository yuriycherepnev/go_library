package author

type CreateAuthorDTO struct {
	Name string `json:"name"`
	Age  int    `json:"year"`
}

type UpdateAuthorDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Age  int    `json:"year"`
}
