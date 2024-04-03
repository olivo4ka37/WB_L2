package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func NewEvent() *Event {
	date, _ := time.Parse("2006-01-02", "2024-29-01")
	return &Event{
		UserID: 1,
		Date:   date,
		Title:  "Birthday",
	}
}

func TestServer_CreateEvent(t *testing.T) {
	s := newServer(
		Config{":8080"},
		newStore(),
	)

	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("user_id", "1")
	form.Add("date", "2023-05-23")
	form.Add("title", "Birthday")

	req := httptest.NewRequest(http.MethodPost, "/create_event", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	s.CreateEvent().ServeHTTP(rec, req)
	assert.Equal(t, "{\"result\":\"created event with id=0\"}\n", rec.Body.String())
}

func TestServer_UpdateEvent(t *testing.T) {
	s := newServer(
		Config{":8080"},
		newStore(),
	)
	s.store.Event().CreateEvent(NewEvent())

	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("id", "0")
	form.Add("user_id", "1")
	form.Add("date", "2023-05-23")
	form.Add("title", "Doctor")

	req := httptest.NewRequest(http.MethodPost, "/update_event", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	s.UpdateEvent().ServeHTTP(rec, req)
	assert.Equal(t, "{\"result\":\"updated event with id=0\"}\n", rec.Body.String())
}

func TestServer_DeleteEvent(t *testing.T) {
	s := newServer(
		Config{":8080"},
		newStore(),
	)
	s.store.Event().CreateEvent(NewEvent())

	rec := httptest.NewRecorder()

	form := url.Values{}
	form.Add("id", "0")

	req := httptest.NewRequest(http.MethodPost, "/delete_event", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	s.DeleteEvent().ServeHTTP(rec, req)
	assert.Equal(t, "{\"result\":\"deleted event with id=0\"}\n", rec.Body.String())
}
