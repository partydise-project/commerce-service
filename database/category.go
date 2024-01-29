package database

func CreateCategory(categoriaEvento *CategoriaEvento) error {
	result := DB.Create(categoriaEvento)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
