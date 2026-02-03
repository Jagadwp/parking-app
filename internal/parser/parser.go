package parser

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	CmdCreateParkingLot = "create_parking_lot"
	CmdPark             = "park"
	CmdLeave            = "leave"
	CmdStatus           = "status"
)

type Command struct {
	Type string
	Args []string
}

func ParseCommand(line string) (*Command, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}

	tokens := strings.Fields(line)
	if len(tokens) == 0 {
		return nil, nil
	}

	cmdType := tokens[0]
	args := tokens[1:]

	switch cmdType {
	case CmdCreateParkingLot:
		if len(args) != 1 {
			return nil, fmt.Errorf("create_parking_lot requires exactly 1 argument (capacity), got %d", len(args))
		}
		// Validate capacity is a valid positive integer
		capacity, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, fmt.Errorf("invalid capacity: %s (must be an integer)", args[0])
		}
		if capacity <= 0 {
			return nil, fmt.Errorf("capacity must be positive, got: %d", capacity)
		}

	case CmdPark:
		if len(args) != 1 {
			return nil, fmt.Errorf("park requires exactly 1 argument (registration number), got %d", len(args))
		}
		// Registration number validation (basic check for non-empty)
		if strings.TrimSpace(args[0]) == "" {
			return nil, fmt.Errorf("registration number cannot be empty")
		}

	case CmdLeave:
		if len(args) != 2 {
			return nil, fmt.Errorf("leave requires exactly 2 arguments (registration number, hours), got %d", len(args))
		}
		// Validate registration number
		if strings.TrimSpace(args[0]) == "" {
			return nil, fmt.Errorf("registration number cannot be empty")
		}
		// Validate hours is a valid non-negative integer
		hours, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, fmt.Errorf("invalid hours: %s (must be an integer)", args[1])
		}
		if hours < 0 {
			return nil, fmt.Errorf("hours cannot be negative, got: %d", hours)
		}

	case CmdStatus:
		if len(args) != 0 {
			return nil, fmt.Errorf("status command takes no arguments, got %d", len(args))
		}

	default:
		return nil, fmt.Errorf("unknown command: %s", cmdType)
	}

	return &Command{
		Type: cmdType,
		Args: args,
	}, nil
}
