package inireader

import (
	"os"
	"time"
)

// Listener is the interface of listener
type Listener interface {
	Listen(string) error
}

// FileListener is the interface of FileListener
type FileListener struct {
}

// Listen implemented by FileListener
func (fl FileListener) Listen(filePath string) error {
	initialStat, err := os.Stat(filePath)
	initialStatSize := initialStat.Size()
	initialStatModTime := initialStat.ModTime()
	if err != nil {
		return err
	}

	for {
		stat, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		if stat.Size() != initialStatSize || stat.ModTime() != initialStatModTime {
			// fmt.Println("A file change detected, how do you want to deal load it again? y/Y for Yes, n/N for No")

			// // var choose string
			// fmt.Scanln(&choose)

			// for choose != "y" && choose != "Y" && choose != "N" && choose != "n" && choose != "exit" {
			// 	fmt.Println("Invalid option! Do you want to load it again? y/Y for Yes, n/N for No")
			// 	fmt.Scanln(&choose)
			// }

			// switch choose {
			// case "y", "Y":
			// 	Load(filePath)
			// 	fmt.Println("Reload!")
			// case "n", "N":
			// 	fmt.Println("Fine")
			// case "exit":
			// 	return nil
			// }

			break

		}
		initialStatSize = stat.Size()
		initialStatModTime = stat.ModTime()
		time.Sleep(1 * time.Second)
	}
	return nil
}
