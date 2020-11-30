package main

import (
	"log"
	"encoding/json"
	"net"
	"fmt"
	"bufio"
	"os"
	"math/big"

	ec "github.com/Exypte/go_encryption"
)

func main() {
	pubKey := ec.CouplePublic(11)
	priKey := ec.CouplePrivate(pubKey)

	conn, err := net.Dial("tcp", "localhost:8280")

	if err != nil {
		log.Fatal(err.Error())
	}

	/*
	* Send its public key
	*/
	response, err := json.Marshal(pubKey)

	if err != nil {
		log.Fatal(err.Error())
	}

	conn.Write(response)

	/*
	* Wait for the server public key
	*/
	reader := bufio.NewReader(conn)

	pubKeyBytes, err := reader.ReadBytes(byte('}'))

	if err != nil {
		log.Fatal(err.Error())
	}

	var serverPubKey *ec.PublicKey

	err = json.Unmarshal(pubKeyBytes, &serverPubKey)

	if err != nil {
		log.Fatal(err.Error())
	}

	/*
	* Wait for an input
	*/
	buf := bufio.NewReader(os.Stdin)

	for {
		/*
		* Wait for user input
		*/
		fmt.Print("User name : ")

		sentence, err := buf.ReadBytes('\n')

		if err != nil {
			log.Fatal(err.Error())
		} else {
			msg := string(sentence)

			msgEncry := ec.Encryption(msg, serverPubKey)

			response, err := json.Marshal(msgEncry)

			if err != nil {
				log.Fatal(err.Error())
			}
		
			conn.Write(response)
		}

		/*
		* Listen for the name of one user
		*/
		msgBytes, err := reader.ReadBytes(byte(']'))

		if err != nil {
			break
		}

		var msgEncry []*big.Int
	
		err = json.Unmarshal(msgBytes, &msgEncry)
	
		if err != nil {
			break
		}

		validation := ec.Decryption(msgEncry, priKey)

		log.Println(validation)
	}
}