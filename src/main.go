package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
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

type userHandler struct {
	client *redis.Client
}

const keyPrefix = "user:"

var appIP string

var vmUsername string = os.Getenv("VM_USERNAME")
var vmPassword string = os.Getenv("VM_PASSWORD")

// Existing code from above
func handleRequests() {
	redisHost := os.Getenv("REDIS_HOST")
    redisPassword := os.Getenv("REDIS_PASSWORD")
	op := &redis.Options{Addr: redisHost, Password: redisPassword, TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12}, WriteTimeout: 5 * time.Second, MaxRetries: 3}
	client := redis.NewClient(op)

	ctx := context.Background()
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("failed to connect with redis instance at %s - %v", redisHost, err)
	}

	uh := userHandler{client: client}

	myRouter := mux.NewRouter().StrictSlash(true)
	// creates a new instance of a mux router
	myRouter.HandleFunc("/test", testFunc).Methods("POST")

	// admin commands
	myRouter.HandleFunc("/init", initNet).Methods("POST")
	myRouter.HandleFunc("/clear", clearNet).Methods("POST")

	// student commands
	myRouter.HandleFunc("/creategroup", createGrp).Methods("POST")
	myRouter.HandleFunc("/registernumber", uh.registerNr).Methods("POST")
	myRouter.HandleFunc("/users/{Author}", uh.getUser).Methods("GET")
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
		User: vmUsername,
		Auth: []ssh.AuthMethod{
			ssh.Password(vmPassword)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	
	appIP = cp.IP + ":22"
	log.Println(appIP)
	conn, err := ssh.Dial("tcp", appIP, config)
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
		User: vmUsername,
		Auth: []ssh.AuthMethod{
			ssh.Password(vmPassword)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", appIP, config)
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
		User: vmUsername,
		Auth: []ssh.AuthMethod{
			ssh.Password(vmPassword)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", appIP, config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	runCommand("sudo /var/lib/waagent/custom-script/download/0/project/bloc-server/commands/createchannel.sh "+cp.Author+" "+cp.Group+" "+cp.Commit, conn, w)
}

func (uh userHandler) registerNr(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var u map[string]interface{}
	err = json.Unmarshal([]byte(reqBody), &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userid := u["Author"].(string)
	_, err = uh.client.HSet(r.Context(), keyPrefix+userid, u).Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uh userHandler) getUser(w http.ResponseWriter, r *http.Request) {
	userid := mux.Vars(r)["Author"]
	info, err := uh.client.HGetAll(r.Context(), keyPrefix+userid).Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(info)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Header().Del("Content-Type")
	}
}

func pushHash(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	var cp ContentPost
	json.Unmarshal(reqBody, &cp)

	config := &ssh.ClientConfig{
		User: vmUsername,
		Auth: []ssh.AuthMethod{
			ssh.Password(vmPassword)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", appIP, config)
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
		User: vmUsername,
		Auth: []ssh.AuthMethod{
			ssh.Password(vmPassword)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", appIP, config)
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
