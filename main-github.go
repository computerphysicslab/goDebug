/*
When debugging your code in goLang you may need a pretty print function not just for int or strings,
but capable to render any kind of data structure. Here it is, simple and useful. Enjoy it!
*/

package main

import "github.com/computerphysicslab/goDebug"

func main() {
	link1 := "https://hn.algolia.com/?dateRange=pastWeek&page=0&prefix=false&query=golang&sort=byDate&type=story"
	link2 := "https://www.serverlab.ca/tutorials/linux/administration-linux/how-to-configure-proxy-on-ubuntu-18-04/"
	link3 := "https://www.expressvpn.com/what-is-my-ip?gclid=CjwKCAjw34n5BRA9EiwA2u9k3-hTxM73D-v-MsEX21RKlAucrnrXkJ5aWR-xduY4V-_i-gcJEdjKCBoCgsMQAvD_BwE"

	type ALink struct {
		Url     string
		Count   int
		Visited bool
	}

	var LPool []ALink

	LPool = append(LPool, ALink{Url: link1, Count: 0, Visited: false})
	LPool = append(LPool, ALink{Url: link2, Count: 0, Visited: false})
	LPool = append(LPool, ALink{Url: link3, Count: 0, Visited: false})

	// goDebug.Print("LPool", LPool)
	goDebug.Print(LPool)
}
