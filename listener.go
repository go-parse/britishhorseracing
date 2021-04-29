package main

import (
	"flag"
)

func listener() {

	flags.db = flag.Bool("db", false, "database configuration")
	flags.proxy = flag.Bool("proxy", false, "proxy configuration")

	flag.Parse()
}
