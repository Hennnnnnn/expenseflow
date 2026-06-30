package service

import "github.com/Hennnnnnn/expenseflow/internal/transport/http/dto"

type DashboardService interface {
	GetDashboard() (*dto.DashboardResponse, error)
}
