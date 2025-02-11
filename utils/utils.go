package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		errMsg := fmt.Sprintf(`{"error": "Internal server error",  "details": "%v"}`, err)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
}

func WriteError(w http.ResponseWriter, status int, errorMessage string, err error) {
	var errMessage string
	if err != nil {
		errMessage = fmt.Sprintf("%s: %v", errorMessage, err)
	} else {
		errMessage = errorMessage
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": strings.TrimSpace(errMessage)}); err != nil {
		http.Error(w, `{"error":"err new encoder"}`, http.StatusBadRequest)
	}
}

func ParseJSON(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	defer r.Body.Close()

	r.Body = http.MaxBytesReader(nil, r.Body, 1048576)

	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}
	return nil
}

type BaseModel struct {
	ID        uint       `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at,omitempty"`
}

// func StripGormModel(input interface{}) (interface{}, error) {
// 	val := reflect.ValueOf(input)
// 	if val.Kind() == reflect.Ptr {
// 		val = val.Elem()
// 	}

// 	if val.Kind() != reflect.Struct {
// 		return nil, errors.New("input must be a struct")
// 	}

// 	typ := val.Type()
// 	output := reflect.New(typ).Elem()

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)
// 		fieldType := typ.Field(i)
// 		if fieldType.Anonymous && fieldType.Type == reflect.TypeOf(BaseModel{}) {
// 			continue
// 		}
// 		output.Field(i).Set(field)
// 	}

// 	return output.Interface(), nil
// }
