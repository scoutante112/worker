package statsd

type Key string

const (
	KeyMessages      Key = "messages"
	KeyTickets       Key = "tickets"
	KeyCommands      Key = "commands"
	KeyJoins         Key = "joins"
	KeyLeaves        Key = "leaves"
	KeyRest          Key = "rest"
	KeyReconnect     Key = "reconnect"
	KeyIdentify      Key = "identify"
	KeySlashCommands Key = "slash_commands"
	KeyEvents        Key = "events"
	AutoClose        Key = "autoclose"
	KeyUnavailable   Key = "unavailable_guilds"
	KeyDirectMessage Key = "direct_message"
	KeyOpenCommand   Key = "open_command"
)

func (k Key) String() string {
	return string(k)
}
