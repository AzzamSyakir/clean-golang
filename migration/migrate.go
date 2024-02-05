package migration

import (
	"database/sql"

	"github.com/rs/zerolog/log"
)

func MigrationDb(db *sql.DB) error {
	// Memeriksa koneksi ke database
	err := db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Gagal melakukan ping ke database")
	}
	// Panggil fungsi migration untuk inisialisasi migrasi database
	if err := Usermigration(db); err != nil {
		log.Error().Err(err).Msg("Gagal melakukan migrasi user")
	}

	if err := Tokenmigration(db); err != nil {
		log.Error().Err(err).Msg("Gagal melakukan migrasi token")
	}

	return err
}