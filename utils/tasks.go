package tasks

import (
	"fmt"
	"http"
	"io/ioutil"
    "log"
	"os"
)

func GetTasks() {
	response, err := http.Get("https://api.fikas.io/api/v1/tasks")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}
