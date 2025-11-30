package handler

import (
	"net/http"
	"strconv"
)

func UpdateHandle(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	metricType := req.PathValue("type")
	metricName := req.PathValue("name")
	metricValue := req.PathValue("value")

	// Проверяем наличие имени метрики
	if metricName == "" {
		http.Error(res, "Metric name is required", http.StatusNotFound)
		return
	}

	// Устанавливаем Content-Type для всех ответов
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")

	switch metricType {
	case "gauge":
		_, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(res, "Invalid gauge value", http.StatusBadRequest)
			return
		}

		res.WriteHeader(http.StatusOK)
	case "counter":
		_, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(res, "Invalid counter value", http.StatusBadRequest)
			return
		}

		res.WriteHeader(http.StatusOK)

	default:
		http.Error(res, "Invalid metric type", http.StatusBadRequest)
		return
	}
}
