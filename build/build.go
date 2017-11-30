package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var (
	goarch  string
	goos    string
	debug   = os.Getenv("BUILDDEBUG") != ""
	version string
)

var targets = map[string]target{
	"dep-agent": {
		name:       "deployment-agent",
		binaryName: "dep-agent",
		buildPkg:   "cgit.dcpri.me/deployment-agent/dep-agent",
		archiveFiles: []archiveFile{
			{src: "{{binary}}", dst: "{{binary}}", perm: 0755},
			{src: "../README.md", dst: "README.txt", perm: 0644},
			// {src: "../LICENSE", dst: "LICENSE.txt", perm: 0644},
			// {src: "../AUTHORS", dst: "AUTHORS.txt", perm: 0644},
		},
	},
}

type target struct {
	name         string
	description  string
	buildPkg     string
	binaryName   string
	archiveFiles []archiveFile
}

type archiveFile struct {
	src  string
	dst  string
	perm os.FileMode
}

func main() {
	parseFlags()

	err := updateVersion(version)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
	flag.StringVar(&version, "version", "0.0.0", "Define new version")
	flag.Parse()
}

func updateVersion(version string) error {
	versionFile := "../dep-agent/cmd/version.go"
	var outputBytes []byte
	buildDate := time.Now().Format("2006-01-02 15:04:05 MST")

	rv, rverr := regexp.Compile("version = \".*\"")
	if rverr != nil {
		return rverr
	}
	rb, rberr := regexp.Compile("buildDate = \".*\"")
	if rberr != nil {
		return rberr
	}
	outputBytes, err := ioutil.ReadFile(versionFile)
	if err != nil {
		return err
	}
	replaceStringVersion := fmt.Sprintf("version = \"%s\"", version)
	replaceStringDate := fmt.Sprintf("buildDate = \"%s\"", buildDate)
	outputBytes = rv.ReplaceAll(outputBytes, []byte(replaceStringVersion))
	outputBytes = rb.ReplaceAll(outputBytes, []byte(replaceStringDate))
	f, err := os.OpenFile(versionFile, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	n, err := f.Write(outputBytes)
	if err != nil {
		return err
	}
	fmt.Printf("%d bytes written\n", n)
	f.Close()
	return nil
}

func build(t target) {
	args := []string{"build"}
	args = append(args, t.buildPkg)
	os.Setenv("GOOS", goos)
	os.Setenv("GOARCH", goarch)
	runPrint("go", args...)
}

func buildTar(t target) {
	name := archiveName(t)
	filename := name + ".tar.gz"
	build(t)
	for i := range t.archiveFiles {
		t.archiveFiles[i].src = strings.Replace(t.archiveFiles[i].src, "{{binary}}", t.BinaryName(), 1)
		t.archiveFiles[i].dst = strings.Replace(t.archiveFiles[i].dst, "{{binary}}", t.BinaryName(), 1)
		t.archiveFiles[i].dst = name + "/" + t.archiveFiles[i].dst
	}
	tarGz(filename, t.archiveFiles)
	fmt.Println(filename)
}

func buildArch() string {
	os := goos
	if os == "darwin" {
		os = "macosx"
	}
	return fmt.Sprintf("%s-%s", os, goarch)
}

func archiveName(target target) string {
	return fmt.Sprintf("%s-%s-%s", target.name, buildArch(), version)
}

func (t target) BinaryName() string {
	if goos == "windows" {
		return t.binaryName + ".exe"
	}
	return t.binaryName
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
	}
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
