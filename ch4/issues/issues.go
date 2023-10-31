// Issues prints a table of GitHub issues matching the search terms.

package main

import (
	"fmt"
	"gobook/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	thisMonth := time.Now().AddDate(0, -1, 0)
	thisYear := time.Now().AddDate(-1, 0, 0)

	young := make([]int, 0)
	middle := make([]int, 0)
	old := make([]int, 0)

	for i, item := range result.Items {
		if item.CreatedAt.Compare(thisMonth) >= 0 {
			young = append(young, i)
		} else if item.CreatedAt.Compare(thisYear) >= 0 {
			middle = append(middle, i)
		} else {
			old = append(old, i)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	fmt.Println("Less than a month old :")
	for _, k := range young {
		fmt.Printf("#%-5d %9.9s %.55s\n", result.Items[k].Number,
			result.Items[k].User.Login, result.Items[k].Title)
	}

	fmt.Println("Less than a year old :")
	for _, k := range middle {
		fmt.Printf("#%-5d %9.9s %.55s\n", result.Items[k].Number,
			result.Items[k].User.Login, result.Items[k].Title)
	}

	fmt.Println("More than a year old :")
	for _, k := range old {
		fmt.Printf("#%-5d %9.9s %.55s\n", result.Items[k].Number,
			result.Items[k].User.Login, result.Items[k].Title)
	}
}
