# 
☁  config-repo [master] brew install siege                                                                            [master|]
==> Downloading https://homebrew.bintray.com/bottles-portable/portable-ruby-2.3.3.leopard_64.bottle.1.tar.gz
Already downloaded: /Users/fqc/Library/Caches/Homebrew/portable-ruby-2.3.3.leopard_64.bottle.1.tar.gz
==> Pouring portable-ruby-2.3.3.leopard_64.bottle.1.tar.gz
Updating Homebrew...
==> Auto-updated Homebrew!
Updated 2 taps (homebrew/core, pivotal/tap).
==> Updated Formulae
azure-cli                 filebeat                  packetbeat                snakemake                 wireshark
dovecot                   kibana                    pivotal/tap/tcserver      swiftlint
elasticsearch             metricbeat                s6                        syncthing
faad2                     node-build                saltstack                 unbound

==> Installing dependencies for siege: openssl
==> Installing siege dependency: openssl
==> Downloading https://homebrew.bintray.com/bottles/openssl-1.0.2l.sierra.bottle.tar.gz
######################################################################## 100.0%
==> Pouring openssl-1.0.2l.sierra.bottle.tar.gz
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
🍺  /usr/local/Cellar/openssl/1.0.2l: 1,709 files, 12.2MB
==> Installing siege
==> Downloading https://homebrew.bintray.com/bottles/siege-4.0.4.sierra.bottle.tar.gz
######################################################################## 100.0%
==> Pouring siege-4.0.4.sierra.bottle.tar.gz
==> Caveats
macOS has only 16K ports available that won't be released until socket
TIME_WAIT is passed. The default timeout for TIME_WAIT is 15 seconds.
Consider reducing in case of available port bottleneck.

You can check whether this is a problem with netstat:

    # sysctl net.inet.tcp.msl
    net.inet.tcp.msl: 15000

    # sudo sysctl -w net.inet.tcp.msl=1000
    net.inet.tcp.msl: 15000 -> 1000

Run siege.config to create the ~/.siegerc config file.
==> Summary
🍺  /usr/local/Cellar/siege/4.0.4: 16 files, 291.7KB

