const { Events } = require('discord.js');

const GuildService = require('../services/guildService.js')

module.exports = {
    name: Events.GuildCreate,
    once: true,
    execute(guild) {
        GuildService.createGuild(guild.id, guild.ownerId, guild.name, guild.joinedAt)
    },
};