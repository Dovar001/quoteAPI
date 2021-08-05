package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	//"sync/atomic"
	"time"

	"github.com/gorilla/mux"
)

//Структура цитаты
// Quote stuct
type Quote struct {
	ID      string  `json:"id,omitempty"`
	Content string  `json:"content,omitempty"`
	Author  *Author `json:"author,omitempty"`
}

//Автор цитаты
// Author is the person who said so
type Author struct {
	Name string `json:"name,omitempty"`
}

//Функция для отображении цитаты с помощью id
//Function for displaying a quote
func quoteById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range quotes {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Quote{})
}

//Функция для отображении всех цитат
//Function for displaying all quotes
func allQuotes(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(quotes)
}

//Функция для удалении цитаты с помощью id
//Function for remove quote by id
func deleteQuoteById(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range quotes {
		if item.ID == params["id"] {
			id,err:=strconv.Atoi(item.ID)
		     id-=1
			fmt.Println("id equal to",id)
		
			if err != nil {
				fmt.Println("wrong id",err)
			}
			
			remove(id,quotes)
		
		}
	}

	json.NewEncoder(writer).Encode(&Quote{})
}

//Функция для рандомной цитаты
//Function for sudden quote
func suddenQuote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(request)
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(randInt(1,4)) 

    s:=(randInt(1, 4))
	
	id:=strconv.Itoa(s)
	
	for _, item := range quotes {
		if item.ID == id{
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Quote{})
}

//Функция для удалении из массива
//Function for removing from array
func remove(i int, list []Quote)[]Quote{
	if i < len(list)-1{
		list = append(list[:i],list[i+1:]... )
	}else{
		log.Print("wrong id")
		return list
	}
	return list
}

//Функция для генерации рандомного числа
//Function for generating random number
func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}


//worker сделан не до конца, не хватило времени)))


/*
func Worker(d time.Duration, f func()) {
	var reEntranceFlag int64
	for range time.Tick(d) {
	 go func() {
	  if atomic.CompareAndSwapInt64(&reEntranceFlag, 0, 1) {
	   defer atomic.StoreInt64(&reEntranceFlag, 0)
	  } else {
	   log.Println("Previous worker in process now")
	   return
	  }
	  f()
	 }()
	}
   }
   
   Worker(5,deleteExpiredQuote)


   func deleteExpiredQuote(){
	   
   }
*/


//Массив цитат
//array of quotes
var quotes []Quote

func main() {
	router := mux.NewRouter()

    fmt.Println("Server started")

	quotes = append(quotes, Quote{ID: "1", Content: "Don't cry because it's over, smile because it happened.", Author: &Author{Name: "Dr.Dovar"}})
	quotes = append(quotes, Quote{ID: "2", Content: "So many books, so little time.", Author: &Author{Name: "Dr.Daler"}})
	quotes = append(quotes, Quote{ID: "3", Content: "You know you're in love when you can't fall asleep because reality is finally better than your dreams.", Author: &Author{Name: "Dr.Safar"}})

	router.HandleFunc("/quotes",allQuotes).Methods("GET")
	router.HandleFunc("/quote/{id}", quoteById).Methods("GET")
	router.HandleFunc("/sudden", suddenQuote).Methods("GET")
	router.HandleFunc("/delete/{id}",deleteQuoteById).Methods("DELETE")

	
	log.Fatal(http.ListenAndServe(":8080", router))
	
}