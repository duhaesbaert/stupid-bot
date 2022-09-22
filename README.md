## The Stupid Bot
This is a small Discord bot written in Go for the sole purpose of having a bot of my own.

## Configuration
For this bot to work, it is necessary to add a ``config.json`` file on the same directory, which will load the configuration from. It is necessary to add three specific fields into this file:

1. Token: This is your bot token, that can be found at ``discord.gg/Developers``
2. BotPrefix: The prefix of your bot
3. BotListening: Indicator if the bot will be listening for messages upon startup. This can be changed by sending ``/start_listening`` or ``/stop_listenint`` directly to your bot.

Example:
````
{
  "Token": "my-bot-token",
  "BotPrefix": "stupid-bot",
  "BotListening": true
}
````

One important setting when creating your bot on discord, is to make sure the Message Content Intent is flagged, so the bot can read messages from your channels.

## How the bot works
This bot uses a pseudo-random algorithm, which will 50% of the time reply back into your messages. There are some prefixed messages which can be used or changed, but as of this moment it is all hardcoded into the bot package.


