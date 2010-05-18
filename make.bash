# Copyright (c) 2010, Jonathan Wills (runningwild@gmail.com)
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
gotgo -o btree.go -package-name=main btree.got int

6g -o bench.6 bench.go btree.go
6l -o bench bench.6

6g -o test.6 test.go btree.go
6l -o test test.6

g++ -o c_bench -O3 ctree.cpp
