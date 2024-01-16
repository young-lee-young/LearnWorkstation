//package main
//
//import (
//	"crypto/hmac"
//	"crypto/sha256"
//	"encoding/hex"
//	"fmt"
//	"net/url"
//)
//
//func getSignature(appKey, appId, nonce, timestamp, path string) string {
//	fmt.Printf("appKey: %s\n", appKey)
//	fmt.Printf("appId: %s\n", appId)
//	fmt.Printf("nonce: %s\n", nonce)
//	fmt.Printf("timestamp: %s\n", timestamp)
//	fmt.Printf("path: %s\n", path)
//	sig := hmac.New(sha256.New, []byte(appKey))
//	sig.Write([]byte(fmt.Sprintf("%s%s%s%s", appId, nonce, timestamp, path)))
//	sigResult := hex.EncodeToString(sig.Sum(nil))
//	return sigResult
//}
//
//func main() {
//	appId := "Ib8thze167CZi1pl"
//	appKey := "Wqo_tP315pkTK3EaODX!K9PnLGGIS6mI"
//
//	nonce := "A4BCA6A2-4BDF-4114-8823-1A6169C0A7D7"
//
//	timestamp := "1685933311817"
//
//	uri := "https://test-cg.ppio.cloud/openapi/v1/security_groups?regionId=CN-test-office-01"
//	path, err := url.Parse(uri)
//	if err != nil {
//		return
//	}
//
//	// path.Path 类似：/openapi/v1/security_groups
//	ret := getSignature(appKey, appId, nonce, timestamp, path.Path)
//
//	fmt.Println("ret: ", ret)
//}

package main

import (
	"fmt"
	"sync"
	"time"
)

var a = 1

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go add()
	time.Sleep(time.Second * 1)
	go reduce()

	time.Sleep(time.Second * 5)

	fmt.Println("a is ", a)
}

func add() {
	var mu sync.Mutex

	mu.Lock()
	fmt.Println("lock")
	a = a + 1
	time.Sleep(time.Second * 2)
	fmt.Println("unlock")
	mu.Unlock()
}

func reduce() {
	a = a - 1
}
