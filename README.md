# Rubiks Cube Last Step

This program is to help solve the last step of the rubiks cube. Or rather, help investigate some shuffles on that last step.

The last step is when you have solved the cube upto the last remaining cornercases which are in the right place of the cube however these are turned or swapped like on the following image example.

## Examples last step

In the last step you might find either two or four of the corners twisted. Below are two examples.
<img src="/images/twistedcornercopy.png" width="400" height="350"/>

In this example with blue as the top-facecolor, all four corners are 'at their right spot' only need to be flipped properly.

<img src="/images/twotwistedcorners.png" width="400" height="350"/>

Also, in this example, now with green as the top-facecolor, only two corner cubies are still twisted.

In both cases: all cubes are in the right place, but only the last remaining corner cubes are still a bit twisted.

There is an algorithm to solve these final corners "re-orient corners" which can be found here:
https://how-to-solve-a-rubix-cube.com/last-step/

The first step in the above link describes how to position the corners, and then the last step is to re-orient those corners; it's this last step of re-orienting that pretty much garbles up the whole cube up in the process which is kind of 'magical'. Also it is somewhat confusing because almost the entire cube gets mixed up and it's easy to make a mistake.

## Alternative solution?
The goal of this small program is to see if a combination of an earlier used algorithm which just positions three cubes in the right places would be usuable for re-orientation of the corners, and eventually solving the cube.

### What algorithms 

The program supports four different cube algorithms, and has some code to assign the colors to the faces of the corner cubies after each of these four different steps. These are the four steps supported:

1. LEFT_TO_FRONT: move the cube whereby the left side of cube becomes the front. No shuffles are done for this move. It is just re-orienting the entire cube 'counter clockwise' where top stays top.
2. LEFTSTART_SWAP: an algorithm whereby three corners switch places, viewn from top as per image:
<img src="/images/topviewleftsideswap.png" />
shuffle algorithm as per known shortcuts: L' U R U' L U R' U'

3. RIGHT_TO_FRONT: move the cube whereby the right side of the cube becomes the front. No shuffles are done for this move. It is just re-orienting the entire cube 'clockwise' where top stays top.
4. RIGHTSTART_SWAP: an algorithm whereby three corners switch places, viewn from top as per image:
<img src="/images/topviewrightsideswap.png" />
shuffle algorithm as per known shortcuts: R U' L' U R' U' L U

The above steps are pretty easy to learn. When you do several of these in random order it can be seen that only the orientation of the four corner cubies in the top layer are shuffled around a bit. The colorfaces also shuffle and twist.

# Problem description

My problem with the Rubiks cube is that to get from 'solved' state to 'two corners twisted' or other way round always is kind
of a lucky shot after doing several of the above four shuffles, in some random order. And if finally the 'two corners twisted state' is reached, getting the cube back in 'solved' state with the above four shuffles gets rather tedious. Usually I ended up suddenly having four
corners shuffled instead of two. Or, two other corners were twisted while the original corners where placed perfectly well. So I just needed to get the solution for going from the two corner shuffled to the solved state by brute-forcing
the above four shuffles. Technically, in the program this is done using iterative deepening, which just means to try out all possible permutations of shuffles first for 1 shuffle, than 2, than 3 and so forth until a solution is found.

# Result?
Yes!

The program found an alternative solution with the limit of above steps for the two-corner case with following six steps:
1. LEFTSTART_SWAP
2. RIGHT_TO_FRONT
3. LEFTSTART_SWAP
4. LEFTSTART_SWAP
5. LEFT_TO_FRONT
6. RIGHTSTART_SWAP

You have to position the cube with the two mis-aligned corners on the right and then execute the above steps.

Of course, it also goes the other way round: You can also use the above algorithm to change a solved cube into the 'two corners twisted' state, and ask a friend if she can solve it from that state.