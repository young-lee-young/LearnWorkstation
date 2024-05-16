package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"strings"
	"time"
)

const (
	RegionId = "region-4Z3EM74rrxdbKwt6qXmAKk" // 张江
	// RegionId = "region-NfuPuzZFzVQDwKzau5byRm" // 办公室
)

const (
	VOLUME_FILE = "/Users/lee/Downloads/test.csv"
	PATH_FILE   = "/Users/lee/Downloads/test_path.csv"
	JSON_FILE   = "/Users/lee/Downloads/volume_map.json"
)

const (
	GB = 1024 * 1024 * 1024
)

const (
	ETCD_HOST   = "https://10.1.161.104:2379"
	ETCD_CERT   = "-----BEGIN CERTIFICATE-----\nMIIDpDCCAoygAwIBAgIIcujhad39QXQwDQYJKoZIhvcNAQELBQAwZzELMAkGA1UE\nBhMCQ04xEjAQBgNVBAgTCUd1YW5nZG9uZzERMA8GA1UEBxMIU2hlbnpoZW4xFjAU\nBgNVBAoTDVRlbmNlbnQgQ2xvdWQxGTAXBgNVBAMTEGV0Y2QtN2kzenV3azMtY2Ew\nHhcNMjQwMjAxMDgzNzQ2WhcNNDQwMjAxMDgzNzQ2WjBbMQswCQYDVQQGEwJDTjES\nMBAGA1UECBMJR3Vhbmdkb25nMREwDwYDVQQHEwhTaGVuemhlbjEWMBQGA1UEChMN\nVGVuY2VudCBDbG91ZDENMAsGA1UEAxMEcm9vdDCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAKyCSadubmmxQapR9d45O5atB9tvCERx7c7enVmQ49vC9HaK\nscZrjrnIR5TC85cPiWzZ6VLvRimuImnYYnE5c4TiwjqJ3PQk9qA80wFK9yj59dWd\nXY/Difzs/r61MymZv6CPDRY9KygnF+kJPyYqzLDKhbAVV7XEc2DokkDIceqXP3He\nh4DVGewy9th7oNjFRHhV3jTujwbAGRRvV3J5L1gW3mrbWZ+qokbIvUNAfc8zHLh4\nLgfDUP7ZkiM7fTYodD/u7beO/Ui05iZXnVk0s86BRceE42Pp/emG5jjZ0S4lVyZV\nO16j6MYZ1p3mpjhlbba4C0oBs6Vvp7IchNxoLAkCAwEAAaNgMF4wDgYDVR0PAQH/\nBAQDAgKEMB0GA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAMBgNVHRMBAf8E\nAjAAMB8GA1UdIwQYMBaAFAeyqS28ckPGmRwpdFIZixwGnbKJMA0GCSqGSIb3DQEB\nCwUAA4IBAQCREShye0pFTrNkesbPcneugKNSBzzgPRc3MtIIcWHD/K0IWbynDN0J\n/j923Q7nzyqXP+3M8X61RfqyXVhUJHWfPgAUv7s3DxbCE/GoHo5yEponzGCQ5Xy2\n2B6b2/YdO0S8THnKGyn3BlWptxkrcGMUqCVVqP3cwNfRXQvMTSnHUjqpupcxv5eG\nh5wfIQCgSSXl0A3oJqSg1uajMlHrMpNto3+fpzwZOvOC5mrBTzSTEdah7gLlr/0g\n1RtB+qxzLWj/KtHRzC20hRFbetGqr+ZH9hoFCC356vjSZPGpJces9UrFAgmqFOnt\njr4TXCt3MEXTdFSLLZwbTFe084fxpyOR\n-----END CERTIFICATE-----"
	ETCD_KEY    = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEArIJJp25uabFBqlH13jk7lq0H228IRHHtzt6dWZDj28L0doqx\nxmuOuchHlMLzlw+JbNnpUu9GKa4iadhicTlzhOLCOonc9CT2oDzTAUr3KPn11Z1d\nj8OJ/Oz+vrUzKZm/oI8NFj0rKCcX6Qk/JirMsMqFsBVXtcRzYOiSQMhx6pc/cd6H\ngNUZ7DL22Hug2MVEeFXeNO6PBsAZFG9XcnkvWBbeattZn6qiRsi9Q0B9zzMcuHgu\nB8NQ/tmSIzt9Nih0P+7tt479SLTmJledWTSzzoFFx4TjY+n96YbmONnRLiVXJlU7\nXqPoxhnWneamOGVttrgLSgGzpW+nshyE3GgsCQIDAQABAoIBAQCZosN0ET7WdJ5V\n0P7W9kPisqGfxWgohQrVRZ//h7HXI+7SsD0ZyOro/Lz7eRC2hYjP39cUHLNxJXvS\npnst+YuvWsC+n8+fxFe9Z85IhXpQDSXrJGaSRgTbtDpAoGI7d4mZti1ilAGelgqO\nacTK+i0oeoXLx4ZBBP/428h9V8dUs/6cexv+s/+xlr3AQhIqrl3dcZOew+3juCNm\nO0ou//Tr1Wg0OpH6ceqrccmQnkfHH7dOTKVh3a3/7TIkhBMV/aWIgaszh2Jcji8z\n/Q2NfxIdjKUitD2p2TtvFbyDf8l5/YVAe6ZKfRG5WUfGPI7Wqn6lbSiViL08immV\n4FdZqlwBAoGBAMosQ2JDn8JrweTzbDrsKHyFPfhSC6Zh5gpH6Fu026E+zm5HbIQq\no2GNYsThNjW4BAw+3sioI74c9rEgubHkBpMOreMN/SAuEUT7BQ9f8m+YgVrUKiER\nOUgwhG+uy+LRix0+xV1k643oaGvmP5oJ54tpSZLezWsazHoUkxekDZyBAoGBANpw\nLrvO6+H2SK1n/RWlzUZtH8wGUZrrpHE80w2Ss6paw92MviVKsdoLoVOUgtEoLFa3\npPNBWuFUE3DdST8rRNEJtolCmtwkztnFM4DhAXXEHgDFBew0dS+HyMrylMntiOZ4\nCMBxk2xuJJJh6aq0jLG8dwlJYK8rVa5BtOVL9euJAoGAM/lNGWXpHEI3hnRNzpgH\nkHtwgvYUqKFiUiJKSckk7T8IMrGJ5DBR8DI4wgRju0ujPsKq40rWZXlssfyHwIOu\nbP/JFCX0fOFa7zDZDCrOWIIMPx1r9zRGVZijc1Ksd5MHAG2yr1/Sbpjgv1xh3WfO\nbriSgX9lr7Mj/f5sEhfKqAECgYAP8trmqX/9eo7p4krPNtDbEIKe0eNJbWUPQ4Kh\nXhGyIXhlh64maOK9adX4lUJL5SP/nN71R9NL2mik9/MuIomevZfPQ8asxh84NTbC\ndSs3Hv+VzMzW5ymbM83MS5Pjfm1dbmwtKN27r/sJmQ6HgTE4lOsYCXx+rapO08tS\nZo5pCQKBgCY6oqS35JPVhShjotNgE6mOEqZ9TNxsrGpvdvcmGtqh1xJVa+sssMGc\n9XXBp5Tf9HKnWjjQfMCxaGjP9HSndj401RzGfv63vy+va/j/9BjkXqZjOu0Gw4rD\nhDf6iYAjx4WXYYIq8IOWVC1lJtgEqTSOicH2kGBo6meqDYpb3Fv/\n-----END RSA PRIVATE KEY-----"
	ETCD_CACERT = "-----BEGIN CERTIFICATE-----\nMIIDkjCCAnqgAwIBAgIILPrfOyu3aUUwDQYJKoZIhvcNAQELBQAwZzELMAkGA1UE\nBhMCQ04xEjAQBgNVBAgTCUd1YW5nZG9uZzERMA8GA1UEBxMIU2hlbnpoZW4xFjAU\nBgNVBAoTDVRlbmNlbnQgQ2xvdWQxGTAXBgNVBAMTEGV0Y2QtN2kzenV3azMtY2Ew\nHhcNMjQwMjAxMDgzNjM5WhcNNDQwMjAxMDgzNjM5WjBnMQswCQYDVQQGEwJDTjES\nMBAGA1UECBMJR3Vhbmdkb25nMREwDwYDVQQHEwhTaGVuemhlbjEWMBQGA1UEChMN\nVGVuY2VudCBDbG91ZDEZMBcGA1UEAxMQZXRjZC03aTN6dXdrMy1jYTCCASIwDQYJ\nKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKuNaO5P+Rt1qyboMLtO4uXP7HdoQ3R/\noGgOOPZsZpFC1WS0/cp1aBxuki0IrPLZv/S21yzc4+QCqqkF60ZkD+4kQyoJCUSf\nbk3Z3PrdJpikGp0m/JYH4Rjd2gU4+KwGvlyxffJ2Hu87cnMMcmnJiI9qmWHzYhkK\nA+VFrXbw5JT4+MlmV26Rebklk+p0miLsZhgM0TKCqik17E83QwqdJsZL9RLW4XfG\n7UgtrN2Sn07EVgD11WsvJXSDt8fEFs5cU9q/i7Is3SRAyxdpEzUti/L5IKvbtpiP\nHEqsZxnAmM9MjOZc7RrqG7cRdLlJ3DNiwpMkQTZ5pVRjyhzA77VUUBkCAwEAAaNC\nMEAwDgYDVR0PAQH/BAQDAgKEMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFAey\nqS28ckPGmRwpdFIZixwGnbKJMA0GCSqGSIb3DQEBCwUAA4IBAQAk9KZbCiIst+P5\nYK4z9ffJ9m1c+iih2T4FELpc/N7jspddV/k3bMLZ0t9/bMcOHXC2bvKWqXUzzzOU\nm/QuVq0pQYriu0szfpcKJAlK8VvQWynrR9cZtG52VOk3DOL2L8R5nKj0B38ve3Wd\nSfAHc+GpC+/1W3ooEmYHbXV8oP4y17OC5hQOJMGmY7gErZsSpdbPuL7f3wj24NQ8\nPgxsCu9WYTq2NXKQYb0EPQQa70amJOXXHTfJPBoVHTz+JqeS97l/iNc74Y+OZZm/\n8BCzY1PKUi2/5Xolg9xhsr6nLs8lPuwPzRi+YYMq2jdGfQZ0ZD8Yw7e5lnfbsENP\nUknzLlT0\n-----END CERTIFICATE-----"
)

