# fake-github-webhook

# What is?
This is an application that simulates github webhooks. You choose a file containing JSON Payload Data and then sends it to a given target.

# In which scenarios I can use this thing?
- You can use it to test Continous Integration tools on localhost. This is good because you can test your tool and put it to run on internet only when it's ready.
- Also, you can use it to avoid useless commits on a test repository.

# How I use this thing?
    # start your CI tool listening on 127.0.0.1:8080
    # then test it triggering a payload to 127.0.0.1:8080
    $ fake-github-webhook -host 127.0.0.1:8080 -file payload.json

