# RPG Damage Calculator (Go CLI)

Um projeto em Go para simular cálculo de dano em um sistema de RPG utilizando linha de comando.

O objetivo do projeto é treinar:

- Structs
- Métodos com receiver
- Lógica de negócio
- Leitura de JSON
- Manipulação de argumentos (`os.Args`)
- Separação de responsabilidades

---

## 🎯 Funcionalidades

O programa calcula o dano total de um ataque baseado em:

- Força do personagem
- Tipo de arma (d4, d6, d8, d10)
- Se o ataque foi crítico

### Fórmula de Dano

Dano Base:
rolagem_do_dado + força

Se for crítico:
dano_total = dano_base * 2

---

## Sintaxe de Uso

Após compilar:

./rpg 5 d8 sim

Ou usando go run:

go run main.go 5 d8 sim


### Argumentos

| Argumento | Descrição |
|-----------|-----------|
| Classe | String representando a classe do personagem |
| Arma | Tipo de dado: d4, d6, d8 ou d10 |
| Crítico | `sim` ou `nao` |

---

## Estrutura do Projeto


rpg-damage-calculator/
├── main.go
├── go.mod
├── data/
│ └── Dices.json
└── README.md


---

## 📦 Fonte de Dados

As armas são carregadas a partir de um arquivo JSON:

`data/Dices.json`

Exemplo:

```
json
{
  "d4": 4,
  "d6": 6,
  "d8": 8,
  "d10": 10
}
```

Isso permite adicionar novos tipos de arma sem modificar o código.

## Conceitos Praticados

1. Modelagem de domínio com Struct
2. Métodos associados a Struct
3. Encapsulamento
4. Tratamento de erros idiomático em Go
5. Leitura de arquivos JSON
6. Conversão de tipos
7. Uso do pacote math/rand
8. Separação entre regra de negócio e infraestrutura

## Desafios Extras

1. Suporte a múltiplos dados (ex: 2d6)
2. Multiplicador de crítico customizável
3. Sistema de classes (Warrior, Mage, etc)
4. Interface Attacker
5. Seed configurável para testes determinísticos
6. Testes unitários

##Exemplo de Saída
Dano causado: 17