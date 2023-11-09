package models

type Guild struct {
	GuildID   string `json:"guild_id"`
	OwnerID   string `json:"owner_id"`
	GuildName string `json:"guild_name"`
	JoinedAt  string `json:"joined_at"`
	InGuild   bool   `json:"in_guild"`
}
