package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// creating a deck type which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	//underscore is used in place of i and j as iterators
	//error for not using i and j will disappear
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	//has a receiver as deck
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//returning multiple values from the function
// (deck,deck) --> tells deck two different vals are to be returned of type deck (string here)

func deal(d deck, handSize int) (deck, deck) {
	//deck as an argument ; can be used as methods later on
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	//default permission - 0666 - lets everyone read and write the file ;
	//vscode automatically imports io/ioutil - subpackage
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// func func_name arguments return_type --> syntax
func newDeckFromFile(filename string) deck {
	// received bs - byteslice and error - err
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//option1 - log the error and return a call to newDeck()
		//option2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// func receiver_deck func_name no_return
func (d deck) shuffle() {
	//creating a random no generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	//for index in range of d
	for i := range d {
		// rand.Intn is a func - length of slice(actual size) - 1 - resulting in a random no.
		newPosition := r.Intn(len(d) - 1)
		//one line swapping [a,b=b,a] ; math/rand - import stm
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
