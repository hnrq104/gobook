package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(term []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(term, " "))
	// fmt.Println(q)
	/*
		Escaping here is of utmost importance, as it will make the query possible
		we've sent "repo:golang/go is:open json decoder" in term
		Escaping to "repo%3Agolang%2Fgo+is%3Aopen+json+decoder"
	*/
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
