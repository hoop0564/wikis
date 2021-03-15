package main

import (
	"fmt"
	"strconv"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

// fetcher是填充后的fakeFetcher。
// 真正的爬虫是用http去抓取网页分析content中的URL。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// Crawl使用fetcher从某个URL开始递归的爬取页面，直到找到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// todo：并行的抓取URL。
	// todo: 不重复抓取页面。
	// 下面并没有实现上面两种情况。
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s, %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

type User struct {
	Id int
	Name string
}

type Profile struct {
	Users []User
}

func initUser()  {
	var p Profile
	p.Users = append(p.Users, User{1, ""})
	p.Users = append(p.Users, User{2, ""})
	updateUser(&p)
	fmt.Printf("111111>>> %+v\n", p)

	var users []User
	users = append(users, User{100, ""})
	users = append(users, User{101, ""})
	update2(users)
	fmt.Printf("222222>>> %+v\n", users)
}

func update2(users []User)  {
	size := len(users)
	for i:=0; i<size; i++ {
		users[i].Name = "gzc"+strconv.Itoa(i+100)
	}
}

func updateUser(p *Profile) {
	size := len(p.Users)
	for i:=0; i<size; i++ {
		p.Users[i].Name = "gzc"+strconv.Itoa(i+1)
	}
}

func main() {
	initUser()
	return
	Crawl("https://golang.org/", 4, fetcher)
}
