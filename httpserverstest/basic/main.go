package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/seferen/basic/taskstore"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

func (ts *taskServer) taskHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/task/" {

		switch r.Method {
		case http.MethodPost:
			ts.createTaskHandler(w, r)
		case http.MethodGet:
			ts.getAllTaskHandler(w, r)
		case http.MethodDelete:
			ts.deleteAllTaskHandler(w, r)
		default:
			http.Error(w, fmt.Sprintf("%v method don't allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}

	} else {
		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")

		if len(pathParts) < 2 {
			http.Error(w, "expect /task/<id> in task handler", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(pathParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			ts.getTaskHandler(w, r, id)
		case http.MethodDelete:
			ts.deleteTaskHandler(w, r, id)
		default:
			http.Error(w, fmt.Sprintf("%v method don't allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

func (ts *taskServer) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling task create at %s\n", r.URL.Path)

	type RequestTask struct {
		Text string    `json:"text"`
		Tags []string  `json:"tags"`
		Date time.Time `json:"due"`
	}

	type ResponseId struct {
		Id int `json:"id"`
	}

	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "Don't supported this mediatype", http.StatusUnsupportedMediaType)
		return
	}

	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var rt RequestTask
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ts.store.CreateTask(rt.Text, rt.Tags, rt.Date)
	js, err := json.Marshal(ResponseId{id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func (ts *taskServer) getAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling get all tasks at %s\n", r.URL.Path)

	allTask := ts.store.GetAllTasks()
	js, err := json.Marshal(allTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ts *taskServer) getTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	log.Printf("handling get task at %s\n", r.URL.Path)

	task, err := ts.store.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ts *taskServer) deleteTaskHandler(w http.ResponseWriter, r *http.Request, id int) {
	log.Printf("handling delete task at %s\n", r.URL.Path)

	if err := ts.store.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

}

func (ts *taskServer) deleteAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling delete all tasks at %s\n", r.URL.Path)
	if err := ts.store.DeleteAllTasks(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (ts *taskServer) tagHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling tasks by tag at %s\n", r.URL.Path)

	switch r.Method {
	case http.MethodGet:

		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
			http.Error(w, "expect /tag/<tag> path", http.StatusBadRequest)
			return
		}

		tag := pathParts[1]

		tasks := ts.store.GetTaskByTag(tag)
		js, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(js)

	default:
		http.Error(w, fmt.Sprintf("Unexpected method %s", r.Method), http.StatusMethodNotAllowed)
		return
	}

}

func (ts *taskServer) dueHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling tasks by due time at %s\n", r.URL.Path)

	switch r.Method {
	case http.MethodGet:

		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")

		badRquestError := func() {
			http.Error(w, fmt.Sprintf("expect /due/<year>/<month>/<day>, got %v", r.URL.Path), http.StatusBadRequest)
		}

		year, err := strconv.Atoi(pathParts[1])
		if err != nil {
			badRquestError()
			return
		}
		month, err := strconv.Atoi(pathParts[2])
		if err != nil {
			badRquestError()
			return
		}
		day, err := strconv.Atoi(pathParts[3])
		if err != nil {
			badRquestError()
			return
		}

		tasks := ts.store.GetTaskByDueDate(year, time.Month(month), day)
		js, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(js)

	default:
		http.Error(w, fmt.Sprintf("Don't expected method %s", r.Method), http.StatusMethodNotAllowed)
		return
	}

}

func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()

	mux.HandleFunc("/task/", server.taskHandler)
	mux.HandleFunc("/tag/", server.tagHandler)
	mux.HandleFunc("/due/", server.dueHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
