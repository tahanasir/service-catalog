package transport

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"tahanasir/service-catalog/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type dbStorage interface {
	GetService(uuid.UUID) (models.SingleService, error)
	GetTotalServicesCount(string) (int, error)
	GetAllServices(int, int, string) ([]models.Service, error)
}

func GetService(storage dbStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serviceID := chi.URLParam(r, "id")
		id, err := uuid.Parse(serviceID)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		service, err := storage.GetService(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(service); err != nil {
			http.Error(w, "Problem encoding JSON response", http.StatusInternalServerError)
			return
		}
	}
}

func GetAllServices(storage dbStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		search := r.URL.Query().Get("search")

		// default page
		if page < 1 {
			page = 1
		}

		// default limit
		if limit < 1 {
			limit = 10
		}

		totalServices, err := storage.GetTotalServicesCount(search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		offset := (page - 1) * limit // calculate offset
		services, err := storage.GetAllServices(limit, offset, search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// calcuate totalPages by taking the ceiling
		totalPages := int(math.Ceil(float64(totalServices) / float64(limit)))

		nextPage := page + 1
		if nextPage > totalPages {
			nextPage = 0
		}
		prevPage := page - 1
		if prevPage < 1 {
			prevPage = 0
		}

		pagination := models.Pagination{
			TotalServices: totalServices,
			TotalPages:    totalPages,
			CurrentPage:   page,
			NextPage:      nextPage,
			PrevPage:      prevPage,
		}

		apiResponse := models.Response{
			Pagination: pagination,
			Services:   services,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
			http.Error(w, "Problem encoding JSON response", http.StatusInternalServerError)
			return
		}
	}
}
