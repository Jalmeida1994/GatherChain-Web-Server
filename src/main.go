package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/ssh"
)

// ContentPost is the json format of the content for all post calls
type ContentPost struct {
	Author string
	Group  string
	Commit string
	IP     string
}

// create a data structure that can hold the response from the script
type scriptResponse struct {
	Response string
}

// Existing code from above
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// creates a new instance of a mux router
	myRouter.HandleFunc("/test", testFunc).Methods("POST")

	// admin commands
	myRouter.HandleFunc("/init", initNet).Methods("POST")
	myRouter.HandleFunc("/clear", clearNet).Methods("POST")

	// student commands
	myRouter.HandleFunc("/creategroup", createGrp).Methods("POST")
	myRouter.HandleFunc("/push", pushHash).Methods("POST")

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8010", myRouter))
}

func initNet(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	var cp ContentPost
	json.Unmarshal(reqBody, &cp)

	config := &ssh.ClientConfig{
		User: "adminUsername",
		Auth: []ssh.AuthMethod{
			ssh.Password("adminPassword2020")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", cp.IP, config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// TODO: check if user is admin
	runCommand("sudo /var/lib/waagent/custom-script/download/0/project/bloc-server/commands/init.sh "+cp.Author+" "+cp.Group+" "+cp.Commit, conn, w)
}

func clearNet(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	var cp ContentPost
	json.Unmarshal(reqBody, &cp)

	config := &ssh.ClientConfig{
		User: "adminUsername",
		Auth: []ssh.AuthMethod{
			ssh.Password("adminPassword2020")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", cp.IP, config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// TODO: check if user is admin
	runCommand("sudo /var/lib/waagent/custom-script/download/0/project/bloc-server/commands/clear.sh", conn, w)
}

func createGrp(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	var cp ContentPost
	json.Unmarshal(reqBody, &cp)

	config := &ssh.ClientConfig{
		User: "adminUsername",
		Auth: []ssh.AuthMethod{
			ssh.Password("adminPassword2020")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", cp.IP, config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	runCommand("sudo /var/lib/waagent/custom-script/download/0/project/bloc-server/commands/createchannel.sh "+cp.Author+" "+cp.Group+" "+cp.Commit, conn, w)
}

func pushHash(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	var cp ContentPost
	json.Unmarshal(reqBody, &cp)

	config := &ssh.ClientConfig{
		User: "adminUsername",
		Auth: []ssh.AuthMethod{
			ssh.Password("adminPassword2020")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", cp.IP, config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	runCommand("sudo /var/lib/waagent/custom-script/download/0/project/bloc-server/commands/push.sh "+cp.Author+" "+cp.Group+" "+cp.Commit, conn, w)
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	log.Println(reqBody)

	var cp ContentPost
	json.Unmarshal(reqBody, &cp)

	config := &ssh.ClientConfig{
		User: "adminUsername",
		Auth: []ssh.AuthMethod{
			ssh.Password("adminPassword2020")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", cp.IP, config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Println(cp)
	log.Println(cp.Author)
	log.Println(cp.Group)
	log.Println(cp.Commit)

	// Call Run method with command you want to run on remote server.
	runCommand("sudo /var/lib/waagent/custom-script/download/0/project/bloc-server/commands/test.sh "+cp.Author+" "+cp.Group+" "+cp.Commit, conn, w)
}

func runCommand(cmd string, conn *ssh.Client, w http.ResponseWriter) {
	sess, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	results, err := sess.Output(cmd)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		panic("Can't run remote command: " + err.Error())
	}
	// convert results into string and populate an instance of
	// the scriptResponse struct
	response := scriptResponse{string(results)}
	// encode response into JSON and deliver back to user
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		panic("Can't return remote command output: " + err.Error())
	}
}

func main() {
	port := 8010
	log.Printf("Starting webserver on port %d\n", port)
	handleRequests()
}
