/*
   Copyright 2015 Brian McCallister

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/brianm/mdns"
)

var helpText = `Usage: nickname NAME IP

Publishes NAME.local pointing to IP.

So to create a record for something.local. pointing to 127.3.4.5,
which would look like "something.local. 5 IN A 127.3.4.5", you
would run the command as:

  $ nickname something 127.3.4.5
`
var help = false

func init() {
	flag.BoolVar(&help, "h", false, "Show help")
	flag.Parse()
}

func main() {
	if help || len(os.Args) != 3 {
		fmt.Println(helpText)
		os.Exit(1)
	}

	name := os.Args[1]
	ip := os.Args[2]

	record := fmt.Sprintf("%s.local. 5 IN A %s", name, ip)
	err := mdns.Publish(record)
	if err != nil {
		fmt.Printf("Unable to publish record '%s': %s\n", record, err)
		os.Exit(1)
	}

	select {}
}
