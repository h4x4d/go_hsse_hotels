CREATE TABLE IF NOT EXISTS hotels
(
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    city        TEXT NOT NULL,
    address     TEXT NOT NULL,
    hotel_class INTEGER CHECK ( hotel_class >= 0 and hotel_class <= 5 )
);

CREATE TABLE IF NOT EXISTS rooms
(
    id          SERIAL PRIMARY KEY,
    hotel_id    SERIAL REFERENCES hotels (id),
    cost        INTEGER NOT NULL CHECK ( cost >= 0 ),
    person_count INTEGER NOT NULL CHECK ( person_count > 0 )
);

CREATE TABLE IF NOT EXISTS Tags (
    room_id    SERIAL REFERENCES rooms (id),
    tag        TEXT NOT NULL,
    UNIQUE (room_id, tag)
);