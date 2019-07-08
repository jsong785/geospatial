# Geospatial

Quick unification of various programs I've written under one repo (monorepo).
This was mainly written after I learned some Go.

The tiling code has incorrectly behavior on a split (quad tree); from memory, it is not moving data from the current level to the 4 children below.
I also do not 100% know the details of a quad tree, so it may just be wrong.

I wrote GoTest since at the time, I was familiar with unit testing with Google Test in C++, and wanted to "mimic" that behavior for some reason.

