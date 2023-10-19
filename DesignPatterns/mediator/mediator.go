package main

type Mediator interface {
	AddUser(user *User)
	SendMessage(message string, from *User)
}

type User struct {
	Name    string
	Mediate *ChatRoom
}

type ChatRoom struct {
	Users []*User
}

func (c *ChatRoom) AddUser(user *User) {
	c.Users = append(c.Users, user)
}

func (c *ChatRoom) SendMessage(message string, from *User) {
	for _, user := range c.Users {
		if user != from {
			println(user.Name + " received message: " + message)
		}
	}
}

func NewUser(name string, mediator *ChatRoom) *User {
	user := &User{Name: name, Mediate: mediator}
	mediator.AddUser(user)
	return user
}

func (user *User) Send(message string) {
	user.Mediate.SendMessage(message, user)
}

func main() {
	chatRoom := &ChatRoom{}
	user1 := NewUser("DevUser", chatRoom)
	user2 := NewUser("Bot", chatRoom)
	user1.Send("Hi there!")
	user2.Send("Hey! Thanks for contacting DevRev. How can I help you?")

	user1.Send("I need help with the Go programming language.")
	user2.Send("Sure, I can help you with that. What do you need help with?")

	user1.Send("Could you tech me about the Mediator pattern?")
	user2.Send("Mediator pattern is used to reduce communication complexity between multiple objects or classes.")

	user1.Send("Thanks!")
	user2.Send("You're welcome!")

	user1.Send("Bye!")
	user2.Send("Bye!")
}
