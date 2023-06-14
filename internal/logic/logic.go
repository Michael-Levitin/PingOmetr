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
	db *PingsData
}

type PingsData struct {
	data    map[string]ob.PingUser
	lock    sync.RWMutex
	wg      *sync.WaitGroup
	fastest ob.PingUser
	slowest ob.PingUser
	admin   ob.PingAdmin
}

// NewPingLogic - создаем новую логику, подключаем список сайтов, обновление пингов/админа, ...
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
	db.setUserData()
	log.Print("logic: done")

	go db.updateAdminDataDB() // периодическое обновление данных для админов, на диске
	go db.updateUserData()    // периодическое обновление данных ping
	defer db.updateAdminDataDBOnce()

	return &PingLogic{&db}, nil
}

func (l PingLogic) GetFastest(ctx context.Context) (*ob.PingUser, error) {
	l.db.lock.RLock()
	user := l.db.fastest
	l.db.lock.RUnlock()

	l.db.lock.Lock()
	l.db.admin.Fastest++
	l.db.lock.Unlock()
	return &user, nil
}

func (l PingLogic) GetSlowest(ctx context.Context) (*ob.PingUser, error) {
	l.db.lock.RLock()
	user := l.db.slowest
	l.db.lock.RUnlock()

	l.db.lock.Lock()
	l.db.admin.Slowest++
	l.db.lock.Unlock()
	return &user, nil
}

func (l PingLogic) GetSpecific(ctx context.Context, site string) (*ob.PingUser, error) {
	l.db.lock.RLock()
	user, exist := l.db.data[site]
	if !exist {
		return nil, fmt.Errorf("site is not on the list")
	}
	l.db.lock.RUnlock()

	l.db.lock.Lock()
	l.db.admin.Specific++
	l.db.lock.Unlock()
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

// setUserData() обновляет значения ping один раз
func (d *PingsData) setUserData() {
	d.wg.Add(len(d.data))
	for site := range d.data {
		go d.ping(site)
	}
	d.wg.Wait()
	d.setFastSlowData()
}

// updateUserData() обновляет значения ping каждые userDataUpdateCoolDown секунд
func (d *PingsData) updateUserData() {
	for {
		time.Sleep(time.Second * userDataUpdateCoolDown)
		d.setUserData()
	}
}

// setFastSlowData() обновляет самый быстрый и самый медленный сайт
func (d *PingsData) setFastSlowData() {
	fast := ob.PingUser{
		Msec:  10000,
		Site:  "",
		Error: "",
	}
	slow := ob.PingUser{
		Msec:  0,
		Site:  "",
		Error: "",
	}

	d.lock.RLock()
	for _, user := range d.data {
		if user.Error == "" && user.Msec < fast.Msec {
			fast = user
		}
		if user.Error == "" && user.Msec > slow.Msec {
			slow = user
		}
	}
	d.lock.RUnlock()

	d.lock.Lock()
	d.fastest = fast
	d.slowest = slow
	d.lock.Unlock()
}

// setAdminData() считывает прошлые значение инфы для админов
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
		Slowest:  int64(adminI[0]),
		Fastest:  int64(adminI[1]),
		Specific: int64(adminI[2]),
	}
	return nil
}

// updateAdminDataDB() обновляет значения инфы для админов на диске каждые adminDataUpdateCoolDown секунд
func (d *PingsData) updateAdminDataDB() {
	for {
		time.Sleep(time.Second * adminDataUpdateCoolDown)
		d.updateAdminDataDBOnce()
	}
}

func (d *PingsData) updateAdminDataDBOnce() {
	d1 := []byte(fmt.Sprintf("%d %d %d", d.admin.Fastest, d.admin.Slowest, d.admin.Specific))
	err := os.WriteFile("./../../internal/logic/db/admin.txt", d1, 0644)
	if err != nil {
		log.Print("logic: failed updating admin db", err)
	}
}

// resetAdminDataDB() обнуляет значения инфы для админов на диске
func (d *PingsData) resetAdminDataDB() error {
	d1 := []byte("0 0 0")
	err := os.WriteFile("./../../internal/logic/db/admin.txt", d1, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (d *PingsData) ping(url string) {
	// используем ping cmd/терминала, 1 раз
	output, err := exec.Command("ping", "-c", "1", url).CombinedOutput()

	// вытаскиваем значение ping из ответа
	s := strings.Split(string(output), "\n")
	matches := re.FindStringSubmatch(s[1])
	if err == nil && len(matches) != 2 {
		err = fmt.Errorf("regex error")
	}

	var errString string
	if err != nil {
		switch err.Error() {
		case "exit status 2":
			errString = "no address associated with hostname"
		case "exit status 1":
			errString = "no reply from host - timeout"
		}
		d.lock.Lock()
		d.data[url] = ob.PingUser{
			Msec:  0,
			Site:  url,
			Error: errString,
		}
		d.lock.Unlock()
		d.wg.Done()
		return
	}

	ping, err := strconv.Atoi(matches[1])
	if err == nil && ping < 100 {
		ping++
	}

	if err != nil {
		errString = err.Error()
	}

	d.lock.Lock()
	d.data[url] = ob.PingUser{
		Msec:  int32(ping),
		Site:  url,
		Error: errString,
	}
	d.lock.Unlock()
	d.wg.Done()
	//fmt.Println(ping, url)
}
