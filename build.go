package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	goarch      string
	goos        string
	debug       = os.Getenv("BUILDDEBUG") != ""
	newVersion  string
	version     string
	versionFile = "dep-agent/cmd/version.go"
)

var targets = map[string]target{
	"dep-agent": {
		name:        "deployment-agent",
		binaryName:  "dep-agent",
		buildPkg:    "github.com/dtchanpura/deployment-agent/dep-agent",
		buildDir:    "bin",
		description: "Deployment Agent target details.",
		archiveFiles: []archiveFile{
			{src: "{{binary}}", dst: "{{binary}}", perm: 0755},
			{src: "README.md", dst: "README.txt", perm: 0644},
			// {src: "LICENSE", dst: "LICENSE.txt", perm: 0644},
			// {src: "AUTHORS", dst: "AUTHORS.txt", perm: 0644},
		},
	},
}

type target struct {
	name         string
	description  string
	buildPkg     string
	binaryName   string
	buildDir     string
	archiveFiles []archiveFile
}

func (t target) BinaryName() string {
	if goos == "windows" {
		return t.binaryName + ".exe"
	}
	return t.binaryName
}

type archiveFile struct {
	src  string
	dst  string
	perm os.FileMode
}

func init() {
	// for targetName := range targets {
	// 	if targets[targetName].buildDir != "" {
	// 		for i := range targets[targetName].archiveFiles {
	// 			targets[targetName].archiveFiles[i].src = filepath.Join(targets[targetName].buildDir, targets[targetName].archiveFiles[i].src)
	// 		}
	// 	}
	// }
	if version == "0.0.0" || version == "" {
		versionFileBytes, errF := ioutil.ReadFile(versionFile)
		if errF != nil {
			fmt.Println(errF)
		}
		rv, rverr := regexp.Compile("version = \".*\"")
		if rverr != nil {
			fmt.Println(rverr)
		}
		versionBytes := rv.Find(versionFileBytes)
		rvt, _ := regexp.Compile("\".*\"")
		versionBytes = rvt.Find(versionBytes)
		if debug {
			log.Printf("Current version: %s", versionBytes)
		}
		version = string(versionBytes[1 : len(versionBytes)-1])
	}
}

func main() {
	parseFlags()

	targetName := "dep-agent"
	if flag.NArg() == 0 {
		runCommand("build", targets[targetName])
	} else {
		if flag.NArg() > 1 {
			targetName = flag.Arg(1)
		}

		runCommand(flag.Arg(0), targets[targetName])
	}
	os.Exit(0)
}

func parseFlags() {
	flag.StringVar(&goarch, "goarch", runtime.GOARCH, "GOARCH")
	flag.StringVar(&goos, "goos", runtime.GOOS, "GOOS")
	flag.StringVar(&newVersion, "version", "0.0.0", "Define new version")
	flag.Parse()
}

func buildStamp() string {
	layout := "2006-01-02 15:04:05 MST"
	// If SOURCE_DATE_EPOCH is set, use that.
	if s, _ := strconv.ParseInt(os.Getenv("SOURCE_DATE_EPOCH"), 10, 64); s > 0 {
		return time.Unix(s, 0).Format(layout)
	}

	// Try to get the timestamp of the latest commit.
	bs, err := runError("git", "show", "-s", "--format=%ct")
	if err != nil {
		// Fall back to "now".
		return time.Now().Format(layout)
	}

	s, _ := strconv.ParseInt(string(bs), 10, 64)
	return time.Unix(s, 0).Format(layout)
}

func updateVersion() error {
	var outputBytes []byte

	outputBytes, err := ioutil.ReadFile(versionFile)
	if err != nil {
		return err
	}
	// Replace build version
	if newVersion != "" && newVersion != "0.0.0" && newVersion != version {
		rv, rverr := regexp.Compile("version = \".*\"")
		if rverr != nil {
			return rverr
		}

		replaceStringVersion := fmt.Sprintf("version = \"%s\"", newVersion)
		outputBytes = rv.ReplaceAll(outputBytes, []byte(replaceStringVersion))

		// Replace build date
		rb, rberr := regexp.Compile("buildDate = \".*\"")
		if rberr != nil {
			return rberr
		}
		currentFileBytes, errF := ioutil.ReadFile(versionFile)
		if errF != nil {
			return errF
		}
		replaceStringDate := fmt.Sprintf("buildDate = \"%s\"", buildStamp())
		outputBytes = rb.ReplaceAll(outputBytes, []byte(replaceStringDate))
		if !bytes.Equal(currentFileBytes, outputBytes) {
			f, err := os.OpenFile(versionFile, os.O_WRONLY, 0644)
			if err != nil {
				return err
			}
			n, err := f.Write(outputBytes)
			if err != nil {
				return err
			}
			if debug {
				log.Printf("%d bytes written\n", n)
			}
			f.Close()
		}

	}
	return nil
}

