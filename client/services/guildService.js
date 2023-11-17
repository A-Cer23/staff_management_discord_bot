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

    logger.info(`Joining Guild ID: ${guild_id}`)

    fetch(request)
        .then(res => {
            if (!res.ok) {
                logger.error(`Guild create response was not okay \nResponse: ${res.statusText}`)
                return
            }
            logger.info(`Joined Guild ID: ${guild_id}`)
        })
}


const leaveGuild = (guild_id) => {
    const options = {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
    }

    const request = new Request(GUILD_URL + `/${guild_id}`, options)

    logger.info(`Leaving Guild ID: ${guild_id}`)

    fetch(request)
        .then(res => {
            if (!res.ok) {
                logger.error(`Leaving guild response was not okay \nResponse: ${res.statusText}`)
                return
            }
            logger.info(`Left Guild ID: ${guild_id}`)
        })

}

GuildService = {
    createGuild,
    leaveGuild
}

module.exports = GuildService