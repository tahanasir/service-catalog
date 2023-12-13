package database

import (
	"tahanasir/service-catalog/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	Database *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Database: db,
	}
}

// Query to get service info from services table
// Second query to get version info if service exists
func (s *Storage) GetService(id uuid.UUID) (models.SingleService, error) {
	var service models.SingleService
	err := s.Database.Get(&service, "SELECT id, name, description FROM services WHERE id = ?", id)
	if err != nil {
		return service, err
	}

	err = s.Database.Select(&service.Versions, "SELECT name, changelog FROM versions WHERE service_id = ?", id)
	if err != nil {
		return service, err
	}

	return service, nil
}

// Query to count the total number of services, filtered by search
func (s *Storage) GetTotalServicesCount(search string) (int, error) {
	var totalServices int
	query := "SELECT COUNT(*) FROM services"

	var searchParam []interface{}
	if search != "" {
		query += " WHERE name LIKE ?"
		searchParam = append(searchParam, "%"+search+"%")
	}

	err := s.Database.Get(&totalServices, query, searchParam...)
	if err != nil {
		return 0, err
	}

	return totalServices, nil
}

// Query to get all services build using optional query params
// Join to get the version count for each service id
// Filter by search and return result with a limit and offset
func (s *Storage) GetAllServices(limit int, offset int, search string) ([]models.Service, error) {
	var services []models.Service
	query := `
		SELECT s.id, s.name, s.description, COUNT(v.service_id) AS version_count
		FROM services s
		LEFT JOIN versions v ON s.id = v.service_id
	`
	var queryParams []interface{}

	if search != "" {
		query += "WHERE s.name LIKE ?"
		queryParams = append(queryParams, "%"+search+"%")
	}

	query += " GROUP BY s.id, s.name, s.description ORDER BY s.created_date DESC"

	// cannot offset without a specified limit
	if limit > 0 {
		query += " LIMIT ?"
		queryParams = append(queryParams, limit)
		if offset > 0 {
			query += " OFFSET ?"
			queryParams = append(queryParams, offset)
		}
	}

	var err error
	if len(queryParams) > 0 {
		err = s.Database.Select(&services, query, queryParams...)
	} else {
		err = s.Database.Select(&services, query)
	}

	if err != nil {
		return nil, err
	}

	return services, nil
}
