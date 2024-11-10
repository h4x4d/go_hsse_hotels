CREATE TABLE IF NOT EXISTS bookings
(
    id          SERIAL PRIMARY KEY,
    date_from   DATE NOT NULL,
    date_to     DATE NOT NULL,
    room_id     INTEGER NOT NULL,
    hotel_id    INTEGER NOT NULL,
    user_id     INTEGER NOT NULL,
    status      TEXT NOT NULL CHECK ( status in ('Unpayed', 'Confirming', 'Confirmed', 'Canceled') )
);

CREATE TABLE IF NOT EXISTS users
(
    id          INTEGER PRIMARY KEY,
    name        TEXT NOT NULL,
    telegram    TEXT NOT NULL CHECK ( telegram LIKE '@%' ),
    role        TEXT NOT NULL CHECK ( role in ('customer', 'hotelier') )
);
