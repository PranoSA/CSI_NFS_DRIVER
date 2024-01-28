package main

/**
 *
 *
 * Something stupid, just create an HTTPS Server and listen on port 443
 * THen take Requests for creating NFS Shares
 * Then create the NFS Shares
 * Should I Add Authentication ?
 *
 */

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PranoSA/NFS_API_CSI/backend"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

/**
 * Path will be /nfs_shares/{uuid}
 */

func CreateVolume(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req backend.CreateVolumeRequest
	// Get the Name of the Volume

	//Json Unmarshal the Request Body into the CreateVolumeRequest Struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
	}

	var res backend.CreateVolumeResponse

	var uuid_string string = uuid.New().String()

	// Create Rnadom ID NF Share
	path := "/nfs_shares/" + uuid_string

	// Create the NFS Share By Opening File and ???
	// Create the NFS Share By Creating a Directory and ???
	err = os.MkdirAll(path, 0777)

	if err != nil {
		log.Println(err)
	}

	// Add an entry to /etc/exports
	exports, err := os.OpenFile("/etc/exports", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer exports.Close()

	_, err = exports.WriteString(path + " *(rw,sync,no_subtree_check)\n")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Capacity = req.RequestedBytes
	res.Path = path
	res.Host = "nfs://nfs_sevice"

	// Return the Response
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

}

func DeleteVolume(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("shareid")

	// Delete the NFS Share
	err := os.RemoveAll("/nfs_shares/" + id)

	if err != nil {
		log.Println(err)
	}

	// Remove the entry from /etc/exports
	exports, err := os.OpenFile("/etc/exports", os.O_RDWR, 0644)
	if err != nil {
		log.Println(err)
	}

	defer exports.Close()

	// Read the File

	// Remove the Line
	// Write the File

}

func UpdateVolume(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Ignore For Now
	return
}

func GetVolume(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Ignore For Now
	return
}

var port int

func main() {

	flag.IntVar(&port, "port", 443, "Port to listen on")

	flag.Parse()
	router := httprouter.New()
	router.POST("/volume", CreateVolume)
	router.DELETE("/volume/{shareid}", DeleteVolume)
	router.PUT("/volume/{shareid}", UpdateVolume)
	router.GET("/volume/{shareid}", GetVolume)

	server := &http.Server{
		Addr:     fmt.Sprintf(":%d", port),
		Handler:  router,
		ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	log.Fatal(server.ListenAndServe())
}
