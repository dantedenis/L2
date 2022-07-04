package server

import (
	"encoding/json"
	"errors"
	"ex11_calendar/model"
	"net/http"
)

// хелпер для обработки пост запроса
func (s *Server) postHelperEvents(r *http.Request) (*model.Model, error) {
	if r.Method != http.MethodPost {
		return nil, errors.New("invalid method: " + r.Method)
	}

	// парсим необходимые данные из тела запроса
	return s.jsonDecode(r)
}

// хэлпер для обработки гет запроса
func getHelperEvent(r *http.Request) (*model.Model, error) {
	timeReq, err := model.Parse(r.URL.Query().Get("date"))
	if err != nil {
		return nil, err
	}

	// собираем модель эвента из queryString
	m := model.Model{
		Date:   timeReq,
		UserID: r.URL.Query().Get("user_id"),
	}
	return &m, nil
}

// метод для декодера тела запроса в модель
func (s *Server) jsonDecode(r *http.Request) (*model.Model, error) {
	event := model.Model{}

	// преобразуем и возвращаем результат
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// метод для вывода ошибок
func writeError(w http.ResponseWriter, msg string, statusCode int) {
	// анонимная структура которую маршалим и отпрвляет ответом
	err := struct {
		Err string `json:"error"`
	}{Err: msg}

	input, _ := json.MarshalIndent(&err, "", "  ")
	http.Error(w, string(input), statusCode)
}

// метод для печати результата
func writeResult(w http.ResponseWriter, res []model.Model) {
	// анонимная структра для маршала
	result := struct {
		Result []model.Model `json:"result"`
	}{res}

	input, _ := json.MarshalIndent(&result, "", "  ")
	_, err := w.Write(input)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadGateway)
	}
}
