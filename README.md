# gload
Golang Load Testing.  Closely named after my favorite Producer.

# Two Binaries
* ***Harness*** The main client orchestrating all of the remote agents, sending commands and processing results
* ***Agent***  The remote process that listens on an available machine port & registers itself with a discovery service.  Just uses gRPC to invoke functionality.  In this case the Agent is the service implementation and the Harness is the client.


# Ideal Workflow

1) Developer creates any new commands as necessary

2) Developer builds new remote agents

3) Remote agents are distributed

3a) The main harness can fork child agent processes?

4) Remote agents register themselves to the cluster?  The harness needs to know about each of them?

4a) Just use etcd?  Just use: https://discovery.etcd.io (allow any etcd cluster to be used)

4b) Remote Agents must be pre-built with the correct discovery token & url.  They will crash if they are unable to register

5) The harness discovers all the remote agents via etcd

6) The harness connects with all the remote agents & establishes a gRPC connection to send commands

7) The harness can:

* Execute commands from source (static harness build -- only runs 1 kind of test)

* Execute in a REPL

* Capture response times 

** Request -> Response

** All Commands should have Errors
