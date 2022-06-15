# Kickoff


### Design notes

The library provides a runner for a series of handlers.

The handlers are of two types:

- initial, takes no state, generates starting state (title, etc)
- continuation, takes state, appends to state
- all handlers delegate to either internally, or out to the next handler sequentially in the parent
- when we're at the top level, we send the the final state as a create issue request to GitHub

Yes, CJ is very correct in saying this is a _fine_ fit for the Chain of Responsibility pattern.

### What I don't know how to solve yet

#### Questions, Questioner, Multiline

Say I want the Questioner to understand that sometimes a question requires multi-line answers. Long form stuff.
This is easy enough when the Questioner has a web implementation (no real change), but harder with a CLI. I would
need to change the way the Questioner handles the input from the io.Reader - finish building on the first
empty line, for instance. But to tell the Questioner that the basic method would be to flag the Question - `IsMultiline`
or similar. But this feels like I'm falling into a bad polymorphism trap. But any other way of doing it - 
getting the Question to handle input...

I guess at the moment the `Question` is dumb. The `Questioner` is gobbling up the logic of how to ask questions.
Would it be better to have the `Questioner` just feed the `Question` lines (or similar), and then ask it if it's happy at the end?
Then the `Question` itself is responsible for converting that simple input into an `Answer`.

This still doesn't solve the multiline problem.