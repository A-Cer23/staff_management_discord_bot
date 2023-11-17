const { Events } = require('discord.js');
const GuildService = require('../services/guildService.js')

module.exports = {
    name: Events.GuildDelete,
    async execute(guild) {
        GuildService.leaveGuild(guild.id)
    },
};