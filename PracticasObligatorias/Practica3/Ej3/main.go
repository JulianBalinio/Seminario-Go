package main

import (
    "Ej3/agenda"
    "Ej3/contact"
    "fmt"
    "sync"
)

func main() {
    agenda := agenda.NewAgenda()

    contacts := []*contact.Contact{
        contact.NewContact("Julian", "Ba", "julianba@email.com", 123456789),
        contact.NewContact("Nailuj", "Ab", "nailujab@email.com", 987654321),
        contact.NewContact("Lijuan", "Abba", "lijuan@email.com", 112233445),
    }

    var (
        wg  sync.WaitGroup
        cwg sync.WaitGroup
    )

    cwg.Add(len(contacts))
    for _, c := range contacts {
        go func(c *contact.Contact) {
            defer cwg.Done()
            agenda.AddContact(c)
        }(c)
    }
    cwg.Wait()

    fmt.Println("\nLista de contactos:")
    for _, c := range agenda.ListContacts() {
        fmt.Println(*c)
    }

    deleteContact := func(emailToDelete string) {
        err := agenda.DeleteContact(emailToDelete)
        if err != nil {
            fmt.Printf("Error al eliminar contacto %s: %s\n", emailToDelete, err)
        } else {
            fmt.Println("Contacto eliminado correctamente.")
        }
    }

    searchContact := func(emailToSearch string) {
        contact, err := agenda.SearchContact(emailToSearch)
        if err != nil {
			fmt.Printf("Error al buscar contacto %s: %s\n", emailToSearch, err)
            
        } else {
            fmt.Println("Contacto encontrado:", *contact)
        }
    }

    wg.Add(2)
    go func() {
		defer wg.Done()
        deleteContact("julianba@email.com")
    }()

    go func() {
		defer wg.Done()
        searchContact("nailujab@email.com")
    }()
    wg.Wait()

    fmt.Println("\nLista de contactos actualizada:")
    for _, c := range agenda.ListContacts() {
        fmt.Println(*c)
    }
}