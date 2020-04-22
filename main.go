package main

import (
	"os"
	"log"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	passwordPromptString = "Password"
	hashPromptString     = "Hash"
)

const (
	iterations = 14
)

func usage(scriptName string) {
	fmt.Printf("usage: %s <command>\n", scriptName)
	fmt.Println("\tavailable commands:\n")
	fmt.Println("\t\thelp")
	fmt.Println("\t\tcreate")
	fmt.Println("\t\tverify")
}

func getUserInputAsBytes(message string) ([]byte, error) {
	fmt.Println(message)

	var i string

	_, err := fmt.Scan(&i)
	if err != nil {
		return nil, err
	}

	return []byte(i), nil
}

func promptUserForPassword(message string) ([]byte, error) {
	return getUserInputAsBytes(message)
}

func promptUserForHash(message string) ([]byte, error) {
	return getUserInputAsBytes(message)
}

func hashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, iterations)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func verifyMatch(hash, pwd []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, pwd) == nil
}

func main() {
	argv := os.Args[1:]
	if len(argv) == 0 {
		log.Fatal("error: at least one argument expected")
	}

	switch argv[0] {
	case "help":
		usage(os.Args[0])

		os.Exit(1)
	case "create":
		password, err := promptUserForPassword(passwordPromptString)
		if err != nil {
			log.Fatal(err)
		}

		hash, err := hashAndSalt(password)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(hash)

		os.Exit(0)
	case "verify":
		/*params := argv[1:]

		if len(params) < 2 {
			log.Fatal("error: missing arguments, password and hashed password")
		}

		password := params[0]
		hash := params[1]*/

		password, err := promptUserForPassword(passwordPromptString)
		if err != nil {
			log.Fatal(err)
		}

		hash, err := promptUserForHash(hashPromptString)
		if err != nil {
			log.Fatal(err)
		}

		if verifyMatch(hash, password) {
			fmt.Println("ok")
		} else {
			fmt.Printf("provided password \"%s\" does not match hash \"%s\"\n", password, hash)
		}

		os.Exit(0)
	}

	os.Exit(0)
}
