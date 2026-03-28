package repository

import (
	"context"
	"database/sql"
	"errors"
	"hotel-booking-backend/internal/models"
)

type RoomReservationRepository struct {
	db *sql.DB
}

func NewRoomReservationRepository(db *sql.DB) *RoomReservationRepository {
	return &RoomReservationRepository{db: db}
}

var ErrRoomReservationNotFound = errors.New("room reservation not found")

func (r *RoomReservationRepository) GetByIDReservation(ctx context.Context, idReservation int) ([]models.RoomReservation, error) {
	query := `
		SELECT ID, [IDReservation], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [Status]
		FROM room_reservations
		WHERE IDReservation=@p1
		ORDER BY ID
	`

	rows, err := r.db.QueryContext(ctx, query, idReservation)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var room_reservations []models.RoomReservation

	for rows.Next() {
		var rv models.RoomReservation

		err := rows.Scan(
			&rv.ID,
			&rv.IDReservation,
			&rv.IDRoom,
			&rv.BookingStart,
			&rv.BookingEnd,
			&rv.ExtraCharges,
			&rv.Status,
		)
		if err != nil {
			return nil, err
		}

		room_reservations = append(room_reservations, rv)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return room_reservations, nil
}

func (r *RoomReservationRepository) GetByIDRoom(ctx context.Context, idRoom int) ([]models.RoomReservation, error) {
	query := `
		SELECT ID, [IDReservation], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [Status]
		FROM room_reservations
		WHERE IDRoom=@p1
		ORDER BY ID
	`

	rows, err := r.db.QueryContext(ctx, query, idRoom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var room_reservations []models.RoomReservation

	for rows.Next() {
		var rv models.RoomReservation

		err := rows.Scan(
			&rv.ID,
			&rv.IDReservation,
			&rv.IDRoom,
			&rv.BookingStart,
			&rv.BookingEnd,
			&rv.ExtraCharges,
			&rv.Status,
		)
		if err != nil {
			return nil, err
		}

		room_reservations = append(room_reservations, rv)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return room_reservations, nil
}

func (r *RoomReservationRepository) GetByID(ctx context.Context, id int) (*models.RoomReservation, error) {
	query := `
		SELECT ID, [IDReservation], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [Status]
		FROM room_reservations
		WHERE id = @p1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	var rv models.RoomReservation
	err := row.Scan(
		&rv.ID,
		&rv.IDReservation,
		&rv.IDRoom,
		&rv.BookingStart,
		&rv.BookingEnd,
		&rv.ExtraCharges,
		&rv.Status,
	)

	if err == sql.ErrNoRows {
		return nil, ErrRoomReservationNotFound
	}
	if err != nil {
		return nil, err
	}

	return &rv, nil
}

func (r *RoomReservationRepository) Create(ctx context.Context, rv *models.RoomReservation) error {
	query := `
		INSERT INTO room_reservations 
		([IDReservation], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [Status])
		OUTPUT INSERTED.ID
		VALUES (@IDReservation,@IDRoom,@BookingStart, @BookingEnd, @ExtraCharges, @Status)
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		sql.Named("IDReservation", rv.IDReservation),
		sql.Named("IDRoom", rv.IDRoom),
		sql.Named("BookingStart", rv.BookingStart),
		sql.Named("BookingEnd", rv.BookingEnd),
		sql.Named("ExtraCharges", rv.ExtraCharges),
		sql.Named("Status", rv.Status),
	).Scan(&rv.ID)
}

func (r *RoomReservationRepository) Update(ctx context.Context, id int, rv *models.RoomReservation) (*models.RoomReservation, error) {
	query := `
        UPDATE room_reservations
        SET IDReservation=@IDReservation,
			IDRoom=@IDRoom,
            BookingStart=@BookingStart,
            BookingEnd=@BookingEnd,
			ExtraCharges=@ExtraCharges,
			Status=@Status
        OUTPUT INSERTED.ID, INSERTED.IDReservation, INSERTED.IDRoom, INSERTED.BookingStart
               ,INSERTED.BookingEnd, INSERTED.ExtraCharges, INSERTED.Status
        WHERE ID=@id
    `

	row := r.db.QueryRowContext(ctx, query,
		sql.Named("IDReservation", rv.IDReservation),
		sql.Named("IDRoom", rv.IDRoom),
		sql.Named("BookingStart", rv.BookingStart),
		sql.Named("BookingEnd", rv.BookingEnd),
		sql.Named("ExtraCharges", rv.ExtraCharges),
		sql.Named("Status", rv.Status),
		sql.Named("id", id),
	)

	var updated models.RoomReservation
	err := row.Scan(
		&updated.ID,
		&updated.IDReservation,
		&updated.IDRoom,
		&updated.BookingStart,
		&updated.BookingEnd,
		&updated.ExtraCharges,
		&updated.Status,
	)
	if err == sql.ErrNoRows {
		return nil, ErrRoomReservationNotFound
	}
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *RoomReservationRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM room_reservations WHERE id=@p1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrRoomReservationNotFound
	}

	return nil
}
