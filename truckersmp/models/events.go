package models

type EventList struct {
	Featured []Event `json:"featured"`
	Today    []Event `json:"today"`
	Now      []Event `json:"now"`
	Upcoming []Event `json:"upcoming"`
}

type Event struct {
	ID        int `json:"id"`
	EventType struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"event_type"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Game   string `json:"game"`
	Server struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"server"`
	Language  string `json:"language"`
	Departure struct {
		Location string `json:"location"`
		City     string `json:"city"`
	} `json:"departure"`
	Arrive struct {
		Location string `json:"location"`
		City     string `json:"city"`
	} `json:"arrive"`
	MeetupAt     string `json:"meetup_at"`
	StartAt      string `json:"start_at"`
	Banner       string `json:"banner"`
	Map          string `json:"map"`
	Description  string `json:"description"`
	Rule         string `json:"rule"`
	VoiceLink    string `json:"voice_link"`
	ExternalLink string `json:"external_link"`
	Featured     string `json:"featured"`
	VTC          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"vtc"`
	User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
	Attendances struct {
		Confirmed int `json:"confirmed,string"`
		Unsure    int `json:"unsure,string"`
		VTCs      int `json:"vtcs"`
	}
	DLCs      map[string]string `json:"dlcs"`
	URL       string            `json:"url"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}
