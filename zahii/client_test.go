package zahii

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClient_PerRequestAuth(t *testing.T) {
	authCount := 0
	apiCount := 0
	sessionToken := "session-token-123"
	superAppToken := "super-app-master-token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.Path, "/super-app/authenticate") {
			authCount++
			// Check that it's using the correct master token
			if !strings.Contains(r.URL.Path, superAppToken) {
				t.Errorf("auth call missing correct master token in path")
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"message": "success", "body": {"token": "%s"}}`, sessionToken)
			return
		}

		apiCount++
		// Check that the session token is attached
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer "+sessionToken {
			t.Errorf("API call missing fresh session token, got: %s", authHeader)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "success"}`)
	}))
	defer server.Close()

	client, _ := NewClient(Config{
		BaseURL:       server.URL,
		SuperAppToken: superAppToken,
	})

	// First call to some API (e.g., Wishlist.List)
	_, err := client.Customer.Wishlist.List()
	if err != nil {
		t.Fatalf("first API call failed: %v", err)
	}

	// Second call to some API
	_, err = client.Customer.Wishlist.List()
	if err != nil {
		t.Fatalf("second API call failed: %v", err)
	}

	if authCount != 2 {
		t.Errorf("expected 2 auth calls (one per request), got %d", authCount)
	}
	if apiCount != 2 {
		t.Errorf("expected 2 API calls, got %d", apiCount)
	}
}