const (
	CEPH_HOST   = "10.110.4.211,10.110.4.212,10.110.4.213"
	CEPH_USER   = "admin"
	CEPH_SECRET = "AQDe35tlWRZHORAAq74dDE2+ZJWCPbkNRjib3Q=="
)

const (
	VOLUME_STATUS_CREATING  = "creating"
	VOLUME_STATUS_READY     = "ready"
	VOLUME_STATUS_EXPANDING = "expanding"
)

type VolumeState struct {
	State        string
	Error        string
	ErrorMessage string
}

type Volume struct {
	Id       string
	Version  int64
	Type     string
	Size     int64
	State    *VolumeState
	Metadata map[string]string
	CreateAt time.Time
	UpdateAt time.Time
}

type VolumeCSV struct {
	Id          string `csv:"id"`
	CreatedTime string `csv:"created_time"`
	UpdatedTime string `csv:"updated_time"`
	Size        int64  `csv:"size"`
	Datacenter  string `csv:"datacenter"`
	User        string `csv:"user"`
	Name        string `csv:"name"`
}

type PathCSV struct {
	Id   string `csv:"id"`
	Path string `csv:"path"`
}

type StorageMap map[string]string

func main() {
	volumes, err := readFromCSV()
	if err != nil {
		panic(err)
	}

	paths, err := readPathFromCSV()
	if err != nil {
		panic(err)
	}

	pathMap := make(map[string]string)
	for _, path := range paths {
		pathMap[path.Id] = path.Path
	}

	cli, err := newEtcdCli()
	if err != nil {
		panic(err)
	}

	format := "2006-01-02 15:04:05"

	count := 0

	volumeMap := StorageMap{}

	for _, item := range volumes {
		if item.Datacenter != RegionId {
			continue
		}

		create, _ := time.Parse(format, item.CreatedTime)
		update, _ := time.Parse(format, item.UpdatedTime)

		path, ok := pathMap[item.Id]
		if !ok {
			fmt.Printf("volume %s not found path\n", item.Id)
			continue
		}

		id, _ := strings.CutPrefix(path, "/")

		volumeMap[item.Id] = id

		volume := &Volume{
			Id:      id,
			Version: 1,
			Type:    "cephfs",
			Size:    20 * GB,
			State: &VolumeState{
				State: VOLUME_STATUS_EXPANDING,
			},
			Metadata: map[string]string{
				"type":    "cephfs",
				"path":    path,
				"servers": CEPH_HOST,
				"name":    CEPH_USER,
				"secret":  CEPH_SECRET,
			},
			CreateAt: create,
			UpdateAt: update,
		}

		err = writeETCE(cli, volume)
		if err != nil {
			fmt.Printf("volume %s write etcd error\n", volume.Id)
			continue
		}

		count++
	}

	fmt.Println("write count: ", count)

	err = writeJson(volumeMap)
	if err != nil {
		fmt.Printf("write json file error: %v\n", err)
	}
}

