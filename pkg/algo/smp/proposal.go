package smp

/*
 * Implements the 1st (and only) phase of the Gale-Shapley (1962) algorithm.
 *
 * This calculates a stable match between two groups (alpha and beta) given their
 * individual preference lists for members of the other group
 *
 * See:
 *   https://en.wikipedia.org/wiki/Stable_marriage_problem#Solution
 *   https://www.youtube.com/watch?v=GsBf3fJFpSw
 *
 * Each unmatched member "proposes" to their top remaining preference
 * Each recipient of a proposal can take one of 3 actions -
 *
 * 1. The recipient has not received a previous proposal and immediately accepts
 *
 * 2. The recipient prefers this new proposal over an existing one.
 *    The recipient "rejects" it's initial proposl and accepts this new one
 *
 * 3. The recipient prefers the existing proposal over the new one.
 *    The recipient "rejects" the new proposal
 *
 * Note: Rejections are mutual. If `i` removes `j` from their preference list,
 *       then `j` must also remove `i` from its list
 *
 * This cycle continues until every alpha/beta has a match.
 *
 * Mathematically, every participant is guranteed a match so this algorithm
 * always converges on a solution.
 *
 *
 * === EXAMPLE
 *
 * Take the following preference lists
 *
 * alpha preferences:
 *   A => [O, M, N, L, P]
 *   B => [P, N, M, L, O]
 *   C => [M, P, L, O, N]
 *   D => [P, M, O, N, L]
 *   E => [O, L, M, N, P]
 *
 * beta preferences:
 *   L => [D, B, E, C, A]
 *   M => [B, A, D, C, E]
 *   N => [A, C, E, D, B]
 *   O => [D, A, C, B, E]
 *   P => [B, E, A, C, D]
 *
 * We can begin the proposal sequence with any unmatched member. Different starting members
 * may yield different results, however. In our example we start with 'A'.
 *
 * The sequence of events are -
 *
 *   'A' proposes to 'O'
 *   'O' accepts 'A'
 *   'B' proposes to 'P'
 *   'P' accepts 'B'
 *   'C' proposes to 'M'
 *   'M' accepts 'C'
 *   'D' proposes to 'P'
 *   'P' rejects 'D'
 *   'E' proposes to 'O'
 *   'O' rejects 'E'
 *   'D' proposes to 'M'
 *   'M' accepts 'D', rejects 'C'
 *   'E' proposes to 'L'
 *   'L' accepts 'E'
 *   'C' proposes to 'P'
 *   'P' rejects 'C'
 *   'C' proposes to 'L'
 *   'L' rejects 'C'
 *   'C' proposes to 'O'
 *   'O' rejects 'C'
 *   'C' proposes to 'N'
 *   'N' accepts 'C'
 *
 * At this point there are no alpha users left unmatched (and by definition, no
 * corresponding beta users left unmatched). All alpha members have had their
 * proposals accepted by a beta user.
 *
 * The resulting solution is
 *
 *   A => O
 *   B => P
 *   C => N
 *   D => M
 *   E => L
 *   L => E
 *   M => D
 *   N => C
 *   O => A
 *   P => B
 */

import (
	"github.com/abhchand/libmatch/pkg/core"
)

func phase1Proposal(ptA, ptB *core.PreferenceTable) {
	for true {
		unmatchedMembers := ptA.UnmatchedMembers()

		if len(unmatchedMembers) == 0 {
			break
		}

		for i := range unmatchedMembers {
			member := unmatchedMembers[i]
			topChoice := member.FirstPreference()
			simulateProposal(member, topChoice)
		}
	}
}

func simulateProposal(proposer, proposed *core.Member) {
	if !proposed.HasAcceptedProposal() {
		proposed.AcceptMutually(proposer)
	} else if proposed.WouldPreferProposalFrom(*proposer) {
		proposed.RejectMutually(proposed.CurrentProposer())
		proposed.AcceptMutually(proposer)
	} else {
		proposed.RejectMutually(proposer)
	}
}
