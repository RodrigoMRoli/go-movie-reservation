package helpers

import (
	"database/sql"
	"time"
)

func StringPointerToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false} // SQL ignora (COALESCE mantém o valor antigo)
	}
	return sql.NullString{String: *s, Valid: true} // Atualiza (mesmo se for "")
}

func IntPointerToNullInt32(i *int) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: int32(*i), Valid: true}
}

func TimePointerToNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func SafeString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func SafeInt(ptr *int) int {
	if ptr == nil {
		return 0
	}
	return *ptr
}

func SafeTime(ptr *time.Time) time.Time {
	if ptr == nil {
		return time.Time{}
	}
	return *ptr
}