func readFromCSV() ([]*VolumeCSV, error) {
	file, err := os.OpenFile(VOLUME_FILE, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var volumes []*VolumeCSV
	err = gocsv.UnmarshalFile(file, &volumes)
	return volumes, err
}

func readPathFromCSV() ([]*PathCSV, error) {
	file, err := os.OpenFile(PATH_FILE, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var paths []*PathCSV
	err = gocsv.UnmarshalFile(file, &paths)
	return paths, err
}

func writeETCE(cli *clientv3.Client, volume *Volume) error {
	ctx := context.Background()

	v, err := json.Marshal(volume)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("/nexus/volumes/%s", volume.Id)

	fmt.Println("write volume: ", volume.Id, volume.Size, volume.CreateAt, volume.UpdateAt, volume.Metadata)

	_, err = cli.Put(ctx, key, string(v))

	return err
}

func newEtcdCli() (*clientv3.Client, error) {
	// read cert and private key
	certificate, err := tls.X509KeyPair(
		[]byte(ETCD_CERT),
		[]byte(ETCD_KEY),
	)
	if err != nil {
		return nil, err
	}
	// read ca cert
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(ETCD_CACERT))
	// gen tls config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      caCertPool,
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{ETCD_HOST},
		TLS:         tlsConfig,
		DialTimeout: 3 * time.Second,
	})

	return cli, err
}

func writeJson(volumeMap StorageMap) error {
	file, err := os.OpenFile(JSON_FILE, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	v, _ := json.Marshal(volumeMap)

	fmt.Println("volume map: ", string(v))

	_, err = file.Write(v)

	return err
}
