# Advent of code 2019

Some quick and dirty code to solve the [Advent of Code Challenges 2019](https://adventofcode.com/).

Thought I would try doing it in Go to learn some of that too.


## Running Tests


In the project root:
```bash
unset GOROOT
export GOPATH=$(pwd)

go test challenge_<number>
```


## Running Challenge Code


There are scripts in the root directory to run each challenge.

e.g. `./run_challenge_three.sh`


## Code Structure / Design Principles


* KISS
* DRY
* modularisation
* encapsulation
* Did I mention trying to keep it simple?
* ALl challenge and helper code TDD'd
* Most invocation (main) code manually tested.

That said, this code was written for fun to solve programming challenges, so is not production standard. There are missing tests, unaddressed techincal debt, and slow algorithms in here. It was just for me to have fun tbh!


## Notes


* Challenge 3 runs very slowly. Would be good to profile it and understand why when I have some time.