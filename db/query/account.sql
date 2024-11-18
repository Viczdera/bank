-- name: CreateAccount :one
INSERT INTO accounts(
    owner,
    balance,
    currency
) VALUES (
    $1, $2 ,$3
) RETURNING * ;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE _id = $1
LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY _id
LIMIT $1
OFFSET $2;

-- name: DeleteAccount :exec
DELETE FROM accounts 
WHERE _id = $1;

-- name: UpdateAccount :one
UPDATE accounts 
SET balance = $2
WHERE _id = $1
RETURNING *;