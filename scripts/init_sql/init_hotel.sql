CREATE TABLE IF NOT EXISTS hotels
(
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    city        TEXT NOT NULL,
    address     TEXT NOT NULL,
    hotel_class INTEGER CHECK ( hotel_class >= 0 and hotel_class <= 5 ),
    cost       INT  NOT NULL
);