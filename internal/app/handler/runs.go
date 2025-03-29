package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/nilhiu/srleaderboard/internal/app/ajax"
	"github.com/nilhiu/srleaderboard/internal/app/htmx"
	"github.com/nilhiu/srleaderboard/internal/app/middleware"
	"github.com/nilhiu/srleaderboard/internal/service/leaderboard"
	"github.com/nilhiu/srleaderboard/internal/service/user"
	"github.com/nilhiu/srleaderboard/internal/view/component"
)

// RunsHandler provides methods to handle HTTP requests about runs.
type RunsHandler struct {
	ctx            context.Context
	userSvc        user.Service
	leaderboardSvc leaderboard.Service
}

func NewRunsHandler(
	ctx context.Context,
	userSvc user.Service,
	leaderboardSvc leaderboard.Service,
) *RunsHandler {
	return &RunsHandler{
		ctx:            ctx,
		userSvc:        userSvc,
		leaderboardSvc: leaderboardSvc,
	}
}

// GetRuns is a handler, which returns runs from the leaderboard database based
// on the query parameters: `offset` and `limit`. Supports both AJAX and HTMX
// requests (`Hx-Request`).
func (h *RunsHandler) GetRuns() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		offset, err := strconv.Atoi(q.Get("offset"))
		if err != nil {
			offset = 0
		}

		limit, err := strconv.Atoi(q.Get("limit"))
		if err != nil {
			limit = 5
		}

		if offset < 0 || limit > 100 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Header.Get("Hx-Request") != "true" {
			h.getRunsAJAX(offset, limit).ServeHTTP(w, r)
		} else {
			h.getRunsHTMX(offset, limit).ServeHTTP(w, r)
		}
	}
}

// GetUserRuns is a handler, which returns runs from the user database based
// on the path value `user` and query parameters: `offset` and `limit`. Supports
// both AJAX and HTMX requests (`Hx-Request`).
func (h *RunsHandler) GetUserRuns() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.PathValue("user")
		q := r.URL.Query()
		offset, err := strconv.Atoi(q.Get("offset"))
		if err != nil {
			offset = 0
		}

		limit, err := strconv.Atoi(q.Get("limit"))
		if err != nil {
			limit = 5
		}

		if offset < 0 || limit > 100 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Header.Get("Hx-Request") != "true" {
			h.getUserRunsAJAX(user, offset, limit).ServeHTTP(w, r)
		} else {
			h.getUserRunsHTMX(user, offset, limit).ServeHTTP(w, r)
		}
	}
}

// AddRun is a handler, which adds a run to both the user and leaderboard
// databases. This route requires the user to be authenticated, a.k.a. the
// request context must have the value for [middleware.ContextValueKeyUser].
// Supports both AJAX and HTMX requests (`Hx-Request`).
func (h *RunsHandler) AddRun() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ajax.AddRunRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, ok := r.Context().Value(middleware.ContextValueKeyUser).(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if r.Header.Get("Hx-Request") != "true" {
			h.addRunAJAX(req, user).ServeHTTP(w, r)
		} else {
			h.addRunHTMX(req, user).ServeHTTP(w, r)
		}
	}
}

func (h *RunsHandler) getRunsAJAX(offset, limit int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.getRunsDB(offset, limit)
		if err != nil {
			WriteErrorHeader(w, err)
			return
		}

		ajax.MustEncode(w, res)
	}
}

func (h *RunsHandler) getRunsHTMX(offset, limit int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.getRunsDB(offset, limit)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		curPage := int(math.Ceil(float64(offset)/5)) + 1
		pages := int(math.Ceil(float64(res.FullAmount) / 5))
		htmx.MustRender(h.ctx, w, component.Leaderboard(res.Runs, curPage, pages))
	}
}

func (h *RunsHandler) getRunsDB(offset, limit int) (ajax.GetRunsResponse, error) {
	runs, err := h.leaderboardSvc.GetRuns(offset, limit)
	if err != nil {
		return ajax.GetRunsResponse{}, HTTPErrInternalServerError
	}

	fullAmount, err := h.leaderboardSvc.GetRunCount()
	if err != nil {
		return ajax.GetRunsResponse{}, HTTPErrInternalServerError
	}

	return ajax.GetRunsResponse{
		Runs:       runs,
		Amount:     len(runs),
		Offset:     offset,
		Limit:      limit,
		FullAmount: fullAmount,
	}, nil
}

func (h *RunsHandler) getUserRunsAJAX(user string, offset, limit int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.getUserRunsDB(user, offset, limit)
		if err != nil {
			WriteErrorHeader(w, err)
			return
		}

		ajax.MustEncode(w, res)
	}
}

func (h *RunsHandler) getUserRunsHTMX(user string, offset, limit int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.getUserRunsDB(user, offset, limit)
		if err != nil {
			WriteErrorHeader(w, err)
			return
		}

		curPage := int(math.Ceil(float64(offset)/5)) + 1
		pages := int(math.Ceil(float64(res.FullAmount) / 5))
		htmx.MustRender(h.ctx, w, component.UserBoard(user, res.Runs, curPage, pages))
	}
}

func (h *RunsHandler) getUserRunsDB(
	user string,
	offset, limit int,
) (ajax.GetUserRunsResponse, error) {
	userID, err := h.userSvc.GetUserID(user)
	if err != nil {
		return ajax.GetUserRunsResponse{}, HTTPErrNotFound
	}

	runs, err := h.userSvc.GetUserRuns(userID, offset, limit)
	if err != nil {
		return ajax.GetUserRunsResponse{}, HTTPErrInternalServerError
	}

	fullAmount, err := h.userSvc.GetUserRunCount(userID)
	if err != nil {
		return ajax.GetUserRunsResponse{}, HTTPErrInternalServerError
	}

	return ajax.GetUserRunsResponse{
		Runs:       runs,
		Amount:     len(runs),
		Offset:     offset,
		Limit:      limit,
		FullAmount: fullAmount,
	}, nil
}

func (h *RunsHandler) addRunAJAX(req ajax.AddRunRequest, user string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.addRunDB(req, user)
		if err != nil {
			WriteErrorHeader(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		ajax.MustEncode(w, res)
	}
}

func (h *RunsHandler) addRunHTMX(req ajax.AddRunRequest, user string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.addRunDB(req, user)
		if err != nil {
			WriteErrorHeader(w, err)
			return
		}

		htmx.NewTrigger().
			Add("add-run").
			AlertOK(fmt.Sprintf("Added the new run, you got position: %d!", res.Placement)).
			Write(w)
		w.WriteHeader(http.StatusCreated)
		htmx.MustRender(h.ctx, w, component.AddRunForm())
	}
}

func (h *RunsHandler) addRunDB(req ajax.AddRunRequest, user string) (ajax.AddRunResponse, error) {
	userID, err := h.userSvc.GetUserID(user)
	if err != nil {
		return ajax.AddRunResponse{}, HTTPErrInternalServerError
	}

	run, err := h.userSvc.AddRun(userID, req.Time.Duration)
	if err != nil {
		return ajax.AddRunResponse{}, HTTPErrInternalServerError
	}

	err = h.leaderboardSvc.AddRun(user, req.Time.Duration)
	if err != nil {
		return ajax.AddRunResponse{}, HTTPErrInternalServerError
	}

	rank, err := h.leaderboardSvc.GetRank(user)
	if err != nil {
		return ajax.AddRunResponse{}, HTTPErrInternalServerError
	}

	return ajax.AddRunResponse{
		Username:  user,
		Placement: rank + 1,
		Time:      req.Time,
		DateAdded: run.CreatedAt,
	}, nil
}
