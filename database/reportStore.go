package database

import (
	"github.com/dhax/go-base/models"
	"github.com/go-pg/pg"
	//"database/sql"
)

// ReportStore implements database operations for profile management.
type ReportStore struct {
	db *pg.DB
}

// NewReportStore returns a ReportStore implementation.
func NewReportStore(db *pg.DB) *ReportStore {
	return &ReportStore{
		db: db,
	}
}

// Get gets a report by account ID.
func (s *ReportStore) Get(accountID int) (*models.Report, error) {
	p := models.Report{AccountID: accountID}
	err := s.db.Model(&p).
		Where("account_id = ?", accountID).
		First()

	return &p, err
}

// Update report .
func (s *ReportStore) Update(p *models.Report) error {
	err := s.db.Update(p)
	return err
}

// Create inserts the Reports to the database.
func (s *ReportStore) Create(p *models.Report) error {
	return s.db.Insert(p)
}
