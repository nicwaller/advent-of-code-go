--- Part Two ---

Next, you need to find the largest basins so you know what areas are most important to avoid.

A basin is all locations that eventually flow downward to a single low point. Therefore, every low point has a basin, although some basins are very small. Locations of height 9 do not count as being in any basin, and all other locations will always be part of exactly one basin.

The size of a basin is the number of locations within the basin, including the low point. The example above has four basins.

The top-left basin, size 3:

```
21
3
```

The top-right basin, size 9:

```
43210
 4 21
    2
```

The middle basin, size 14:

```
  878     
 85678 
87678
 8
```

The bottom-right basin, size 9:

```
2199943210
3987894921
       8
      678 
     65678
```

Find the three largest basins and multiply their sizes together. In the above example, this is 9 * 14 * 9 = 1134.

What do you get if you multiply together the sizes of the three largest basins?
