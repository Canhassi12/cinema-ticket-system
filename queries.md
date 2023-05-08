DROP TABLE tickets;

CREATE INDEX tickets_by_status ON tickets (status);

CREATE INDEX tickets_by_movie ON tickets (movie);

INSERT INTO tickets (ticket_id, movie, price, seat, status, user_id, timestamp) VALUES (1a76c0b5-08b0-416d-8ca5-255f92a0c77B, 'opa', 30, 12, 'free', '', '');

SELECT * FROM tickets;

POST  http://localhost:3000/reserve/oie
Content-Type: application/json

{
    "price": "half",
    "seat": 2,
    "movie": "canhassi",
    "ticketId": "7b527ca2-e377-11ed-b5ea-0242ac120002"
}

GET http://localhost:3000/ticket/1a76c0b5-08b0-416d-8ca5-255f92a0c76e
Content-Type: application/json