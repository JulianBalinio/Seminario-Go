package agenda

import (
	"sync"
	"Ej3/contact"
	"errors"
)

var notFound = errors.New("No se encontro el contacto")

type Schedule struct {
	mu       sync.Mutex
	Contacts map[string]*contact.Contact
}

func NewAgenda() *Schedule {
	return &Schedule{
		Contacts: make(map[string]*contact.Contact),
	}
}

func (a *Schedule) AddContact(contact *contact.Contact) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Contacts[contact.Email()] = contact
}

func (a *Schedule) DeleteContact(email string) error {
    a.mu.Lock()
    defer a.mu.Unlock()
    if _, ok := a.Contacts[email]; ok {
        delete(a.Contacts, email)
		return nil
    } else {
        return notFound
    }
}

// BuscarContacto busca y devuelve un contacto dado su correo electrónico
func (a *Schedule) SearchContact(email string) (*contact.Contact, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	contact, ok := a.Contacts[email]
	if !ok {
		return nil, notFound
	} else {
		return contact, nil
	}
}

// ListarContactos devuelve todos los contactos en la agenda
func (a *Schedule) ListContacts() []*contact.Contact {
    a.mu.Lock()
    defer a.mu.Unlock()

    contacts := make([]*contact.Contact, 0, len(a.Contacts)) // Pre-allocate for efficiency
    for _, c := range a.Contacts {
        contacts = append(contacts, c)
    }
    return contacts
}
