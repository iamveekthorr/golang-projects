// Package main provides TODO: add description
package main

import (
	"errors"
	"fmt"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameFromID(userID string) (string, bool) {

	name, ok := sds.userData[userID]
	return name, ok
}

// NewSimpleDataStore - Abstract Factory - Creates an instance of SimpleDataStore
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

type DataStore interface {
	UserNameFromID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

// Log - The lg() function creates a closure, so any function can be passed into
// Log(). Therefore making it a clever way to decouple dependencies!
func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameFromID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameFromID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}

	return "Goodbye, " + name, nil
}
