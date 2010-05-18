// Copyright (c) 2010, Jonathan Wills (runningwild@gmail.com)
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import "rand"
import "fmt"

func CheckInsertsAndDeletes(d []int) bool {
  t := NewTree(func(a,b int) bool { return a < b })
  count := 0
  for _,v := range d {
    t.Insert(v)
    count++
    if t.Len() != count { return false }
    if !t.fsck() { return false }

    // Repeat the insert, nothing should changes
    t.Insert(v)
    if t.Len() != count { return false }
    if !t.fsck() { return false }
  }

  for _,v := range d {
    t.Remove(v)
    count--
    if t.Len() != count { return false }
    if !t.fsck() { return false }

    // Repeat the delete, nothing should change
    t.Remove(v)
    if t.Len() != count { return false }
    if !t.fsck() { return false }
  }

  return true
}

func CheckQueries(full []int) bool {
  t := NewTree(func(a,b int) bool { return a < b })
  divs := 4
  d := make([][]int, divs)
  for i := range d {
    d[i] = full[(len(full)*i) / divs : (len(full)*(i+1)) / divs]
  }

  for i := range d {
    for _,v := range d[i] {
      t.Insert(v)
    }
    for j := range d {
      for _,v := range d[j] {
        if t.Contains(v) != (j <= i) { return false}
      }
    }
  }

  for i := range d {
    for _,v := range d[i] {
      t.Remove(v)
    }
    for j := range d {
      for _,v := range d[j] {
        if t.Contains(v) != (j > i) { return false}
      }
    }
  }


  return true
}

func main() {
  n := 10000
  vrand := rand.Perm(n)
  vinc := make([]int, n)
  for i := range vinc { vinc[i] = i }
  vdec := make([]int, n)
  for i := range vdec { vdec[i] = i }

  if !CheckInsertsAndDeletes(vrand) {
    fmt.Printf("Failed inserts and deletes when done in random order.\n")
  }
  if !CheckQueries(vrand) {
    fmt.Printf("Failed queries when done in random order.\n")
  }

  if !CheckInsertsAndDeletes(vinc) {
    fmt.Printf("Failed inserts and deletes when done in increasing order.\n")
  }
  if !CheckQueries(vinc) {
    fmt.Printf("Failed queries when done in increasing order.\n")
  }

  if !CheckInsertsAndDeletes(vdec) {
    fmt.Printf("Failed inserts and deletes when done in decreasing order.\n")
  }
  if !CheckQueries(vdec) {
    fmt.Printf("Failed queries when done in decreasing order.\n")
  }
}
