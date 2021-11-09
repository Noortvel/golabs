package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lab1()
	lab2()
	lab3()
	lab4()
	lab5()
	lab6()
}

func lab1() {
	println("lab1")
	xs := []float32{1, 2, 3, 4, 5}
	result := Sum(xs)
	println(result)
	println()
}

//1.1
func Sum(xs []float32) float32 {
	var sum float32 = 0
	for i := 0; i < len(xs); i++ {
		sum += xs[i]
	}

	return sum
}

//2.1
func lab2() {
	println("lab2")
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		matched, err := regexp.MatchString(`.*\.(jpg|png|bmp)`, f.Name())
		if err != nil {
			log.Fatal((err))
		}
		if matched {
			fmt.Println(f.Name())
		}
	}
	println()
}

//3.3

type digital_address struct {
	Id             int
	FirstName      string
	LastName       string
	MiddleName     string
	Emails         []string
	Social_url_map map[string]string
	Age            int
}

// var database []digital_address

func appendEmail(val digital_address, email string) {
	val.Emails = append(val.Emails, email)
}

func removeEmail(val digital_address, email string) bool {
	r_i := -1
	for i, val := range val.Emails {
		if val == email {
			r_i = i
		}
	}

	if r_i == -1 {
		return false
	}

	val.Emails = append(val.Emails[:r_i], val.Emails[r_i+1:]...)

	return true
}

func setSocial(val digital_address, name string, url string) {
	val.Social_url_map[name] = url
}

func findByLastName(database []digital_address, lastName string) (digital_address, error) {
	for _, val := range database {
		if val.LastName == lastName {
			return val, nil
		}
	}

	return digital_address{}, errors.New("Not found")
}

func findBySocial(database []digital_address, social string, url string) (digital_address, error) {
	for _, val := range database {
		_url, prs := val.Social_url_map[social]
		if prs && _url == url {
			return val, nil
		}
	}

	return digital_address{}, errors.New("Not found")
}

func saveDb(database []digital_address) (bool, error) {
	bt, err := json.Marshal(database)
	if err != nil {
		return false, err
	}

	if err := os.WriteFile("database.json", bt, 0644); err != nil {
		return false, err
	}

	return true, nil
}

func loadDb() ([]digital_address, error) {
	var out []digital_address

	dat, err := os.ReadFile("database.json")
	if err != nil {
		return out, err
	}

	if err := json.Unmarshal(dat, &out); err != nil {
		return out, err
	}

	return out, nil
}

func lab3() {
	println("lab3")
	currDb, err1 := loadDb()
	println("current db: ")
	if err1 == nil {
		fmt.Printf("%+v\n", currDb)
	}

	database := []digital_address{
		{
			FirstName: "Bob",
			LastName:  "Mercer",
			Age:       12,
			Emails: []string{
				"somemail@mail.com",
				"my@mail.com"},
			Social_url_map: map[string]string{
				"telegram":  "@example",
				"instagram": "@insta",
			},
		},
	}

	println("new db: ")
	fmt.Printf("%+v\n", database)

	_, err := saveDb(database)
	if err != nil {
		log.Fatal(err.Error())
	}

	println()
}

// 4.1
type Sqare struct {
	a float32
	b float32
}

func (sqr Sqare) Area() float32 {
	return sqr.a * sqr.b
}

func lab4() {
	println("lab4")
	sqr := Sqare{
		a: 1,
		b: 2,
	}

	println(sqr.Area())

	println()
}

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

// 5.1
func minEl(arr []int, minValCh chan int) {
	min := MaxInt
	for _, val := range arr {
		if val < min {
			min = val
		}
	}

	minValCh <- min
}

func maxEl(arr []int, maxValCh chan int) {
	max := MinInt
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	maxValCh <- max
}

func lab5() {
	println("lab5")

	arr := []int{1, 2, 3, 4, 5, 6}
	minValCh := make(chan int)
	maxValCh := make(chan int)

	var minVal int
	var maxVal int

	go minEl(arr, minValCh)
	go maxEl(arr, maxValCh)

	for i := 0; i < 2; i++ {
		select {
		case minVal = <-minValCh:
		case maxVal = <-maxValCh:
		}
	}

	println(fmt.Sprintf("minVal %d", minVal))
	println(fmt.Sprintf("maxVal %d", maxVal))

	println()
}

//6.1

func findMaxSizeFile(file_path string, maxSize int64, maxFileName string) (int64, string) {
	files, err := ioutil.ReadDir(file_path)
	if err != nil {
		return maxSize, maxFileName
	}

	for _, f := range files {
		fSize := f.Size()
		name := f.Name()
		if fSize > maxSize {
			maxSize = fSize
			maxFileName = name
		}

		maxSize, maxFileName = findMaxSizeFile(file_path+f.Name()+"/", maxSize, maxFileName)
	}

	return maxSize, maxFileName
}

func lab6() {
	println("lab6")
	file_path := "./"
	maxSize, maxFileName := findMaxSizeFile(file_path, -1, "./")
	println(strconv.FormatInt(maxSize, 10) + " " + maxFileName)

	println()
}
