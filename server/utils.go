package server

import (
	"crypto/rand"
	"fmt"
	"kode-task/models"
)

func pseudo_uuid() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	uuid = fmt.Sprintf("%X%X%X%X%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}
func reverse(input []models.Note) []models.Note {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}
