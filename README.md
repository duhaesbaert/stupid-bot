## The Stupid Bot
This is a Discord bot written in Go, with some predefined replies according to the messages identified on the channel, to make it somewhat fun, and also some actions that can be executed by the members of the channel. 

## Configuration
For the bot to work, it is necessary to add a ``config.json`` file on the same directory, which will load the configuration from. It is necessary to add three specific fields into this file:

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

IMPORTANT: When seting up the bot on discord portal, make sure the Message Content Intent is enalbed, so the bot can read messages from your channels. If not, the bot will not be able to read any messages from the channel, and will only interact when messaged directly.

## How the bot works
This bot uses a pseudo-random algorithm, which will 50% of the time reply back into your messages, to avoid spamming on every single conversation. There are some prefixed messages which can be used or changed, but as of this moment it is all hardcoded into the bot package.

## Commands Available
List of commands currently available:
````
/stop_listening -> this will make the bot stop listening(reading) the conversations and will not longer interact with it.
/start_listening -> makes the bot start listening(reading) the conversations and interact with it accordingly.
/play {game_name} -> for a pre-set of games, the bot will react with some specific message and put mention @here at the discord channel, inviting the users to play.
/poll {topic} -> This will add a poll, that will last for 5 minutes and show up up and down. After the 5 minutes have passed, the bot will post back the results into the channel.
````


