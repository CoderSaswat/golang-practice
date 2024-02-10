package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardSymbols := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen"}
	deck := deck{}
	for i := 0; i < len(cardSymbols); i++ {
		for j := 0; j < len(cardValues); j++ {
			deck = append(deck, cardValues[j]+" of "+cardSymbols[i])
		}
	}
	return deck
}

func (d deck) print() {
	for i, s := range d {
		fmt.Printf("%v : %v\n", i, s)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	fistDeck := d[:handSize]
	secondDeck := d[handSize:]
	return fistDeck, secondDeck
}

func (d deck) saveToFile(fileName string) error {
	deckStr := strings.Join(d, "\n")
	//deckByteArr := []byte(deckStr)
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.WriteString(deckStr)
	if err != nil {
		return err
	}
	//return ioutil.WriteFile(fileName, deckByteArr, 0644)
	return nil
}

func newDeckFromFile(fileName string) (deck, error) {
	contentByteArr, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	contentStr := string(contentByteArr)
	contentStrArr := strings.Split(contentStr, "\n")
	deck := deck(contentStrArr)
	return deck, nil
}

func (d deck) suffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
