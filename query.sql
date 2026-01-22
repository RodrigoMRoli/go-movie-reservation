-- name: GetPerson :one
SELECT * FROM mv_people
WHERE id = $1 LIMIT 1;

-- name: GetPeople :many
SELECT * FROM mv_people
ORDER BY first_name;

-- name: CreatePerson :one
INSERT INTO mv_people
(id, first_name, last_name, gender)
VALUES
($1, $2, $3, $4)
RETURNING *;

-- name: DeletePerson :exec
DELETE FROM mv_people
WHERE id = $1;

-- name: TotalMalePeople :one
SELECT COUNT(*) FROM mv_people
WHERE gender = 'M';