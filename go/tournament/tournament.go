package tournament

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type teamResult struct {
	name   string
	played int
	won    int
	drawn  int
	lost   int
	points int
}

func (tr teamResult) String() string {
	return fmt.Sprintf(
		"% -31s| % 2d | % 2d | % 2d | % 2d | % 2d\n",
		tr.name, tr.played, tr.won, tr.drawn, tr.lost, tr.points,
	)
}

// Tally calculates scores from the input reader
// and writes a results table to the input writer.
func Tally(r io.Reader, w io.Writer) error {
	var err error

	var buf []byte
	if buf, err = ioutil.ReadAll(r); nil != err {
		return err
	}
	matches := strings.Split(string(buf), "\n")

	var resultMap map[string]teamResult
	if resultMap, err = scoreMatches(matches); nil != err {
		return err
	}

	if err = writeResults(w, resultMap); nil != err {
		return err
	}

	return nil
}

// scoreMatches takes in a slice of lines representing matches
// and returns a map of the results by team name.
func scoreMatches(matches []string) (map[string]teamResult, error) {
	resultMap := make(map[string]teamResult)

	for _, match := range matches {
		if strings.TrimSpace(match) == "" || strings.HasPrefix(match, "#") {
			continue
		}

		parts := strings.Split(match, ";")

		if len(parts) != 3 {
			return resultMap, fmt.Errorf(
				"match %s has %d parts, should have 3", match, len(parts),
			)
		}

		t1, t2, matchRes := parts[0], parts[1], parts[2]
		t1res, t2res := resultMap[t1], resultMap[t2]
		t1res.played++
		t2res.played++

		switch matchRes {
		case "win":
			t1res.won++
			t2res.lost++
		case "loss":
			t1res.lost++
			t2res.won++
		case "draw":
			t1res.drawn++
			t2res.drawn++
		default:
			return resultMap, fmt.Errorf("unknown match result %s", matchRes)
		}

		resultMap[t1], resultMap[t2] = t1res, t2res
	}

	return resultMap, nil
}

// writeResults writes the results as a table to the
// io.Writer passed in.
func writeResults(w io.Writer, resultMap map[string]teamResult) error {
	var err error

	var teamResults []teamResult
	for tn, tr := range resultMap {
		tr.name = tn
		tr.points = (tr.won * 3) + tr.drawn
		teamResults = append(teamResults, tr)
	}

	sort.Slice(teamResults, func(i, j int) bool {
		if teamResults[i].points == teamResults[j].points {
			return teamResults[i].name < teamResults[j].name
		}
		return teamResults[i].points > teamResults[j].points
	})

	for i, tr := range teamResults {
		if i == 0 {
			header := fmt.Sprintf("% -31s| MP |  W |  D |  L |  P\n", "Team")
			if _, err = w.Write([]byte(header)); nil != err {
				return err
			}
		}

		if _, err = w.Write([]byte(tr.String())); nil != err {
			return err
		}
	}

	return nil
}
