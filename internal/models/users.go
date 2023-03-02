// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// Lets model the question table
type User struct {
	UserID    int32
	Name      string
	Email     string
	CreatedAt time.Time
}

// setup dependency injection
type UserModel struct {
	DB *sql.DB //connection pool
}

// write SQL code to access the database
// TODO
func (m *UserModel) Get() (*User, error) { //we use QuestionModel because it has acces to connection pool
	var q User

	statement := `
			SELECT user_id, name, email 
			FROM users
			ORDER BY RANDOM()
			LIMIT 1
			`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.UserID, &q.Email) //m is the instNCE, DB. connectionpool, and we want to send query row context
	if err != nil {
		return nil, err
	}
	return &q, err

}
