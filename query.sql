-- name: GetPerson :one
SELECT * FROM mv_person
WHERE id = $1 LIMIT 1;

-- name: ListPeople :many
SELECT * FROM mv_person
ORDER BY first_name;

-- name: CreatePerson :one
INSERT INTO mv_person
(id, first_name, last_name, gender)
VALUES
($1, $2, $3, $4)
RETURNING *;

-- name: DeletePerson :exec
DELETE FROM mv_person
WHERE id = $1;

-- name: TotalMalePeople :one
SELECT COUNT(*) FROM mv_person
WHERE gender = 'M';