package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {

}

func LoadFile(filename string) (map[string]int64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var dices map[string]int64
	err = json.Unmarshal(data, &dices)

	if err != nil {
		return nil, err
	}

	return dices, nil
}

func GetDice(dice string, dices map[string]int64) (int64, error) {
	diceValue, exists := dices[dice]

	if !exists {
		return 0, fmt.Errorf("Dado %s nao encontrado", dice)
	}

	return diceValue, nil
}

func RollDice(numDice, diceSize int64) int64 {
	var total int64

	for range numDice {
		roll := rand.Int64N(diceSize) + 1
		total += roll
	}

	return total
}
