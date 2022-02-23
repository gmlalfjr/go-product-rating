package main

import "server/app"

func main() {
	r := app.AppLancher()
	err := r.Run()
	if err != nil {
		return
	}

}
