CREATE TABLE IF NOT EXISTS go.tickets 
(
    ticket_id uuid PRIMARY KEY,
    movie text,
    price int,
    seat int,
    status text,
    timestamp timestamp,
    user_id uuid,
    PRIMARY KEY (ticket_id)
);

CREATE INDEX tickets_by_status ON tickets (status);

CREATE INDEX tickets_by_movie ON tickets (movie);
