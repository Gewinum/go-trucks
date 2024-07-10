package models

type GameChecksum struct {
	DLL string `json:"dll"`
	ADB string `json:"adb"`
}

type GameInfo struct {
	Name                     string       `json:"name"`
	Numeric                  string       `json:"numeric"`
	Stage                    string       `json:"stage"`
	ETS2MPChecksum           GameChecksum `json:"ets2mp_checksum"`
	ATSMPChecksum            GameChecksum `json:"atsmp_checksum"`
	Time                     string       `json:"time"`
	SupportedETS2GameVersion string       `json:"supported_game_version"`
	SupportedATSGameVersion  string       `json:"supported_ats_game_version"`
}

type RulesInfo struct {
	Rules    string `json:"rules"`
	Revision int64  `json:"revision"`
}
