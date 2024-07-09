package models

type PlayerPatreonInformation struct {
	IsPatron       bool   `json:"is_patron"`
	Active         bool   `json:"active"`
	Color          string `json:"color"`
	TierID         int    `json:"tier_id"`
	CurrentPledge  int    `json:"current_pledge"`
	LifetimePledge int    `json:"lifetime_pledge"`
	NextPledge     int    `json:"next_pledge"`
	Hidden         bool   `json:"hidden"`
}

type PlayerPermissions struct {
	IsStaff               bool `json:"isStaff"`
	IsUpperStaff          bool `json:"isUpperStaff"`
	IsGameAdmin           bool `json:"isGameAdmin"`
	ShowDetailedOnWebMaps bool `json:"showDetailedOnWebMaps"`
}

type PlayerVTC struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Tag      string `json:"tag"`
	InVTC    bool   `json:"inVTC"`
	MemberID int64  `json:"member_id"`
}

type PlayerVTCHistory struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Verified bool   `json:"verified"`
	JoinDate string `json:"join_date"`
	LeftDate string `json:"left_date"`
}

type Player struct {
	ID                int64                    `json:"id"`
	Name              string                   `json:"name"`
	Avatar            string                   `json:"avatar"`
	SmallAvatar       string                   `json:"smallAvatar"`
	JoinDate          string                   `json:"joinDate"`
	SteamID64         int64                    `json:"steamID64"`
	SteamID           string                   `json:"steamID"`
	DiscordSnowFlake  string                   `json:"discordSnowFlake"`
	DisplayVTCHistory bool                     `json:"displayVTCHistory"`
	GroupName         string                   `json:"groupName"`
	GroupColor        string                   `json:"groupColor"`
	GroupID           int64                    `json:"groupID"`
	Banned            bool                     `json:"banned"`
	BannedUntil       string                   `json:"bannedUntil"`
	BansCount         int                      `json:"bansCount"`
	DisplayBans       bool                     `json:"displayBans"`
	Patreon           PlayerPatreonInformation `json:"patreon"`
	Permissions       PlayerPermissions        `json:"permissions"`
	VTC               PlayerVTC                `json:"vtc"`
	VTCHistory        []PlayerVTCHistory       `json:"vtcHistory"`
}

type PlayerBan struct {
	Expiration string `json:"expiration"`
	TimeAdded  string `json:"timeAdded"`
	Active     bool   `json:"active"`
	Reason     string `json:"reason"`
	AdminName  string `json:"adminName"`
	AdminID    int64  `json:"adminID"`
}
