package app

import (
	"net/http"
	"strings"

	"github.com/gocopper/copper/cerrors"
	"github.com/gocopper/copper/chttp"
	"github.com/gocopper/copper/clogger"
	"github.com/isuquo/fortune3-pickaxe/pkg/goals"
)

type NewRouterParams struct {
	RW     *chttp.ReaderWriter
	Logger clogger.Logger
	Goal   *goals.Queries
}

func NewRouter(p NewRouterParams) *Router {
	return &Router{
		rw:     p.RW,
		logger: p.Logger,
		goal:   p.Goal,
	}
}

type Router struct {
	rw     *chttp.ReaderWriter
	logger clogger.Logger
	goal   *goals.Queries
}

func (ro *Router) Routes() []chttp.Route {
	return []chttp.Route{
		{
			Path:    "/goal",
			Methods: []string{http.MethodPost},
			Handler: ro.HandleGoalSubmit,
		},

		{
			Path:    "/goal",
			Methods: []string{http.MethodGet},
			Handler: ro.HandleGoalPage,
		},

		{
			Path:    "/",
			Methods: []string{http.MethodGet},
			Handler: ro.HandleIndexPage,
		},
	}
}

func (ro *Router) HandleIndexPage(w http.ResponseWriter, r *http.Request) {

	ro.rw.WriteHTML(w, r, chttp.WriteHTMLParams{
		PageTemplate: "index.html",
	})
}

func (ro *Router) HandleGoalPage(w http.ResponseWriter, r *http.Request) {

	ro.rw.WriteHTML(w, r, chttp.WriteHTMLParams{
		PageTemplate: "goal.html",
	})
}

func (ro *Router) HandleGoalSubmit(w http.ResponseWriter, r *http.Request) {
	var (
		name                = strings.TrimSpace(r.PostFormValue("name"))
		totalNeeded         = strings.TrimSpace(r.PostFormValue("total-needed"))
		currentSavings      = strings.TrimSpace(r.PostFormValue("current-savings"))
		monthlyContribution = strings.TrimSpace(r.PostFormValue("monthly-contribution"))
		monthsToGoal        = strings.TrimSpace(r.PostFormValue("months-to-goal"))
		startDate           = strings.TrimSpace(r.PostFormValue("start-date"))
		endDate             = strings.TrimSpace(r.PostFormValue("end-date"))
	)

	goal, err := goals.Validate(name, totalNeeded, currentSavings, monthlyContribution, startDate, endDate, monthsToGoal)
	if err != nil {
		ro.rw.WriteHTMLError(w, r, cerrors.New(err, "invalid goal data", map[string]interface{}{
			"form": r.Form,
		}))
		return
	}

	err = ro.goal.CreateGoal(r.Context(), &goal)
	if err != nil {
		ro.rw.WriteHTMLError(w, r, cerrors.New(err, "failed to create goal", map[string]interface{}{
			"form": r.Form,
		}))
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
