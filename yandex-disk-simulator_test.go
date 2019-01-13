package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	SyncDirPath    = "$HOME/TeSt_Yandex.Disk_TeSt"
	ConfigFilePath = "$HOME/.config/TeSt_Yandex.Disk_TeSt"
)

func TestMain(m *testing.M) {
	flag.Parse()

	exec.Command("go", "build", "yandex-disk-simulator.go").Run()
	cwd, _ := os.Getwd()
	os.Setenv("PATH", cwd+":"+os.Getenv("PATH"))
	SyncDirPath = os.ExpandEnv(SyncDirPath)
	os.Setenv("Sim_SyncDir", SyncDirPath)
	ConfigFilePath = os.ExpandEnv(ConfigFilePath)
	os.Setenv("Sim_ConfDir", ConfigFilePath)

	// Run tests
	e := m.Run()

	// Clearance
	exec.Command("yandex-disk-simulator", "stop").Run()
	os.RemoveAll(ConfigFilePath)
	os.RemoveAll(SyncDirPath)

	os.Exit(e)
}

func getLogEvent(timeout time.Duration) error {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watch.Close()
	err = watch.Add(filepath.Join(SyncDirPath, ".sync/cli.log"))
	if err != nil {
		//log.Debug("Watch path error:", err)
		return err
	}
	select {
	case err := <-watch.Errors:
		return err
	case <-watch.Events:
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("Timeout")
	}
}

func Example01StartWithoutCommand() {
	cmd := exec.Command("yandex-disk-simulator")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Error: command hasn't been specified. Use the --help command to access help
	// or setup to launch the setup wizard.
}

func Example05StartUnconfigured() {
	cmd := exec.Command("yandex-disk-simulator", "start")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Error: option 'dir' is missing --
}

func Example10StartSetup() {
	cmd := exec.Command("yandex-disk-simulator", "setup")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	//
}

func Example20StartSuccess() {
	cmd := exec.Command("yandex-disk-simulator", "start")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Starting daemon process...Done
}

func Example25SecondStart() {
	cmd := exec.Command("yandex-disk-simulator", "start")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Daemon is already running.
}

