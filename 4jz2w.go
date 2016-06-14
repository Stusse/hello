// SimplePush project SimplePush.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
)

func main() {

	// Set the access token.
	f, err := os.Open("token.txt")

	tokenFile := bufio.NewReader(f)
	token, _, err := tokenFile.ReadLine()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read from file error: %s\n", err)
		return
	}

	fmt.Println("api key: " + string(token))

	raid, tmpRaid := "", ""
	for raid == "" {
		raidFile, err := os.Open("raid.txt")
		checkErr(err)
		raidReader := bufio.NewReader(raidFile)
		raid, err = raidReader.ReadString('\n')
		checkErr(err)
		raidFile.Close()
		time.Sleep(time.Second)
	}

	for {
		raidFile, err := os.Open("raid.txt")
		if err != nil {
			fmt.Println("error: " + err.Error())
			tmpRaid = raid
		}
		raidReader := bufio.NewReader(raidFile)
		tmpRaid, err = raidReader.ReadString('\n')
		if err != nil {
			fmt.Println("error: " + err.Error())
			tmpRaid = raid
		}
		raidFile.Close()

		if tmpRaid != raid {
			raid = tmpRaid
			fmt.Print(raid)
			pb := pushbullet.New(string(token))

			// Create a push. The following codes create a note, which is one of push types.
			n := requests.NewNote()
			n.Title = raid
			n.Body = time.Now().Format("Mon Jan 2 15:04:05")
			// Send the note via Pushbullet.
			if _, err := pb.PostPushesNote(n); err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)

			}
		}
		time.Sleep(time.Second)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
}
