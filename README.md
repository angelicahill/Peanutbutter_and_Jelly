<h1>Peanutbutter and Jelly: Practicing Concurrency and Channels</h1>

**What is this program and what is its purpose?**

A program I wrote to Practice using Channels and running them concurrently using go routines. 

**How would you use this?**

This is a practice excercise, not a user-facing tool.

**Why did I personally write this program?**

To practice using Channels and running them concurrently using go routines. 

**How does it work?** 

I created two string variables, as well as a channel beofre writing two functions (getPB and getJelly) that sleep for 2 seconds and then send the string "Jelly" and the string "Peanut Butter" to the channel I created. I then wrote two go routines calling the two funtions I wrote and used the time package to record how long it took for both functions to return their values and therefore enable the terminal to print the string: 

```
I got some Jelly and Peanut Butter to make myself a scrummy sandwhich and it only took me 2.000504095s time to get it!

```

**How to run program?**

If you would like to run this program yourself in the terminal all you have to do is: 
- Copy/paste my code into your text editor of choice, save it as a Go file within a directory on your desktop. 
- Open your terminal and navigate to this directory before running “go run main.go” 
- The terminal should then display the following:

```
I got some Jelly and Peanut Butter to make myself a scrummy sandwhich and it only took me 2.000504095s time to get it!

```