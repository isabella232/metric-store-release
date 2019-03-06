// Code generated by go-bindata.
// sources:
// certs/metric-store-ca.crl
// certs/metric-store-ca.crt
// certs/metric-store-ca.key
// certs/metric-store.crt
// certs/metric-store.csr
// certs/metric-store.key
// DO NOT EDIT!

package testing

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _metricStoreCaCrl = []byte(`-----BEGIN X509 CRL-----
MIICiTBzAgEBMA0GCSqGSIb3DQEBCwUAMBoxGDAWBgNVBAMTD21ldHJpYy1zdG9y
ZS1jYRcNMTgwOTI0MTcyNzUxWhcNMjAwMzI0MTcyNzUwWjAAoCMwITAfBgNVHSME
GDAWgBSJRk6y0KBGbLwtzu6GvlEIv+RJATANBgkqhkiG9w0BAQsFAAOCAgEAhMo5
rygIYvVZIve5P/DsV7XBdIEeS/CBl36UWCJqSu45u33MSKJpBqBGEiDNaPoh2uHa
c29kqnWT4tNI9TG18GezmtySIKPXM0pa1ynNCdW2meERbCACuBKzqN2fqfqhK8Sr
WtZ/fgia3B6UZow46Sy5I77R9aDNK6c2gLBLgfoUBuyhQzAl7Vftct0JNpBGFcb3
8coYAUneXlqH2Z2YVKEndXwRJ9WPG2swE3pa5yAu0gr7I3FKSwvInlpGfM3MFmht
RuwAqV3S27G1jl9zcaQNoh5kHaJwOZMqTtbSF9TWfcFhxKeF0JiyKilW/mjHn2Wd
+WVCiOrwxcrBPxgZ6wSg9IJEB6bDg7gurDYsa0tJ6hSIh2u+1Y5Se0w8eDGxzA31
55Q3Fh2UQ91vuP5BZR5f5GBV4v8z/7uTILfHkhno1BQk7r2w04xJqddDfZmfzCco
C5GRjkCTutSMAveCBsz2fogWNJNpAgWcqI615iP3kWHpPPxkX5kyXAKVY8ZQF34R
VK3Pw4War+tqJdy2//fnxAodM4Yw9AC8PcUsmcB0NuJZwPvjnYcoE4Lr+ZnCX9Bl
fhGro2Ze7PtIWYmTSbr46dnOpCBaZrYWrfZjlsBPDlwl0B8R+Z5ovH8SFnOZKkZw
/+9m1LK9OUWLXKt8mNMnAz1Z7iwM1PQxJ0RDjQY=
-----END X509 CRL-----
`)

func metricStoreCaCrlBytes() ([]byte, error) {
	return _metricStoreCaCrl, nil
}

