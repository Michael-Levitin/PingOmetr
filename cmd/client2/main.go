package main

import (
	"context"
	"fmt"
	pb "github.com/Michael-Levitin/PingOmetr/proto"
	"log"
	"math/rand"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	host = "localhost"
	port = "5000"
)

func main() {
	// подключаемся к grpc серверу
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("could not connect to grpc server: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	grpcClient := pb.NewPingOmetrClient(conn) // передаем подключение в сервис
	pSearch := []string{"google.com", "youtube.com", "facebook.com", "baidu.com", "wikipedia.org", "qq.com", "taobao.com", "yahoo.com", "tmall.com", "amazon.com", "google.co.in", "twitter.com", "sohu.com", "jd.com", "live.com", "instagram.com", "sina.com.cn", "weibo.com", "google.co.jp", "reddit.com", "vk.com", "360.cn", "login.tmall.com", "blogspot.com", "yandex.ru", "google.com.hk", "netflix.com", "linkedin.com", "pornhub.com", "google.com.br", "twitch.tv", "pages.tmall.com", "csdn.net", "yahoo.co.jp", "mail.ru", "aliexpress.com", "alipay.com", "office.com", "google.fr", "google.ru", "google.co.uk", "microsoftonline.com", "google.de", "ebay.com", "microsoft.com", "livejasmin.com", "t.co", "bing.com", "xvideos.com", "google.ca"}

	pingDialer(&pSearch, grpcClient)
}

func pingDialer(s *[]string, grpcClient pb.PingOmetrClient) {
	sel := [3]string{"Fast", "Slow", "Specific"}
	for {
		selector := sel[randInt(len(sel))]
		siteRes := &pb.GetResponse{}
		var err error
		switch selector {
		case "Fast":
			siteRes, err = grpcClient.GetFastest(context.TODO(), &pb.GetFastestRequest{})
		case "Slow":
			siteRes, err = grpcClient.GetSlowest(context.TODO(), &pb.GetSlowestRequest{})
		case "Specific":
			site := (*s)[randInt(len((*s)))]
			siteRes, err = grpcClient.GetSpecific(context.TODO(), &pb.GetSpecificRequest{
				SiteName: site,
			})
		}
		log.Println(selector, siteRes, err)
	}
}

func randInt(num int) int {
	rand.Seed(time.Now().UnixNano() + rand.Int63()) // for truly? random
	return rand.Intn(num)
}

//// печатаем результат
//func printAnswer(books []*pb.Book, err error) {
//	if err != nil {
//		log.Println("failed to execute request: ", err)
//	}
//	if books != nil && len(books) != 0 {
//		fmt.Println("Author\t\tTitle")
//		fmt.Println("---------------------------------------")
//		for _, book := range books {
//			fmt.Println(book.Name, "-", book.Title)
//		}
//	} else {
//		fmt.Println(" - Nothing's found")
//	}
//	fmt.Println("==============================================")
//}
