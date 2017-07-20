// updater
package main

import (
	"bufio"
	"crypto"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"runtime"
	"strings"

	"github.com/inconshreveable/go-update"
)

var updatservers = [...]string{
	// development servers
	"http://localhost:20163",
	"http://localhost:5555",
}

type Release struct {
	Exec   string
	SHA256 string
	sig    string
}

func checkUpdate() error {
	log.Println("checking for updates...")
	var err error
	var resp *http.Response
	err = errors.New("No new version")
	for _, updateserverurl := range updatservers {
		resp, err = http.Get(updateserverurl + "/index.txt")
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()

		// find release files (binaries, checksum, signature)
		var release Release
		release, err = getLatestRelease(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}

		// retrieve checksum
		resp, err = http.Get(updateserverurl + "/" + release.SHA256)
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()
		var checksum []byte
		checksum, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}

		// retrieve signature
		resp, err = http.Get(updateserverurl + "/" + release.sig)
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()
		var signature []byte
		signature, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}

		// retrieve binaries
		resp, err = http.Get(updateserverurl + "/" + release.Exec)
		if err != nil {
			log.Println(err)
			continue
		}
		defer resp.Body.Close()

		// Do the update
		if err = doUpdate(resp.Body, string(checksum), signature); err != nil {
			log.Println(err)
			continue
		}
		err = nil
	}
	return err
}

func getLatestRelease(reader io.Reader) (Release, error) {
	// create Release object
	var release Release

	// build regex to find all files for this system
	regexs := "^" + SOFTWARE + "_(.*)_" + runtime.GOOS + "_" + runtime.GOARCH + "\\.\\w+$"
	log.Println(regexs)
	regex, err := regexp.Compile(regexs)
	if err != nil {
		log.Fatal(err)
	}

	// scan text file for new version
	var files []string
	scanner := bufio.NewScanner(reader)
	latest := VERSION
	for scanner.Scan() {
		r := scanner.Text()
		files = append(files, r)
		if regex.MatchString(r) {
			if v := regex.FindAllStringSubmatch(r, 1)[0][1]; strings.Compare(v, latest) > 0 {
				latest = v
			}
		}
	}
	if strings.Compare(latest, VERSION) < 1 {
		return release, errors.New("No new version")
	}
	log.Println("New release available: " + latest)

	for _, s := range files {
		regex, err := regexp.Compile(SOFTWARE + "_" + latest + "_" + runtime.GOOS + "_" + runtime.GOARCH + "\\.\\w+$")
		if err != nil {
			log.Fatal(err)
		}
		if regex.MatchString(s) {
			release.Exec = s
		}
		regex, err = regexp.Compile(SOFTWARE + "_" + latest + "_" + runtime.GOOS + "_" + runtime.GOARCH + "\\.\\w+\\.sha256$")
		if err != nil {
			log.Fatal(err)
		}
		if regex.MatchString(s) {
			release.SHA256 = s
		}
		regex, err = regexp.Compile(SOFTWARE + "_" + latest + "_" + runtime.GOOS + "_" + runtime.GOARCH + "\\.\\w+\\.sha256\\.asc$")
		if err != nil {
			log.Fatal(err)
		}
		if regex.MatchString(s) {
			release.sig = s
		}
	}

	return release, nil
}

func doUpdate(binary io.Reader, hexChecksum string, signature []byte) error {
	checksum, err := hex.DecodeString(hexChecksum)
	if err != nil {
		return err
	}
	opts := update.Options{
		Hash:      crypto.SHA256,
		Verifier:  update.NewRSAVerifier(),
		Checksum:  checksum,
		Signature: signature,
	}
	err = opts.SetPublicKeyPEM(publicKey)
	if err != nil {
		return err
	}
	err = update.Apply(binary, opts)
	if err != nil {
		log.Println(err)
	}
	return err
}

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuzV/lr3rRW5t3twUs7aB
6UEri+d+eZk/P2XjC+jw8I1ZDTxZqHM8dXXyXJhP+ug4QG4yc6gwOMI0pEphLApw
RdODLDFlFQTrWgqTK0dkrGlVWVbewK/xTw8sAMqzfrrNqNmWvbH58MxKDH+VuX6h
jsSMlmK8UjlwUKxK319Gw0yqc7uKmG0QLPB/8dLFM+xR5MZ+Wc2csyJO1pzXUrfs
wkblMw2EBxX+NVnD9x6pjTrFJ/W1mzRU68cZGxlNu0fMZHLnW/gLXkuCQCfU0sl2
71T0o5ZvmUNqBfKKWYejBkNK2yMM8tKeHZ9DuQoMIBWnmPK28iKxy6j+fL5+Wzs+
wwIDAQAB
-----END PUBLIC KEY-----
`)
