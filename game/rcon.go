package game

import (
	"log"
	"net"
	"regexp"
	"time"
)

func SendRCON(host string, pass string, rconCommand string) {
	conn, err := net.Dial("udp", host)
	if err != nil {

		return
	}
	defer conn.Close()

	send := prepareCommand("challenge rcon\n")
	_, err = conn.Write(send)
	if err != nil {
		log.Fatalf("Error")
		return
	}

	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Fatalf("Error")
		return
	}

	getChallenge := make([]byte, 1024)
	n, err := conn.Read(getChallenge)
	if err != nil {
		log.Fatalf("Error trying to receive challenge from %v ==> %v",
			conn.RemoteAddr().String(), err.Error())
		return
	}
	getChallenge = getChallenge[:n]

	challengeSplit := regexp.MustCompile(`[^\d]`).Split(string(getChallenge), -1)
	var challenge string
	for i := 0; i < len(challengeSplit); i++ {
		challenge += challengeSplit[i]
	}

	send = prepareCommand("rcon \"" + challenge + "\" " + pass + " " + rconCommand + "\n")

	_, err = conn.Write(send)
	if err != nil {
		log.Fatalf("Error trying to send the command to %v ==> %v",
			conn.RemoteAddr().String(), err.Error())
		return
	}
}

func prepareCommand(command string) []byte {
	var sequence []byte
	sequence = append(sequence, []byte{255, 255, 255, 255}...)
	sequence = append(sequence, []byte(command)...)
	return sequence
}
