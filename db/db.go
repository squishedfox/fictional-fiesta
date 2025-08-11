package db

type (
	FormsRepositoryContextKeyType string
	// Repository is the object to manage form entities and related entities
	FormsRepository interface {
		// GetForms filters and fetches forms
		GetForms(model *GetFormsModel) (*FormsModel, error)
		// CreateForm appends a new form to the database and returns the new identifier
		CreateForm(model *CreateFormModel) (any, error)
	}
)

var (
	FormsRepositoryContextKey FormsRepositoryContextKeyType = "FormsRepository"
)
