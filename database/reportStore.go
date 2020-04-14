package database

import (
	"github.com/dhax/go-base/models"
	"github.com/go-pg/pg"
	//"database/sql"
)

// ProfileStore implements database operations for profile management.
type ReportStore struct {
	db *pg.DB
}

// NewProfileStore returns a ProfileStore implementation.
func NewReportStore(db *pg.DB) *ReportStore {
	return &ReportStore{
		db: db,
	}
}

// Get gets an profile by account ID.
func (s *ReportStore) Get(accountID int) (*models.Report, error) {
	p := models.Report{AccountID: accountID}
	_, err := s.db.Model(&p).
		Where("account_id = ?", accountID).
		SelectOrInsert()

	return &p, err
}

// Create inserts the Reports to the database.
func (s *ReportStore) Create(p *models.Report) error {
	err := s.db.Insert(p)
	if err != nil {
		return err
	}
	return err
}