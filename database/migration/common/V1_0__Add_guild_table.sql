
CREATE TABLE guild
(
    guild_id    bigserial   PRIMARY KEY,
    owner_id    bigserial   NOT NULL,
    guild_name  text        NOT NULL,
    joined_at  timestamp   NOT NULL,
    in_guild    boolean       NOT NULL
)
