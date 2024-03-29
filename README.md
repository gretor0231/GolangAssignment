# Golang Tinder

Design an HTTP server for the Tinder matching system. The HTTP server must support the
following three APIs:

1. AddSinglePersonAndMatch : Add a new user to the matching system and find any
possible matches for the new user.

2. RemoveSinglePerson : Remove a user from the matching system so that the user
cannot be matched anymore.

3. QuerySinglePeople : Find the most N possible matched single people, where N is a
request parameter.

Structured project layout:

![Screenshot from 2024-03-08 09-36-05.png](photos%2FScreenshot%20from%202024-03-08%2009-36-05.png)


how to use:

1.start build docker image in local machine

sudo make up_build
sudo make start


2.stop the docker image

sudo make down
sudo make stop


time complexity of APIs:
\
1.AddSinglePersonAndMatch: O(n)

2.RemoveSinglePerson: O(n)

3.QuerySinglePeople:O(n)
