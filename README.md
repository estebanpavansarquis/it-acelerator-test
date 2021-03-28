# it-acelerator-test

[![codecov](https://codecov.io/gh/estebanpavansarquis/it-acelerator-test/branch/develop/graph/badge.svg?token=31CSEE9ZRH)](https://codecov.io/gh/estebanpavansarquis/it-acelerator-test)

Solution for the Back-End IT-Acelerator Test 2019/2020 & Bootcamp-IT 2021

## Fun with Anagrams

>Two strings are anagrams if they are permutations of each other.

Given an array of strings, remove each string that is an anagram of an earliear string, then return the remaning array in sorted order.

 **Function description:**
Develop the funWithAnagrams() function. It must return a list of strings in alphabetical order, ascending.

> Benchmarking results of the different algorithms implemented -> [Link](https://docs.google.com/spreadsheets/d/188v-5ZI-q7DxTML3ozlEDL9dRH8XwsQeCA1-SJ3pnIY/edit?usp=sharing)

## Authentication Tokens

In a standard transaction system, users are authenticated using authentication tokens. Each token has an expiry, denoted by expiryLimit. Each token expires automatically after expiryLimit minutes pass after it was reset. When reset, the expiry time is resets to expiryLimit minutes from the time of reset.
- A token can be reset any number of times.
- Once a token expires it can no longer be accessed, and a reset issued to it will be ignored.

**Note:**
- Get is represented by 0 and Reset by 1.
- The current time, in time end, is the maximun time of all commands.

Given a sequence of commands with no pre-existing tokens, and these 
commands sorted by ascending by their T parameter, find the number of tokens that are active just after all commands have been executed.

**Function description:**
Develop the numberOfTokens() function. It must return an integer that denotes the number of tokens that exist at the end of the command stream.

## Most repeted letter
Given a string as input, create a function that prints a message with the most repeted letter and the number of times repeated.
- The input can contain spaces and special characters.
- The result must ignore casing, spaces and special charaters.
