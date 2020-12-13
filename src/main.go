package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/SophisticaSean/easyssh"
	"github.com/gorilla/mux"
)

// ContentPost is the json format of the content for all post calls
type ContentPost struct {
	Author string
	Group  string
	Commit string
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

	// TODO: check if user is admin
	cmd := exec.Command("/vagrant/commands/init.sh", cp.Author, cp.Group, cp.Commit)
	results, err := cmd.Output()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	// convert results into string and populate an instance of
	// the scriptResponse struct
	response := scriptResponse{string(results)}
	// encode response into JSON and deliver back to user
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		log.Println(err)
		return
	}
}

func clearNet(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	var cp ContentPost
	json.Unmarshal(reqBody, &cp)

	// TODO: check if user is admin
	cmd := exec.Command("/vagrant/commands/clear.sh")
	results, err := cmd.Output()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	// convert results into string and populate an instance of
	// the scriptResponse struct
	response := scriptResponse{string(results)}
	// encode response into JSON and deliver back to user
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		log.Println(err)
		return
	}
}

func createGrp(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	var cp ContentPost
	json.Unmarshal(reqBody, &cp)

	cmd := exec.Command("/vagrant/commands/createchannel.sh", cp.Author, cp.Group, cp.Commit)
	results, err := cmd.Output()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	// convert results into string and populate an instance of
	// the scriptResponse struct
	response := scriptResponse{string(results)}
	// encode response into JSON and deliver back to user
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		log.Println(err)
		return
	}
}

func pushHash(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	var cp ContentPost
	json.Unmarshal(reqBody, &cp)

	cmd := exec.Command("/vagrant/commands/push.sh", cp.Author, cp.Group, cp.Commit)
	results, err := cmd.Output()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
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
		return
	}
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	log.Println(reqBody)

	//TODO: change values
	ssh := &easyssh.SSHConfig{
		User:   "vagrant",
		Server: "127.0.0.1",
		// Optional key or Password without either we try to contact your agent SOCKET
		//Password: "password",
		Key:  "/Users/jalmeida/FCT/Tese/Thesis_Test/Online_Folder/server/.vagrant/machines/default/virtualbox/private_key",
		Port: "2222",
	}

	var cp ContentPost
	json.Unmarshal(reqBody, &cp)
	log.Println(cp)
	log.Println(cp.Author)
	log.Println(cp.Group)
	log.Println(cp.Commit)

	// Call Run method with command you want to run on remote server.
	stdout, stderr, done, err := ssh.Run("/vagrant/commands/test.sh "+cp.Author+" "+cp.Group+" "+cp.Commit, 60)
	// Handle errors
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println("don is :", done, "stdout is :", stdout, ";   stderr is :", stderr)
		encoder := json.NewEncoder(w)
		err = encoder.Encode(stdout)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	port := 8010
	log.Printf("Starting webserver on port %d\n", port)
	handleRequests()
}
