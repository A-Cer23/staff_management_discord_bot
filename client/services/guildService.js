const logger = require('../utlis/logger.js')


const GUILD_URL = process.env.API_URL + "/guild"

const createGuild = async (guild_id, owner_id, guild_name, joined_at) => {

    const data = {
        "guild_id": guild_id,
        "owner_id": owner_id,
        "guild_name": guild_name,
        "joined_at": joined_at,
        "in_guild": true,
    }


    const options = {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }

    const request = new Request(GUILD_URL, options);

    logger.info(`Sending create guild request for Guild ID: ${guild_id}`)

    fetch(request)
        .then(res => res.json())
        .then(jsonResponse => logger.info(JSON.stringify(jsonResponse)))
}

GuildService = {
    createGuild,
}

module.exports = GuildService