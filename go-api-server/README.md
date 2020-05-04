# GO API Server Stack

This stack provide a simple API server developped in GO language. It provides source code tempalte inclding the defintion of a http server, basic url (/, health, liveness) handlers. 

The content of the stack is :
```
$ ls - al
total 16
drwxr-xr-x  6 jtarte  staff  192 21 Oct 00:14 .
drwxr-xr-x  3 jtarte  staff   96 21 Oct 00:14 ..
-rw-r--r--  1 jtarte  staff  621 21 Oct 00:14 README.md
drwxr-xr-x  7 jtarte  staff  224 21 Oct 00:14 image
-rw-r--r--  1 jtarte  staff  297 21 Oct 00:14 stack.yaml
drwxr-xr-x  3 jtarte  staff   96 21 Oct 00:14 templates
```

Acting as an application developer, you can now use the CLI to initialize a project using this stack:

```bash
$ mkdir ~/test
$ cd ~/test
$ appsody init dev-local/gp-api-server
$ appsody run
```
