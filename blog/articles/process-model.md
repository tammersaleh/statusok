# Process Model

Web application process types
can be defined in a manifest named `Procfile`
such as this one for a Rails app:

```
web: bin/rails server
webpacker: ./bin/webpack-dev-server
workerdefault: QUEUES=default bundle exec rake jobs:work
workermailer: QUEUES=mailers bundle exec rake jobs:work
```

In development,
[Foreman](http://blog.daviddollar.org/2011/05/06/introducing-foreman.html)
can be used as a process manager.

Start Foreman processes:

```
foreman start
```

In this mode,
Foreman interleaves output streams,
responds to crashed processes,
and handles user-initiated restarts and shutdowns.
