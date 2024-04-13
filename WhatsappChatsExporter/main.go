package main

import (
	"bufio"
	. "fmt"
	"os"
	"regexp"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

func ReadText(testo string) string {
	Print("Insert text (ends with 'CTRL+D' or '//'): ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "//" {
			break
		} else {
			testo += scanner.Text()
		}
	}
	return testo
}

func DuplicatesRemover(slice []string) []string {
	encountered := map[string]bool{}
	var result []string
	for _, v := range slice {
		if encountered[v] == true {
			continue
		}
		encountered[v] = true
		result = append(result, v)
	}
	return result
}

func DateFinder(testo string) []string {
	var newT = []string{"0", testo}
	var newT2 []string
	format := "([0-9]{2})/([0-9]{2})/([0-9]{2}), ([0-9]{2}):([0-9]{2}) - "
	campione, _ := regexp.Compile(format)
	dates := campione.FindAllString(testo, -1)
	dates = DuplicatesRemover(dates)
	for _, k := range dates {
		newT = strings.Split(newT[1], k)
		newT2 = append(newT2, newT[:len(newT)-1]...)
		newT = []string{" ", newT[len(newT)-1]}
	}
	newT2 = append(newT2, newT[1])
	format2 := ": "
	campione2, _ := regexp.Compile(format2)
	var newT3 []string
	var newT4 []string
	for _, s := range newT2 {
		newT3 = append(newT3, campione2.FindAllString(s, 1)...)
		if len(newT3) > 0 {
			newT4 = append(newT4, s)
			newT3 = nil
		}
	}
	return newT4
}

func UserColorer(names []string) map[string][]string {
	//Color map names --> Colors: 1) [R] red, 2) [G] green, 3) [B] blue, 4) [M] magenta, 5) [C] cyan, 6) [Y] yellow, 7) [W] white
	nameMap := make(map[string][]string)
	count := 1
	for i := 0; i < len(names); i++ {
		switch count {
		case 1:
			nameMap["R"] = append(nameMap["R"], names[i])
		case 2:
			nameMap["G"] = append(nameMap["G"], names[i])
		case 3:
			nameMap["B"] = append(nameMap["B"], names[i])
		case 4:
			nameMap["M"] = append(nameMap["M"], names[i])
		case 5:
			nameMap["C"] = append(nameMap["C"], names[i])
		case 6:
			nameMap["Y"] = append(nameMap["Y"], names[i])
		case 7:
			nameMap["W"] = append(nameMap["W"], names[i])
		}
		if count >= 7 {
			count = 0
		}
		count++
	}
	return nameMap
}

func NameFinder(cNoDate []string) ([][]string, []string) {
	var chat [][]string
	var names []string
	for i := 0; i < len(cNoDate); i++ {
		x := strings.Split(cNoDate[i], ": ")
		x[0] = x[0] + ": "
		chat = append(chat, x)
		names = append(names, chat[i][0])
	}
	return chat, names
}

func ColoredPrintName(chat [][]string, nameMap map[string][]string, i int, j int) {
	//Color map names --> Colors: 1) [R] red, 2) [G] green, 3) [B] blue, 4) [M] magenta, 5) [C] cyan, 6) [Y] yellow, 8) [W] white
	c1 := color.New(color.FgRed)
	c2 := color.New(color.FgGreen)
	c3 := color.New(color.FgBlue)
	c4 := color.New(color.FgMagenta)
	c5 := color.New(color.FgCyan)
	c6 := color.New(color.FgYellow)
	c7 := color.New(color.FgWhite)
	s := chat[i][j]
	for k, v := range nameMap {
		for _, n := range v {
			if s == n {
				switch k {
				case "R":
					c1.Print(s)
				case "G":
					c2.Print(s)
				case "B":
					c3.Print(s)
				case "M":
					c4.Print(s)
				case "C":
					c5.Print(s)
				case "Y":
					c6.Print(s)
				case "W":
					c7.Print(s)
				}
			}
		}
	}
}

func PrintChat(chat [][]string, nameMap map[string][]string) {
	for i := 0; i < len(chat); i++ {
		for j := 0; j < 2; j++ {
			if j == 0 {
				ColoredPrintName(chat, nameMap, i, j)
			}
			if j > 0 {
				Print(chat[i][j])
			}
		}
		Println()
	}
}

func readFromFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func writeToFile(matrix [][]string) error {
	file, err := os.Create("Chat.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < 2; j++ {
			Fprint(w, matrix[i][j])
		}
		Fprintln(w)
	}
	return w.Flush()
}

func showMenu(options []string) int {
	c1 := color.New(color.FgRed)
	selectedOption := 0
	// Inizializza la tastiera
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	println(" Choose you input option: ")
	Println()
	// Loop per leggere l'input dell'utente
	for {
		for i, option := range options {
			if i == selectedOption {
				c1.Printf(" >> %d. %s << \n", i+1, option)
			} else {
				Printf("    %d. %s\n", i+1, option)
			}
		}
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		// Esegui il codice in base all'input dell'utente
		if key == keyboard.KeyArrowUp || char == 'w' {
			selectedOption--
			if selectedOption < 0 {
				selectedOption = len(options) - 1
			}
		} else if key == keyboard.KeyArrowDown || char == 's' {
			selectedOption++
			if selectedOption >= len(options) {
				selectedOption = 0
			}
		} else if key == keyboard.KeyEnter || key == keyboard.KeySpace {
			Println("You chose:", options[selectedOption])
			return selectedOption
		}
		println()
		os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
		println(" Choose you input option: ")
		Println()
	}
}

func main() {
	os.Stdout.WriteString("\033[2J\033[1;1H") // pulisco la board
	testo := ""
	/*Println("Choose your input method:")
	Println("(Use \"w\" & \"s\" to navigate menu, and choose the option with Enter)")*/
	options := []string{"Read from file", "Paste text", "Quit"}
	selectedOption := showMenu(options)
	switch selectedOption {
	case 0:
		var fileName string
		Print("Enter file name: ")
		Scan(&fileName)
		content, err := readFromFile(fileName)
		testo = content
		if err != nil {
			Println(err)
			return
		}
	case 1:
		testo = ReadText(testo)
	case 2:
		return
	}
	Printf("\n\n")
	cNoDate := DateFinder(testo)
	chat, names := NameFinder(cNoDate)
	names = DuplicatesRemover(names)
	nameMap := UserColorer(names)
	Println("Start of the Chat: ")
	PrintChat(chat, nameMap)
	err2 := writeToFile(chat)
	if err2 != nil {
		Println(err2)
	}
	Print("Press Enter to end the Program: ")
	Scanln()
}
