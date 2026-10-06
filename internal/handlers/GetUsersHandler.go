package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// user struct for response
type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Login string `json:"login"`
}

func (h *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	//query request users list
	rows, err := h.DB.QueryContext(
		r.Context(),
		`
			SELECT id, name, email, role, login
			FROM users
			ORDER BY id
		`,
	)
	if err != nil {
		log.Println("[GetUsersHandler]: Query error: ", err)
	}
	defer rows.Close()

	users := make([]User, 0)

	//adding users list to users[] from rows
	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Role,
			&user.Login,
		)
		if err != nil {
			log.Println("[GetUsersHandler]: Scan error: ", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("[GetUsers]: Rows iteration error:", err)
	}

	//sending response with users list data
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(users)
}
