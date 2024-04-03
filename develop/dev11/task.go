package main

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Event хранит данные о событии
type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}

// String возвращает событие в виде строки
func (e *Event) String() string {
	year, month, day := e.Date.Date()
	return fmt.Sprintf("%s on %v %v %v", e.Title, year, month, day)
}

// Error используется для отправки ответа с описанием ошибки
type Error struct {
	Error string `json:"error"`
}

// Result используется для отправки ответа с описанием результата обработки запроса
type Result struct {
	Result string `json:"result"`
}

// Server описывает структуру сервера
type Server struct {
	config Config
	store  Store
	router *http.ServeMux
}

// Config описывает конфигурацию сервера
type Config struct {
	addr string
}

// Store описывает абстрактную базу данных
type Store interface {
	Event() EventRepository
}

// MyStore представляет конкретную базу данных
type MyStore struct {
	eventRepository *MyEventRepository
}

// Event возвращает хранилище событий
func (s *MyStore) Event() EventRepository {
	if s.eventRepository != nil {
		return s.eventRepository
	}

	s.eventRepository = &MyEventRepository{
		eventRepository: make(map[int]*Event),
	}

	return s.eventRepository
}

// EventRepository описывает абстрактное хранилище событий
type EventRepository interface {
	CreateEvent(*Event) error
	UpdateEvent(*Event) error
	DeleteEvent(*Event) error
	FindEventsForDay(userID int, date time.Time) []*Event
	FindEventsForWeek(userID int, date time.Time) []*Event
	FindEventsForMonth(userID int, date time.Time) []*Event
}

// MyEventRepository представляет конкретное хранилище событий
type MyEventRepository struct {
	eventRepository map[int]*Event
}

// CreateEvent создает событие
func (r *MyEventRepository) CreateEvent(event *Event) error {
	event.ID = len(r.eventRepository)
	r.eventRepository[event.ID] = event
	return nil
}

// UpdateEvent обновляет событие по его id
func (r *MyEventRepository) UpdateEvent(event *Event) error {
	if _, ok := r.eventRepository[event.ID]; !ok {
		return errors.New("event doesn't exist")
	}
	r.eventRepository[event.ID] = event
	return nil
}

// DeleteEvent удаляет событие
func (r *MyEventRepository) DeleteEvent(event *Event) error {
	if _, ok := r.eventRepository[event.ID]; !ok {
		return errors.New("event doesn't exist")
	}
	delete(r.eventRepository, event.ID)
	return nil
}

// FindEventsForDay возвращает события на день для заданного пользователя
func (r *MyEventRepository) FindEventsForDay(userID int, date time.Time) []*Event {
	var events []*Event

	for _, event := range r.eventRepository {
		if event.UserID == userID &&
			event.Date.Equal(date) {
			events = append(events, event)
		}
	}
	return events
}

// FindEventsForWeek возвращает события на неделю для заданного пользователя
func (r *MyEventRepository) FindEventsForWeek(userID int, date time.Time) []*Event {
	var events []*Event

	for _, event := range r.eventRepository {
		if event.UserID == userID &&
			(event.Date.Equal(date) ||
				event.Date.After(date) &&
					event.Date.Before(date.Add(7*24*time.Hour))) {
			events = append(events, event)
		}
	}
	return events
}

// FindEventsForMonth возвращает события на месяц для заданного пользователя
func (r *MyEventRepository) FindEventsForMonth(userID int, date time.Time) []*Event {
	var events []*Event

	for _, event := range r.eventRepository {
		if event.UserID == userID &&
			(event.Date.Equal(date) ||
				event.Date.After(date) &&
					event.Date.Before(date.Add(30*24*time.Hour))) {
			events = append(events, event)
		}
	}
	return events
}

// configureRouter регистрирует обработчики
func (s *Server) configureRouter() {
	s.router.HandleFunc("/create_event", s.middleware(s.CreateEvent()))
	s.router.HandleFunc("/update_event", s.middleware(s.UpdateEvent()))
	s.router.HandleFunc("/delete_event", s.middleware(s.DeleteEvent()))
	s.router.HandleFunc("/events_for_day", s.middleware(s.EventsForDay()))
	s.router.HandleFunc("/events_for_week", s.middleware(s.EventsForWeek()))
	s.router.HandleFunc("/events_for_month", s.middleware(s.EventsForMonth()))
}

// middleware логирует запросы
func (s *Server) middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t%s\n", r.Method, r.URL)
		next(w, r)
	}
}

// CreateEvent обрабатывает создание события
func (s *Server) CreateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}

			userID, err := strconv.Atoi(
				r.PostFormValue("user_id"),
			)
			if err != nil {
				sendError(w, http.StatusBadRequest, "missing or invalid user_id")
				return
			}

			date, err := time.Parse(
				"2006-01-02",
				r.PostFormValue("date"),
			)
			if err != nil {
				sendError(w, http.StatusBadRequest, "missing or invalid date")
				return
			}

			title := r.PostFormValue("title")
			if len(title) == 0 {
				sendError(w, http.StatusBadRequest, "missing or empty title")
				return
			}

			event := Event{
				UserID: userID,
				Date:   date,
				Title:  title,
			}

			if err := s.store.Event().CreateEvent(&event); err != nil {
				sendError(w, http.StatusServiceUnavailable, err.Error())
				return
			}

			sendResult(w, http.StatusCreated, fmt.Sprintf("created event with id=%d", event.ID))

		default:
			sendError(w, http.StatusNotFound, "Status Not Found")
		}
	}
}

