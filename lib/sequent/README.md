# sequent

primarily intended for AoC, but could have other applications

includes some experimental syntatical "sugar" to create seqents from other sources

streaming iterators, for processing things that might not fit in memory

builds on Go native iter.Seq iterators

adds some filter chaining stuff

and some map/reduce stuff

the syntax is imperfect, especially `.Map()` because of how Go generics are limited

avoids channels because they are too slow

avoids my own non-standard iterator implementation

the name "sequent"? well come on, it's hard to come up with new names for an iterator package.

## Not Included

- `.Fork()` forking an iterator is not safe because we cannot guarantee that the child iterators will be consumed at the same rate, meaning we need unbounded buffering.
- `.Partition()` has the same problem.
