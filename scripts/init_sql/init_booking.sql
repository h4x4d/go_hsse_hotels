CREATE TABLE IF NOT EXISTS bookings
(
    id        SERIAL PRIMARY KEY,
    date_from DATE    NOT NULL,
    date_to   DATE    NOT NULL,
    hotel_id  INTEGER NOT NULL,
    full_cost INTEGER                                                                     DEFAULT 0,
    status    TEXT CHECK ( status in ('Waiting', 'Confirmed', 'Canceled') ) DEFAULT 'Waiting',
    user_id   TEXT    NOT NULL
);