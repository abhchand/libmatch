/*
Package smp implements the solution to the "Stable Marriage Problem".

It implements the Gale-Shapley (1962) algorithm, which calculates a stable match
between two groups (alpha and beta) given their individual preferences for
members of the other group.

ALGORITHM

See: https://en.wikipedia.org/wiki/Stable_marriage_problem#Solution

As part of the algorithm, each unmatched member "proposes" to the top
remaining preference on their list. Each recipient of a proposal can take one of
3 actions -

1. The recipient has not received a previous proposal and immediately accepts

2. The recipient prefers this new proposal over an existing one.
The recipient "rejects" its initial proposl and accepts this new one

3. The recipient prefers the existing proposal over the new one.
The recipient "rejects" the new proposal

NOTE: Rejections are mutual. If `i` removes `j` from their preference list,
then `j` must also remove `i` from its list

This cycle continues until every member is matched.

STABILITY AND DETERMINISM

Mathematically, every participant is guranteed a match because the algorithm
always converges on a stable solution.

Notes:

1. Multiple stable matchings may exist, and this process will only return one
possible stable matching.

2. The algorithm itself prioritizes the preferences of the first specified
preference table over the second.

3. While the algorithm is deterministic, the `libmatch` implementation is not.
That is, it is not guranteed to return the same matching everytime. This is due
to the underlying storage mechanism, which uses a `map` type. Go randomizes the
access keys of a `map` and so the algorithm may produce different results on
different runs.

ALGORITHM EXAMPLE

See: https://www.youtube.com/watch?v=GsBf3fJFpSw

Take the following preference tables

	// alpha preferences
	A => [O, M, N, L, P]
	B => [P, N, M, L, O]
	C => [M, P, L, O, N]
	D => [P, M, O, N, L]
	E => [O, L, M, N, P]

	// beta preferences
	L => [D, B, E, C, A]
	M => [B, A, D, C, E]
	N => [A, C, E, D, B]
	O => [D, A, C, B, E]
	P => [B, E, A, C, D]

We can begin the proposal sequence with any unmatched member. As noted above,
different starting members may yield different results.

In our example we start with 'A'. The sequence of events are -

	'A' proposes to 'O'
	'O' accepts 'A'
	'B' proposes to 'P'
	'P' accepts 'B'
	'C' proposes to 'M'
	'M' accepts 'C'
	'D' proposes to 'P'
	'P' rejects 'D'
	'E' proposes to 'O'
	'O' rejects 'E'
	'D' proposes to 'M'
	'M' accepts 'D', rejects 'C'
	'E' proposes to 'L'
	'L' accepts 'E'
	'C' proposes to 'P'
	'P' rejects 'C'
	'C' proposes to 'L'
	'L' rejects 'C'
	'C' proposes to 'O'
	'O' rejects 'C'
	'C' proposes to 'N'
	'N' accepts 'C'

At this point there are no members of the "alpha" group left unmatched (and by
definition, no corresponding members of "beta" left unmatched). All members have
had their proposals accepted by a member of the other group.

The matching result is:

	A => O
	B => P
	C => N
	D => M
	E => L
	L => E
	M => D
	N => C
	O => A
	P => B
*/
package smp
