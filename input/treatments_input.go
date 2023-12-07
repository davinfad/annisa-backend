package input

type InputTreatments struct {
	TreatmentName string `json:"treatment_name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Price         int    `json:"price" binding:"required"`
}