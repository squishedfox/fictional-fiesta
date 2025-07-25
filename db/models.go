package db

type (
	// Operation is a database operation that can be applied to a value or column when filtering data
	Operation string
	// SearchFieldFilter provides context to the field and value which should be compared or used and the operation that should be applied to them
	SearchFieldFilter struct {
		Operation Operation
		// Field is the name of the field to apply the operation on
		Field string
		// Value should be the value to compare the field value for when searching
		Value any
	}

	// FormModel is the singular form object
	FormModel struct {
		// ID is the unique identifier of the form
		ID string
		// Name is the user friendly name of the form
		Name string
	}

	// FormsModel is the model returned when there is a search being performed for one or many different forms
	FormsModel struct {
		Forms []*FormModel
	}

	// GetFormsModel is a request object for fetching and filtering form data
	GetFormsModel struct {
		// Name (Optional) is the form name to search for
		Filters []*SearchFieldFilter
		// Limit should be the number of recrods that should be limited to the results for returning
		Limit int
		// Skip is the number of records to skip over before fetching more records
		Skip int
		// OrderBy is a highly recommended field to ensure your sorting occurs as it should
		OrderBy []string
	}

	CreateFormModel struct {
		// Name is the user friendly name of the form
		Name string
	}
)

var (
	EqualsOperation      Operation = "equals"
	NotEqualsOperation   Operation = "notEquals"
	LessThanOperation    Operation = "lessThan"
	GreaterThanOperation Operation = "greaterThan"
	StartsWithOperation  Operation = "startsWith"
	EndsWithOperation    Operation = "endsWith"
	containsOperation    Operation = "contains"
	notContainsOperation Operation = "doesNotContain"
)
