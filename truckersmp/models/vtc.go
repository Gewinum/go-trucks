package models

type VTCList struct {
	Recent        []ShortVTCInfo `json:"recent"`
	Featured      []ShortVTCInfo `json:"featured"`
	FeaturedCover []ShortVTCInfo `json:"featured_cover"`
}

type VTCSocials struct {
	Twitter  string `json:"twitter"`
	Facebook string `json:"facebook"`
	Twitch   string `json:"twitch"`
	Discord  string `json:"discord"`
	Youtube  string `json:"youtube"`
}

type VTCGameList struct {
	ATS bool `json:"ats"`
	ETS bool `json:"ets"`
}

type ShortVTCInfo struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	OwnerId       int         `json:"owner_id"`
	OwnerUsername string      `json:"owner_username"`
	Slogan        string      `json:"slogan"`
	Tag           string      `json:"tag"`
	Website       string      `json:"website"`
	Socials       VTCSocials  `json:"socials"`
	Games         VTCGameList `json:"games"`
	MembersCount  int         `json:"members_count"`
	Recruitment   string      `json:"recruitment"`
	Language      string      `json:"language"`
	Verified      bool        `json:"verified"`
	Validated     bool        `json:"validated"`
	Created       string      `json:"created"`
}

type DetailedVTCInfo struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	OwnerID       int         `json:"owner_id"`
	OwnerUsername string      `json:"owner_username"`
	Slogan        string      `json:"slogan"`
	Tag           string      `json:"tag"`
	Logo          string      `json:"logo"`
	Cover         string      `json:"cover"`
	Information   string      `json:"information"`
	Rules         string      `json:"rules"`
	Requirements  string      `json:"requirements"`
	Website       string      `json:"website"`
	Socials       VTCSocials  `json:"socials"`
	Games         VTCGameList `json:"games"`
	MembersCount  int         `json:"members_count"`
	Recruitment   string      `json:"recruitment"`
	Language      string      `json:"language"`
	Verified      bool        `json:"verified"`
	Validated     bool        `json:"validated"`
	Created       string      `json:"created"`
}

type VTCNewsList struct {
	News []struct {
		ID             int    `json:"id"`
		Title          string `json:"title"`
		ContentSummary string `json:"content_summary"`
		AuthorID       int    `json:"author_id"`
		Author         string `json:"author"`
		Pinned         bool   `json:"pinned"`
		UpdatedAt      string `json:"updated_at"`
		PublishedAt    string `json:"published_at"`
	} `json:"news"`
}

type VTCDetailedNews struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	ContentSummary string `json:"content_summary"`
	Content        string `json:"content"`
	AuthorID       int    `json:"author_id"`
	Author         string `json:"author"`
	Pinned         bool   `json:"pinned"`
	UpdatedAt      string `json:"updated_at"`
	PublishedAt    string `json:"published_at"`
}

type VTCRoleList struct {
	Roles []VTCRole `json:"roles"`
}

type VTCRole struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
	Owner     bool   `json:"owner"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// VTCPlayerInfo Warning: ETS2 has duplicate of SteamID which is presented as a string
type VTCPlayerInfo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	SteamID   int    `json:"steam_id"`
	SteamID64 int64  `json:"steamID64"`
	RoleID    int    `json:"role_id"`
	Role      string `json:"role"`
	IsOwner   bool   `json:"is_owner"`
	JoinDate  string `json:"join_date"`
}

type VTCMemberList struct {
	Members []VTCPlayerInfo `json:"members"`
}
