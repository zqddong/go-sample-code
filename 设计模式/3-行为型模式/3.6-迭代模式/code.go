package main

import "fmt"

type iterator interface {
	hasNext() bool
	getNext() *User
}

type collection interface {
	createIterator() iterator
}

type User struct {
	name string
	age  int
}

type userCollection struct {
	users []*User
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

type userIterator struct {
	index int
	users []*User
}

func (u *userIterator) hasNext() bool {
	return u.index < len(u.users)
}

func (u *userIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

func main() {
	userK := &User{
		name: "Kevin",
		age:  33,
	}
	userD := &User{
		name: "Diamond",
		age:  38,
	}

	userCollection := &userCollection{users: []*User{userK, userD}}

	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %v \n", user)
	}
}
