/* Map is another data type like structure and dictionary(python). A key is an objects that use to
retrive a value later. */

package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	/* Create A Map */
	countryCapitalMap = make(map[string]string)

	/* insert Key and values */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* print map using key */
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is", capital)
	}
	/* Test if the country present in the map or not */

	capital, ok := countryCapitalMap["India"]

	/* if ok is true, entry is presnt */
	if ok {
		fmt.Println(" ________________________________")
		fmt.Println("| Capital of India is", capital, " |")
		fmt.Println(" --------------------------------")
	} else {
		fmt.Println("___________________________________")
		fmt.Println("Capital of India is not in the Map")
	}

}