func metricStoreCaCrl() (*asset, error) {
	bytes, err := metricStoreCaCrlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metric-store-ca.crl", size: 934, mode: os.FileMode(292), modTime: time.Unix(1537810071, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricStoreCaCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIE9DCCAtygAwIBAgIBATANBgkqhkiG9w0BAQsFADAaMRgwFgYDVQQDEw9tZXRy
aWMtc3RvcmUtY2EwHhcNMTgwOTI0MTcyNzUwWhcNMjAwMzI0MTcyNzUwWjAaMRgw
FgYDVQQDEw9tZXRyaWMtc3RvcmUtY2EwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAw
ggIKAoICAQC2dGCNK0eXL6NMtYlBGJ7aCvUmsbMUpyGpFCC86LIGlQ9AqdpIadxz
5L4qC194iqxvi66UreZVAPXGwH0UIDmexRUtHZ/EVTKxVvDfRaLlDa0UtqbEb08D
U6mj1HEbCfT6HP56GKySn7Q22zuANU50x0w/m0g002xC7gYuGdGIeQMS9KZg/evb
PKdsEjuooiniBaH25EB4ZZnzB4qg+4Hyl+TSUWTv3mNZEYroJ2iHkNO6aMPZ9drs
7mKWJaHr2+JDMg8/wKF5d0TwfYZu+i1Ipin9PB2nKi7FW2RgeemjoqPinpbr0qPC
e657RDt1+7qzhFpDXbkhEdMj7MKpDH7S/mdf9oNArdX0HM0glhK2wG+9Dvu5oaT5
Zpq6jily3yhqbK8TVqJ//+vzsRjtXpsQwHhSg6l6ulAWMUGe5x3uKfyhyk4CjJdE
BBqmFOYCR++XEDZEedMm+EU7juLYLA11icJQ682JaUPZv765vFRZgn1JiPzCY4E4
0ge8VzfmuJX7IlDtN+LHPi1xmDl7qyUMxYi0byxU/I7wm66ujMpOJuQs39Ty889s
1pCoChNu9x/pZPW4++ydQZSVNzkk3hd4CqojpJUQsbT5A7XBCGBuZ9ZGLA1bFCMb
0Wc9xnnGSojQT/kemC9Hrxz61NJjfd7GSP8+AHiPv73mYKUWOPEGNQIDAQABo0Uw
QzAOBgNVHQ8BAf8EBAMCAQYwEgYDVR0TAQH/BAgwBgEB/wIBADAdBgNVHQ4EFgQU
iUZOstCgRmy8Lc7uhr5RCL/kSQEwDQYJKoZIhvcNAQELBQADggIBAHY1rti3ipjp
SrkTmXAxjUqk94SHjthAjcSqZHSKwc1SJLGLdOO1xBZCDEUFXsHm1Yo315gBaa3n
haF8f7C0xWz9QknItisvW1X+tPUvNe4JIXr6lGLts41vaUfKlZr5l/DCDUIte4eT
Dcya/E/YaWeKp9tdEg3zqRrVS04VaHgei/5uvdP9lATkYqPdLLb+1KOIs0MdpRr/
ZVTXVqbdUg4Xkn07osu8u6TMjRFrxcYi6QjkgW8wHUuErNiKJ4KETlvAw00utUWj
2oSSo5O8vifHjA3g1xq7lsQs7+iACx28klrCucY6vvJ/6si8wRsB3l/f/a5MQnKe
13K5jnko/Ox+DohBxbdsleq6gAvtsyHgTmH6Rumddzh26UY+PIgg5/ezfE1BKG4i
EaUzhTyLqCUw8QATg6NBuMoDtxWZwWtEVlMWUzj6knLgtPHoQcchcRxl3Ce/JZUb
ZnczMbwBbLkF7PjTDaDZEyoOaK8r/gJOHy0t/vhodXXyiZMeB+v7hVs/5H8BTtWt
d5OB7s4mD41tOCk/99rbAbYIMBF42bVOXQoNOZHiQmPT0cHCh56Kjmr3n9qJoNde
pzIYpFbu7syrzfne/po2kp4RNDZ+1m4vjquw+WTTAd0hqXL8B9aYs17O9+HOfAUi
GWKhFuVCFrPBQH2siL+Kx4JsNDWlLGXh
-----END CERTIFICATE-----
`)

func metricStoreCaCrtBytes() ([]byte, error) {
	return _metricStoreCaCrt, nil
}

func metricStoreCaCrt() (*asset, error) {
	bytes, err := metricStoreCaCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metric-store-ca.crt", size: 1777, mode: os.FileMode(292), modTime: time.Unix(1537810071, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricStoreCaKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`)

func metricStoreCaKeyBytes() ([]byte, error) {
	return _metricStoreCaKey, nil
}

