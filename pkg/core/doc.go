/*
Package core implements the common data models to execute matching algorithms

DATA MODEL

Most matching algorithm rely on the following shared concept:

- A "member" is an individual element that is to be matched with another member.
- Each member typically has a "preference list" containing an ordered list of
other members it prefers to be matched with.
- A mapping of all members to their preference lists is called a "preference
table"

A preference table can be represented as follows:

	A => [B, D, F, C, E]
	B => [D, E, F, A, C]
	C => [D, E, F, A, B]
	D => [F, C, A, E, B]
	E => [F, C, D, B, A]
	F => [A, B, D, C, E]

Most algorithms implement an iterative process to reduce each member's preference
list by eliminating less preferred options.

The resulting mapping is the mathematically "stable" solution, where no two
members would prefer each other more than their existing matches.

	A => F
	B => E
	C => D
	D => C
	E => B
	F => A
*/
package core
