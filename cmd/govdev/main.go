package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/ucladevx/govdev"
)

const (
	logo = `
_________            ________             
__  ____/________   ____  __ \_______   __
_  / __ _  __ \_ | / /_  / / /  _ \_ | / /
/ /_/ / / /_/ /_ |/ /_  /_/ //  __/_ |/ / 
\____/  \____/_____/ /_____/ \___/_____/  
`
	GOARCH string = runtime.GOARCH // Architecture running on
	GOOS   string = runtime.GOOS   // Operating system running on
)

var (
	Version   string
	GitHash   string
	BuildTime string
	GoVersion string = runtime.Version()
)

var mainCmd = &cobra.Command{
	Use:   "govdev",
	Short: "UCLA DevX management system",
	Run:   serve,
}

func init() {
	flags := mainCmd.Flags()

	flags.BoolP("version", "v", false, "print version of autofresh")
	// flags.Bool("hidebanner", false, "hide autofresh banner")
	// flags.IntP("port", "p", 0, "Serve directory files on localhost:<port>")
}

func version() {
	fmt.Printf("govdev %s %s %s/%s\n", Version, GoVersion, GOOS, GOARCH)
	fmt.Printf("git hash: %s\n", GitHash)
	fmt.Printf("built at: %s\n", BuildTime)
	os.Exit(0)
}

func main() {
	mainCmd.Execute()
}

func serve(c *cobra.Command, args []string) {
	v, _ := c.Flags().GetBool("version")
	if v {
		version()
	}
	conf, _ := govdev.LoadConfig()

	if !conf.Hidebanner {
		fmt.Println(logo)
	}

	govdev.Start(conf)
}
