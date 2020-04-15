package app

import (
	"errors"
	"github.com/dhax/go-base/auth/jwt"

	//"github.com/dhax/go-base/auth/pwdless"
	"github.com/dhax/go-base/database"
	"net/http"

	"github.com/dhax/go-base/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
)

// The list of error types returned from account resource.
var (
	ErrReportValidation = errors.New("report validation error")
)

// ReportStore defines database operations for a report.
type ReportStore interface {
	//List(accountID int) ([]models.Report, error)
	List(filter *database.ReportFilter) ([]models.Report, int, error)
	Create(*models.Report) error
}

// ReportResource implements report management handler.
type ReportResource struct {
	Store ReportStore
}

// NewReportResource creates and returns a report resource.
func NewReportResource(store ReportStore) *ReportResource {
	return &ReportResource{
		Store: store,
	}
}

func (rs *ReportResource) router() *chi.Mux {
	r := chi.NewRouter()
	//r.Use(rs.reportCtx)
	r.Get("/", rs.list)
	r.Post("/", rs.create)
	return r
}
/*
func (rs *ReportResource) reportCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//claims := jwt.ClaimsFromCtx(r.Context())
		f, err := database.NewReportFilter(r.URL.Query())
		p, count, err := rs.Store.List(f)
		if err != nil {
			log(r).WithField("reportCtx", claims.Sub).Error(err)
			render.Render(w, r, ErrInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), ctxReport, p)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
*/
type reportRequest struct {
	*models.Report
	ProtectedID int `json:"id"`
}

func (d *reportRequest) Bind(r *http.Request) error {
	return nil
}

type reportResponse struct {
	Reports []models.Report `json:"reports"`
	Count    int               `json:"count"`
}


func newReportResponse(a []models.Report, count int) *reportResponse {
	resp := &reportResponse{
		Reports: a,
		Count: count,
	}
	return resp
}

func (rs *ReportResource) list(w http.ResponseWriter, r *http.Request) {
	claims := jwt.ClaimsFromCtx(r.Context())

	f, err := database.NewReportFilter(claims)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	al, count, err := rs.Store.List(f)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, newReportResponse(al,count))
}

func (rs *ReportResource) create(w http.ResponseWriter, r *http.Request) {
	data := &reportRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	}

	if err := rs.Store.Create(data.Report); err != nil {
		switch err.(type) {
		case validation.Errors:
			render.Render(w, r, ErrValidation(ErrReportValidation, err.(validation.Errors)))
			return
		}
		render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, http.StatusOK)
}
/*
func (rs *ReportResource) update(w http.ResponseWriter, r *http.Request) {
	p := r.Context().Value(ctxReport).(*models.Reports)
	data := &reportRequest{Reports: p}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	}

	if err := rs.Store.Update(p); err != nil {
		switch err.(type) {
		case validation.Errors:
			render.Render(w, r, ErrValidation(ErrReportValidation, err.(validation.Errors)))
			return
		}
		render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, newReportResponse(p))
}
*/