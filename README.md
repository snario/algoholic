# Algoholic #

## Introduction ##

Algoholic is a collection of many algorithms in many different languages.

The idea is for me to have a project in which to practice the design and implementation of
algorithms, perhaps others will find it useful.

No claim is made as to the efficiency of the algorithms, and emphasis is put on making the
implementation of the algorithms as clear as possible so there is heavy commenting
throughout. Production implementations of these algorithms should probably include fewer
comments :-)

## Instructions ##

To run code testing the algorithm implementations, execute the following in the appropriate
directory:-

* Go - `go test`
* Coffeescript - Ensure you've run `npm install` and have mocha installed via `npm install -g mocha`. Then run `./runCoffeeTests.sh` in the project root directory.

To benchmark go code run `go test -bench .` in the appropriate directory.

__Note__ in `sort/`, `go test -bench .` will take a long while to complete the insertion sort
benchmark. Use `go test -benchtime 0.3s -bench .` or similar (varying the time to taste) to
reduce this time. Interestingly reducing the time significantly can highlight cases where
insertion sort beats alternatives :-)

## Sorting ##

### Insertion Sort ###

* [go][isort_go]
* [coffeescript][isort_cs]

### Merge Sort ###

* [go][msort_go]

## Searching ##

### Binary Search ###

* [go][bsearch_go]

### Linear Search ###

* [go][lsearch_go]

[isort_go]:/sort/isort.go
[isort_cs]:/sort/isort.coffee

[msort_go]:/sort/msort.go

[bsearch_go]:/search/bsearch.go

[lsearch_go]:/search/lsearch.go
