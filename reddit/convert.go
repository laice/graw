package reddit

func (p *Post) ToComment() *Comment {
	return &Comment{
		ID:        p.ID,
		Name:      p.Name,
		Permalink: p.Permalink,

		CreatedUTC: p.CreatedUTC,
		Deleted:    p.Deleted,

		Ups:   p.Ups,
		Downs: p.Downs,
		Likes: p.Likes,

		Author:              p.Author,
		AuthorFlairText:     p.AuthorFlairText,
		AuthorFlairCSSClass: p.AuthorFlairCSSClass,

		LinkAuthor: p.Author,
		LinkURL:    p.URL,
		LinkTitle:  p.Title,

		Subreddit:   p.Subreddit,
		SubredditID: p.SubredditID,

		Body:     p.SelfText,
		BodyHTML: p.SelfTextHTML,

		Replies: p.Replies,

		Gilded:        p.Gilded,
		Distinguished: p.Distinguished,
	}
}

func (c *Comment) ToPost() *Post {
	return &Post{
		ID:        c.ID,
		Name:      c.Name,
		Permalink: c.Permalink,

		CreatedUTC: c.CreatedUTC,
		Deleted:    c.Deleted,

		Ups:   c.Ups,
		Downs: c.Downs,
		Likes: c.Likes,

		Author:              c.Author,
		AuthorFlairText:     c.AuthorFlairText,
		AuthorFlairCSSClass: c.AuthorFlairCSSClass,

		URL:   c.LinkURL,
		Title: c.LinkTitle,

		Subreddit:   c.Subreddit,
		SubredditID: c.SubredditID,

		SelfText:     c.Body,
		SelfTextHTML: c.BodyHTML,

		Replies: c.Replies,

		Gilded:        c.Gilded,
		Distinguished: c.Distinguished,
	}
}
