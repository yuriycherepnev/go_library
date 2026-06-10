package author

type CreateAuthorDTO struct {
	Name string `json:"name"`
}

type UpdateAuthorDTO struct {
	Id   string `json:"uuid"`
	Name string `json:"name"`
}

type DeleteAuthorDto struct {
	Id string `json:"uuid"`
}
