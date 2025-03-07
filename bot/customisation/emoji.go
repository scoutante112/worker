package customisation

import (
	"fmt"

	"github.com/rxdn/gdl/objects"
	"github.com/rxdn/gdl/objects/guild/emoji"
)

type CustomEmoji struct {
	Name     string
	Id       uint64
	Animated bool
}

func NewCustomEmoji(name string, id uint64, animated bool) CustomEmoji {
	return CustomEmoji{
		Name: name,
		Id:   id,
	}
}

func (e CustomEmoji) String() string {
	if e.Animated {
		return fmt.Sprintf("<a:%s:%d>", e.Name, e.Id)
	} else {
		return fmt.Sprintf("<:%s:%d>", e.Name, e.Id)
	}
}

func (e CustomEmoji) BuildEmoji() *emoji.Emoji {
	return &emoji.Emoji{
		Id:       objects.NewNullableSnowflake(e.Id),
		Name:     e.Name,
		Animated: e.Animated,
	}
}

var (
	EmojiId         = NewCustomEmoji("id", 1334818151036358666, false)
	EmojiOpen       = NewCustomEmoji("open", 1334818181612703775, false)
	EmojiOpenTime   = NewCustomEmoji("opentime", 1334824967942049792, false)
	EmojiClose      = NewCustomEmoji("close", 1334824949885566976, false)
	EmojiCloseTime  = NewCustomEmoji("closetime", 1334824929161510932, false)
	EmojiReason     = NewCustomEmoji("reason", 1334824897683259447, false)
	EmojiSubject    = NewCustomEmoji("subject", 1334824880528691251, false)
	EmojiTranscript = NewCustomEmoji("transcript", 1334824861385887814, false)
	EmojiClaim      = NewCustomEmoji("claim", 1334824841336848406, false)
	EmojiPanel      = NewCustomEmoji("panel", 1334824826128564244, false)
	EmojiRating     = NewCustomEmoji("rating", 1334824756687536200, false)
	EmojiStaff      = NewCustomEmoji("staff", 1334824738165620799, false)
	EmojiThread     = NewCustomEmoji("thread", 1334824714497032223, false)
	EmojiBulletLine = NewCustomEmoji("bulletline", 1334825028532961281, false)
	EmojiPatreon    = NewCustomEmoji("patreon", 1334825012720570388, false)
	EmojiDiscord    = NewCustomEmoji("discord", 1334824993414185041, false)
	//EmojiTime       = NewCustomEmoji("time", 974006684622159952, false)
)

// PrefixWithEmoji Useful for whitelabel bots
func PrefixWithEmoji(s string, emoji CustomEmoji, includeEmoji bool) string {
	if includeEmoji {
		return fmt.Sprintf("%s %s", emoji, s)
	} else {
		return s
	}
}