func metricStoreCaKey() (*asset, error) {
	bytes, err := metricStoreCaKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metric-store-ca.key", size: 3243, mode: os.FileMode(288), modTime: time.Unix(1537810071, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricStoreCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIELDCCAhSgAwIBAgIQKuEbF5iPQVO1M41r71q2kTANBgkqhkiG9w0BAQsFADAa
MRgwFgYDVQQDEw9tZXRyaWMtc3RvcmUtY2EwHhcNMTgwOTI0MTcyNzUyWhcNMjAw
MzI0MTcyNzQ5WjAXMRUwEwYDVQQDEwxtZXRyaWMtc3RvcmUwggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDiGlctqzlXnR0mFhaIYDCSrYxF+jA+2ejKhlGX
RhQ5RSApj2V1BrkxlXuzog5weunIb9Xgi1b0llT7Rym1MZbkfhoZumtSdheIPlui
1mXkAGhCFV6u9VB0tvaBt1RllO3qshDdwxrtDcMo377VAVeuJ6bmLy2bmDa4Viio
VD3GsjFBjvvM7lyu75JGt9q1Pzox5GkWgOjulm8pHmw+tvc7Z92ezuJ6g8yQAvuq
n35mSsRJ6W3V+VSG8J6j2DZeIScEHqahRIuHrLQUHjBr2WoQFucyOf4hMhlNV6rl
7J1mTyuXqueb+2iGyZ6YTGkLDm/Pt6XkRV+0ODFXrp1eDoOPAgMBAAGjcTBvMA4G
A1UdDwEB/wQEAwIDuDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwHQYD
VR0OBBYEFB1LKgiwKMQiyBGfetjQCDB1iWzXMB8GA1UdIwQYMBaAFIlGTrLQoEZs
vC3O7oa+UQi/5EkBMA0GCSqGSIb3DQEBCwUAA4ICAQBZQh3qT47xU1yWUSMqRSex
rWSc4GifIfTWmdfdGj9ux0erHQKG+vm/nJ2sVTfuBMcLSgNtNMbo0+unZQshtXfM
wF86gQKBSkv+Dnys9LTMjgvzsiJr/dVe/TlmsG0GzUR+3dzqriVYUeOHcUt1Frhp
pp/3lXIpVKDZaJYts1oS3zCt7cUFpb3P5pGfkbeAOhIdSr78NpPVHGX/KBev7qi3
75bysOC4B7Fod0vxPVVTSkQL9CK5ZXozF45QLCBpCT9N60BDj0XnTbEdg/hKswiT
ygA7v8gx9lB+IP1dHqh9zHhSkn+WHmANMmkvNyuQ8upJDLuPA9zHrji6H7t+ofr+
/sz/iIcQgdn0KXtaunjFDjjFq46nV873ApYIfw/bq+briM/Um0bI09QNRCBXOJZP
h9vGw7DdH1sfIQsZtRAEyoKdTpgqrt+TsM2O2OUA6ZnOd37l9Gr+CeK7p50xgNGG
ce6Pa6U/iNcO1Ip6cDYVrSdJyrliHHctrTDkGoeY6mURDrPMYpmW4rxhhoZWOXpS
sKLm2E/x+AEMCIn+pRN31HxofMSzC1ByElf4hWoj58F8UHSHbeqE8qT92WjEAOyT
rQbDrRqkb/fZ46u6T6YKaVo3HuEMY8OvvpxaNfnXwZrvqpyEhUPwK6gOUjK6y7L8
cJxGIAtsnYQGoFuMhirubQ==
-----END CERTIFICATE-----
`)

func metricStoreCrtBytes() ([]byte, error) {
	return _metricStoreCrt, nil
}

func metricStoreCrt() (*asset, error) {
	bytes, err := metricStoreCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metric-store.crt", size: 1509, mode: os.FileMode(292), modTime: time.Unix(1537810072, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricStoreCsr = []byte(`-----BEGIN CERTIFICATE REQUEST-----
MIICXDCCAUQCAQAwFzEVMBMGA1UEAxMMbWV0cmljLXN0b3JlMIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4hpXLas5V50dJhYWiGAwkq2MRfowPtnoyoZR
l0YUOUUgKY9ldQa5MZV7s6IOcHrpyG/V4ItW9JZU+0cptTGW5H4aGbprUnYXiD5b
otZl5ABoQhVervVQdLb2gbdUZZTt6rIQ3cMa7Q3DKN++1QFXriem5i8tm5g2uFYo
qFQ9xrIxQY77zO5cru+SRrfatT86MeRpFoDo7pZvKR5sPrb3O2fdns7ieoPMkAL7
qp9+ZkrESelt1flUhvCeo9g2XiEnBB6moUSLh6y0FB4wa9lqEBbnMjn+ITIZTVeq
5eydZk8rl6rnm/tohsmemExpCw5vz7el5EVftDgxV66dXg6DjwIDAQABoAAwDQYJ
KoZIhvcNAQELBQADggEBADWpCSCJQ7jxMuqqc6IlWRAXBnpNmQbACn5UkEXX4BxM
zmvXGJ+plgCexGXtiT3yg/Q6xsQGnxwyt56rhwc2NgbtYMElYGh6kGdQHK4J9A/d
EBA4vHJRyXk+haszFNizNfJA3+oE/xJt4D1k18Ibcg9ZOMVSRms/BvuZElKhKEHP
RkRop4aCpjM8wBuTEirqJCWYAoRD45yJ4JKIA1xpDP/dyRCfsoiOc22QM0iUy+s/
kEhYgHD4nN21zYL7eJTqa+Jcpme3klMLFFcnFf9tRLOotXKxoclFLa7f8bQDzRI2
ZdSjQ7bMO/Y0VQwlLEX2rmmup3YuKnFfX0aYxw1/JBY=
-----END CERTIFICATE REQUEST-----
`)

func metricStoreCsrBytes() ([]byte, error) {
	return _metricStoreCsr, nil
}

func metricStoreCsr() (*asset, error) {
	bytes, err := metricStoreCsrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metric-store.csr", size: 895, mode: os.FileMode(292), modTime: time.Unix(1537810072, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metricStoreKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA4hpXLas5V50dJhYWiGAwkq2MRfowPtnoyoZRl0YUOUUgKY9l
dQa5MZV7s6IOcHrpyG/V4ItW9JZU+0cptTGW5H4aGbprUnYXiD5botZl5ABoQhVe
rvVQdLb2gbdUZZTt6rIQ3cMa7Q3DKN++1QFXriem5i8tm5g2uFYoqFQ9xrIxQY77
zO5cru+SRrfatT86MeRpFoDo7pZvKR5sPrb3O2fdns7ieoPMkAL7qp9+ZkrESelt
1flUhvCeo9g2XiEnBB6moUSLh6y0FB4wa9lqEBbnMjn+ITIZTVeq5eydZk8rl6rn
m/tohsmemExpCw5vz7el5EVftDgxV66dXg6DjwIDAQABAoIBACj2Pq7+YzAVPa+l
tsVUL9iS6mPPFKh+T+dI+EUgpA9QD6iX2vidlDWAcF210UndarcuU6APflxnU9QG
K28xIbYZhl783+6biF3ddjqE/waUcE7wkiu+M6pBb11GulqA5a1sPxP+AcXKhX3F
M0xhpuHdOyZei7OxmtBAg7MjN+HkFUhdIAMzd782IlipzE6a1MzHgmM9WLrV5qJ9
gFG4S1/qWRszJDQyHxnyJB+bSWOEfVvxLnZpLiawECxM5sGU6jwapMOvHrpANqnx
SVHEQ58mr6I0T7/IbNYL7jIgWp/pks0VnqY7Mu4DAl7SQzcri1uCqLPYFIEBnxLx
X2hzZhkCgYEA9g0mhTSQZklc2dt9wZqixyoLZyOSHtsF3kIVk+Ilr1ma5qB1TOXt
Sj4hq54eHbGdt3VKkV6Y78iiuhfR4jBzxooL5uLr+h5uYDT1Wk1N/uPQofPDOHQn
IWPLCgb6DvZsNRwTC2+N0MmZ6snE4ztXoqsxgGY/HxN52uHnrtaWGfsCgYEA6z60
p1E0Il4diwPABHB3VU5CznGiT7B0T27pkEFM7cAVmun9kpN7NKzi15GH+cZ3g/hh
JLTKnXP6IRV6VB/4oWt1B35EgcVgjC/jDsKjBSgKs1xvh0viAn0knEFoORc+LJwO
SB0faDwDNnOiR55dat0yuGbUM12eqYC+BO93PH0CgYEAg6bBp6BaawBx9/djMEYH
Nr2eYE8+DdhvKV7+oKPuOgadxSyx1rVn48OezG7L+mNg2hqeDW4qMpKNzziTio0W
RXLzr1RXovrJYBy18t1OEEXhCead3AT6MvlsWC8neP4NI9Wjswi7Pq2/90qCWHsW
BunGkMckmwIWvzEEgB49LiUCgYEAyJEZ7V2qmXKstY1o9V/+HlkvVFxGCrNjNyZV
NIy4TixrPz0o2QOtE+gjL7AAwtCXrYjjKiyKY1wycmqhdYAct2Oqz8y0FAB4pI1f
hYIlA4x0MyAoZq4n7/9Ka37IoYRTmi0jcBCEapZgFtjYDz/SXf5h9B6X5YN4dwQw
/nw8qPECgYBBRrXxIvva30kSse0KGDFuv0sPphyRjgy7xS31fOA/Qum7SXtmIH7A
xakJOThKeJUCCGnYs2R04czWu0Mk+nuqBREaWFdPslKqMl0GZu6bloSLFxCr4NB3
TaQg9pxMEDDUXKbOV1ChhnWt525dn0mQl20Cv++ZklxiUmvVwXHiFg==
-----END RSA PRIVATE KEY-----
`)

func metricStoreKeyBytes() ([]byte, error) {
	return _metricStoreKey, nil
}

func metricStoreKey() (*asset, error) {
	bytes, err := metricStoreKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metric-store.key", size: 1679, mode: os.FileMode(288), modTime: time.Unix(1537810072, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"metric-store-ca.crl": metricStoreCaCrl,
	"metric-store-ca.crt": metricStoreCaCrt,
	"metric-store-ca.key": metricStoreCaKey,
	"metric-store.crt":    metricStoreCrt,
	"metric-store.csr":    metricStoreCsr,
	"metric-store.key":    metricStoreKey,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"metric-store-ca.crl": {metricStoreCaCrl, map[string]*bintree{}},
	"metric-store-ca.crt": {metricStoreCaCrt, map[string]*bintree{}},
	"metric-store-ca.key": {metricStoreCaKey, map[string]*bintree{}},
	"metric-store.crt":    {metricStoreCrt, map[string]*bintree{}},
	"metric-store.csr":    {metricStoreCsr, map[string]*bintree{}},
	"metric-store.key":    {metricStoreKey, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
