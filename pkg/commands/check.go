package commands

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/valyala/fastjson"
)

type RustEnv struct {
	Branch           Branch
	LinuxManifestID  string
	CommonManifestID string
}

type Branch struct {
	BuildID   string
	UpdatedAt string
}

func Check(c *cli.Context) error {
	appId := c.String("app-id")
	linuxDepotId := c.String("linux-depot-id")
	commonDepotId := c.String("common-depot-id")
	branch := c.String("branch")
	export := c.Bool("export")

	res, err := http.Get(fmt.Sprintf("https://api.steamcmd.net/v1/info/%s", appId))
	if err != nil {
		return errors.Wrap(err, "couldnt get app information")
	}

	buf := &bytes.Buffer{}
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return errors.Wrap(err, "couldnt read steamcmd response")
	}
	if err := res.Body.Close(); err != nil {
		panic(err)
	}

	v, err := fastjson.Parse(buf.String())
	if err != nil {
		return errors.Wrap(err, "invalid json received from steamcmd")
	}

	// Validate the requested branch is present in results
	if !v.Exists("data", appId, "depots", "branches", branch) {
		return fmt.Errorf("requests branch '%s' does not exist", branch)
	}

	buildId := ""
	if len(c.String("build-id")) != 0 {
		buildId = c.String("build-id")
	} else {
		buildId = strings.Trim(v.Get("data", appId, "depots", "branches", branch, "buildid").String(), `"`)
	}

	buildUpdatedTime := strings.Trim(v.Get("data", appId, "depots", "branches", branch, "timeupdated").String(), `"`)
	linuxManifestId := strings.Trim(v.Get("data", appId, "depots", linuxDepotId, "manifests", branch, "gid").String(), `"`)
	commonManifestId := strings.Trim(v.Get("data", appId, "depots", commonDepotId, "manifests", branch, "gid").String(), `"`)


	if !export {
		log.Printf("'%s' branch last updated at %s.", branch, buildUpdatedTime)
		log.Printf("Build ID: %s", buildId)
		log.Printf("Linx Manifest ID: %s", linuxManifestId)
		log.Printf("Common Manifest ID: %s", commonManifestId)
	} else {
		fmt.Printf("BUILD_ID=%s\n", buildId)
		fmt.Printf("LINUX_MANIFEST_ID=%s\n", linuxManifestId)
		fmt.Printf("COMMON_MANIFEST_ID=%s\n", commonManifestId)
	}

	return nil
}
