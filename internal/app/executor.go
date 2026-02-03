package app

import (
	"fmt"
	"strconv"

	"parking-lot/internal/models"
	"parking-lot/internal/parser"
)

type Executor struct {
	parkingLot *models.ParkingLot
}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) ExecuteCommand(cmd *parser.Command) error {
	switch cmd.Type {

	case parser.CmdCreateParkingLot:
		return e.handleCreateParkingLot(cmd.Args)

	case parser.CmdPark:
		return e.handlePark(cmd.Args)

	case parser.CmdLeave:
		return e.handleLeave(cmd.Args)

	case parser.CmdStatus:
		return e.handleStatus()

	default:
		return fmt.Errorf("unsupported command: %s", cmd.Type)
	}
}

func (e *Executor) handleCreateParkingLot(args []string) error {
	capacity, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid capacity: %w", err)
	}

	pl, err := models.NewParkingLot(capacity)
	if err != nil {
		return err
	}

	e.parkingLot = pl
	fmt.Printf("Created a parking lot with %d slots\n", capacity)
	return nil
}

func (e *Executor) handlePark(args []string) error {
	if e.parkingLot == nil {
		return fmt.Errorf("parking lot has not been created")
	}

	return e.parkingLot.Park(args[0])
}

func (e *Executor) handleLeave(args []string) error {
	if e.parkingLot == nil {
		return fmt.Errorf("parking lot has not been created")
	}

	hours, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid hours: %w", err)
	}

	return e.parkingLot.Leave(args[0], hours)
}

func (e *Executor) handleStatus() error {
	if e.parkingLot == nil {
		return fmt.Errorf("parking lot has not been created")
	}

	e.parkingLot.Status()
	return nil
}
