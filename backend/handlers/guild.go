package handlers

import (
	"context"
	"net/http"

	"github.com/A-Cer23/adminbot-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

const (
	insertGuildQuery = "INSERT INTO guild (guild_id, owner_id, guild_name, joined_at, in_guild) VALUES ($1, $2, $3, $4, $5)"
)

func CreateGuild(c *gin.Context, db *pgxpool.Pool) {
	var guild models.Guild

	if err := c.ShouldBindJSON(&guild); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Info().Str("GuildID", guild.GuildID).Msg("Creating new guild")

	_, err := db.Exec(
		context.Background(),
		insertGuildQuery,
		guild.GuildID, guild.OwnerID, guild.GuildName, guild.JoinedAt, guild.InGuild,
	)

	if err != nil {
		log.Error().Msg(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Problem creating guild", "guild_id": guild.GuildID})
	} else {
		log.Info().Str("GuildID", guild.GuildID).Msg("Guild created")
		c.JSON(http.StatusCreated, gin.H{"message": "Guild created successfully"})
	}

}
