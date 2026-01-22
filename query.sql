-- name: GetMovie :one
SELECT * FROM mv_movie
WHERE id = $1 LIMIT 1;

-- name: GetMovies :many
SELECT * FROM mv_movie
ORDER BY title;

-- name: CreateMovie :one
INSERT INTO mv_movie
(title, description, poster_image, poster_ext, release_date, language, country_origin)
VALUES
($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateMovie :exec
UPDATE mv_movie
SET 
    title = CASE WHEN @title_do_update::boolean
        THEN @title::VARCHAR(50) ELSE title END
WHERE
    id = @id
RETURNING *;

-- name: DeleteMovie :exec
DELETE FROM mv_movie
WHERE id = $1;
