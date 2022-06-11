# Kickoff


### Design notes

The library provides a runner for a series of handlers.

The handlers are of two types:

- initial, takes no state, generates starting state (title, etc)
- continuation, takes state, appends to state
- all handlers delegate to either internally, or out to the next handler sequentially in the parent
- when we're at the top level, we send the the final state as a create issue request to GitHub

Yes, CJ is very correct in saying this is a _fine_ fit for the Chain of Responsibility pattern.