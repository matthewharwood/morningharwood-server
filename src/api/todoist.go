package api
import (
	"net/http"
	"strings"
	"fmt"
	"encoding/json"
	"log"
)

type Items struct {
	Name string
	Phone string
}
type TodoistData struct {
	Items []struct {
		Content string `json:"content"`
		MetaData interface{} `json:"meta_data"`
		UserID int `json:"user_id"`
		TaskID int `json:"task_id"`
		ProjectID int `json:"project_id"`
		CompletedDate string `json:"completed_date"`
		ID int `json:"id"`
	} `json:"items"`
	Projects struct {
					Num193815246 struct {
												 Name string `json:"name"`
												 Color int `json:"color"`
												 Collapsed int `json:"collapsed"`
												 InboxProject bool `json:"inbox_project"`
												 ParentID interface{} `json:"parent_id"`
												 ItemOrder int `json:"item_order"`
												 Indent int `json:"indent"`
												 ID int `json:"id"`
												 IsDeleted int `json:"is_deleted"`
												 IsArchived int `json:"is_archived"`
											 } `json:"193815246"`
				} `json:"projects"`
}
func Todoist() TodoistData {
	body := strings.NewReader(`token=12387c33bc508f9d5d14b78ea9270778f8d38c5d`)
	req, err := http.NewRequest("POST", "https://todoist.com/API/v7/completed/get_all", body)
	if err != nil {
		fmt.Println("errz", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("err", err)
	}

	defer resp.Body.Close()

	var record TodoistData

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record
}