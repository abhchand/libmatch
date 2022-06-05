/*
Package smp implements the solution to the "Stable Roommates Problem".

It implements the Irving (1985) algorithm, which calculates a stable match
within an even-sized set given each member's preference of the other members.

ALGORITHM

See: https://en.wikipedia.org/wiki/Stable_roommates_problem

The algorithm runs in 3 phases.


Phase 1

In this round each member who has not had their proposal accepted "proposes"
to their top remaining preference

Each recipient of a proposal can take one of 3 actions -

1. The recipient has not received a previous proposal and immediately accepts

2. The recipient prefers this new proposal over an existing one.
The recipient "rejects" it's initial proposl and accepts this new one

3. The recipient prefers the existing proposal over the new one.
The recipient "rejects" the new proposal

In the above situations, every rejection is mutual - if `i` removes `j` from
its preference list, then `j` must also remove `i` from its list

This cycle continues until one of two things happens:

A. Every member has had their proposal accepted (Move on to Phase 2)

B. At least one member has exhausted their preference list (No solution exists)


Phase 2

In this phase, each member that has accepted a proposal will remove those they
prefer less than their current proposer.

At the end of one iteration, one of 3 states are possible -

A. At least one member has exhausted their preference list (No solution exists)

B. Some members have multiple preferences remaining (Proceed to Phase 3)

C. All members have one preference remaining (Solution has been found, no need
to run Phase 3)


Phase 3

In this last phase we attempt to find any preference cycles and reject them.

This is done by building a pair of members (Xi, Yi) as follows

	* Xi is the first member with at least 2 preferences
	* Yi is null
	* Yi+1 is the 2nd preference of Xi
	* Xi+1 is the last preference of Yi+1

Continue calculating pairs (Xi, Yi) until Xi repeats values. At this point a
cycle has been found.

We then mutually reject every pair (Xi, Yi)

After this one of 3 possiblities exists -

A. At least one member has exhausted their preference list (No solution exists)

B. Some members have multiple preferences remaining (Repeat the above process to
eliminate further cycles)

C. All members have one preference remaining (Solution has been found)

STABILITY AND DETERMINISM

A stable solution is NOT guranteed. If no further possible proposal exists in
Phase 1 or the reduced preference table results in an empty list in Phase 3,
then no stable solution will exist.

However, if a solution does exist it is guranteed to be deterministic. That is,
a solution will always converge on one and only one optimial mapping between
members.

ALGORITHM EXAMPLE

See: https://www.youtube.com/watch?v=9Lo7TFAkohE

Take the following preference table

	A => [B, D, F, C, E]
	B => [D, E, F, A, C]
	C => [D, E, F, A, B]
	D => [F, C, A, E, B]
	E => [F, C, D, B, A]
	F => [A, B, D, C, E]

We always start with the first unmatched user. Initially this is "A".

The sequence of events are -

	'A' proposes to 'B'
	'B' accepts 'A'
	'B' proposes to 'D'
	'D' accepts 'B'
	'C' proposes to 'D'
	'D' accepts 'C', rejects 'B'
	'B' proposes to 'E'
	'E' accepts 'B'
	'D' proposes to 'F'
	'F' accepts 'D'
	'E' proposes to 'F'
	'F' rejects
	'E' proposes to 'C'
	'C' accepts 'E'
	'F' proposes to 'A'
	'A' accepts 'F'

The result of Phase 1 is shown below. A "-" indicates a proposal made and a "+"
indicates a proposal accepted. Rejected members are removed.

	A => [-B, D, +F, C, E]
	B => [-E, F, +A, C]
	C => [-D, +E, F, A,B]
	D => [-F, +C, A, E]
	E => [-C, D, +B, A]
	F => [-A, B, +D, C]

Phase 2 rejects occur as follows. Note that all rejections are
mutual - if `i` removes `j` from its preference list, then `j` must also
remove `i` from its list

	'A' accepted by 'B'. 'B' rejecting members less preferred than 'A': ["C"]
	'B' accepted by 'E'. 'E' rejecting members less preferred than 'B': ["A"]
	'C' accepted by 'D'. 'D' rejecting members less preferred than 'C': ["A", "E"]
	'D' accepted by 'F'. 'F' rejecting members less preferred than 'D': ["C"]
	'E' accepted by 'C'. 'C' rejecting members less preferred than 'E': ["A"]
	'F' accepted by 'A'. 'A' rejecting members less preferred than 'F': []

The output of this phase is a further reduced table is as follows

	A => [B, F]
	B => [E, F, A]
	C => [D, E]
	D => [F, C]
	E => [C, B],
	F => [A, B, D]

Since at least one member has multiple preferences remaining, we proceed to
Phase 3.

Phase 3 starts with "A" since it is the first member with at least two
preferences.

Build (Xi, Yi) pairs as follows

	i      1   2   3   4
	-----+---+---+---+----
	x:   | A | D | E | A
	y:   | - | F | C | B

Where -

	'F' is the 2nd preference of 'A'
	'D' is the last preference of 'F'
	'C' is the 2nd preference of 'D'
	etc...

As soon as we see "A" again, we stop since we have found a cycle.

Now we will mutually reject the following pairs, as definied by the inner
pairings

	(D, F)
	(E, C)
	(A, B)

At this point, no preference list is exhausted and some have more than one
preference remaining. We need to find and eliminate more preference cycles.

	i      1   2
	-----+---+---
	x:   | B | B
	y:   | - | F

Now we will mutually reject

	(F, B)

This gives us the stable solution below, since each roommate has exactly one
preference remaining

	A => F
	B => E
	C => D
	D => C
	E => B
	F => A
*/
package srp
