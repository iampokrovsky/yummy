package api

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	menu_model "yummy/internal/app/menu/model"
	menu_repo "yummy/internal/app/menu/repo"
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrNotFound   = errors.New("not found")
)

// TODO: reduce code duplication

func (rt *Router) createMenuItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var items []menu_model.MenuItem
	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ids, err := rt.menuRepo.Create(r.Context(), items...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		Ids []uint64 `json:"ids"`
	}{Ids: ids}

	encodedResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(encodedResp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rt *Router) getMenuItemByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := rt.menuRepo.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, menu_repo.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rt *Router) listMenuItemsByRestaurantID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	restId, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	items, err := rt.menuRepo.ListByRestaurantID(r.Context(), restId)
	if err != nil {
		if errors.Is(err, menu_repo.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rt *Router) listMenuItemsByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if name == "" {
		http.Error(w, ErrBadRequest.Error(), http.StatusBadRequest)
		return
	}

	items, err := rt.menuRepo.ListByName(r.Context(), name)
	if err != nil {
		if errors.Is(err, menu_repo.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rt *Router) updateMenuItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var items []menu_model.MenuItem
	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ok, err := rt.menuRepo.Update(r.Context(), items...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, ErrNotFound.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rt *Router) deleteMenuItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ok, err := rt.menuRepo.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, ErrNotFound.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rt *Router) restoreMenuItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ok, err := rt.menuRepo.Restore(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, ErrNotFound.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
