DROP TABLE IF EXISTS ST_ticket CASCADE;
DROP TABLE IF EXISTS ST_room_seats CASCADE;
DROP TABLE IF EXISTS ST_showtime_rooms CASCADE;
DROP TABLE IF EXISTS ST_showtime CASCADE;
DROP TABLE IF EXISTS SY_user_roles CASCADE;
DROP TABLE IF EXISTS MV_movie_genres CASCADE;
DROP TABLE IF EXISTS ST_seat CASCADE;
DROP TABLE IF EXISTS ST_room CASCADE;
DROP TABLE IF EXISTS SY_role CASCADE;
DROP TABLE IF EXISTS SY_user CASCADE;
DROP TABLE IF EXISTS MV_genre CASCADE;
DROP TABLE IF EXISTS MV_movie CASCADE;

CREATE TABLE MV_movie (
    id UUID PRIMARY KEY,
    title VARCHAR(50),
    description VARCHAR(255),
    poster_image VARCHAR(255),
    poster_ext VARCHAR(5),
    minutes INTEGER,
    release_date DATE,
    language VARCHAR(50),
    country_origin VARCHAR(5)
);

COMMENT ON COLUMN MV_movie.poster_image IS 'filename';
COMMENT ON COLUMN MV_movie.poster_ext IS 'extension';

CREATE TABLE MV_genre (
    id UUID PRIMARY KEY,
    title VARCHAR(50)
);

CREATE TABLE SY_user (
    id UUID PRIMARY KEY,
    email VARCHAR UNIQUE,
    password VARCHAR(255),
    salt_rounds INTEGER,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    birthdate DATE
);

CREATE TABLE SY_role (
    id UUID PRIMARY KEY,
    title VARCHAR(50)
);

CREATE TABLE ST_room (
    id UUID PRIMARY KEY,
    title VARCHAR(10)
);

CREATE TABLE ST_seat (
    id UUID PRIMARY KEY,
    code VARCHAR(5)
);

CREATE TABLE MV_movie_genres (
    movie_id UUID,
    genre_id UUID,
    CONSTRAINT fk_mv_genres_movie FOREIGN KEY (movie_id) REFERENCES MV_movie(id),
    CONSTRAINT fk_mv_genres_genre FOREIGN KEY (genre_id) REFERENCES MV_genre(id)
);

CREATE TABLE SY_user_roles (
    user_id UUID,
    role_id UUID,
    CONSTRAINT fk_sy_user_roles_user FOREIGN KEY (user_id) REFERENCES SY_user(id),
    CONSTRAINT fk_sy_user_roles_role FOREIGN KEY (role_id) REFERENCES SY_role(id)
);

CREATE TABLE ST_showtime (
    id UUID PRIMARY KEY,
    movie_id UUID,
    price DECIMAL(10,2),
    start_time TIMESTAMP, 
    end_time TIMESTAMP,   
    available_seats INTEGER,
    CONSTRAINT fk_st_showtime_movie FOREIGN KEY (movie_id) REFERENCES MV_movie(id)
);

CREATE TABLE ST_showtime_rooms (
    showtime_id UUID,
    room_id UUID,
    CONSTRAINT fk_st_showtime_rooms_showtime FOREIGN KEY (showtime_id) REFERENCES ST_showtime(id),
    CONSTRAINT fk_st_showtime_rooms_room FOREIGN KEY (room_id) REFERENCES ST_room(id)
);

CREATE TABLE ST_room_seats (
    room_id UUID,
    seat_id UUID,
    CONSTRAINT fk_st_room_seats_room FOREIGN KEY (room_id) REFERENCES ST_room(id),
    CONSTRAINT fk_st_room_seats_seat FOREIGN KEY (seat_id) REFERENCES ST_seat(id)
);

CREATE TABLE ST_ticket (
    id UUID PRIMARY KEY,
    showtime_id UUID,
    room_id UUID,
    seat_id UUID,
    user_id UUID,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    paid DECIMAL(10,2),
    
    CONSTRAINT fk_st_ticket_showtime FOREIGN KEY (showtime_id) REFERENCES ST_showtime(id),
    CONSTRAINT fk_st_ticket_room FOREIGN KEY (room_id) REFERENCES ST_room(id),
    CONSTRAINT fk_st_ticket_seat FOREIGN KEY (seat_id) REFERENCES ST_seat(id),
    CONSTRAINT fk_st_ticket_user FOREIGN KEY (user_id) REFERENCES SY_user(id),
    
    CONSTRAINT uq_ticket_selection UNIQUE (showtime_id, room_id, seat_id)
);