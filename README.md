# GraphTheory - Regular Expressions
A go project for my third year project in Graph Theory.
## Description
A program written in Go Programming language that can build a non-deterministic finite automaton from a regular expression and can use the NFA to check if the regular expression matches any given string of text.

## Prerequisites

I used github for my project so it would not be lost and be easy for other people to access.

### Push to Github:

In order to submit my project changes to github from my github folder i used the following commands:
git add .
git commit -m "Initial commit"
git push

### Download from github:
For you to download my project you must clone my repository link from the command promp:

git clone "example.github/project"

### You can then run my code using:
nfa.go is my main project
"go run nfa.go" 
You do this from the command once you are in the github folder that contains the project.

My Program will run in the terminal.

## Coding Syle

In my project i used Go programming Language.

Give an example

I have one go file that runs the program.
I use println to show the testing of different outcomes.
```
	fmt.Println(pomatch("ab.c*|", "ccc")) //return true
	fmt.Println(pomatch("ab.c*|", "def")) //return false
```

I have a user input for you to enter the string and regular expression against your string.
for example: 
String: ccc
Regular expression: ab.c*|
Output: True

## Resources i used for my project.

https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e

https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b

https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b

https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line

https://tutorialedge.net/golang/reading-console-input-golang/
