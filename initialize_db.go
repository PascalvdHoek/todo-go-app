package main

import "database/sql"

func createAlbumTable(db *sql.DB) error {
	_, err := db.Exec(
		`CREATE TABLE albums (
    id INTEGER NOT NULL UNIQUE ,
    title VARCHAR NOT NULL,
    artist VARCHAR NOT NULL,
    price VARCHAR NOT NULL,
);`)

	if err != nil {
		return err
	}
	return nil
}

func insertAlbums(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO albums (
                    id,
    title,
    artist,
    price
)
VALUES
    (1,
        'Blue Train',
        'John Coltrane',
        '12.50'
    ),(2,
        'Blue truck',
        'John 123',
        '25.00'
    ),(3,
        'Green Train',
        'John 345',
        '50.00'
    )`)

	if err != nil {
		return err
	}
	return nil
}
