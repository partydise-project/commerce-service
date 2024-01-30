package database

func CreatePlanEvent(plan *PlanEvento) error {
	result := DB.Create(plan)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
