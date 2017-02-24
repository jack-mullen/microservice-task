Sipsynergy  Task
==================

Prerequisites
-------------

- Fork this repository as a starting point
- Use docker container to run your project in
- Include Dockerfile configuration files, so we can test your code

Task
----

You will be creating a small API microservice : "URL to PDF service".

Create Go Lang project.


The end result should be allowing your application to accept URL via HTTP call and return PDF 
(return 'application/pdf' content type and PDF containing requested url contents)

e.g.:
 
`GET http://localhost:8000/api/to-pdf?url=http://www.google.com`


The main focus of this task is to demonstrate architectural decisions you undertake.

You may complete the task in any way you see fit (it does not have to be finished). 

Conclusion
----------

We would suggest that you use git commits in a way that allows the reviewer to understand the thought process taken to achieve each given goal.

In your repository *PROGRESS* file, you should note your experience with the task, and provide some critique to the codebase as a whole.

