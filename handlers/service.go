package handlers

import "crypto-service/internal/client"

// Service struct holds all variables common to all handlers.
type Service struct {
	Client client.HITBTC
}
