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

// Character interface defines the methods that any character class should implement
type Character interface {
	GetName() string
	Attack(diceSize int64) RollResult
	TestAttribute(attribute string, diceSize int64, difficulty int64) (RollResult, bool)
}

// Warrior struct represents the Warrior class and implements the Character interface
type Warrior struct {
	RPGClass
}

func (w Warrior) GetName() string {
	return w.Name
}

func (w Warrior) Attack(diceSize int64) RollResult {
	return RollDice(w.Strength, diceSize)
}

// Mage struct represents the Mage class and implements the Character interface
type Mage struct {
	RPGClass
}

func (m Mage) GetName() string {
	return m.Name
}

func (m Mage) Attack(diceSize int64) RollResult {
	return RollDice(m.Intelligence, diceSize)
}

// Archer struct represents the Archer class and implements the Character interface
type Archer struct {
	RPGClass
}

func (a Archer) GetName() string {
	return a.Name
}

func (a Archer) Attack(diceSize int64) RollResult {
	return RollDice(a.Dexterity, diceSize)
}

// BardBarbarian struct represents the Bard/Barbarian class and implements the Character interface
type BardBarbarian struct {
	RPGClass
}

func (bb BardBarbarian) GetName() string {
	return bb.Name
}

func (bb BardBarbarian) Attack(diceSize int64) RollResult {
	average := (bb.Strength + bb.Dexterity + bb.Intelligence) / 3
	return RollDice(average, diceSize)
}

// Factory functions to create instances of each class
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

	//var quantidadeDado int64
	//var valorDado string
	var player Character

	//LER O ARQUIVO JSON
	dices, err := LoadFile("dices.json")
	if err != nil {
		fmt.Printf("Erro ao carregar o arquivo: %v\n", err)
		return
	}

	// //ESCOLHA SUA CLASSE
	fmt.Println("Escolha sua classe: warrior, mage, archer, bardBarian")
	var choice string
	fmt.Scan(&choice)
	choice = strings.ToLower(choice)

	switch choice {
	case "warrior":
		player = Warrior{NewWarrior()}
	case "mage":
		player = Mage{NewMage()}
	case "archer":
		player = Archer{NewArcher()}
	case "bardbarian":
		player = BardBarbarian{NewBardBarian()}
	default:
		fmt.Println("Classe invalida!")
		return
	}

	GameMenu(player, dices)

	// fmt.Println("Usando Interfaces e Structs para criar classes de RPG")
	// player = Warrior{NewWarrior()}
	// resultInterface := player.Attack(6)
	// fmt.Println("Classe: ", player.GetName())
	// fmt.Println("Total Dados: ", resultInterface.Total)
	// fmt.Println("Fim do uso de interfaces e structs")

	// fmt.Println("===============================================")
	// fmt.Println("Bem Vindo ao RPG Dice Roller!")
	// // fmt.Println("Quantos dados voce deseja jogar?")
	// // fmt.Scan(&quantidadeDado)
	// //VALIDAR QUANTIDADE DE DADOS
	// // if quantidadeDado <= 0 {
	// // 	fmt.Println("Quantidade invalida de dados!")
	// // 	return
	// // }

	//var playerClass RPGClass

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

func (c RPGClass) TestAttribute(attibute string, diceSize int64, difficult int64) (RollResult, bool) {
	var numDice int64

	switch strings.ToLower(attibute) {
	case "strength":
		numDice = c.Strength
	case "dexterity":
		numDice = c.Dexterity
	case "intelligence":
		numDice = c.Intelligence
	default:
		numDice = 1
	}

	result := RollDice(numDice, diceSize)
	success := result.Total >= difficult

	return result, success
}

func HandleAttack(player Character, dices map[string]int64) {
	fmt.Println("Escolha o dado para ataque: D4, D6, D8, D10, D12")
	var dice string
	fmt.Scan(&dice)

	diceValue, _ := GetDice(dice, dices)
	result := player.Attack(diceValue)
	fmt.Println("Ataque Total: ", result.Total)
}

func HandleAttributeTest(player Character, dices map[string]int64) {
	fmt.Println("Escolhe o atributo que sera testado: Strength, Dexterity, Intelligence")
	var atrib string
	fmt.Scan(&atrib)

	fmt.Println("Qual a dificuldade do desafio? 1->20")
	var diff int64
	fmt.Scan(&diff)
	if diff < 1 {
		diff = 1
	}
	if diff > 20 {
		diff = 20
	}

	resultTest, success := player.TestAttribute(atrib, 20, diff)

	fmt.Println("Resultado dos dados: ")
	for i, roll := range resultTest.Rolls {
		fmt.Printf("Dado %d: %d\n", i+1, roll)
	}

	fmt.Println("Total: ", resultTest.Total)
	if success {
		fmt.Println("Sucesso no teste")
	} else {
		fmt.Println("Falhou no teste")
	}
}

func HandleDiceRoll(dices map[string]int64) {
	fmt.Println("Quanto dados Deseja rolar?")
	var diceAmount int64
	fmt.Scan(&diceAmount)

	if diceAmount <= 0 {
		fmt.Println("Quantidade invalida")
		return
	}

	fmt.Println("Qual dado voce deseja jogar? (ex: d4, d6, d8, d10, d12 d20, etc.)")
	var dice string
	fmt.Scan(&dice)

	diceValue, err := GetDice(dice, dices)
	if err != nil {
		fmt.Printf("Erro ao obter o valor do dado: %v\n", err)
		return
	}
	//ROLAR DADO
	result := RollDice(diceAmount, diceValue)

	fmt.Printf("Resultados dos dados:\n")

	for i, roll := range result.Rolls {
		fmt.Printf(" Dado %d: %d\n", i+1, roll)
	}
	fmt.Printf(" Total: %d\n", result.Total)

}

func GameMenu(player Character, dices map[string]int64) {
	for {
		fmt.Println("======= Bem vindo ao RPG Dice Roler =======")
		fmt.Println("1 - Rolar Dado")
		fmt.Println("2 - Atacar")
		fmt.Println("3 - Teste de atributo")
		fmt.Println("4 - Classe atual")
		fmt.Println("0 - Sair")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			HandleDiceRoll(dices)
		case 2:
			HandleAttack(player, dices)
		case 3:
			HandleAttributeTest(player, dices)
		case 4:
			fmt.Println("Classe atual: ", player.GetName())
		case 0:
			fmt.Println("Ate a proxima aventura...")
			return
		default:
			fmt.Println("Opcao Invalida")
		}
	}
}
