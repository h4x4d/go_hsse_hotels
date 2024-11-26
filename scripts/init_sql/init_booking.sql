CREATE TABLE IF NOT EXISTS bookings
(
    id        SERIAL PRIMARY KEY,
    date_from DATE        NOT NULL,
    date_to   DATE        NOT NULL,
    room_id   INTEGER     NOT NULL,
    hotel_id  INTEGER     NOT NULL,
    user_id   VARCHAR(32) NOT NULL,
    full_cost INTEGER DEFAULT 0,
    status    TEXT        NOT NULL CHECK ( status in ('Unpayed', 'Confirming', 'Confirmed', 'Canceled') )
);