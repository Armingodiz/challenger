package RandomData

import (
	crand "crypto/rand"
	"fmt"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"math/rand"
	"time"
)

var DataBroker []models.BrokerData
var DataCache []models.CacheData

func GetData() ([]models.BrokerData, []models.CacheData) {
	return DataBroker, DataCache
}

func init() {
	DataBroker = make([]models.BrokerData, 200)
	DataCache = make([]models.CacheData, 200)
	for i := 0; i < 200; i++ {
		min, max := getMinMax(i % 4)
		data := models.BrokerData{
			UserName:     getName(4),
			ID:           1000 + i,
			TrafficUsage: rand.Intn((max - min) + min),
			Ip:           genIpAddr(),
			Port:         "8080",
		}
		DataBroker = append(DataBroker, data)
		DataCache = append(DataCache, models.CacheData{Ip: data.Ip, Mac: getMacAdd()})
	}
}

func getMacAdd() string {
	buf := make([]byte, 6)
	_, err := crand.Read(buf)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	buf[0] |= 2
	return fmt.Sprintf("Random MAC address: %02x:%02x:%02x:%02x:%02x:%02x\n", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
}

func genIpAddr() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}

func getMinMax(n int) (int, int) {
	switch n {
	case 0:
		return 0, 100
	case 1:
		return 101, 500
	case 2:
		return 501, 1000
	case 3:
		return 1001, 1500
	default:
		return 0, 0
	}
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func getName(length int) string {
	return StringWithCharset(length, charset)
}
