package types

import (
	"fmt"
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type NetSpeed struct {
	Upload   int64 `json:"upload"`
	Download int64 `json:"download"`
}

type Location struct {
	Latitude  int64  `json:"latitude"`
	Longitude int64  `json:"longitude"`
	City      string `json:"city"`
	Country   string `json:"country"`
}

type VPNDetails struct {
	Address    csdkTypes.AccAddress
	APIPort    int64
	Location   Location
	NetSpeed   NetSpeed
	EncMethod  string
	PricePerGB int64
	Version    string
	Status     string
	LockerID   string
}

type SessionDetails struct {
	VPNID         string
	ClientAddress csdkTypes.AccAddress
	GBToProvide   int64
	PricePerGB    int64
	Upload        int64
	Download      int64
	StartTime     *time.Time
	EndTime       *time.Time
	Status        string
}

func SessionsCountKey(accAddress csdkTypes.AccAddress) string {
	return fmt.Sprintf("sessions_count/%s", accAddress.String())
}

func VPNsCountKey(accAddress csdkTypes.AccAddress) string {
	return fmt.Sprintf("vpns_count/%s", accAddress.String())
}

const (
	MinLatitude  = -90 * 100000
	MinLongitude = -180 * 100000
	MaxLatitude  = 90 * 100000
	MaxLongitude = 180 * 100000

	KeyActiveNodeIDs    = "ACTIVE_NODE_IDS"
	KeyActiveSessionIDs = "ACTIVE_SESSION_IDS"
	KeySession          = "session"
	KeyVPN              = "vpn"

	StatusActive     = "ACTIVE"
	StatusDeregister = "DEREGISTERED"
	StatusEnd        = "ENDED"
	StatusInactive   = "INACTIVE"
	StatusRegister   = "REGISTERED"
	StatusStart      = "STARTED"
)