package main

const body string = `<!--The purpose of the issue template is to help us figure out a technical approach to a task, so we understand what the problem is and how to solve it. It should lead us to a solution that not only meets the story's requirements, but also maintains high standards of operability and internal quality. Finally, the approach should be documented with a list of checkmarks, so we can understand our progress and assign parts of the task as we see fit.-->

## What?

<!--🚨_What are you trying to achieve?_ What's the big picture at the end of the rainbow? Does this really help our user do their job? Now is a good time to decide. -->

JIRA Link: https://saltpayco.atlassian.net/browse/SAG-xxx

### Example mapping

<!--
You should add checkboxes to each example to ensure it's implemented and tested
For example:

- [ ] Happy path
Given we have X information
When Y happens
Then we do Z

- [ ] Unhappy path - invalid data
Given we don't have X
When Y happens
Then we give back a 400 bad request
-->

### Rules

<!--
You should add checkboxes to each rule to ensure they're implemented and tested
For example:

- [ ] If a store exists in the db, and a subscription to Terminal service has not been triggered then the TerminalSubscriptionSignupStatus should be NOT_INITIATED.
- [ ] When a subscription to Terminal service has been triggered but still processing then the TerminalSubscriptionSignupStatus should be INITIATED.
- [ ] When a subscription to Terminal service has failed then the TerminalSubscriptionSignupStatus should be FAILED.
-->

### Questions/uncertainties

<!--
These should also have checkboxes. If the answer to the question is found and it requires us to implement something, then
add a checkbox somewhere to this KO so that we don't forget about it. If the question or answer is being descoped, we should
link to the new JIRA ticket.

For example:

- [ ] What rules govern how we generate TIDs?
-->

## How?

<!--
- How are you going to proceed?
- Prefer a top-down approach, starting with an **acceptance test**; unless this is an iteration on existing functionality, which could just be driven with unit tests.
- What's the smallest change _that can be integrated_, and that gets you towards your final goal? Think about "steel thread" through the system. Could you maybe hard-code some values first to move forward?
- What's the next change after that? And after that?
- Can you see a path made of small steps that get us to our destination?
🚨🚨🚨 **List out your steps that you figured out below, and then delete the leading questions, so we have a clear plan documented!** 🚨🚨🚨 -->

## Are there any code-review tasks related to this work that could speed up the ticket?

-   [ ] Check the backlog for any code review tasks that might help speed up the ticket. ([Find them here](https://saltpayco.atlassian.net/jira/software/c/projects/SAG/issues/?jql=%22Epic%20Link%22%3DSAG-234%20AND%20%28%22Status%22%3D%22TO%20DO%22%20OR%20%22Status%22%3D%22In%20Progress%22%29%20ORDER%20BY%20rank))

## When?

Give an ETA of when you think it will be _done_ (including DOD finished)

### How will you know it works, and it works well?

<!--🚨
- What alerts do you need ?
- What metrics will you need to look at to know it's working well?
- What logging do you need for when things go wrong?-->

### Have you integrated with a new HTTP API as part of this story?

<!--🚨You should be asking yourself how will you know that our interactions with another API are working correctly. Luckily our HTTP clients already give you some standard metrics for free, such as status codes, response times and logging. In the bright future when we have a dashboard you should be thinking of visualising these metrics so we have a sense of how well we're interacting with the new system.-->

### What knowledge do you need to be able to work on this ticket?

<!--🚨e.g:
- basic sql queries
- openapi specs
- http server middleware-->

## Definition of Done (DOD)

<!--Please delete items that are not relevant to your story AT KICK-OFF 🚨🚨🚨-->

### Testing

-   [ ] Unit
-   [ ] Integration
-   [ ] Acceptance
-   [ ] e2e

### Operability

-   [ ] Metrics
-   [ ] Logging
-   [ ] Alerts

### Documentation

-   [ ] Do you need an ADR? (Think about the context in which the decision is/was made)
-   [ ] Have you added a new external dependency? If yes, add to the context diagram
-   [ ] Have you changed the way the system is built or requires developers to do extra setup? Document in the README.md
-   [ ] Have you considered how this change may affect on-call? If so, update the runbook.

### Internal quality

-   [ ] Are you happy with the code?
-   [ ] Is someone else happy with the code?

### Delivering value

-   [ ] Is it deployed to dev?
-   [ ] Is it deployed to prod?

### Retrospective notes

-   What went well?
-   What didn't go well?
-   Did you finish the story when you think you would?

### Lessons Learned

-   What have we learned about how our system behaves?
-   Have we made any changes/improvements to the design of the sytem?
-   If yes, should this be captured as an ADR ?
`
