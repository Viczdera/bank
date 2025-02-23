-- name: CreateTransfer :one
INSERT INTO transfers(
    from_account,
    to_account,
    amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE _id = $1
LIMIT 1;
