# advent-of-code-2021

**I am doing this to get more comfortable with Go.** It's not perfect, and I compare my answers to others (though primarily with coworkers and friends and Reddit if I think my coding was subpar) _after_ I've completed the puzzle.

## Running

To try a day from the root of this repository, run

```
go run <dayXX> <1|2>
```

Each folder is its own module with a main.go that takes either a '1' or '2' and runs the corresponding part.

## Notes

### In spite of trying to learn Go, I got pretty familiar with bash

For starters, I created a [script that got private leaderboard times](https://github.com/ndd7xv/aoc-times) and initially had shell scripts to run modules from a parent directory and copy/paste module templates. Then I got to the point where I learned about workspaces and started moving things over with

```bash
# Removed the run script I use to have, since workspaces allow me to do `go run dayXX`
# Just needed to change inputs in each of my module files
$ for i in $(git grep -l inputs/ | xargs -i echo {} | cut -d'/' -f1 | uniq); do sed -i "s|inputs/input1.txt|$i/inputs/input1.txt|g" $i/lib/*; done

# Removed the replace script I use to have after I set up the templating for all the following days
$ for day in 0{8..9} {10..25}; do cp -r dayXX day${day}; done
$ for day in 0{8..9} {10..25}; do ./replace dayXX day${day} day${day}/; done
```
### Notable Progression
 - Early days were just learning syntax and how to use external libraries (and realizing Go doesn't have built in sets/stacks)
 - Day 7 is when I created a tree and is really the first time implemented a sort of class like thing in Go
 - Day 11 was interesting; learned structs generally copy when being assigned to another variable or looping with the `range` keyword; to avoid this, we can use `&` and index into the array with a traditional loop, respectively. Also learned that closures are a thing