# Advent of Code - GoLang template

A repository to automate some tasks related to Advent of Code.

## Get your session token

First thing you need is your session token from adventofcode.com. This is used to fetch your personalized input data from each puzzle. It can be found under the network tab in your dev tools. [Check this post for assistance](https://github.com/wimglenn/advent-of-code-wim/issues/1).

Once you have your token, copy .env.example, rename it to .env, and include your token in the AOC_SESSION spot. In your .env file you can also adjust what year of Advent of Code you want to work with. Defaults to 2022.

## Usage

The binary for fetching your input data and setting up a template is located in `/bin`.

- Select the year you want to work with in your .env file.
- Run for example `./bin/aocgo --day=5` to start working with day 5 for the selected year.
