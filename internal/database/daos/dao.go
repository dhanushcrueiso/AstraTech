package daos

import (
	"AstraTech/internal/database/db"
	"AstraTech/internal/database/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

const directory = "/tmp/astra/files/"

func ProcessData(filename string) {
	fmt.Println("checking data", filename)

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println("Error reading directory:", err)
		return
	}

	if len(files) == 0 {
		// No files to process, wait for a while before checking again
		time.Sleep(5 * time.Second)
		return
	}

	fmt.Println("", len(files))
	var wg sync.WaitGroup
	for _, file := range files {

		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			fmt.Println("checking 3 file name", filename)
			fileData, err := ioutil.ReadFile(directory + filename)
			if err != nil {
				log.Println("Error reading file:", err)
				return
			}
			// fmt.Println("filedata", fileData)
			fmt.Println("checking 3")
			var data models.DataDao
			if err := json.Unmarshal(fileData, &data); err != nil {
				log.Println("Error decoding JSON:", err)
				return
			}
			// fmt.Println("checking 1", data.Data)
			data.CreatedAt = time.Now()
			// Process the data asynchronously
			fmt.Println("checking 2", data.CreatedAt)
			fmt.Println("checking 3", data.Id)
			// fmt.Println("checking 2", data.Data)
			if err := InsertData(data); err != nil {
				log.Println("Error inserting data into database:", err)
				return
			}
			fmt.Println("checking 2")
			if err := os.Remove(directory + filename); err != nil {
				log.Println("Error deleting file:", err)
			}
		}(file.Name())
	}
	wg.Wait()

}

func InsertData(req models.DataDao) error {
	fmt.Println("inside the daos", req)
	err := db.DB.Unscoped().Table("astra").Create(req).Error
	if err != nil {
		return err
	}
	return nil
}