func build(t target) {
	err := updateVersion()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	args := []string{"build"}
	if t.buildDir != "" {
		args = append(args, "-o", filepath.Join(t.buildDir, t.BinaryName()))
	}
	args = append(args, t.buildPkg)
	os.Setenv("GOOS", goos)
	os.Setenv("GOARCH", goarch)
	runPrint("go", args...)
}

func install(target target) {

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("GOBIN", filepath.Join(cwd, "bin"))
	args := []string{"install", "-v"}
	//"-ldflags", ldflags()}
	// if pkgdir != "" {
	// 	args = append(args, "-pkgdir", pkgdir)
	// }
	// if len(tags) > 0 {
	// 	args = append(args, "-tags", strings.Join(tags, " "))
	// }
	// if installSuffix != "" {
	// 	args = append(args, "-installsuffix", installSuffix)
	// }
	// if race {
	// 	args = append(args, "-race")
	// }
	args = append(args, target.buildPkg)

	os.Setenv("GOOS", goos)
	os.Setenv("GOARCH", goarch)
	runPrint("go", args...)
}

func buildTar(t target) {
	build(t)
	name := archiveName(t)
	filename := name + ".tar.gz"
	for i := range t.archiveFiles {
		t.archiveFiles[i].src = strings.Replace(t.archiveFiles[i].src, "{{binary}}", "bin/"+t.BinaryName(), 1)
		t.archiveFiles[i].dst = strings.Replace(t.archiveFiles[i].dst, "{{binary}}", t.BinaryName(), 1)
		t.archiveFiles[i].dst = strings.TrimLeft(name, t.buildDir+"/") + "/" + t.archiveFiles[i].dst
	}
	tarGz(filename, t.archiveFiles)
	fmt.Println(filename)
}

func clean() {
	rmr("bin", "")
}

func buildArch() string {
	os := goos
	if os == "darwin" {
		os = "macosx"
	}
	return fmt.Sprintf("%s-%s", os, goarch)
}

func archiveName(target target) string {
	filename := fmt.Sprintf("%s-%s-%s", target.name, buildArch(), version)
	if target.buildDir != "" {
		return filepath.Join(target.buildDir, filename)
	}
	return filename
}

func tarGz(out string, files []archiveFile) {
	fd, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}

	gw, err := gzip.NewWriterLevel(fd, gzip.BestCompression)
	if err != nil {
		log.Fatal(err)
	}
	tw := tar.NewWriter(gw)

	for _, f := range files {
		sf, verr := os.Open(f.src)
		if debug {
			log.Printf("Source: %s, Dest: %s with Perm: %s\n", f.src, f.dst, f.perm)
		}
		if verr != nil {
			log.Fatal(verr)
		}

		info, verr := sf.Stat()
		if verr != nil {
			log.Fatal(verr)
		}
		h := &tar.Header{
			Name:    f.dst,
			Size:    info.Size(),
			Mode:    int64(info.Mode()),
			ModTime: info.ModTime(),
		}

		verr = tw.WriteHeader(h)
		if verr != nil {
			log.Fatal(verr)
		}
		_, verr = io.Copy(tw, sf)
		if verr != nil {
			log.Fatal(verr)
		}
		sf.Close()
	}

	err = tw.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = gw.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = fd.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func runCommand(cmd string, t target) {
	switch cmd {
	case "build":
		build(t)
	case "tar":
		buildTar(t)
	case "install":
		install(t)
	case "clean":
		clean()
	}
}

func runError(cmd string, args ...string) ([]byte, error) {
	if debug {
		t0 := time.Now()
		log.Println("runError:", cmd, strings.Join(args, " "))
		defer func() {
			log.Println("... in", time.Since(t0))
		}()
	}
	ecmd := exec.Command(cmd, args...)
	bs, err := ecmd.CombinedOutput()
	return bytes.TrimSpace(bs), err
}

func runPrint(cmd string, args ...string) {
	if debug {
		t0 := time.Now()
		log.Println("runPrint:", cmd, strings.Join(args, " "))
		defer func() {
			log.Println("... in", time.Since(t0))
		}()
	}
	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	err := ecmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func rmr(paths ...string) {
	for _, path := range paths {
		if debug {
			log.Println("rm -r", path)
		}
		os.RemoveAll(path)
	}
}
