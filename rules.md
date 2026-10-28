# Quiz Arena Rules (v1)

Any change to this file gets a line in the Change Log of QUESTS.md.

## Match setup
- A match has 2 or more players. Every rule below works for any number of players.
- Every player starts with 100 health.
- Every match loads exactly 20 questions (a named constant in the code). 20 is the maximum length of a match, and questions are played in order. All players get each question at the same time.

## Answer time
- Answer time is measured in milliseconds (ms), as an integer.
- The clock starts when the answer options appear, which is after the 3 second countdown.
- The question ends at 10,000 ms (10 seconds). A valid answer arrives between 0 and 9,999 ms.
- An answer at 10,000 ms or later never counts. It is treated as no answer.

## Damage
- A correct answer costs between 0 and 10 health.
- Damage grows linearly with answer time and rounds UP: `damage = ceil(answerTimeMs * 10 / 10000)`. In integer math that is `(answerTimeMs * 10 + 9999) / 10000`.
  - 0 ms costs 0, 1 to 1,000 ms costs 1, 5,000 ms costs 5, 9,001 to 9,999 ms costs 10.
- A wrong answer costs a flat 15 health, whatever the answer time.
- No answer (including a late one) counts as a wrong answer and costs 15 health.
- Every player takes damage every round. Players are resolved at the same time, not one after another.

## Health and knockout
- Health is a plain integer and is not clamped. It can go below 0.
- A player is knocked out when health is 0 or below.
- A knocked-out player takes no further damage.

## End of match
- The match ends when one player or fewer is still standing, or when the last question has been played.
- Players who die in different rounds are placed in reverse order of death: the player who died last places best.
- Players who die in the same round are ordered by remaining health. The least negative health places better.
- Players with exactly equal health in that case share a placement.
- If the questions run out with more than one player alive, they are ordered by remaining health, and equal health shares a placement.
