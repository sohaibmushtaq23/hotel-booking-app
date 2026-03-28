package repository

import (
	"context"
	"database/sql"
	"errors"
	"hotel-booking-backend/internal/models"
)

type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

var ErrReservationNotFound = errors.New("reservation not found")

func (r *ReservationRepository) GetAll(ctx context.Context) ([]models.Reservation, error) {
	query := `
		SELECT ID, [IDCustomer], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [AmountPaid], [ReservedAt], [IDReservedBy], [Status]
		FROM reservations
		ORDER BY ID
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservation

	for rows.Next() {
		var rv models.Reservation

		err := rows.Scan(
			&rv.ID,
			&rv.IDCustomer,
			&rv.IDRoom,
			&rv.BookingStart,
			&rv.BookingEnd,
			&rv.ExtraCharges,
			&rv.AmountPaid,
			&rv.ReservedAt,
			&rv.IDReservedBy,
			&rv.Status,
		)
		if err != nil {
			return nil, err
		}

		reservations = append(reservations, rv)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationRepository) GetByID(ctx context.Context, id int) (*models.Reservation, error) {
	query := `
		SELECT ID, [IDCustomer], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [AmountPaid], [ReservedAt], [IDReservedBy], [Status]
		FROM reservations
		WHERE id = @p1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	var rv models.Reservation
	err := row.Scan(
		&rv.ID,
		&rv.IDCustomer,
		&rv.IDRoom,
		&rv.BookingStart,
		&rv.BookingEnd,
		&rv.ExtraCharges,
		&rv.AmountPaid,
		&rv.ReservedAt,
		&rv.IDReservedBy,
		&rv.Status,
	)

	if err == sql.ErrNoRows {
		return nil, ErrReservationNotFound
	}
	if err != nil {
		return nil, err
	}

	return &rv, nil
}

func (r *ReservationRepository) GetByIDClient(ctx context.Context, idClient int) ([]models.Reservation, error) {
	query := `
		SELECT ID, [IDCustomer], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [AmountPaid], [ReservedAt], [IDReservedBy], [Status]
		FROM reservations
		WHERE IDCustomer=@p1
		ORDER BY ID
	`

	rows, err := r.db.QueryContext(ctx, query, idClient)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservation

	for rows.Next() {
		var rv models.Reservation

		err := rows.Scan(
			&rv.ID,
			&rv.IDCustomer,
			&rv.IDRoom,
			&rv.BookingStart,
			&rv.BookingEnd,
			&rv.ExtraCharges,
			&rv.AmountPaid,
			&rv.ReservedAt,
			&rv.IDReservedBy,
			&rv.Status,
		)
		if err != nil {
			return nil, err
		}

		reservations = append(reservations, rv)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationRepository) GetByIDRoom(ctx context.Context, idRoom int) ([]models.Reservation, error) {
	query := `
		SELECT ID, [IDCustomer], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [AmountPaid], [ReservedAt], [IDReservedBy], [Status]
		FROM reservations
		WHERE IDRoom=@p1
		ORDER BY ID
	`

	rows, err := r.db.QueryContext(ctx, query, idRoom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservation

	for rows.Next() {
		var rv models.Reservation

		err := rows.Scan(
			&rv.ID,
			&rv.IDCustomer,
			&rv.IDRoom,
			&rv.BookingStart,
			&rv.BookingEnd,
			&rv.ExtraCharges,
			&rv.AmountPaid,
			&rv.ReservedAt,
			&rv.IDReservedBy,
			&rv.Status,
		)
		if err != nil {
			return nil, err
		}

		reservations = append(reservations, rv)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationRepository) Create(ctx context.Context, rv *models.Reservation) error {
	query := `
		INSERT INTO reservations 
		([IDCustomer], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [AmountPaid], [IDReservedBy], [Status])
		OUTPUT INSERTED.ID
		VALUES (@IDCustomer, @IDRoom, @BookingStart, @BookingEnd, @ExtraCharges, @AmountPaid, @IDReservedBy, @Status)
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		sql.Named("IDCustomer", rv.IDCustomer),
		sql.Named("IDRoom", rv.IDRoom),
		sql.Named("BookingStart", rv.BookingStart),
		sql.Named("BookingEnd", rv.BookingEnd),
		sql.Named("ExtraCharges", rv.ExtraCharges),
		sql.Named("AmountPaid", rv.AmountPaid),
		sql.Named("IDReservedBy", rv.IDReservedBy),
		sql.Named("Status", rv.Status),
	).Scan(&rv.ID)
}

func (r *ReservationRepository) Update(ctx context.Context, id int, rv *models.Reservation) (*models.Reservation, error) {
	query := `
        UPDATE reservations
        SET IDCustomer=@IDCustomer,
			IDRoom=@IDRoom,
			BookingStart=@BookingStart,
			BookingEnd=@BookingEnd,
			ExtraCharges=@ExtraCharges,
            AmountPaid=@AmountPaid,
            IDReservedBy=@IDReservedBy,
			Status=@Status
        OUTPUT INSERTED.ID, INSERTED.IDCustomer, INSERTED.IDRoom, INSERTED.BookingStart
			,INSERTED.BookingEnd, INSERTED.ExtraCharges, INSERTED.AmountPaid
            ,INSERTED.IDReservedBy, INSERTED.Status
        WHERE ID=@id
    `

	row := r.db.QueryRowContext(ctx, query,
		sql.Named("IDCustomer", rv.IDCustomer),
		sql.Named("IDRoom", rv.IDRoom),
		sql.Named("BookingStart", rv.BookingStart),
		sql.Named("BookingEnd", rv.BookingEnd),
		sql.Named("ExtraCharges", rv.ExtraCharges),
		sql.Named("AmountPaid", rv.AmountPaid),
		sql.Named("IDReservedBy", rv.IDReservedBy),
		sql.Named("Status", rv.Status),
		sql.Named("id", id),
	)

	var updated models.Reservation
	err := row.Scan(
		&updated.ID,
		&updated.IDCustomer,
		&updated.IDRoom,
		&updated.BookingStart,
		&updated.BookingEnd,
		&updated.ExtraCharges,
		&updated.AmountPaid,
		&updated.IDReservedBy,
		&updated.Status,
	)
	if err == sql.ErrNoRows {
		return nil, ErrReservationNotFound
	}
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *ReservationRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM reservations WHERE id=@p1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrReservationNotFound
	}

	return nil
}
