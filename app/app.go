package app

import (
	"encoding/json"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"termux-ssh-scripts/config"
	"termux-ssh-scripts/termux"
	"termux-ssh-scripts/util"
)

type App struct {
	config *config.Config
	api *cloudflare.API
	zone *cloudflare.Zone
}

func New(config *config.Config) *App {
	a := new(App)
	a.config = config

	// Create new Cloudflare API client
	api, err := cloudflare.NewWithAPIToken(a.config.ApiToken)
	if err != nil {
		log.Fatal(err)
	}
	a.api = api

	// Get details of the given zone
	zone, err := api.ZoneDetails(config.ZoneId)
	if err != nil {
		log.Fatal(err)
	}
	a.zone = &zone

	return a
}

func (a *App) Install() {
	a.createConf()
	a.createUpdateScript()
	termux.InstallRequirements()
	termux.ScheduleJob()
}

func (a *App) Run() {
	h := a.hostnames()
	s := a.scripts(h)
	fmt.Println(s)
	a.updateShortcuts(s)
}

func (a *App) createConf() {
	fmt.Println("Creating configuration file...")
	h := util.HomeDir()
	f := h + "/.config/termux-ssh-scripts/config.json"

	_ = os.Mkdir(h + "/.config",0700)
	_ = os.Mkdir(h + "/.config/termux-ssh-scripts",0700)
	_ = os.Remove(f)

	j, err := json.MarshalIndent(a.config, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = ioutil.WriteFile(f, j, 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Configuration file created.")
}

func (a *App) createUpdateScript() {
	fmt.Println("Creating update script...")
	c := "#!/data/data/com.termux/files/usr/bin/bash\n" +
		"termux-ssh-scripts update"
	err := ioutil.WriteFile(
		"/data/data/com.termux/files/usr/bin/termux-ssh-scripts-update",
		[]byte(c),
		0700)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Update script created.")
}

func (a *App) hostnames() (hostnames []string) {
	// Get A records for zone
	records, err := a.api.DNSRecords(a.config.ZoneId, cloudflare.DNSRecord{Type: "A"})
	if err != nil {
		log.Fatal(err)
	}

	// Add names of A records to hostnames slice
	for _, r := range records {
		hostnames = append(hostnames, r.Name)
	}

	return
}

func (a *App) scripts(hostnames []string) map[string]string {
	m := make(map[string]string)

	for _, h := range hostnames {

		// Remove root domain from hostname
		t := strings.ReplaceAll(h, "." + a.zone.Name, "")

		// Split the rest of hostname into a slice
		s := strings.Split(t, ".")

		// Don't include this hostname if the slice length is less than 2 and
		// the last char of the last element is numeric
		l := len(s)
		if l < 2 {
			continue
		}
		last := s[l-1]
		char := last[len(last)-1:]
		if util.IsNumeric(char) {
			continue
		}

		// Reverse the slice
		r := util.Reverse(s)

		// Join the reverted slice with "-"
		n := strings.Join(r, "-")

		// Add result string to return map
		m[n] = h
	}

	return m
}

func (a *App) updateShortcuts(scripts map[string]string) {
	p := util.ShortcutsDir()

	_ = os.Mkdir(p, 0700)

	existingScripts := make(map[string]string)

	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fn := f.Name()
		if strings.HasPrefix(fn, "ssh--") {
			if util.MapContainsKey(scripts, fn[5:]) {
				existingScripts[fn[5:]] = ""
			} else {
				_ = os.Remove(p + "/" + f.Name())
			}
		}
	}

	for n, h := range scripts {
		if !util.MapContainsKey(existingScripts, n) {
			c := "#!/data/data/com.termux/files/usr/bin/bash\nssh root@" + h
			_ = ioutil.WriteFile(p + "/" + "ssh--" + n, []byte(c), 0700)
		}
	}
}
