# rubikscubelaststep

This program is to help solve the last step of the rubiks cube.

The last step is when you have solved the cube upto the last remaining cornercases which are in the right place of the cube however these are turned like on the following image example.

![twistedcornercopy](/images/twistedcornercopy.png "Twisted Corner")

In other words: all cubes are in the right place, but only the last remaining corner cubes are still a bit twisted.

There is a somewhat simple algorithm to solve these final corners which can be found here:
https://how-to-solve-a-rubix-cube.com/last-step/

The first step described is to position the corners, and then the last step is to re-orient those corners; it's this last step that pretty much garbles up the whole cube up in the process and I thought there should be a combination of an earlier used algorithm which just positions the cubes in the right places again, that is, by re-using the repositioning algorithm a few times.
