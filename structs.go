package main

// Activity Created, Read, Shredded etc
type Activity struct {
}

// Note The Core Object
type Note struct {
	Title      string `json:"title"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  string `json:"created_at"`
	Message    string `json:"message"`
	Token      string `json:"token"`
	Content    string `json:"content"`
	IsShredded bool   `json:"is_shredded"`
}
