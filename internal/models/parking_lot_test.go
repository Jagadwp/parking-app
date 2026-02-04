package models

import (
	"fmt"
	"testing"
)

// TestNewParkingLot tests parking lot creation
func TestNewParkingLot(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		wantErr  bool
	}{
		{"Valid capacity", 6, false},
		{"Large capacity", 10000, false},
		{"Zero capacity", 0, true},
		{"Negative capacity", -5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl, err := NewParkingLot(tt.capacity)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewParkingLot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && pl == nil {
				t.Error("NewParkingLot() returned nil parking lot")
			}
			if !tt.wantErr && pl.capacity != tt.capacity {
				t.Errorf("NewParkingLot() capacity = %v, want %v", pl.capacity, tt.capacity)
			}
		})
	}
}

// TestParkAndLeave tests basic park and leave operations
func TestParkAndLeave(t *testing.T) {
	pl, err := NewParkingLot(3)
	if err != nil {
		t.Fatalf("Failed to create parking lot: %v", err)
	}

	// Test parking
	cars := []string{"KA-01-HH-1234", "KA-01-HH-9999", "KA-01-BB-0001"}
	for i, car := range cars {
		err := pl.Park(car)
		if err != nil {
			t.Errorf("Park(%s) failed: %v", car, err)
		}
		// Verify car is tracked
		if _, exists := pl.carToSlot[car]; !exists {
			t.Errorf("Car %s not found in tracking after parking", car)
		}
		// Verify slot is occupied
		expectedSlot := i + 1
		if pl.occupiedSlots[expectedSlot] != car {
			t.Errorf("Slot %d should contain %s, got %s", expectedSlot, car, pl.occupiedSlots[expectedSlot])
		}
	}

	// Test parking lot full
	err = pl.Park("DL-12-AA-9999")
	if err != nil {
		t.Errorf("Park should not return error when full, got: %v", err)
	}

	// Test leave
	err = pl.Leave("KA-01-HH-9999", 4)
	if err != nil {
		t.Errorf("Leave failed: %v", err)
	}

	// Verify slot is freed
	if _, exists := pl.carToSlot["KA-01-HH-9999"]; exists {
		t.Error("Car should be removed from tracking after leaving")
	}

	// Test leave non-existent car
	err = pl.Leave("NON-EXISTENT", 2)
	if err != nil {
		t.Errorf("Leave should not return error for non-existent car, got: %v", err)
	}
}

// TestSlotReuse tests that freed slots are reused (nearest first)
func TestSlotReuse(t *testing.T) {
	pl, err := NewParkingLot(5)
	if err != nil {
		t.Fatalf("Failed to create parking lot: %v", err)
	}

	// Park cars in slots 1-5
	cars := []string{"CAR1", "CAR2", "CAR3", "CAR4", "CAR5"}
	for _, car := range cars {
		pl.Park(car)
	}

	// Leave cars from slot 2 and 4
	pl.Leave("CAR2", 2)
	pl.Leave("CAR4", 2)

	// Park new car - should get slot 2 (smallest available)
	pl.Park("CAR6")
	if pl.carToSlot["CAR6"] != 2 {
		t.Errorf("CAR6 should be in slot 2 (nearest), got slot %d", pl.carToSlot["CAR6"])
	}

	// Park another car - should get slot 4
	pl.Park("CAR7")
	if pl.carToSlot["CAR7"] != 4 {
		t.Errorf("CAR7 should be in slot 4 (next nearest), got slot %d", pl.carToSlot["CAR7"])
	}
}

// TestCalculateCharge tests the charge calculation logic
func TestCalculateCharge(t *testing.T) {
	tests := []struct {
		hours int
		want  int
	}{
		{1, 10},  // First 2 hours
		{2, 10},  // Exactly 2 hours
		{3, 20},  // 2 hours base + 1 additional
		{4, 30},  // 2 hours base + 2 additional
		{5, 40},  // 2 hours base + 3 additional
		{6, 50},  // 2 hours base + 4 additional
		{10, 90}, // 2 hours base + 8 additional
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d_hours", tt.hours), func(t *testing.T) {
			got := calculateCharge(tt.hours)
			if got != tt.want {
				t.Errorf("calculateCharge(%d) = %d, want %d", tt.hours, got, tt.want)
			}
		})
	}
}

func TestZeroHours(t *testing.T) {
	charge := calculateCharge(0)
	if charge != 10 {
		t.Errorf("0 hours should cost $10, got $%d", charge)
	}
}

// TestDuplicateParking tests that same car cannot park twice
func TestDuplicateParking(t *testing.T) {
	pl, err := NewParkingLot(5)
	if err != nil {
		t.Fatalf("Failed to create parking lot: %v", err)
	}

	// Park car first time
	err = pl.Park("KA-01-HH-1234")
	if err != nil {
		t.Errorf("First park failed: %v", err)
	}

	// Try to park same car again
	err = pl.Park("KA-01-HH-1234")
	if err == nil {
		t.Error("Expected error when parking same car twice, got nil")
	}
}
