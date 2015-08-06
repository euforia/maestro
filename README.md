Maestro
=======

There are 2 types of performers:

#### Local
This option can be either `local` or `localhost`.  It specifies that the composition should be played on the local machine.

#### Remote
This option can be a single host or a comma delimited list of hosts.


##### user
The user options specifies which use the composition will run as.

##### composition
A list of commands to run.

##### parallel
This is a boolean value and specifies whether the symphony should be run in parallel if multiple hosts are provided:

    foo1.bar.org,foo2.bar.org:
        parallel: true
        composition:
        - date
        - hostname

In the above example, all performers (i.e. foo1.bar.org and foo2.bar.org) will run in parallel.