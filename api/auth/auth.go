package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/lazarcloud/pocketbase-dashboard/api/functions"
	"github.com/lazarcloud/pocketbase-dashboard/api/globals"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func EnsurePathDirectoryExists(path string) error {
	if !FileExists(path) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get Authorizations Header
		//check if password is correct
		//if not return 401

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			functions.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{
				"errortype": "auth",
				"error":     "Unauthorized",
			})
			return
		}

		isBearer := authHeader[:7] == "Bearer "
		if !isBearer {
			functions.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{
				"errortype": "auth",
				"error":     "Unauthorized, use Bearer Auth",
			})
			return
		}

		token := authHeader[7:]

		type Password struct {
			Password string `json:"password"`
		}

		passwordJSON, err := os.ReadFile(globals.AuthFilePath)
		if err != nil {
			functions.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{
				"errortype": "auth",
				"error":     fmt.Sprintf("failed reading file: %s, try recreating the dashboard container.", err.Error()),
			})
		}

		passwordData := Password{}
		err = json.Unmarshal(passwordJSON, &passwordData)
		if err != nil {
			functions.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{
				"errortype": "auth",
				"error":     fmt.Sprintf("error unmarshalling json: %s", err.Error()),
			})
			return
		}

		password := passwordData.Password

		if err != nil {
			functions.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{
				"errortype": "auth",
				"error":     err.Error(),
			})
			return
		}

		if string(password) != token {
			functions.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{
				"errortype": "auth",
				"error":     "Unauthorized, wrong password.",
			})
			return
		}

		next(w, r)
	}
}

func WriteJSONToFile(path string, data []byte) error {
	file, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func PrepareDefaultAuth(defaultPassword string) error {

	if !FileExists(globals.AuthFilePath) {
		err := EnsurePathDirectoryExists(globals.AuthFilePath)
		if err != nil {
			return err
		}
		err = WriteJSONToFile(globals.AuthFilePath, []byte(fmt.Sprintf(`{"password": "%s"}`, defaultPassword)))
		if err != nil {
			return err
		}
	}
	return nil
}
