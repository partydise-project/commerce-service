package database

func CreatePlanEvent(plan *PlanEvento) error {
	result := DB.Create(plan)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ReadPlanEventPublished(idCategory uint) ([]PlanEvento, error) {
	var plans []PlanEvento
	if err := DB.Where("categoria_id = ? AND estado = ?", idCategory, true).Find(&plans).Error; err != nil {
		return nil, err
	}

	return plans, nil
}
