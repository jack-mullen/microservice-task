Technology stack
------------------
First step in the design of this Url to PDF service was to think of the technology
and dependencies required. After some consideration I decided to:

- Use `rabbitMq` to queue up requests to the service. This will allow the service to
cope with high traffic.
- Use `wkhtmltopdf` - a free tool for html to PDF conversion. This was a time
saving decision as building a html to pdf converter from scratch would take
longer than the time allowed for this test.

Docker Design
----------------
With docker I found myself trying to decide between 2 options
- build app on host, deploy binary to container
- deploy code to container and build binary in container

Both options have pros and cons. However, for this particular exercise I
choose to go with option 2 as it makes the application easier to test and
does not require the tester to have `go` installed on their machine; making
docker the only dependency.

Also at this stage, I had the option between using a standard DockerFile or
using docker-compose to connect containers.

I decided to use `docker-compose` as it mirrors closely what the final solution
will look like in production. i.e there would be a dedicated rabbitMq cluster
or containers providing queue/message brokering service

High Availability
------------------
As I built this application, I made it an objective to make sure the service
is able to handle a lot of requests and added a caching layer to it. Which
means that the service does not have to convert a PDF for the same request
over and over again. Instead, it checks to see if it has done a similar job
before.

The caching layer could be improved to ignore small changes to markups that
cause it to think it is a different page even though it is same content.
For example using this service against a URL like http://google.com will
cause the service to generate a fresh pdf each time because each page request
generates a unique set of javascript code. Even though the page reads and
looks the same, the caching layer does not know any better as it just a simple
implementation that hashes the content of the web page using md5.

A much robust approach would be to strip out any javascript, css and
comments before md5ing. A fully working version of this cannot be completed
within the time allowed for this test but should be considered for future work.

Code
---------
For the queuing system I applied the provider pattern to allow for
other queue providers apart from rabbit to be used. The queue package was
designed with reuse-ability in mind. It can be used in any microservice
that requires a queuing system. The `main` package was split into separate files
just to organize the code better and for readability

I also decided to adopt a fail quick approach with this microservice -
this means that we set a maximum timeout of 30 secs and we always ack
messages from the rabbit queue whether conversion was successful or not.
This is because I think this service is not mission-critical and retrying
failed jobs is expected in this kind of application

Config
---------
I had the option between using a config file or pass config values to
the app through environment variable. For simplicity, and because of the
small size of the configurable items I decided to go with using
environment variables.


Serving PDF real time
----------------------
To reliably serve PDF under 30 seconds. When a request hits the service, it
creates a job and queues it up in rabbit and sits in a loop for 30 seconds
checking if the job is complete by calling `ioutil.Readfile()` and
checking for errors. Because `wkhtmltopdf` creates pdf progressively, it
is possible for `ioutil.Readfile()` to read a partial PDF which the browser
cannot render. To solve this problem, I changed the way the app uses `wkhtmltopdf`
to generate PDF. It was changed to create a temp `.pdfx` file and renaming it
to a `.pdf`. This way we achieve an atomic write, which fixed the corrupt PDF
issue





