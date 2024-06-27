package contact

type Contact struct{
	fname string
	lname string
	email string
	phoneNum int
}

func NewContact(fn, ln, em string, pn int) *Contact {
	return &Contact{
		fname: fn,
		lname: ln,
		email: em,
		phoneNum: pn,
	}
}

func (c *Contact) FirstName() string {
	return c.fname 
}

func (c *Contact) LastName() string {
	return c.lname 
}

func (c *Contact) Email() string {
	return c.email 
}

func (c *Contact) PhoneNumber() int {
	return c.phoneNum
}