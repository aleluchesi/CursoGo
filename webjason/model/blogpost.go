package model

type Blogpost struct {
	UserID		int		`json:"userid"`
	ID			int		`json:"id"`
	Titulo		string	`json:"titulo"`
	Texto		string	`json:"body"`
}