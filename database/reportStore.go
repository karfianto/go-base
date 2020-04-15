package database

import (
	"github.com/dhax/go-base/models"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/go-pg/pg/urlvalues"
	"net/url"

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

// Create inserts the Reports to the database.
func (s *ReportStore) Create(p *models.Report) error {
	err := s.db.Insert(p)
	if err != nil {
		return err
	}
	return err
}


// ReportFilter provides pagination and filtering options on accounts.
type ReportFilter struct {
	Pager  *urlvalues.Pager
	Filter *urlvalues.Filter
	Order  []string
}

// NewReportFilter returns an ReportFilter with options parsed from request url values.
func NewReportFilter(params interface{}) (*ReportFilter, error) {
	v, ok := params.(url.Values)
	if !ok {
		return nil, ErrBadParams
	}
	p := urlvalues.Values(v)
	f := &ReportFilter{
		Pager:  urlvalues.NewPager(p),
		Filter: urlvalues.NewFilter(p),
		Order:  p["order"],
	}
	return f, nil
}

// Apply applies an ReportFilter on an orm.Query.
func (f *ReportFilter) Apply(q *orm.Query) (*orm.Query, error) {
	q = q.Apply(f.Pager.Pagination)
	q = q.Apply(f.Filter.Filters)
	q = q.Order(f.Order...)
	return q, nil
}

// List applies a filter and returns paginated array of matching results and total count.
func (s *ReportStore) List(f *ReportFilter) ([]models.Report, int, error) {
	a := []models.Report{}
	count, err := s.db.Model(&a).
		Apply(f.Apply).
		SelectAndCount()
	if err != nil {
		return nil, 0, err
	}
	return a, count, nil
}
