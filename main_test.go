package main

// func TestPingRoute(t *testing.T) {
// 	router := setupRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/ping", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{"message":"pong"}`, w.Body.String())
// }

// func TestGetUserByIdRoute(t *testing.T) {
// 	router := setupRouter()
// 	w := httptest.NewRecorder()
// 	id := 1
// 	req, _ := http.NewRequest("GET", "/user/"+strconv.Itoa(id), nil)
// 	router.ServeHTTP(w, req)

// 	user, _ := model.
// 	jsonUser, _ := json.Marshal(user)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, string(jsonUser), w.Body.String())
// }
