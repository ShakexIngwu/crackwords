package danmu_rest

import (
	"fmt"
	"bytes"
	"log"
	"os/exec"
)

//MlController ...Controller of Python machine learning scripts.
type MlController struct{}

//getMostSimilarWords get most similar words based on the prediction from word2vec
func (m *MlController) getMostSimilarWords(word string) (okResult bool, stdout, stderr string) {
	log.Println("Running topSynonym.py to get similar words...")
	cmd := exec.Command(
		"/home/ubuntu/danmu/./topSynonym.py",
		"-w", word,
	)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	okResult = true
	err := cmd.Run()
	stdout = outb.String()
	stderr = errb.String()
	if err != nil {
		//log.Fatalln("failed to execute python file topSynonym.py.", err)
		fmt.Println(fmt.Sprint(err) + ": " + stderr)
		log.Fatal(err)
		okResult = false
	}

	log.Println("topSynonym.py executed successfully!")
	return
}

//getRecentTopic
func (m *MlController) getRecentTopic(interval int) (okResult bool, stdout, stderr string) {
	log.Println("Running predictRecentWords.py to get similar words...")
	//in the future will add interval when executing predictRecentWords.py
	cmd := exec.Command(
		"/home/ubuntu/danmu/./predictRecentWords.py",
	)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	okResult = true
	err := cmd.Run()
	stdout = outb.String()
	stderr = errb.String()
	log.Println(stdout)
	if err != nil {
		//log.Fatalln("failed to execute python file topSynonym.py.", err)
		fmt.Println(fmt.Sprint(err) + ": " + stderr)
		log.Fatal(err)
		okResult = false
	}

	log.Println("predictRecentWords.py executed successfully!")
	return
}

