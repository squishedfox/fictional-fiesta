package psql

import (
	"github.com/squishedfox/fictional-fiesta/db"
)

type (
	Repository struct {
	}
)

func NewRepository() *db.Repository {
	return &Repository{}
}
