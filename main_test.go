package traefikkeycloak

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ServeHTTP(t *testing.T) {
	config := CreateConfig()
	config.URL = "http://localhost:7070"
	config.Token = "Auth-Token"
	config.Realm = "dev"
	config.ParsedToken = "Parsed-Token"
	next := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {})
	handler, err := New(context.Background(), next, config, "keycloak-plugin")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "Auth-Token",
		Value: "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJoX2Nsc21iTnBHWFhlVjNzeVhWZ0RnWXFqUmcwTTBKZ3R4X0VIeTFYdndzIn0.eyJleHAiOjE2MDI1NzgzNTIsImlhdCI6MTYwMjU3ODA1MiwiYXV0aF90aW1lIjoxNjAyNTc3MTY2LCJqdGkiOiI4YzM5MmI2MS1lMjA0LTQ1MTAtYjU0ZC02OTU1ZmI4ZjYzY2QiLCJpc3MiOiJodHRwOi8vY2xvdWQtZGV2LmF6aW50ZWxlY29tLmF6OjcwNzAvYXV0aC9yZWFsbXMvZGV2IiwiYXVkIjpbImF6Y2xvdWQtYWRtaW4tdWkiLCJhY2NvdW50Il0sInN1YiI6ImFhMTUyNmE0LTg5MTctNGM5MC05YTA1LThlMTI4YmVkYTIzMSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImF6Y2xvdWQtdWkiLCJub25jZSI6IjIwNGYyMzhhLTQzMmYtNDdiZi04N2FiLWQ2ZWZhZWQ4NDhmZCIsInNlc3Npb25fc3RhdGUiOiI0MzA2Njc2Mi1iZjZiLTRlZGItOWJmNy1kNDZjZDMzN2Y0ZGYiLCJhY3IiOiIwIiwiYWxsb3dlZC1vcmlnaW5zIjpbImh0dHA6Ly9kZXYuYXppbnRlbGVjb20uYXo6ODA4MCIsImh0dHA6Ly9kZXYuYXppbnRlbGVjb20uYXo6ODA4MSJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImF6Y2xvdWQtYWRtaW4tdWkiOnsicm9sZXMiOlsiYWRtaW4iXX0sImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicHJvamVjdHMiOlt7ImlkIjoiN2JjMWYwZjQtN2E4Ni00MTVkLWExNGYtN2JmNzIyNmNkMGZjIiwibmFtZSI6IlByb2plY3QgWSJ9LHsiaWQiOiI1ZmQwNTRiYi01M2QwLTQyNGMtOTNiNi0yZmE4M2Y1ZGQxMTMiLCJuYW1lIjoiTGFiIn1dLCJuYW1lIjoiSWxraW4gTXVzYXlldiIsImNvbXBhbnkiOnsiaWQiOiI4NDE3OTU2Zi0zMWI3LTRlOWMtODU1NS0wZjI3NGMzYWFkMzMiLCJuYW1lIjoiVGVzdCJ9LCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJpbGtpbi5tdXNheWV2QGF6aW50ZWxlY29tLmF6IiwiZ2l2ZW5fbmFtZSI6Iklsa2luIiwiZmFtaWx5X25hbWUiOiJNdXNheWV2IiwiZW1haWwiOiJpbGtpbi5tdXNheWV2QGF6aW50ZWxlY29tLmF6In0.YOsges2DcnjZCX6fgCk2_68xvNFRfkYTwQY4Twbzzv7rylf5j3Mhjgy8mTdRr_CzX1Wct3EYwvjQSod1E-9brB_-n_3i24OGdIvTMAv9TSU6eIw4St_7pnA7WohctjuHohaMX-PFoTRJWs4BTC1RmqT5yBL6ey2QkuZs4c7ZmpX2nAimDGWDxYA7AINxb1V5JKBHTwKl565ze8ZgSogc9r4-ZcNvaC1ZgVk8FKwXC0bVmCk-NeBaq-76QM-yMLZ7TiTAfnolHpr60mgxQNxeUs82XR4AIKSDtLN1UAN5CMWp2C0w89pGog-7w53f765wlmLqni20a09DbiLaH72JOw",
	})
	handler.ServeHTTP(recorder, req)
	parsedToken := recorder.Header().Get("Parsed-Token")
	if parsedToken == "" {
		t.Fatal(err)
	}
}
