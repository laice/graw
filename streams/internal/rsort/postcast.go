// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package rsort

import "github.com/laice/graw/reddit"

type postsThingImpl struct {
	e *reddit.Post
}

func (g postsThingImpl) Name() string { return g.e.Name }

func (g postsThingImpl) Birth() uint64 { return g.e.CreatedUTC }

func postsAsThings(gs []*reddit.Post) []redditThing {
	things := make([]redditThing, len(gs))
	for i, g := range gs {
		things[i] = &postsThingImpl{g}
	}
	return things
}
