package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var quotes = []string{
	"I do not fear computers. I fear lack of them. – Isaac Asimov",
	"A computer once beat me at chess, but it was no match for me at kick boxing. – Emo Philips",
	"Computer Science is no more about computers than astronomy is about telescopes. – Edsger W. Dijkstra",
	"The computer was born to solve problems that did not exist before. – Bill Gates",
	"Software is like entropy: It is difficult to grasp, weighs nothing, and obeys the Second Law of Thermodynamics; i.e., it always increases.  – Norman Augustine",
	"Software is a gas; it expands to fill its container. – Nathan Myhrvold",
	"All parts should go together without forcing.  You must remember that the parts you are reassembling were disassembled by you.  Therefore, if you can’t get them together again, there must be a reason.  By all means, do not use a hammer.    – IBM Manual",
	"Standards are always out of date.  That’s what makes them standards. – Alan Bennett",
	"Physics is the universe’s operating system.  – Steven R Garman",
	"It’s hardware that makes a machine fast.  It’s software that makes a fast machine slow. – Craig Bruce",
	"Imagination is more important than knowledge.  For knowledge is limited, whereas imagination embraces the entire world, stimulating progress, giving birth to evolution. – Albert Einstein",
	"The greatest enemy of knowledge is not ignorance, it is the illusion of knowledge. – Stephen Hawking",
	"The more you know, the more you realize you know nothing. – Socrates",
	"Tell me and I forget.  Teach me and I remember.  Involve me and I learn. – Benjamin Franklin",
	"Real knowledge is to know the extent of one’s ignorance. – Confucius",
	"If people never did silly things, nothing intelligent would ever get done. – Ludwig Wittgenstein",
	"Getting information off the Internet is like taking a drink from a fire hydrant. – Mitchell Kapor",
	"If you think your users are idiots, only idiots will use it. – Linus Torvalds",
	"From a programmer’s point of view, the user is a peripheral that types when you issue a read request. – P. Williams",
	"Where is the ‘any’ key? – Homer Simpson, in response to the message, Press any key",
	"Computers are good at following instructions, but not at reading your mind. – Donald Knuth",
	"There is only one problem with common sense; it’s not very common. – Milt Bryce",
	"Your most unhappy customers are your greatest source of learning. – Bill Gates",
	"Let us change our traditional attitude to the construction of programs: Instead of imagining that our main task is to instruct a computer what to do, let us concentrate rather on explaining to human beings what we want a computer to do. – Donald E. Knuth",
	"The Internet?  We are not interested in it. – Bill Gates, 1993",
	"The best way to get accurate information on Usenet is to post something wrong and wait for corrections. – Matthew Austern",
}

const failWhale = `▄██████████████▄▐█▄▄▄▄█▌
██████▌▄▌▄▐▐▌███▌▀▀██▀▀
████▄█▌▄▌▄▐▐▌▀███▄▄█▌
▄▄▄▄▄██████████████▀
`

func getRandomQuote() string {
	l := len(quotes)
	i := rand.Intn(l)
	return quotes[i]
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	body := os.Getenv("PLUGIN_FAILED")

	s := failWhale
	if body == "" {
		s = getRandomQuote()
	}

	fmt.Printf("\n================\n%s\n================\n", s)
}
