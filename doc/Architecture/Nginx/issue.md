# Issue

## Close Wait Issue
<img src="close_wait.png">

The passive side doesn't send FIN signal to client

- no close()
- busy something, exceed the timeout
- backlog size
    - accept backlog size is too large to allow accept