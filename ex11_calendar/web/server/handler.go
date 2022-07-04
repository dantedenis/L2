package server

import (
	"net/http"
)

// NewRouter конструктор нашего роутера, который перенаправляет запросы методам через мидлвар
func (s *Server) NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/create_event", middleware(http.HandlerFunc(s.create)))
	router.HandleFunc("/update_event", middleware(http.HandlerFunc(s.update)))
	router.HandleFunc("/delete_event", middleware(http.HandlerFunc(s.delete)))
	router.HandleFunc("/events_for_day", middleware(http.HandlerFunc(s.eventsDay)))
	router.HandleFunc("/events_for_week", middleware(http.HandlerFunc(s.eventsWeek)))
	router.HandleFunc("/events_for_month", middleware(http.HandlerFunc(s.eventsMonth)))

	return router
}

// метод обработчик запросов для создания и добавления эвента
func (s *Server) create(w http.ResponseWriter, r *http.Request) {
	event, err := s.postHelperEvents(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// пытается добавить эвент в наш кэш
	result, err := s.Events.CreateEvent(event)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// если все ок мы пишем наш результат, в данном случае выводит все существующие эвенты для этого айди пользователя
	writeResult(w, result)
}

// update обработчик запросов для обновления данных
func (s *Server) update(w http.ResponseWriter, r *http.Request) {
	event, err := s.postHelperEvents(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := s.Events.UpdateEvent(event)

	writeResult(w, result)
}

// метод использует хэлпер и выполняет изменение кэша, и выводит результат
func (s *Server) delete(w http.ResponseWriter, r *http.Request) {
	event, err := s.postHelperEvents(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := s.Events.DeleteEvent(event)

	writeResult(w, result)
}

//    методы ниже используют гетхэлпер анализирует значения и выводит результат
func (s *Server) eventsDay(w http.ResponseWriter, r *http.Request) {
	event, err := getHelperEvent(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := s.Events.EventForDay(event)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeResult(w, result)
}

func (s *Server) eventsWeek(w http.ResponseWriter, r *http.Request) {
	event, err := getHelperEvent(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := s.Events.EventForWeek(event)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeResult(w, result)
}

func (s *Server) eventsMonth(w http.ResponseWriter, r *http.Request) {
	event, err := getHelperEvent(r)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := s.Events.EventForMonth(event)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeResult(w, result)
}
