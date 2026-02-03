package models

import (
	"container/heap"
	"fmt"
)

// ParkingLot represents a parking lot with automated ticketing system.
type ParkingLot struct {
	capacity       int            // Total number of parking slots
	availableSlots *MinHeap       // Min heap to efficiently get nearest available slot
	occupiedSlots  map[int]string // Maps slot number to vehicle registration number
	carToSlot      map[string]int // Maps vehicle registration to slot number (for O(1) lookup)
}

const (
	BaseCharge           = 10 // $10 for first 2 hours
	BaseHours            = 2  // First 2 hours covered by base charge
	AdditionalChargeRate = 10 // $10 per additional hour after first 2 hours
)

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("parking lot capacity must be positive, got: %d", capacity)
	}

	availableSlots := &MinHeap{}
	for i := 1; i <= capacity; i++ {
		*availableSlots = append(*availableSlots, i)
	}

	// Initialize the min heap from the slice
	heap.Init(availableSlots)

	return &ParkingLot{
		capacity:       capacity,
		availableSlots: availableSlots,
		occupiedSlots:  make(map[int]string),
		carToSlot:      make(map[string]int),
	}, nil
}

// Park allocates a parking slot to a vehicle.
// It always allocates the nearest available slot (smallest slot number).
func (p *ParkingLot) Park(registrationNumber string) error {
	// Edge case: Check if car is already parked
	if _, exists := p.carToSlot[registrationNumber]; exists {
		return fmt.Errorf("vehicle %s is already parked", registrationNumber)
	}

	// Check if parking lot is full
	if p.availableSlots.Len() == 0 {
		fmt.Println("Sorry, parking lot is full")
		return nil
	}

	// Get the nearest available slot (minimum from heap) - O(log n)
	slot := heap.Pop(p.availableSlots).(int)

	// Update state: mark slot as occupied
	p.occupiedSlots[slot] = registrationNumber
	p.carToSlot[registrationNumber] = slot

	fmt.Printf("Allocated slot number: %d\n", slot)
	return nil
}

// Leave removes a vehicle from the parking lot and calculates the parking charge.
func (p *ParkingLot) Leave(registrationNumber string, hours int) error {
	// Validate hours
	if hours < 0 {
		return fmt.Errorf("parking hours cannot be negative: %d", hours)
	}

	// Check if car exists in parking lot
	slot, exists := p.carToSlot[registrationNumber]
	if !exists {
		fmt.Printf("Registration number %s not found\n", registrationNumber)
		return nil
	}

	// Calculate parking charge
	charge := calculateCharge(hours)

	// Return the slot to available pool - O(log n)
	heap.Push(p.availableSlots, slot)

	// Update state: remove vehicle from tracking
	delete(p.occupiedSlots, slot)
	delete(p.carToSlot, registrationNumber)

	// Output as per specification (exact format matters!)
	fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n",
		registrationNumber, slot, charge)

	return nil
}

// Status displays the current state of the parking lot.
func (p *ParkingLot) Status() {
	fmt.Println("Slot No. Registration No.")

	// Iterate through all slots in order (1 to capacity)
	// Only print occupied slots
	for slot := 1; slot <= p.capacity; slot++ {
		if registrationNumber, occupied := p.occupiedSlots[slot]; occupied {
			fmt.Printf("%d %s\n", slot, registrationNumber)
		}
	}
}

// calculateCharge computes the parking charge based on hours parked.
func calculateCharge(hours int) int {
	if hours <= BaseHours {
		return BaseCharge
	}

	// Additional hours beyond the base period
	additionalHours := hours - BaseHours
	return BaseCharge + (additionalHours * AdditionalChargeRate)
}
