package main

import (
	"math/rand"
	"time"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
)
const (
	internetConnectionFeature = 1;
    bathroomFeature = 2;
    heatingFeature = 3;
    airConditionerFeature = 4;
    electricalOutletFeature = 5;
    tapFeature = 6;
    bedLinensFeature = 7;
    pillowsFeature = 8;
)
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GenRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
	  b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func isDateValidAndFromTomorrow(d *grpc_gen.Date) bool {
	if d.Year < 1000 && d.Year >= 3000 && d.Month < 1 && d.Month > 12 && d.Day < 1 && d.Day > 31 {
		return false
	}
	date := time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.Local)
	if date.Year() != int(d.Year) || int(date.Month()) != int(d.Month) || date.Day() != int(d.Day) {
		return false
	}
	return time.Now().Before(date)
}

// Assumed that isDateValidAndFromTomorrow(d) is true 
func flatterizeDate(d *grpc_gen.Date) int32 {
	return int32(d.Year*10000+d.Month*100+d.Day)
}
func deflatterizeDate(s int32) *grpc_gen.Date {
	d := s%100
	m := (s%10000)/100
	y := s/10000
	return &grpc_gen.Date{
		Day: uint32(d),
		Month: uint32(m),
		Year: uint32(y),
	}
}
func featuresToInts(s []grpc_gen.Feature) []int32 {
	res := make([]int32, len(s))
	for i, f := range s {
		res[i] = int32(f)
	}
	return res
}
func intsTofeatures(s []int32) []grpc_gen.Feature {
	res := make([]grpc_gen.Feature, len(s))
	for i, f := range s {
		res[i] = grpc_gen.Feature(f)
	}
	return res
}

func allDistinct[T comparable](s []T) bool {
	m := make(map[T]struct{})
	for _, v := range s {
		_, found := m[v]
		if found {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}