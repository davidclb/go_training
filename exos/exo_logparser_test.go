package exos

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	// cas valide
	got, err := ParseLine("192.168.1.10 GET /api/users 200")
	if err != nil {
		t.Fatalf("ParseLine valide a retourné une erreur: %v", err)
	}
	want := LogEntry{IP: "192.168.1.10", Method: "GET", Path: "/api/users", Status: 200}
	if got != want {
		t.Errorf("ParseLine = %+v; want %+v", got, want)
	}

	// ligne malformée (pas assez de champs) -> doit retourner une erreur
	if _, err := ParseLine("oops"); err == nil {
		t.Error("ParseLine sur ligne malformée: attendu une erreur, got nil")
	}

	// status non numérique -> doit retourner une erreur
	if _, err := ParseLine("192.168.1.10 GET /api/users abc"); err == nil {
		t.Error("ParseLine sur status non numérique: attendu une erreur, got nil")
	}
}

func TestTopErrors(t *testing.T) {
	entries := []LogEntry{
		{"10.0.0.1", "GET", "/login", 401},
		{"10.0.0.2", "GET", "/login", 500},
		{"10.0.0.3", "GET", "/users", 200}, // succès -> ne doit PAS compter
		{"10.0.0.4", "POST", "/users", 503},
		{"10.0.0.5", "GET", "/users", 200}, // succès -> ne doit PAS compter
	}
	got := TopErrors(entries, 2)
	want := []PathCount{
		{"/login", 2}, // 2 erreurs (401, 500)
		{"/users", 1}, // 1 erreur (503) — les deux 200 sont exclus
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TopErrors = %+v; want %+v", got, want)
	}
}
