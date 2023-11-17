package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/A-Cer23/adminbot-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

const (
	insertGuildQuery = "INSERT INTO guild (guild_id, owner_id, guild_name, joined_at, in_guild) VALUES ($1, $2, $3, $4, $5)"
	leaveGuildQuery  = "UPDATE guild SET in_guild = $1 where guild_id = $2"
	guildExistsQuery = "SELECT EXISTS (SELECT 1 FROM guild WHERE guild_id = $1)"
	rejoinGuildQuery = "UPDATE guild SET joined_at = $1, in_guild = $2 WHERE guild_id = $3"
)

func CreateGuild(c *gin.Context, db *pgxpool.Pool) {
	var guild models.Guild

	if err := c.ShouldBindJSON(&guild); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var guild_exists bool

	err := db.QueryRow(context.Background(), guildExistsQuery, guild.GuildID).Scan(&guild_exists)

	if err != nil {
		log.Error().Err(err).Msg("Error in guild query row")
	}

	if !guild_exists {
		log.Info().Str("GuildID", guild.GuildID).Msg("Creating new guild")

		_, err := db.Exec(
			context.Background(),
			insertGuildQuery,
			guild.GuildID, guild.OwnerID, guild.GuildName, guild.JoinedAt, guild.InGuild,
		)

		if err != nil {
			log.Error().Msg(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Problem creating guild", "guild_id": guild.GuildID})
			return
		} else {
			log.Info().Str("GuildID", guild.GuildID).Msg("Guild created")
			c.JSON(http.StatusCreated, gin.H{"message": "Guild created successfully"})
			return
		}
	} else {
		log.Info().Str("GuildID", guild.GuildID).Msg("Re-joining guild")
		_, err := db.Exec(context.Background(), rejoinGuildQuery, guild.JoinedAt, true, guild.GuildID)
		if err != nil {
			log.Error().Str("GuildID", guild.GuildID).Msg("Guild ID already exists but failed to update in guild status")
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update in guild status for Guild ID %s", guild.GuildID)})
			return
		}
		log.Info().Str("GuildID", guild.GuildID).Msg("Successfully re-joined guild")
	}
}

func LeaveGuild(c *gin.Context, db *pgxpool.Pool) {

	guildId, guild_id_exists := c.Params.Get("id")

	if !guild_id_exists {
		log.Error().Msg(`Guild ID was not received to update.`)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing param guild id"})
		return
	}

	var guild_exists bool

	err := db.QueryRow(context.Background(), guildExistsQuery, guildId).Scan(&guild_exists)

	if err != nil {
		log.Error().Str(`Error`, "Issue querying to exit guild")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "issue in searching guild"})
		return
	}

	if !guild_exists {
		log.Error().Str("GuildID", guildId).Msg("Guild Id does not exist in the database")
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Guild ID %s does not exist", guildId)})
		return
	}

	_, err = db.Exec(context.Background(), leaveGuildQuery, false, guildId)

	if err != nil {
		log.Error().Msg("Error: Failed to leave a guild")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "issue in leaving guild"})
		return
	}

	log.Info().Msg(fmt.Sprintf("Left Guild %s", guildId))
	c.JSON(http.StatusOK, nil)
}
