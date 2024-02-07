package database

func CreateCategory(categoriaEvento *CategoriaEvento) error {
	result := DB.Create(categoriaEvento)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ReadAllCategory() ([]CategoriaEvento, error) {
	var categories []CategoriaEvento
	if err := DB.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func ReadCategoriesPublished() ([]CategoriaEvento, error) {
	var categories []CategoriaEvento
	if err := DB.Where("publicada = ?", true).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func ReadCategoriesPublishedAdults() ([]CategoriaEvento, error) {
	var categories []CategoriaEvento
	if err := DB.Where("publicada = ? AND categoria_enfocada_para = ?", true, "adultos").Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func ReadCategoriesPublishedChildrens() ([]CategoriaEvento, error) {
	var categories []CategoriaEvento
	if err := DB.Where("publicada = ? AND categoria_enfocada_para = ?", true, "niños").Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
