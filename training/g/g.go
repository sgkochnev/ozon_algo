package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

//go:embed tests_G/14
var input []byte

//go:embed tests_G/14.a
var output []byte

func main() {
	in := bufio.NewReader(bytes.NewBuffer(input))
	// in := bufio.NewReader(os.Stdin)
	// s := time.Now()
	friendsMap := readData(in)

	listOfRecommendations := createListsOfRecomendedFriends(friendsMap)

	print(listOfRecommendations)
	// fmt.Println(time.Since(s))

}

func readData(in io.Reader) []map[int]struct{} {
	var n, m int
	fmt.Fscanln(in, &n, &m)

	var x, y int
	friendsMap := make([]map[int]struct{}, n+1)
	for i := 0; i < n+1; i++ {
		friendsMap[i] = make(map[int]struct{})
	}

	for i := 0; i < m; i++ {
		fmt.Fscanln(in, &x, &y)
		friendsMap[x][y] = struct{}{}
		friendsMap[y][x] = struct{}{}
	}

	return friendsMap
}

func createListsOfRecomendedFriends(friendsMap []map[int]struct{}) [][]int {
	listsOfRecommendedFriends := make([][]int, len(friendsMap))

	for self, friends := range friendsMap {
		mFriends := mapOfRecomendedFrinds(self, friendsMap, friends)
		listRecommendedFriends := listOfRecommendedFriends(mFriends, maxInMap(mFriends))
		listsOfRecommendedFriends[self] = listRecommendedFriends
	}
	return listsOfRecommendedFriends
}

func mapOfRecomendedFrinds(self int, friendsMap []map[int]struct{}, listFriends map[int]struct{}) map[int]int {
	recommend := make(map[int]int)
	for friend := range listFriends {
		for recommendedFriend := range friendsMap[friend] {
			_, ok := friendsMap[self][recommendedFriend]
			if !ok && self != recommendedFriend {
				recommend[recommendedFriend]++
			}
		}
	}
	return recommend
}

func maxInMap(m map[int]int) int {
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func listOfRecommendedFriends(m map[int]int, max int) []int {
	list := []int{}
	for k, v := range m {
		if v == max {
			list = append(list, k)
		}
	}
	sort.Ints(list)
	return list
}

func printListOfRecommendedFriends(l []int) {
	if len(l) == 0 {
		fmt.Println(0)
		return
	}

	str := strings.Builder{}
	for _, v := range l {
		// fmt.Printf("%d ", v)
		str.WriteString(strconv.Itoa(v))
		str.WriteByte(' ')
	}
	fmt.Println(str.String())
	// fmt.Println()
}

func print(listOfRecommendations [][]int) {
	for _, list := range listOfRecommendations[1:] {
		printListOfRecommendedFriends(list)
	}
}
