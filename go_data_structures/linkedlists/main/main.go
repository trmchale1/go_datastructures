package main

import l "github.com/go_data_structures/linkedlists"

func main () {
    linkedList := l.New()
	linkedList.Append(1).Append(2).Append(3).PrintAll()
    linkedList.DeleteWithValue(2).PrintAll()
}
