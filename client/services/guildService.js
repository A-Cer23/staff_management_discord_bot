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


    fetch(request)
    // TODO: check if response is not 201. if its not 201 then log the error
}

GuildService = {
    createGuild,
}

module.exports = GuildService