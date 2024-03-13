package responseHelper

import "database/sql"

func NilOrValue(field interface{}) interface{} {
	switch v := field.(type) {
	case sql.NullString:
		if v.Valid {
			return v.String
		}
	case sql.NullInt32:
		if v.Valid {
			return v.Int32
		}
	case sql.NullInt64:
		if v.Valid {
			return v.Int64
		}
	case sql.NullFloat64:
		if v.Valid {
			return v.Float64
		}
	case sql.NullBool:
		if v.Valid {
			return v.Bool
		}
	case sql.NullTime:
		if v.Valid {
			return v.Time
		}
	}
	return nil
}
