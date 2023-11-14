const { Events } = require('discord.js');

const GuildService = require('../services/guildService.js')

module.exports = {
    name: Events.GuildCreate,
    execute(guild) {
        GuildService.createGuild(guild.id, guild.ownerId, guild.name, guild.joinedAt)
    },
};