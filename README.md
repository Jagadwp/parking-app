# Parking Lot Management System

A parking lot management system with automated slot allocation and charge calculation.

## Algorithm

Uses **Min Heap** for slot allocation instead of simpler alternatives.

**Why Min Heap?**
- **Linear Search**: O(n) - slow for large parking lots
- **Sorted Array**: O(n) insertion - bottleneck when returning slots
- **BST (TreeSet)**: O(log n) but requires complex balancing logic
- **Min Heap**: O(log n) for both operations - optimal & simpler ✓

**Performance:**
- Park: O(log n)
- Leave: O(log n)
- Status: O(n)

For 10,000 slots: ~13 operations vs up to 10,000 with linear search.

## Features

- Nearest slot allocation (always assigns the smallest available slot number)
- Automatic charge calculation ($10 for first 2 hours, $10 per additional hour)
- Efficient slot management using Min Heap

## Requirements

- Go 1.21 or higher

## Installation

```bash
# Clone or download the project
cd parking-app

# Build & run example 
go build -o parking-app ./cmd/main.go
./parking-app testdata/input.txt

# Run without build example
go run cmd/main.go testdata/input.txt
```


### Input File Format

```
create_parking_lot 6
park KA-01-HH-1234
park KA-01-HH-9999
leave KA-01-HH-1234 4
status
```

## Commands

| Command | Description | Example |
|---------|-------------|---------|
| `create_parking_lot <capacity>` | Creates parking lot with given capacity | `create_parking_lot 6` |
| `park <registration_number>` | Parks a vehicle | `park KA-01-HH-1234` |
| `leave <registration_number> <hours>` | Vehicle leaves, calculates charge | `leave KA-01-HH-1234 4` |
| `status` | Shows all currently parked vehicles | `status` |

## Example

```bash
$ go run cmd/main.go testdata/input.txt
Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30
Slot No.    Registration No.
2           KA-01-HH-9999
```

## Charge Calculation

- Hours ≤ 2: $10
- Hours > 2: $10 + ($10 × additional hours)

Examples:
- 2 hours = $10
- 4 hours = $30
- 6 hours = $50

## Project Structure

```
parking-app/
├── cmd/
│   └── main.go                   # Entry point
├── internal/
│   ├── app/
│   │   └── executor.go           # Command executor
│   ├── models/
│   │   ├── min_heap.go           # Min heap implementation
│   │   └── parking_lot.go        # Core business logic
│   │   └── parking_lot_test.go   # Unit Test
│   └── parser/
│       └── parser.go             # Command parser
└── testdata/
    └── input.txt                 # Sample input
```

## Testing

```bash
# Run tests
go test ./internal/models -v
```

## Author
Jagad Wijaya Purnomo