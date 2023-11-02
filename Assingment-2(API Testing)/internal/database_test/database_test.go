package main

import (
	"golang/internal/database"
	"testing"
)

func TestOpen(t *testing.T) {
	testCases := []struct {
		dns string
	}
	{
		{"host=localhost user=postgres password=admin1 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"},
		{"host=localhost user=postgres password=admin2 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"},
		{"host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"},
		{"host=localhost user=postgres password=admin3 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"},
		{"host=localhost user=postgres password=admin4 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"},
	}

	for _, tc := range testCases {
		_, err := database.Open(tc.dns)
		if err != nil {
			t.Errorf("database.Open(%s)", tc.dns)
		}
	}

}
