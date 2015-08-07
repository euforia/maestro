Maestro
=======

Configuration
-------------

##### performer
This specifies which host the composition should be played on.  This option can be either `local` or a remote hostname.  If `localhost` is specified rather than `local` then an ssh connection will be issued to `localhost`.

##### user
The user options specifies which use the composition will run as.

##### composition
A list of commands to run.

##### parallel
This is a boolean value and specifies whether the symphony should be run in parallel if multiple hosts are provided:

        - performer:
          - foo1.bar.org
          - foo2.bar.org
          parallel: true
          composition:
          - date
          - hostname

In the above example, all performers (i.e. foo1.bar.org and foo2.bar.org) will run in parallel.