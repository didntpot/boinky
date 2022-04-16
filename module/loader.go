package module

import (
	"log"

	"github.com/df-mc/atomic"
	"github.com/sandertv/gophertunnel/minecraft"
)

func DefaultLoader(l *log.Logger) *Loader {
	return &Loader{
		bs: &BotSpammer{
			name:  "Bot",
			msg:   "This server has been invaded by the horny squad!",
			limit: 1000,
			spam: *atomic.NewBool(true),
			kill: *atomic.NewBool(false),
			conns: []*minecraft.Conn{},
			l: l,
		},
		cr: &ChatReader{
			kill: *atomic.NewBool(false),
			l: l,
		},
	}
}

type Loader struct {
	bs *BotSpammer
	cr *ChatReader
}

func (l *Loader) BotSpammer() *BotSpammer {
	return l.bs
}

func (l *Loader) ChatReader() *ChatReader {
	return l.cr
}