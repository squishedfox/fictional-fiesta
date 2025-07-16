package db

type (
	// Repository is the object to manage form entities and related entities
	Repository interface {
		// GetForms filters and fetches forms
		GetForms(model *GetFormsModel) (FormsModel, error)
	}
)
