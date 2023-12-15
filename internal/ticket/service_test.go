package ticket

import (
	"context"
	"testing"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

var tickets = []domain.Ticket{
	{
		Id:      "1",
		Name:    "Tait Mc Caughan",
		Email:   "tmc0@scribd.com",
		Country: "Finland",
		Time:    "17:11",
		Price:   785.00,
	},
	{
		Id:      "2",
		Name:    "Padget McKee",
		Email:   "pmckee1@hexun.com",
		Country: "China",
		Time:    "20:19",
		Price:   537.00,
	},
	{
		Id:      "3",
		Name:    "Yalonda Jermyn",
		Email:   "yjermyn2@omniture.com",
		Country: "China",
		Time:    "18:11",
		Price:   579.00,
	},
}

var ticketsByDestination = []domain.Ticket{
	{
		Id:      "2",
		Name:    "Padget McKee",
		Email:   "pmckee1@hexun.com",
		Country: "China",
		Time:    "20:19",
		Price:   537.00,
	},
	{
		Id:      "3",
		Name:    "Yalonda Jermyn",
		Email:   "yjermyn2@omniture.com",
		Country: "China",
		Time:    "18:11",
		Price:   579.00,
	},
}

type stubRepo struct {
	db *DbMock
}

type DbMock struct {
	db  []domain.Ticket
	spy bool
	err error
}

func NewRepositoryTest(dbM *DbMock) TicketRepository {
	return &stubRepo{dbM}
}

// Func for test Average Destination
func (r *stubRepo) AverageDestination(ctx context.Context, destination string) (float64, error) {
	r.db.spy = true
	if r.db.err != nil {
		return 0, r.db.err
	}

	var total float64
	var count float64

	for _, t := range r.db.db {
		if t.Country == destination {
			total += t.Price
			count++
		}
	}

	return total / count, nil
}

// Func for test GetTotalTickets
func (r *stubRepo) GetTotalTickets(ctx context.Context, destination string) (float64, error) {
	r.db.spy = true
	if r.db.err != nil {
		return 0, r.db.err
	}

	var count float64

	for _, t := range r.db.db {
		if t.Country == destination {
			count++
		}
	}

	return count, nil
}

// Func TestGetAll
func TestGetAll(t *testing.T) {
	dbMock := &DbMock{
		db:  tickets,
		spy: false,
		err: nil,
	}

	repo := NewRepositoryTest(dbMock)
	service := NewTicketService(repo)

	tkts, err := service.GetAll(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, tkts)
	assert.True(t, dbMock.spy)
}

func (r *stubRepo) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}
	return r.db.db, nil
}

func (r *stubRepo) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var tkts []domain.Ticket

	r.db.spy = true
	if r.db.err != nil {
		return []domain.Ticket{}, r.db.err
	}

	for _, t := range r.db.db {
		if t.Country == destination {
			tkts = append(tkts, t)
		}
	}

	return tkts, nil
}

// Func TestGetTicketByDestination
func TestGetTicketByDestination(t *testing.T) {
	dbMock := &DbMock{
		db:  tickets,
		spy: false,
		err: nil,
	}

	repo := NewRepositoryTest(dbMock)
	service := NewTicketService(repo)

	tkts, err := service.GetTicketByDestination(ctx, "China")

	assert.Nil(t, err)
	assert.NotNil(t, tkts)
	assert.True(t, dbMock.spy)
	assert.Equal(t, len(ticketsByDestination), len(tkts))
}

// Func TestAverageDestination
func TestAverageDestination(t *testing.T) {
	dbMock := &DbMock{
		db:  tickets,
		spy: false,
		err: nil,
	}

	repo := NewRepositoryTest(dbMock)
	service := NewTicketService(repo)

	avr, err := service.AverageDestination(ctx, "China")

	assert.Nil(t, err)
	assert.NotNil(t, avr)
	assert.True(t, dbMock.spy)
}

func TestGetTotalTickets(t *testing.T) {
	dbMock := &DbMock{
		db:  tickets,
		spy: false,
		err: nil,
	}

	repo := NewRepositoryTest(dbMock)
	service := NewTicketService(repo)

	avr, err := service.AverageDestination(ctx, "China")

	assert.Nil(t, err)
	assert.NotNil(t, avr)
	assert.True(t, dbMock.spy)
}
