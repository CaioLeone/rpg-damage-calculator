package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type RPGClass struct {
	Name         string
	Strength     int64
	Dexterity    int64
	Intelligence int64
}

func NewWarrior() RPGClass {
	return RPGClass{
		Name:         "Warrior",
		Strength:     3,
		Dexterity:    2,
		Intelligence: 1,
	}
}

func NewMage() RPGClass {
	return RPGClass{
		Name:         "Mage",
		Strength:     1,
		Dexterity:    2,
		Intelligence: 3,
	}
}

func NewArcher() RPGClass {
	return RPGClass{
		Name:         "Archer",
		Strength:     2,
		Dexterity:    3,
		Intelligence: 1,
	}
}

func NewBardBarian() RPGClass {
	return RPGClass{
		Name:         "Bard/Barbarian",
		Strength:     2,
		Dexterity:    2,
		Intelligence: 2,
	}
}

type RollResult struct {
	Rolls []int64
	Total int64
}

func main() {

	var quantidadeDado int64
	var valorDado string

	//LER O ARQUIVO JSON
	dices, err := LoadFile("dices.json")
	if err != nil {
		fmt.Printf("Erro ao carregar o arquivo: %v\n", err)
		return
	}

	// fmt.Println("Quantos dados voce deseja jogar?")
	// fmt.Scan(&quantidadeDado)
	//VALIDAR QUANTIDADE DE DADOS
	// if quantidadeDado <= 0 {
	// 	fmt.Println("Quantidade invalida de dados!")
	// 	return
	// }

	//ESCOLHA SUA CLASSE
	fmt.Println("Escolha sua classe: warrior, mage, archer, bardBarian")
	var choice string
	fmt.Scan(&choice)
	choice = strings.ToLower(choice)

	var playerClass RPGClass
	switch choice {
	case "warrior":
		playerClass = NewWarrior()
	case "mage":
		playerClass = NewMage()
	case "archer":
		playerClass = NewArcher()
	case "bardbarian":
		playerClass = NewBardBarian()
	default:
		fmt.Println("Classe invalida!")
		return
	}

	quantidadeDado = playerClass.AttackDice()

	fmt.Println("Qual dado voce deseja jogar? (ex: d4, d6, d8, d10, d12 d20, etc.)")
	valorDado = strings.ToLower(valorDado)
	fmt.Scan(&valorDado)

	//OBTER VALOR DADO
	diceValue, err := GetDice(valorDado, dices)
	if err != nil {
		fmt.Printf("Erro ao obter o valor do dado: %v\n", err)
		return
	}
	//ROLAR DADO
	result := RollDice(quantidadeDado, diceValue)

	fmt.Printf("Resultados dos dados:\n")

	for i, roll := range result.Rolls {
		fmt.Printf(" Dado %d: %d\n", i+1, roll)
	}
	fmt.Printf(" Total: %d\n", result.Total)
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

func RollDice(numDice, diceSize int64) RollResult {
	var total int64
	var rolls []int64

	for i := int64(0); i < numDice; i++ {
		roll := rand.Int64N(diceSize) + 1

		if roll == diceSize {
			//Critico Duplicar
			fmt.Printf("Critico! Dado Rolou %d em um D%d\n", roll, diceSize)
			roll *= 2
		}
		rolls = append(rolls, roll)
		total += roll
	}

	return RollResult{
		Rolls: rolls,
		Total: total,
	}
}

func (c RPGClass) AttackDice() int64 {
	switch c.Name {
	case "Warrior":
		return c.Strength
	case "Mage":
		return c.Intelligence
	case "Archer":
		return c.Dexterity
	case "Bard/Barbarian":
		return (c.Strength + c.Dexterity + c.Intelligence) / 3
	default:
		return 1 // Default attack dice value
	}
}
