# install

```
☁  spring-cloud-microservice-examples [master] ⚡ brew install rabbitmq             [master|✚3
==> Installing dependencies for rabbitmq: openssl, libtiff, erlang
==> Installing rabbitmq dependency: openssl
==> Downloading https://homebrew.bintray.com/bottles/openssl-1.0.2m.sierra.bottle.tar.gz
######################################################################## 100.0%
==> Pouring openssl-1.0.2m.sierra.bottle.tar.gz
==> Caveats
A CA file has been bootstrapped using certificates from the SystemRoots
keychain. To add additional certificates (e.g. the certificates added in
the System keychain), place .pem files in
  /usr/local/etc/openssl/certs

and run
  /usr/local/opt/openssl/bin/c_rehash

This formula is keg-only, which means it was not symlinked into /usr/local,
because Apple has deprecated use of OpenSSL in favor of its own TLS and crypto libraries.

If you need to have this software first in your PATH run:
  echo 'export PATH="/usr/local/opt/openssl/bin:$PATH"' >> ~/.zshrc

For compilers to find this software you may need to set:
    LDFLAGS:  -L/usr/local/opt/openssl/lib
    CPPFLAGS: -I/usr/local/opt/openssl/include
For pkg-config to find this software you may need to set:
    PKG_CONFIG_PATH: /usr/local/opt/openssl/lib/pkgconfig

==> Summary
🍺  /usr/local/Cellar/openssl/1.0.2m: 1,792 files, 12.3MB
==> Installing rabbitmq dependency: libtiff
==> Downloading https://homebrew.bintray.com/bottles/libtiff-4.0.8_5.sierra.bottle.tar.gz
######################################################################## 100.0%
==> Pouring libtiff-4.0.8_5.sierra.bottle.tar.gz
🍺  /usr/local/Cellar/libtiff/4.0.8_5: 245 files, 3.4MB
==> Installing rabbitmq dependency: erlang
==> Downloading https://homebrew.bintray.com/bottles/erlang-20.1.5.sierra.bottle.tar.gz
######################################################################## 100.0%
==> Pouring erlang-20.1.5.sierra.bottle.tar.gz
==> Caveats
Man pages can be found in:
  /usr/local/opt/erlang/lib/erlang/man

Access them with `erl -man`, or add this directory to MANPATH.
==> Summary
🍺  /usr/local/Cellar/erlang/20.1.5: 7,112 files, 276.5MB
==> Installing rabbitmq
==> Downloading https://dl.bintray.com/rabbitmq/binaries/rabbitmq-server-generic-unix-3.6.14.t
######################################################################## 100.0%
==> /usr/bin/unzip -qq -j /usr/local/Cellar/rabbitmq/3.6.14/plugins/rabbitmq_management-3.6.14
==> Caveats
Management Plugin enabled by default at http://localhost:15672

Bash completion has been installed to:
  /usr/local/etc/bash_completion.d

To have launchd start rabbitmq now and restart at login:
  brew services start rabbitmq
Or, if you don't want/need a background service you can just run:
  rabbitmq-server
==> Summary
🍺  /usr/local/Cellar/rabbitmq/3.6.14: 209 files, 5.5MB, built in 18 seconds
```

