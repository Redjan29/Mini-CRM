package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID    int
	Nom   string
	Email string
}

var contacts = make(map[int]Contact)
var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) (string, error) {
	fmt.Print(prompt)
	s, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(s), nil
}

func askInt(prompt string) (int, error) {
	for {
		txt, err := readLine(prompt)
		if err != nil {
			return 0, err
		}
		n, err := strconv.Atoi(txt)
		if err != nil {
			fmt.Println("Entrer un nombre valide.")
			continue
		}
		return n, nil
	}
}

func addContact() {
	id, _ := askInt("ID: ")
	if _, ok := contacts[id]; ok {
		fmt.Println("Cet ID existe déjà.")
		return
	}
	nom, _ := readLine("Nom: ")
	email, _ := readLine("Email: ")
	contacts[id] = Contact{ID: id, Nom: nom, Email: email}
	fmt.Println("Contact ajouté.")
}

func listContacts() {
	if len(contacts) == 0 {
		fmt.Println("(aucun contact)")
		return
	}
	for id, c := range contacts {
		fmt.Printf("[%d] %s <%s>\n", id, c.Nom, c.Email)
	}
}

func deleteContact() {
	id, _ := askInt("ID à supprimer: ")
	if _, ok := contacts[id]; !ok {
		fmt.Println("ID introuvable.")
		return
	}
	delete(contacts, id)
	fmt.Println("Contact supprimé.")
}

func updateContact() {
	id, _ := askInt("ID à mettre à jour: ")
	c, ok := contacts[id]
	if !ok {
		fmt.Println("ID introuvable.")
		return
	}
	nom, _ := readLine(fmt.Sprintf("Nouveau nom (%s): ", c.Nom))
	email, _ := readLine(fmt.Sprintf("Nouvel email (%s): ", c.Email))
	if nom == "" {
		nom = c.Nom
	}
	if email == "" {
		email = c.Email
	}
	contacts[id] = Contact{ID: id, Nom: nom, Email: email}
	fmt.Println("Contact mis à jour.")
}

func handleFlags() bool {
	id := flag.Int("id", -1, "ID du contact")
	nom := flag.String("nom", "", "Nom du contact")
	email := flag.String("email", "", "Email du contact")
	flag.Parse()

	if *id == -1 && *nom == "" && *email == "" {
		return false
	}
	if *id == -1 || *nom == "" || *email == "" {
		fmt.Println("Utilisation: go run . -id=1 -nom=Alice -email=alice@mail.com")
		os.Exit(1)
	}
	if _, exists := contacts[*id]; exists {
		fmt.Println("Cet ID existe déjà.")
		os.Exit(1)
	}
	contacts[*id] = Contact{ID: *id, Nom: *nom, Email: *email}
	fmt.Println("Contact ajouté via flags.")
	return true
}

func main() {
	if handled := handleFlags(); handled {
		return
	}

	for {
		fmt.Println("\n=== Mini CRM ===")
		fmt.Println("1) Ajouter un contact")
		fmt.Println("2) Lister les contacts")
		fmt.Println("3) Supprimer un contact")
		fmt.Println("4) Mettre à jour un contact")
		fmt.Println("5) Quitter")

		choice, err := askInt("Votre choix : ")
		if err != nil {
			continue
		}

		switch choice {
		case 1:
			addContact()
		case 2:
			listContacts()
		case 3:
			deleteContact()
		case 4:
			updateContact()
		case 5:
			fmt.Println("Au revoir.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