// UpdateEvent обрабатывает обновление события
func (s *Server) UpdateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}

			id, err := strconv.Atoi(
				r.PostFormValue("id"),
			)
			if err != nil {
				sendError(w, http.StatusBadRequest, "missing or invalid id")
				return
			}

			userID, err := strconv.Atoi(
				r.PostFormValue("user_id"),
			)
			if err != nil {
				fmt.Println(r.PostFormValue("user_id"))
				sendError(w, http.StatusBadRequest, "missing or invalid user_id")
				return
			}

			date, err := time.Parse(
				"2006-01-02",
				r.PostFormValue("date"),
			)
			if err != nil {
				sendError(w, http.StatusBadRequest, "missing or invalid date")
				return
			}

			title := r.PostFormValue("title")
			if len(title) == 0 {
				sendError(w, http.StatusBadRequest, "missing or empty title")
				return
			}

			event := Event{
				ID:     id,
				UserID: userID,
				Date:   date,
				Title:  title,
			}

			if err := s.store.Event().UpdateEvent(&event); err != nil {
				sendError(w, http.StatusServiceUnavailable, err.Error())
				return
			}

			sendResult(w, http.StatusCreated, fmt.Sprintf("updated event with id=%d", event.ID))

		default:
			sendError(w, http.StatusNotFound, "Status Not Found")
		}
	}
}

// DeleteEvent обрабатывает удаление события
func (s *Server) DeleteEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}

			id, err := strconv.Atoi(
				r.PostFormValue("id"),
			)
			if err != nil {
				sendError(w, http.StatusBadRequest, "missing or invalid id")
				return
			}

			event := Event{
				ID: id,
			}

			if err := s.store.Event().DeleteEvent(&event); err != nil {
				sendError(w, http.StatusServiceUnavailable, err.Error())
				return
			}

			sendResult(w, http.StatusOK, fmt.Sprintf("deleted event with id=%d", event.ID))

		default:
			sendError(w, http.StatusNotFound, "Status Not Found")
		}
	}
}

// EventsForDay обрабатывает списка событий на день
func (s *Server) EventsForDay() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userID, err := strconv.Atoi(
				r.URL.Query().Get("user_id"),
			)
			if err != nil {
				sendError(w, http.StatusBadRequest, err.Error())
				return
			}

			date, err := time.Parse(
				"2006-01-02",
				r.URL.Query().Get("date"),
			)
			if err != nil {
				sendError(w, http.StatusBadRequest, err.Error())
				return
			}

			events := s.store.Event().FindEventsForDay(userID, date)

			sendResult(w, http.StatusOK, toString(events))

		default:
			sendError(w, http.StatusNotFound, "Status Not Found")
		}
	}
}

// EventsForWeek обрабатывает списка событий на неделю
func (s *Server) EventsForWeek() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userID, err := strconv.Atoi(
				r.URL.Query().Get("user_id"),
			)
			if err != nil {
				sendError(w, http.StatusBadRequest, err.Error())
				return
			}

			date, err := time.Parse(
				"2006-01-02",
				r.URL.Query().Get("date"),
			)
			if err != nil {
				sendError(w, http.StatusBadRequest, err.Error())
				return
			}

			events := s.store.Event().FindEventsForWeek(userID, date)

			sendResult(w, http.StatusOK, toString(events))

		default:
			sendError(w, http.StatusNotFound, "Status Not Found")
		}
	}
}

// EventsForMonth обрабатывает списка событий на месяц
func (s *Server) EventsForMonth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userID, err := strconv.Atoi(
				r.URL.Query().Get("user_id"),
			)
			if err != nil {
				sendError(w, http.StatusBadRequest, err.Error())
				return
			}

			date, err := time.Parse(
				"2006-01-02",
				r.URL.Query().Get("date"),
			)
			if err != nil {
				sendError(w, http.StatusBadRequest, err.Error())
				return
			}

			events := s.store.Event().FindEventsForMonth(userID, date)

			sendResult(w, http.StatusOK, toString(events))

		default:
			sendError(w, http.StatusNotFound, "Status Not Found")
		}
	}
}

// toString возвращает список событий в виде строки
func toString(events []*Event) string {
	var s []string

	for _, event := range events {
		s = append(s, event.String())
	}

	return strings.Join(s, ", ")
}

// sendResult отправляет результат
func sendResult(w http.ResponseWriter, code int, msg string) {
	err := Result{
		Result: msg,
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

// sendError отправляет ошибку
func sendError(w http.ResponseWriter, code int, msg string) {
	err := Error{
		Error: msg,
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

// start запускает сервер
func (s *Server) start() error {
	s.configureRouter()
	return http.ListenAndServe(s.config.addr, s.router)
}

// newStore конкретную реализацию базы данных
func newStore() Store {
	return &MyStore{}
}

// newServer возвращает инициализированный сервер
func newServer(config Config, store Store) *Server {
	return &Server{
		config: config,
		router: http.NewServeMux(),
		store:  store,
	}
}

func main() {
	config := Config{
		addr: ":8080",
	}
	store := newStore()
	server := newServer(config, store)

	if err := server.start(); err != nil {
		log.Fatalln(err)
	}
}