func Example40StatusAfterStart() {
	cmd := exec.Command("yandex-disk-simulator", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	//
}

func Example45StatusAfter1stEvent() {
	if getLogEvent(time.Duration(2*time.Second)) != nil {
		return
	}
	cmd := exec.Command("yandex-disk-simulator", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Synchronization core status: paused
	// Path to Yandex.Disk directory: '/home/stc/Yandex.Disk'
	// 	The quota has not been received yet.
	//
	// Last synchronized items:
	// 	file: 'File.ods'
	// 	file: 'downloads/file.deb'
	// 	file: 'downloads/setup'
	// 	file: 'download'
	// 	file: 'down'
	// 	file: 'do'
	// 	file: 'd'
	// 	file: 'o'
	// 	file: 'w'
	// 	file: 'n'
}

func Example50StatusAfter2ndEvent() {
	if getLogEvent(time.Duration(2*time.Second)) != nil {
		return
	}
	cmd := exec.Command("yandex-disk-simulator", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Synchronization core status: index
	// Path to Yandex.Disk directory: '/home/stc/Yandex.Disk'
	// 	The quota has not been received yet.
}

func Example55StatusAfter3rdEvent() {
	if getLogEvent(time.Duration(2*time.Second)) != nil {
		return
	}
	cmd := exec.Command("yandex-disk-simulator", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Synchronization core status: busy
	// Path to Yandex.Disk directory: '/home/stc/Yandex.Disk'
	// 	The quota has not been received yet.
	//
	// Last synchronized items:
	// 	file: 'File.ods'
	// 	file: 'downloads/file.deb'
	// 	file: 'downloads/setup'
	// 	file: 'download'
	// 	file: 'down'
	// 	file: 'do'
	// 	file: 'd'
	// 	file: 'o'
	// 	file: 'w'
	// 	file: 'n'
}

func Example60StatusAfter4rdEvent() {
	if getLogEvent(time.Duration(2*time.Second)) != nil {
		return
	}
	cmd := exec.Command("yandex-disk-simulator", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Synchronization core status: index
	// Path to Yandex.Disk directory: '/home/stc/Yandex.Disk'
	// 	The quota has not been received yet.
	//
	// Last synchronized items:
	// 	file: 'File.ods'
	// 	file: 'downloads/file.deb'
	// 	file: 'downloads/setup'
	// 	file: 'download'
	// 	file: 'down'
	// 	file: 'do'
	// 	file: 'd'
	// 	file: 'o'
	// 	file: 'w'
	// 	file: 'n'
}

func Example65StatusAfter5rdEvent() {
	if getLogEvent(time.Duration(6*time.Second)) != nil {
		return
	}
	cmd := exec.Command("yandex-disk-simulator", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Synchronization core status: idle
	// Path to Yandex.Disk directory: '/home/stc/Yandex.Disk'
	// 	Total: 43.50 GB
	// 	Used: 2.89 GB
	// 	Available: 40.61 GB
	// 	Max file size: 50 GB
	// 	Trash size: 0 B
	//
	// Last synchronized items:
	// 	file: 'File.ods'
	// 	file: 'downloads/file.deb'
	// 	file: 'downloads/setup'
	// 	file: 'download'
	// 	file: 'down'
	// 	file: 'do'
	// 	file: 'd'
	// 	file: 'o'
	// 	file: 'w'
	// 	file: 'n'
}

func Example70Sync() {
	cmd := exec.Command("yandex-disk-simulator", "sync")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	//
}

func Example75StatusAfterSyncStart() {
	cmd := exec.Command("yandex-disk-simulator", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Synchronization core status: index
	// Path to Yandex.Disk directory: '/home/stc/Yandex.Disk'
	// 	Total: 43.50 GB
	// 	Used: 2.89 GB
	// 	Available: 40.61 GB
	// 	Max file size: 50 GB
	// 	Trash size: 0 B
	//
	// Last synchronized items:
	// 	file: 'File.ods'
	// 	file: 'downloads/file.deb'
	// 	file: 'downloads/setup'
	// 	file: 'download'
	// 	file: 'down'
	// 	file: 'do'
	// 	file: 'd'
	// 	file: 'o'
	// 	file: 'w'
	// 	file: 'n'
}

func Example80StatusAfter2ndEvent() {
	if getLogEvent(time.Duration(2*time.Second)) != nil {
		return
	}
	cmd := exec.Command("yandex-disk-simulator", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Sync progress: 0 MB/ 139.38 MB (0 %)
	// Synchronization core status: busy
	// Path to Yandex.Disk directory: '/home/stc/Yandex.Disk'
	// 	Total: 43.50 GB
	// 	Used: 2.89 GB
	// 	Available: 40.61 GB
	// 	Max file size: 50 GB
	// 	Trash size: 0 B
	//
	// Last synchronized items:
	// 	file: 'File.ods'
	// 	file: 'downloads/file.deb'
	// 	file: 'downloads/setup'
	// 	file: 'download'
	// 	file: 'down'
	// 	file: 'do'
	// 	file: 'd'
	// 	file: 'o'
	// 	file: 'w'
	// 	file: 'n'
}

func Example85StatusAfter3rdEvent() {
	if getLogEvent(time.Duration(1*time.Second)) != nil {
		return
	}
	cmd := exec.Command("yandex-disk-simulator", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Sync progress: 65.34 MB/ 139.38 MB (46 %)
	// Synchronization core status: busy
	// Path to Yandex.Disk directory: '/home/stc/Yandex.Disk'
	// 	Total: 43.50 GB
	// 	Used: 2.89 GB
	// 	Available: 40.61 GB
	// 	Max file size: 50 GB
	// 	Trash size: 0 B
	//
	// Last synchronized items:
	// 	file: 'File.ods'
	// 	file: 'downloads/file.deb'
	// 	file: 'downloads/setup'
	// 	file: 'download'
	// 	file: 'down'
	// 	file: 'do'
	// 	file: 'd'
	// 	file: 'o'
	// 	file: 'w'
	// 	file: 'n'
}

func Example87StatusAfter4rdEvent() {
	if getLogEvent(time.Duration(3*time.Second)) != nil {
		return
	}
	cmd := exec.Command("yandex-disk-simulator", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Sync progress: 139.38 MB/ 139.38 MB (100 %)
	// Synchronization core status: index
	// Path to Yandex.Disk directory: '/home/stc/Yandex.Disk'
	// 	Total: 43.50 GB
	// 	Used: 2.89 GB
	// 	Available: 40.61 GB
	// 	Max file size: 50 GB
	// 	Trash size: 0 B
	//
	// Last synchronized items:
	// 	file: 'NewFile'
	// 	file: 'File.ods'
	// 	file: 'downloads/file.deb'
	// 	file: 'downloads/setup'
	// 	file: 'download'
	// 	file: 'down'
	// 	file: 'do'
	// 	file: 'd'
	// 	file: 'o'
	// 	file: 'w'
}

func Example88StatusAfter5rdEvent() {
	if getLogEvent(time.Duration(1*time.Second)) != nil {
		return
	}
	cmd := exec.Command("yandex-disk-simulator", "status")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Synchronization core status: idle
	// Path to Yandex.Disk directory: '/home/stc/Yandex.Disk'
	// 	Total: 43.50 GB
	// 	Used: 2.89 GB
	// 	Available: 40.61 GB
	// 	Max file size: 50 GB
	// 	Trash size: 0 B
	//
	// Last synchronized items:
	// 	file: 'File.ods'
	// 	file: 'downloads/file.deb'
	// 	file: 'downloads/setup'
	// 	file: 'download'
	// 	file: 'down'
	// 	file: 'do'
	// 	file: 'd'
	// 	file: 'o'
	// 	file: 'w'
	// 	file: 'n'
}

func Example90Stop() {
	cmd := exec.Command("yandex-disk-simulator", "stop")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Daemon stopped.
}

func Example90SecondaryStop() {
	cmd := exec.Command("yandex-disk-simulator", "stop")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	// Output:
	// Error: daemon not started
}
