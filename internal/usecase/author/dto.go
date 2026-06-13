package author

type CreateAuthorDTO struct {
	Name string `json:"name"`
}

type UpdateAuthorDTO struct {
	Id   int    `json:"uuid"`
	Name string `json:"name"`
}

type DeleteAuthorDto struct {
	Id int `json:"uuid"`
}
