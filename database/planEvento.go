package database

func CreatePlanEvent(plan *PlanEvento) error {
	result := DB.Create(plan)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ReadPlansEventPublished(idCategory string) ([]PlanEvento, error) {
	var plans []PlanEvento
	if err := DB.Where("categoria_evento_id = ? AND estado = ?", idCategory, true).Find(&plans).Error; err != nil {
		return nil, err
	}

	return plans, nil
}
