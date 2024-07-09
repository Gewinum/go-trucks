package models

// truckersmp api incosistency
type ServerResponseWrapper struct {
	Error      bool    `json:"error,string"`
	Descriptor *string `json:"descriptor"`
	Response   any     `json:"response"`
}

type Server struct {
	ID           int64  `json:"id"`
	Game         string `json:"game"`
	IP           string `json:"ip"`
	Port         int    `json:"port"`
	Name         string `json:"name"`
	ShortName    string `json:"shortname"`
	IdPrefix     string `json:"idprefix"`
	Online       bool   `json:"online"`
	Players      int64  `json:"players"`
	Queue        int64  `json:"queue"`
	MaxPlayers   int64  `json:"maxplayers"`
	MapID        int64  `json:"mapid"`
	DisplayOrder int64  `json:"displayorder"`
	// SpeedLimiter '1' means true and '0' means false
	SpeedLimiter         int   `json:"speedlimiter"`
	Collisions           bool  `json:"collisions"`
	CarsForPlayers       bool  `json:"carsforplayers"`
	PoliceCarsForPlayers bool  `json:"policecarsforplayers"`
	AfkEnabled           bool  `json:"afkenabled"`
	Event                bool  `json:"event"`
	SpecialEvent         bool  `json:"specialEvent"`
	ProMods              bool  `json:"promods"`
	SyncDelay            int64 `json:"syncdelay"`
}

type ServerGameTimeResponse struct {
	Error    bool  `json:"error"`
	GameTime int64 `json:"game_time"`
}
