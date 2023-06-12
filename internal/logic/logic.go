package logic

import (
	"context"
	"encoding/csv"
	"fmt"
	ob "github.com/Michael-Levitin/PingOmetr/internal/objects"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var re = regexp.MustCompile(`time=(\d+).+`)

const (
	adminDataUpdateCoolDown = 600 // seconds
	userDataUpdateCoolDown  = 60  // seconds
)

type PingLogic struct {
	db *PingsData // TODO - change to db
}

type PingsData struct {
	data    map[string]ob.PingUser
	lock    sync.RWMutex
	wg      *sync.WaitGroup
	fastest ob.PingUser
	slowest ob.PingUser
	admin   ob.PingAdmin
}

// NewPingLogic() создаем новую логику, подключаем список сайтов, обновление пингов/админа, ...
func NewPingLogic() (*PingLogic, error) {
	var err error
	db := PingsData{wg: &sync.WaitGroup{}, lock: sync.RWMutex{}}

	err = db.setSites()
	if err != nil {
		return nil, err
	}
	err = db.setAdminData()
	if err != nil {
		return nil, err
	}

	log.Print("logic: updating initial ping data ...")
	db.wg.Add(1)
	db.updateUserData()
	db.wg.Wait()
	log.Print("logic: done")

	go db.updateAdminDataDB() // периодическое обновление данных для админов, на диске
	go db.setUserData()       // периодическое обновление данных ping

	return &PingLogic{&db}, nil
}

func (l PingLogic) GetFastest(ctx context.Context) (*ob.PingUser, error) {
	l.db.lock.RLock()
	user := l.db.fastest
	l.db.lock.RUnlock()
	return &user, nil
}

func (l PingLogic) GetSlowest(ctx context.Context) (*ob.PingUser, error) {
	l.db.lock.RLock()
	user := l.db.slowest
	l.db.lock.RUnlock()
	return &user, nil
}

func (l PingLogic) GetSpecific(ctx context.Context, site string) (*ob.PingUser, error) {
	l.db.lock.RLock()
	user := l.db.data[site]
	l.db.lock.RUnlock()
	return &user, nil
}

func (l PingLogic) GetAdminData(ctx context.Context) (*ob.PingAdmin, error) {
	l.db.lock.RLock()
	admin := l.db.admin
	l.db.lock.RUnlock()
	return &admin, nil
}

// setSites() копирует сайты в мапу логики
func (d *PingsData) setSites() error {
	csvData, err := os.Open("./../../internal/logic/db/sites.csv")
	if err != nil {
		return err
	}
	defer csvData.Close()

	r, err := csv.NewReader(csvData).ReadAll()
	if err != nil {
		return err
	}

	d.data = make(map[string]ob.PingUser)
	for _, row := range r {
		//d.data[`https://www.`+row[0]] = ob.PingUser{}
		d.data[row[0]] = ob.PingUser{}
	}
	return nil
}

func (d *PingsData) updateUserData() {
	d.setUserDataOnce()
	d.wg.Done()
}

func (d *PingsData) setAdminData() error {
	dat, err := os.ReadFile("./../../internal/logic/db/admin.txt")
	if err != nil {
		return err
	}
	//log.Print("logic: opened admin file")

	admin := strings.Split(string(dat), " ")
	if len(admin) != 3 {
		return fmt.Errorf("not enouph data in admin file")
	}

	adminI := [3]int{}
	for i := 0; i < 3; i++ {
		adminI[i], err = strconv.Atoi(admin[i])
		if err != nil {
			return err
		}
	}

	d.admin = ob.PingAdmin{
		Min:      int32(adminI[0]),
		Max:      int32(adminI[1]),
		Specific: int32(adminI[2]),
	}
	return nil
}

func (d *PingsData) updateAdminDataDB() {
	for {
		time.Sleep(time.Second * adminDataUpdateCoolDown)
		d1 := []byte(fmt.Sprintf("%d %d %d", d.admin.Max, d.admin.Min, d.admin.Specific))
		err := os.WriteFile("./../../internal/logic/db/admin.txt", d1, 0644)
		if err != nil {
			log.Print("logic: failed updating admin db", err)
		}
	}
}

func (d *PingsData) resetAdminDataDB() error {
	d1 := []byte("0 0 0")
	err := os.WriteFile("./../../internal/logic/db/admin.txt", d1, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (d *PingsData) setUserDataOnce() {
	d.wg.Add(len(d.data))
	for site := range d.data {
		go d.ping(site, true)
	}
}

func (d *PingsData) setUserData() {
	for {
		time.Sleep(time.Second * userDataUpdateCoolDown)
		for site := range d.data {
			go d.ping(site, false)
		}
	}
}

func (d *PingsData) ping(url string, wgON bool) {
	// используем ping cmd/терминала, 1 раз
	output, err := exec.Command("ping", "-c", "1", url).CombinedOutput()

	// вытаскиваем значение ping из ответа
	s := strings.Split(string(output), "\n")
	matches := re.FindStringSubmatch(s[1])
	if err == nil && len(matches) != 2 {
		err = fmt.Errorf("regex error")
	}

	if err != nil {
		switch err.Error() {
		case "exit status 2":
			err = fmt.Errorf("no address associated with hostname")
		case "exit status 1":
			err = fmt.Errorf("no reply from host - timeout")
		}
		d.lock.Lock()
		d.data[url] = ob.PingUser{
			Msec:  0,
			Site:  url,
			Error: err,
		}
		d.lock.Unlock()
		if wgON {
			d.wg.Done()
		}
		return
	}

	ping, err := strconv.Atoi(matches[1])
	if ping < 100 {
		ping++
	}

	//fmt.Println(ping, url)

	d.lock.Lock()
	d.data[url] = ob.PingUser{
		Msec:  int32(ping),
		Site:  url,
		Error: err,
	}
	d.lock.Unlock()
	if wgON {
		d.wg.Done()
	}
}
