package main 

type Website int

const (
	Google Website = iota
	Facebook
	Amazon
	Stackoverflow
	Golang
)

func (w Website) String() string {
    return []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com", 
		"http://stackoverflow.com",
		"http://golang.org"}[w]
}