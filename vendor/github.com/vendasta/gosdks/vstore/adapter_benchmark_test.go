package vstore

import (
	"testing"
	"time"

	"github.com/vendasta/gosdks/pb/vstorepb"
)

// original
// 20000	     65943 ns/op	    4481 B/op	     390 allocs/op
// 20000	     65301 ns/op	    4482 B/op	     390 allocs/op
// 20000	     67894 ns/op	    4481 B/op	     390 allocs/op
//
// new
// 100000	     10768 ns/op	    2024 B/op	      83 allocs/op
// 100000	     11066 ns/op	    2024 B/op	      83 allocs/op
// 100000	     11118 ns/op	    2024 B/op	      83 allocs/op
func Benchmark_StructPBToModel(b *testing.B) {
	b.ReportAllocs()
	RegisterModel("vstore", "Song", (*Song)(nil))
	now := time.Now().UTC()
	s := Song{
		Name:     "Morbid Dimensions",
		Duration: 636,
		Rating:   4.0,
		RecordedAt: &vstorepb.GeoPoint{
			Latitude:  50.0,
			Longitude: 60.0,
		},
		Released: now,
		Genres:   []string{"death metal", "osdm", "norwegian"},
		Artist: &Artist{
			Name: "Execration",
		},
	}
	spb, err := ModelToStructPB(&s)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := StructPBToModel("vstore", "Song", spb)
		if err != nil {
			b.Fatal(err)
		}
	}

}

// original
// 100000	     13837 ns/op	    3449 B/op	     120 allocs/op
// 100000	     13711 ns/op	    3449 B/op	     120 allocs/op
// 100000	     14472 ns/op	    3449 B/op	     120 allocs/op
//
// new
// 100000	     12714 ns/op	    2746 B/op	     116 allocs/op
// 100000	     12998 ns/op	    2746 B/op	     116 allocs/op
// 100000	     12918 ns/op	    2746 B/op	     116 allocs/op
//
func Benchmark_ModelToStructPB(b *testing.B) {
	RegisterModel("vstore", "Song", (*Song)(nil))
	now := time.Now().UTC()
	s := &Song{
		Name:     "Morbid Dimensions",
		Duration: 636,
		Rating:   4.0,
		RecordedAt: &vstorepb.GeoPoint{
			Latitude:  50.0,
			Longitude: 60.0,
		},
		Released: now,
		Genres:   []string{"death metal", "osdm", "norwegian"},
		Artist: &Artist{
			Name: "Execration",
		},
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m, _ := ModelToStructPB(s)
		if m == nil {
			b.FailNow()
		}
	}
}
