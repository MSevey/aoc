package main

import "testing"

func TestRow(t *testing.T) {
	if row("FBFBBFF") != 44 {
		t.Fatal(row("FBFBBFF"))
	}
}
func TestSeat(t *testing.T) {
	if seat("RLR") != 5 {
		t.Fatal(seat("RLR"))
	}
}
func TestUID(t *testing.T) {
	u := uid("FBFBBFFRLR")
	if u != 357 {
		t.Fatal(u)
	}
	u = uid("BFFFBBFRRR")
	if u != 567 {
		t.Fatal(u)
	}
	u = uid("FFFBBBFRRR")
	if u != 119 {
		t.Fatal(u)
	}
	u = uid("BBFFBBFRLL")
	if u != 820 {
		t.Fatal(u)
	}
}
