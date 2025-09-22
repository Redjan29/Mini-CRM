package main

import "testing"

func reset() {
	contacts = make(map[int]Contact)
}

func TestAddContact_DataLayer(t *testing.T) {
    
    contacts = make(map[int]Contact)

    
    contacts[1] = Contact{ID: 1, Nom: "Alice", Email: "alice@mail.com"}

    if c, ok := contacts[1]; !ok {
        t.Fatal("contact non ajouté")
    } else if c.Nom != "Alice" {
        t.Errorf("Nom attendu 'Alice', obtenu '%s'", c.Nom)
    }
}


func TestDeleteContact_DataLayer(t *testing.T) {
	reset()

	contacts[2] = Contact{ID: 2, Nom: "Bob", Email: "bob@mail.com"}

	
	delete(contacts, 2)

	if _, ok := contacts[2]; ok {
		t.Fatal("contact non supprimé")
	}
}
