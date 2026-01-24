-- name: GetMovie :one
SELECT m.*, 
    COALESCE(
        (
            SELECT array_agg(g.title)
            FROM mv_genre g
            JOIN mv_movie_genres mg ON g.id = mg.genre_id
            WHERE mg.movie_id = m.id
        ), 
        '{}'::text[]
    )::text[] as genres
FROM mv_movie m
WHERE m.id = $1 LIMIT 1;

-- name: GetMovies :many
SELECT m.*, 
    COALESCE(
        (
            SELECT array_agg(g.title)
            FROM mv_genre g
            JOIN mv_movie_genres mg ON g.id = mg.genre_id
            WHERE mg.movie_id = m.id
        ), 
        '{}'::text[]
    )::text[] as genres
FROM mv_movie m
ORDER BY title;

-- name: CreateMovie :one
INSERT INTO mv_movie
(title, description, poster_image, poster_ext, minutes, release_date, language, country_origin)
VALUES
($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateMovie :exec
UPDATE mv_movie
SET 
    title          = COALESCE(sqlc.narg('title'), title),
    description    = COALESCE(sqlc.narg('description'), description),
    poster_image   = COALESCE(sqlc.narg('poster_image'), poster_image),
    poster_ext     = COALESCE(sqlc.narg('poster_ext'), poster_ext),
    minutes        = COALESCE(sqlc.narg('minutes'), minutes),
    release_date   = COALESCE(sqlc.narg('release_date'), release_date),
    language       = COALESCE(sqlc.narg('language'), language),
    country_origin = COALESCE(sqlc.narg('country_origin'), country_origin)
WHERE
    id = @id;

-- name: DeleteMovie :exec
DELETE FROM mv_movie
WHERE id = $1;

-- name: CreateGenre :one
INSERT INTO mv_genre 
(title)
VALUES
($1)
RETURNING *;

-- name: AddGenreToMovie :exec
INSERT INTO mv_movie_genres 
(movie_id, genre_id)
VALUES
($1, (SELECT g.id FROM mv_genre g WHERE g.title = $2));

-- name: RemoveGenreFromMovie :exec
DELETE FROM mv_movie_genres
WHERE movie_id = $1 AND genre_id = (SELECT g.id FROM mv_genre g WHERE g.title = $2);

-- name: DeleteGenre :exec
DELETE FROM mv_genre
WHERE id = $1;