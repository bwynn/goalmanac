package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func buildString(s []string) string {
	return strings.Join(s, "")
}

// buildFilename
//
// the filename will be generated using a combination of
// timestamps and locations (postal code),
// allowing for queries to existing files prior to migration/seeding
// to a db instance
//
// e.g.: 95032_1535589784
//		 <postal_code>_<unix_timestamp>.json
func BuildFilename(unixT int64, zipcode string) string {
	// format unix to string
	timestamp := strconv.FormatInt(unixT, 10)
	filename := []string{
		"data/",
		zipcode,
		"_",
		timestamp,
		".json",
	}
	// join the slice of strings w/ util method
	formattedFilename := buildString(filename)
	return formattedFilename
}

// returns unix timestamp and prints time of initialization
func TimeFormat() int64 {
	// get the current time
	t := time.Now()
	// format and set timestamp to unix time
	// before writing to file
	setToUnixT := t.Unix()
	formatted := t.Format("Monday, Aug 1 15:04:05 -0800 PST 2006")
	// log time for posterity
	fmt.Printf("request sent at: %s", formatted) // ahoy!
	return setToUnixT
}
