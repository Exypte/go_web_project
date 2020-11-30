package main

import (
	"log"
	"net"
	"bufio"
	"encoding/json"
	"math/big"
	"net/http"
	"strings"

	ec "github.com/Exypte/go_encryption"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal(err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	pubKey := ec.CouplePublic(11)
	priKey := ec.CouplePrivate(pubKey)

	reader := bufio.NewReader(conn)

	/*
	* Wait for the client public key
	*/
	pubKeyBytes, err := reader.ReadBytes(byte('}'))

	if err != nil {
		log.Fatal(err.Error())
	}

	var clientPubKey *ec.PublicKey

	err = json.Unmarshal(pubKeyBytes, &clientPubKey)

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
	* Wait for an encrypted message
	*/
	log.Print("Wait for message ...")
	for {
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

		msg := ec.Decryption(msgEncry, priKey)

		/*
		* Call the api to know is the employe is known of the database
		*/
		url := "http://microservice_user:8080/company/name/" + msg
		res, err := http.Get(strings.TrimSuffix(url, "\r\n"))

		if err != nil{
			log.Println(err.Error())
			break
		}

		/*
		* Answer to the client if the employe is known of the database
		*/
		validation := ""

		if res.StatusCode == 200{
			validation = "User is known of the database"
		}else{
			validation = "User is unknown of the database"
		}

		validationEncry := ec.Encryption(validation, clientPubKey)

		response, err := json.Marshal(validationEncry)

		if err != nil {
			log.Fatal(err.Error())
		}
	
		conn.Write(response)
	}
}