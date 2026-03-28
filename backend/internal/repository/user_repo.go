package repository

import (
	"context"
	"database/sql"
	"errors"
	"hotel-booking-backend/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

var ErrUserNotFound = errors.New("client not found")

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	query := `
		SELECT ID, [UserName], [Password], [UserRole]
		FROM users
		ORDER BY ID
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User

		err := rows.Scan(
			&u.ID,
			&u.UserName,
			&u.Password,
			&u.UserRole,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := `
		SELECT ID, [UserName], [Password], [UserRole]
		FROM users
		WHERE id = @p1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	var u models.User
	err := row.Scan(
		&u.ID,
		&u.UserName,
		&u.Password,
		&u.UserRole,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) Create(ctx context.Context, u *models.User) error {
	query := `
		INSERT INTO users 
		([UserName], [Password], [UserRole])
		OUTPUT INSERTED.ID
		VALUES (@UserName,@Password,@UserRole)
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		sql.Named("UserName", u.UserName),
		sql.Named("Password", u.Password),
		sql.Named("UserRole", u.UserRole),
	).Scan(&u.ID)
}

func (r *UserRepository) Update(ctx context.Context, id int, u *models.User) (*models.User, error) {
	query := `
        UPDATE users
        SET UserName=@UserName,
			Password=@Password,
			UserRole=@UserRole
        OUTPUT INSERTED.ID, INSERTED.UserName, INSERTED.Password, INSERTED.UserRole
        WHERE ID=@id
    `

	row := r.db.QueryRowContext(ctx, query,
		sql.Named("UserName", u.UserName),
		sql.Named("Password", u.Password),
		sql.Named("UserRole", u.UserRole),
		sql.Named("id", id),
	)

	var updated models.User
	err := row.Scan(
		&updated.ID,
		&updated.UserName,
		&updated.Password,
		&updated.UserRole,
	)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id=@p1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrUserNotFound
	}

	return nil
}
