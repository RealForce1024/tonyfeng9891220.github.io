## 安装环境
cd workshop_day1
sh prepare.sh

```sh
☁  workshop_day1  sh prepare.sh
==> box: Box file was not detected as metadata. Adding it directly...
==> box: Adding box 'leanmsd1' (v0) for provider:
    box: Unpacking necessary files from: file:///Users/fqc/git-work/workshop_day1/leanms_d1.box
The box you're attempting to add already exists. Remove it before
adding it again or add it with the `--force` flag.

Name: leanmsd1
Provider: virtualbox
Version: 0
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Clearing any previously set forwarded ports...
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Configuring and enabling network interfaces...
==> default: Mounting shared folders...
    default: /vagrant => /Users/fqc/git-work/workshop_day1
==> default: Machine already provisioned. Run `vagrant provision` or use the `--provision`
==> default: flag to force provisioning. Provisioners marked to run always will still run.
Welcome to Ubuntu 16.04.2 LTS (GNU/Linux 4.4.0-51-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

0 packages can be updated.
0 updates are security updates.


Last login: Mon Mar 20 09:08:12 2017 from 10.0.2.2
```
 进入ci目录看到一个docker-compose.yml
vagrant@vagrant:~$ ls
cd  ci  hello  temp
vagrant@vagrant:~$ cd ci
vagrant@vagrant:~/ci$ ls
docker-compose.yml

vim docker-compose.yml
```
version: '2'

services:
  gitlab2:
      image: gitlab/gitlab-ce:latest
      ports:
          - "443:443"
          - "80:80"
      networks:
          - hello
      volumes:
          - ~/temp/gitlab/config:/etc/gitlab
          - ~/temp/gitlab/logs:/var/log/gitlab
          - ~/temp/gitlab/data:/var/opt/gitlab

  jenkins:
      image: leanms/jenkins:0.1
      ports:
          - "8080:8080"
          - "50000:50000"
      networks:
          - hello
      volumes:
          - ~/temp/jenkins:/var/jenkins_home
          - /var/run/docker.sock:/var/run/docker.sock
          - ~/temp/gradle/:/root/.gradle

networks:
  hello:
    driver: bridge
```

vagrant@vagrant:~/ci$ sudo docker-compose up

```
Starting ci_jenkins_1
Starting ci_gitlab2_1
Attaching to ci_gitlab2_1, ci_jenkins_1
gitlab2_1  | Thank you for using GitLab Docker Image!
gitlab2_1  | Current version: gitlab-ce=8.17.3-ce.0
gitlab2_1  |
gitlab2_1  | Configure GitLab for your system by editing /etc/gitlab/gitlab.rb file
gitlab2_1  | And restart this container to reload settings.
gitlab2_1  | To do it use docker exec:
gitlab2_1  |
gitlab2_1  |   docker exec -it gitlab vim /etc/gitlab/gitlab.rb
gitlab2_1  |   docker restart gitlab
gitlab2_1  |
gitlab2_1  | For a comprehensive list of configuration options please see the Omnibus GitLab readme
gitlab2_1  | https://gitlab.com/gitlab-org/omnibus-gitlab/blob/master/README.md
gitlab2_1  |
gitlab2_1  | If this container fails to start due to permission problems try to fix it by executing:
gitlab2_1  |
gitlab2_1  |   docker exec -it gitlab update-permissions
gitlab2_1  |   docker restart gitlab
gitlab2_1  |
jenkins_1  | Running from: /usr/share/jenkins/jenkins.war
jenkins_1  | webroot: EnvVars.masterEnvVars.get("JENKINS_HOME")
jenkins_1  | Apr 09, 2017 4:01:48 AM Main deleteWinstoneTempContents
jenkins_1  | WARNING: Failed to delete the temporary Winstone file /tmp/winstone/jenkins.war
jenkins_1  | Apr 09, 2017 4:01:48 AM org.eclipse.jetty.util.log.JavaUtilLog info
jenkins_1  | INFO: Logging initialized @1409ms
jenkins_1  | Apr 09, 2017 4:01:48 AM winstone.Logger logInternal
jenkins_1  | INFO: Beginning extraction from war file
jenkins_1  | Apr 09, 2017 4:01:48 AM org.eclipse.jetty.util.log.JavaUtilLog warn
jenkins_1  | WARNING: Empty contextPath
jenkins_1  | Apr 09, 2017 4:01:48 AM org.eclipse.jetty.util.log.JavaUtilLog info
jenkins_1  | INFO: jetty-9.2.z-SNAPSHOT
gitlab2_1  | Preparing services...
gitlab2_1  | Starting services...
gitlab2_1  | Configuring GitLab package...
gitlab2_1  | /opt/gitlab/embedded/bin/runsvdir-start: line 24: ulimit: pending signals: cannot modify limit: Operation not permitted
gitlab2_1  | /opt/gitlab/embedded/bin/runsvdir-start: line 37: /proc/sys/fs/file-max: Read-only file system
gitlab2_1  | Configuring GitLab...
jenkins_1  | Apr 09, 2017 4:01:50 AM org.eclipse.jetty.util.log.JavaUtilLog info
jenkins_1  | INFO: NO JSP Support for /, did not find org.eclipse.jetty.jsp.JettyJspServlet
jenkins_1  | Jenkins home directory: /var/jenkins_home found at: EnvVars.masterEnvVars.get("JENKINS_HOME")
jenkins_1  | Apr 09, 2017 4:01:53 AM org.eclipse.jetty.util.log.JavaUtilLog info
jenkins_1  | INFO: Started w.@76a4ebf2{/,file:/var/jenkins_home/war/,AVAILABLE}{/var/jenkins_home/war}
jenkins_1  | Apr 09, 2017 4:01:53 AM org.eclipse.jetty.util.log.JavaUtilLog info
jenkins_1  | INFO: Started ServerConnector@58cbafc2{HTTP/1.1}{0.0.0.0:8080}
jenkins_1  | Apr 09, 2017 4:01:53 AM org.eclipse.jetty.util.log.JavaUtilLog info
jenkins_1  | INFO: Started @6287ms
jenkins_1  | Apr 09, 2017 4:01:53 AM winstone.Logger logInternal
jenkins_1  | INFO: Winstone Servlet Engine v2.0 running: controlPort=disabled
jenkins_1  | Apr 09, 2017 4:01:54 AM jenkins.InitReactorRunner$1 onAttained
jenkins_1  | INFO: Started initialization
jenkins_1  | Apr 09, 2017 4:01:57 AM jenkins.InitReactorRunner$1 onAttained
jenkins_1  | INFO: Listed all plugins
jenkins_1  | Apr 09, 2017 4:01:57 AM jenkins.bouncycastle.api.SecurityProviderInitializer addSecurityProvider
jenkins_1  | INFO: Initializing Bouncy Castle security provider.
jenkins_1  | Apr 09, 2017 4:01:57 AM jenkins.bouncycastle.api.SecurityProviderInitializer addSecurityProvider
jenkins_1  | INFO: Bouncy Castle security provider initialized.
gitlab2_1  | gitlab Reconfigured!
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-03-25_15:11:15.70081 localhost @ - - [2017-03-25 15:11:15.649153142 +0000 UTC] "GET /help HTTP/1.1" 200 18546 "" "curl/7.52.0" 0.051605
gitlab2_1  | 2017-03-25_15:12:15.78791 localhost @ - - [2017-03-25 15:12:15.740309983 +0000 UTC] "GET /help HTTP/1.1" 200 18546 "" "curl/7.52.0" 0.047534
gitlab2_1  | 2017-03-25_15:13:15.89988 localhost @ - - [2017-03-25 15:13:15.830857994 +0000 UTC] "GET /help HTTP/1.1" 200 18546 "" "curl/7.52.0" 0.068972
gitlab2_1  | 2017-03-25_15:14:15.99327 localhost @ - - [2017-03-25 15:14:15.939660504 +0000 UTC] "GET /help HTTP/1.1" 200 18546 "" "curl/7.52.0" 0.053554
gitlab2_1  | 2017-03-25_15:15:16.08499 localhost @ - - [2017-03-25 15:15:16.034488856 +0000 UTC] "GET /help HTTP/1.1" 200 18546 "" "curl/7.52.0" 0.050444
gitlab2_1  | 2017-03-25_15:16:16.17551 localhost @ - - [2017-03-25 15:16:16.121372762 +0000 UTC] "GET /help HTTP/1.1" 200 18546 "" "curl/7.52.0" 0.054083
gitlab2_1  | 2017-03-25_15:17:16.26220 localhost @ - - [2017-03-25 15:17:16.212558271 +0000 UTC] "GET /help HTTP/1.1" 200 18546 "" "curl/7.52.0" 0.049589
gitlab2_1  | 2017-03-25_15:18:16.35249 localhost @ - - [2017-03-25 15:18:16.300171313 +0000 UTC] "GET /help HTTP/1.1" 200 18546 "" "curl/7.52.0" 0.052269
gitlab2_1  | 2017-03-25_15:19:16.45166 localhost @ - - [2017-03-25 15:19:16.396480169 +0000 UTC] "GET /help HTTP/1.1" 200 18546 "" "curl/7.52.0" 0.055132
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/state <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/sidekiq/current <==
gitlab2_1  | 2017-03-25_14:20:12.38760 2017-03-25T14:20:12.387Z 429 TID-ouvjt3xq0 RepositoryCheck::BatchWorker JID-dff081f2d56be9673834ca30 INFO: done: 0.015 sec
gitlab2_1  | 2017-03-25_14:50:07.20081 2017-03-25T14:50:07.200Z 429 TID-ouvm13cn4 ExpireBuildArtifactsWorker JID-6a9ed9fc94d4897b713ca804 INFO: start
gitlab2_1  | 2017-03-25_14:50:07.20447 2017-03-25T14:50:07.204Z 429 TID-ouvlye1dc INFO: Cron Jobs - add job with name: expire_build_artifacts_worker
gitlab2_1  | 2017-03-25_14:50:07.21172 2017-03-25T14:50:07.211Z 429 TID-ouvm13cn4 ExpireBuildArtifactsWorker JID-6a9ed9fc94d4897b713ca804 INFO: done: 0.011 sec
gitlab2_1  | 2017-03-25_15:00:01.82038 2017-03-25T15:00:01.819Z 429 TID-ouvjt3xq0 RepositoryArchiveCacheWorker JID-b85f8a75c7ef275f54fc373f INFO: start
gitlab2_1  | 2017-03-25_15:00:01.82373 2017-03-25T15:00:01.823Z 429 TID-ouvlye1dc INFO: Cron Jobs - add job with name: repository_archive_cache_worker
gitlab2_1  | 2017-03-25_15:00:01.83165 2017-03-25T15:00:01.831Z 429 TID-ouvjt3wpc ImportExportProjectCleanupWorker JID-ab5f233f5dfb05a817d5cfd0 INFO: start
gitlab2_1  | 2017-03-25_15:00:01.83308 2017-03-25T15:00:01.832Z 429 TID-ouvjt3xq0 RepositoryArchiveCacheWorker JID-b85f8a75c7ef275f54fc373f INFO: done: 0.013 sec
gitlab2_1  | 2017-03-25_15:00:01.83528 2017-03-25T15:00:01.835Z 429 TID-ouvlye1dc INFO: Cron Jobs - add job with name: import_export_project_cleanup_worker
gitlab2_1  | 2017-03-25_15:00:01.83976 2017-03-25T15:00:01.839Z 429 TID-ouvjt3wpc ImportExportProjectCleanupWorker JID-ab5f233f5dfb05a817d5cfd0 INFO: done: 0.008 sec
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/sidekiq/state <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-rails/gitlab-rails-db-migrate-2017-03-16-07-30-39.log <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-rails/production.log <==
gitlab2_1  | Started GET "/help" for 127.0.0.1 at 2017-03-25 15:17:16 +0000
gitlab2_1  | Processing by HelpController#index as */*
gitlab2_1  | Completed 200 OK in 45ms (Views: 31.9ms | ActiveRecord: 1.6ms)
gitlab2_1  | Started GET "/help" for 127.0.0.1 at 2017-03-25 15:18:16 +0000
gitlab2_1  | Processing by HelpController#index as */*
gitlab2_1  | Completed 200 OK in 45ms (Views: 33.0ms | ActiveRecord: 2.0ms)
gitlab2_1  | Started GET "/help" for 127.0.0.1 at 2017-03-25 15:19:16 +0000
gitlab2_1  | Processing by HelpController#index as */*
gitlab2_1  | Completed 200 OK in 51ms (Views: 38.9ms | ActiveRecord: 2.2ms)
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-rails/application.log <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-shell/gitlab-shell.log <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/postgresql/current <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/postgresql/state <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/unicorn/unicorn_stdout.log <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/unicorn/current <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/unicorn/state <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/unicorn/unicorn_stderr.log <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/access.log <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/current <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_error.log <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_access.log <==
gitlab2_1  | 127.0.0.1 - - [25/Mar/2017:15:11:15 +0000] "GET /help HTTP/1.1" 200 18574 "-" "curl/7.52.0"
gitlab2_1  | 127.0.0.1 - - [25/Mar/2017:15:12:15 +0000] "GET /help HTTP/1.1" 200 18574 "-" "curl/7.52.0"
gitlab2_1  | 127.0.0.1 - - [25/Mar/2017:15:13:15 +0000] "GET /help HTTP/1.1" 200 18574 "-" "curl/7.52.0"
gitlab2_1  | 127.0.0.1 - - [25/Mar/2017:15:14:15 +0000] "GET /help HTTP/1.1" 200 18574 "-" "curl/7.52.0"
gitlab2_1  | 127.0.0.1 - - [25/Mar/2017:15:15:16 +0000] "GET /help HTTP/1.1" 200 18574 "-" "curl/7.52.0"
gitlab2_1  | 127.0.0.1 - - [25/Mar/2017:15:16:16 +0000] "GET /help HTTP/1.1" 200 18574 "-" "curl/7.52.0"
gitlab2_1  | 127.0.0.1 - - [25/Mar/2017:15:17:16 +0000] "GET /help HTTP/1.1" 200 18574 "-" "curl/7.52.0"
gitlab2_1  | 127.0.0.1 - - [25/Mar/2017:15:18:16 +0000] "GET /help HTTP/1.1" 200 18574 "-" "curl/7.52.0"
gitlab2_1  | 127.0.0.1 - - [25/Mar/2017:15:19:16 +0000] "GET /help HTTP/1.1" 200 18574 "-" "curl/7.52.0"
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/error.log <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/sshd/current <==
gitlab2_1  | 2017-04-09_04:01:51.26630 Server listening on 0.0.0.0 port 22.
gitlab2_1  | 2017-04-09_04:01:51.26637 Server listening on :: port 22.
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/redis/current <==
gitlab2_1  | 2017-03-25_15:18:04.08750 328:M 25 Mar 15:18:04.087 * Background saving started by pid 13270
gitlab2_1  | 2017-03-25_15:18:04.09980 13270:C 25 Mar 15:18:04.099 * DB saved on disk
gitlab2_1  | 2017-03-25_15:18:04.10042 13270:C 25 Mar 15:18:04.100 * RDB: 10 MB of memory used by copy-on-write
gitlab2_1  | 2017-03-25_15:18:04.18874 328:M 25 Mar 15:18:04.188 * Background saving terminated with success
gitlab2_1  | 2017-03-25_15:19:05.03590 328:M 25 Mar 15:19:05.035 * 10000 changes in 60 seconds. Saving...
gitlab2_1  | 2017-03-25_15:19:05.03661 328:M 25 Mar 15:19:05.036 * Background saving started by pid 13339
gitlab2_1  | 2017-03-25_15:19:05.04903 13339:C 25 Mar 15:19:05.048 * DB saved on disk
gitlab2_1  | 2017-03-25_15:19:05.04950 13339:C 25 Mar 15:19:05.049 * RDB: 10 MB of memory used by copy-on-write
gitlab2_1  | 2017-03-25_15:19:05.13728 328:M 25 Mar 15:19:05.137 * Background saving terminated with success
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/redis/state <==
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/logrotate/current <==
gitlab2_1  | tail: '/var/log/gitlab/redis/current' has become inaccessible: No such file or directory
gitlab2_1  | tail: '/var/log/gitlab/sidekiq/current' has become inaccessible: No such file or directory
gitlab2_1  | tail: '/var/log/gitlab/sidekiq/current' has appeared;  following new file
gitlab2_1  | tail: '/var/log/gitlab/redis/current' has appeared;  following new file
gitlab2_1  | tail: '/var/log/gitlab/gitlab-workhorse/current' has become inaccessible: No such file or directory
gitlab2_1  | tail: '/var/log/gitlab/gitlab-workhorse/current' has appeared;  following new file
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/redis/current <==
gitlab2_1  | 2017-04-09_04:02:06.41004                 _._
gitlab2_1  | 2017-04-09_04:02:06.41013            _.-``__ ''-._
gitlab2_1  | 2017-04-09_04:02:06.41013       _.-``    `.  `_.  ''-._           Redis 3.2.5 (00000000/0) 64 bit
gitlab2_1  | 2017-04-09_04:02:06.41014   .-`` .-```.  ```\/    _.,_ ''-._
gitlab2_1  | 2017-04-09_04:02:06.41014  (    '      ,       .-`  | `,    )     Running in standalone mode
gitlab2_1  | 2017-04-09_04:02:06.41014  |`-._`-...-` __...-.``-._|'` _.-'|     Port: 0
gitlab2_1  | 2017-04-09_04:02:06.41014  |    `-._   `._    /     _.-'    |     PID: 376
gitlab2_1  | 2017-04-09_04:02:06.41014   `-._    `-._  `-./  _.-'    _.-'
gitlab2_1  | 2017-04-09_04:02:06.41015  |`-._`-._    `-.__.-'    _.-'_.-'|
gitlab2_1  | 2017-04-09_04:02:06.41015  |    `-._`-._        _.-'_.-'    |           http://redis.io
gitlab2_1  | 2017-04-09_04:02:06.41015   `-._    `-._`-.__.-'_.-'    _.-'
gitlab2_1  | 2017-04-09_04:02:06.41015  |`-._`-._    `-.__.-'    _.-'_.-'|
gitlab2_1  | 2017-04-09_04:02:06.41015  |    `-._`-._        _.-'_.-'    |
gitlab2_1  | 2017-04-09_04:02:06.41015   `-._    `-._`-.__.-'_.-'    _.-'
gitlab2_1  | 2017-04-09_04:02:06.41015       `-._    `-.__.-'    _.-'
gitlab2_1  | 2017-04-09_04:02:06.41016           `-._        _.-'
gitlab2_1  | 2017-04-09_04:02:06.41016               `-.__.-'
gitlab2_1  | 2017-04-09_04:02:06.41016
gitlab2_1  | 2017-04-09_04:02:06.41533 376:M 09 Apr 04:02:06.415 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/unicorn/current <==
gitlab2_1  | 2017-04-09_04:02:06.41978 starting new unicorn master
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/redis/current <==
gitlab2_1  | 2017-04-09_04:02:06.42735 376:M 09 Apr 04:02:06.427 # Server started, Redis version 3.2.5
gitlab2_1  | 2017-04-09_04:02:06.42758 376:M 09 Apr 04:02:06.427 # WARNING overcommit_memory is set to 0! Background save may fail under low memory condition. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
gitlab2_1  | 2017-04-09_04:02:06.42793 376:M 09 Apr 04:02:06.427 # WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
gitlab2_1  | 2017-04-09_04:02:06.45891 376:M 09 Apr 04:02:06.458 * DB loaded from disk: 0.031 seconds
gitlab2_1  | 2017-04-09_04:02:06.46451 376:M 09 Apr 04:02:06.464 * The server is now ready to accept connections at /var/opt/gitlab/redis/redis.socket
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/postgresql/current <==
gitlab2_1  | 2017-04-09_04:02:06.48026 LOG:  database system was interrupted; last known up at 2017-03-25 15:54:13 GMT
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-04-09_04:02:06.53642 2017/04/09 04:02:06 Starting gitlab-workhorse v1.3.0-20170307.220611
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/postgresql/current <==
gitlab2_1  | 2017-04-09_04:02:07.12650 LOG:  database system was not properly shut down; automatic recovery in progress
gitlab2_1  | 2017-04-09_04:02:07.13057 LOG:  redo starts at 1/F5000080
gitlab2_1  | 2017-04-09_04:02:07.13596 LOG:  unexpected pageaddr 1/E7000000 in log file 1, segment 246, offset 0
gitlab2_1  | 2017-04-09_04:02:07.13600 LOG:  redo done at 1/F5000080
gitlab2_1  | 2017-04-09_04:02:07.16182 LOG:  database system is ready to accept connections
gitlab2_1  | 2017-04-09_04:02:07.17119 LOG:  autovacuum launcher started
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-04-09_04:02:09.26011 2017/04/09 04:02:09 error: GET "/": badgateway: failed after 0s: dial unix /var/opt/gitlab/gitlab-rails/sockets/gitlab.socket: connect: connection refused
gitlab2_1  | 2017-04-09_04:02:09.26942 2017/04/09 04:02:09 ErrorPage: serving predefined error page: 502
gitlab2_1  | 2017-04-09_04:02:09.27151 192.168.33.10 @ - - [2017-04-09 04:02:09.237557953 +0000 UTC] "GET / HTTP/1.1" 502 2662 "" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36" 0.033919
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_access.log <==
gitlab2_1  | 192.168.33.1 - - [09/Apr/2017:04:02:09 +0000] "GET / HTTP/1.1" 502 2674 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36"
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-04-09_04:02:09.72479 2017/04/09 04:02:09 Send static file "/opt/gitlab/embedded/service/gitlab-rails/public/favicon.ico" ("") for GET "/favicon.ico"
gitlab2_1  | 2017-04-09_04:02:09.75091 192.168.33.10 @ - - [2017-04-09 04:02:09.724326594 +0000 UTC] "GET /favicon.ico HTTP/1.1" 200 5430 "http://192.168.33.10/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36" 0.026507
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_access.log <==
gitlab2_1  | 192.168.33.1 - - [09/Apr/2017:04:02:09 +0000] "GET /favicon.ico HTTP/1.1" 200 5430 "http://192.168.33.10/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36"
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/unicorn/unicorn_stderr.log <==
gitlab2_1  | I, [2017-04-09T04:02:10.188497 #413]  INFO -- : Refreshing Gem list
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-04-09_04:02:14.84889 2017/04/09 04:02:14 error: GET "/": badgateway: failed after 0s: dial unix /var/opt/gitlab/gitlab-rails/sockets/gitlab.socket: connect: connection refused
gitlab2_1  | 2017-04-09_04:02:14.84892 2017/04/09 04:02:14 ErrorPage: serving predefined error page: 502
gitlab2_1  | 2017-04-09_04:02:14.84892 192.168.33.10 @ - - [2017-04-09 04:02:14.848014488 +0000 UTC] "GET / HTTP/1.1" 502 2662 "" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36" 0.000816
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_access.log <==
gitlab2_1  | 192.168.33.1 - - [09/Apr/2017:04:02:14 +0000] "GET / HTTP/1.1" 502 2674 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36"
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-04-09_04:02:20.85026 2017/04/09 04:02:20 error: GET "/": badgateway: failed after 0s: dial unix /var/opt/gitlab/gitlab-rails/sockets/gitlab.socket: connect: connection refused
gitlab2_1  | 2017-04-09_04:02:20.85031 2017/04/09 04:02:20 ErrorPage: serving predefined error page: 502
gitlab2_1  | 2017-04-09_04:02:20.85032 192.168.33.10 @ - - [2017-04-09 04:02:20.849522513 +0000 UTC] "GET / HTTP/1.1" 502 2662 "" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36" 0.000484
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_access.log <==
gitlab2_1  | 192.168.33.1 - - [09/Apr/2017:04:02:20 +0000] "GET / HTTP/1.1" 502 2674 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36"
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-04-09_04:02:21.08587 2017/04/09 04:02:21 Send static file "/opt/gitlab/embedded/service/gitlab-rails/public/favicon.ico" ("") for GET "/favicon.ico"
gitlab2_1  | 2017-04-09_04:02:21.08590 192.168.33.10 @ - - [2017-04-09 04:02:21.08259619 +0000 UTC] "GET /favicon.ico HTTP/1.1" 200 5430 "http://192.168.33.10/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36" 0.003193
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_access.log <==
gitlab2_1  | 192.168.33.1 - - [09/Apr/2017:04:02:21 +0000] "GET /favicon.ico HTTP/1.1" 200 5430 "http://192.168.33.10/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36"
jenkins_1  | Apr 09, 2017 4:02:26 AM jenkins.InitReactorRunner$1 onAttained
jenkins_1  | INFO: Prepared all plugins
jenkins_1  | Apr 09, 2017 4:02:26 AM jenkins.InitReactorRunner$1 onAttained
jenkins_1  | INFO: Started all plugins
jenkins_1  | Apr 09, 2017 4:02:33 AM hudson.ExtensionFinder$GuiceFinder$FaultTolerantScope$1 error
jenkins_1  | INFO: Failed to instantiate optional component hudson.plugins.build_timeout.operations.AbortAndRestartOperation$DescriptorImpl; skipping
jenkins_1  | Apr 09, 2017 4:02:34 AM jenkins.InitReactorRunner$1 onAttained
jenkins_1  | INFO: Augmented all extensions
jenkins_1  | Apr 09, 2017 4:02:34 AM jenkins.InitReactorRunner$1 onAttained
jenkins_1  | INFO: Loaded all jobs
jenkins_1  | Apr 09, 2017 4:02:35 AM jenkins.util.groovy.GroovyHookScript execute
jenkins_1  | INFO: Executing /var/jenkins_home/init.groovy.d/tcp-slave-agent-port.groovy
jenkins_1  | Apr 09, 2017 4:02:35 AM hudson.model.AsyncPeriodicWork$1 run
jenkins_1  | INFO: Started Download metadata
jenkins_1  | Apr 09, 2017 4:02:35 AM org.jenkinsci.main.modules.sshd.SSHD start
jenkins_1  | INFO: Started SSHD at port 39880
jenkins_1  | Apr 09, 2017 4:02:37 AM jenkins.InitReactorRunner$1 onAttained
jenkins_1  | INFO: Completed initialization
jenkins_1  | Apr 09, 2017 4:02:37 AM org.springframework.context.support.AbstractApplicationContext prepareRefresh
jenkins_1  | INFO: Refreshing org.springframework.web.context.support.StaticWebApplicationContext@475b65fc: display name [Root WebApplicationContext]; startup date [Sun Apr 09 04:02:37 UTC 2017]; root of context hierarchy
jenkins_1  | Apr 09, 2017 4:02:37 AM org.springframework.context.support.AbstractApplicationContext obtainFreshBeanFactory
jenkins_1  | INFO: Bean factory for application context [org.springframework.web.context.support.StaticWebApplicationContext@475b65fc]: org.springframework.beans.factory.support.DefaultListableBeanFactory@3fb9e3c
jenkins_1  | Apr 09, 2017 4:02:37 AM org.springframework.beans.factory.support.DefaultListableBeanFactory preInstantiateSingletons
jenkins_1  | INFO: Pre-instantiating singletons in org.springframework.beans.factory.support.DefaultListableBeanFactory@3fb9e3c: defining beans [authenticationManager]; root of factory hierarchy
jenkins_1  | Apr 09, 2017 4:02:38 AM org.springframework.context.support.AbstractApplicationContext prepareRefresh
jenkins_1  | INFO: Refreshing org.springframework.web.context.support.StaticWebApplicationContext@463e5bf0: display name [Root WebApplicationContext]; startup date [Sun Apr 09 04:02:38 UTC 2017]; root of context hierarchy
jenkins_1  | Apr 09, 2017 4:02:38 AM org.springframework.context.support.AbstractApplicationContext obtainFreshBeanFactory
jenkins_1  | INFO: Bean factory for application context [org.springframework.web.context.support.StaticWebApplicationContext@463e5bf0]: org.springframework.beans.factory.support.DefaultListableBeanFactory@2294613e
jenkins_1  | Apr 09, 2017 4:02:38 AM org.springframework.beans.factory.support.DefaultListableBeanFactory preInstantiateSingletons
jenkins_1  | INFO: Pre-instantiating singletons in org.springframework.beans.factory.support.DefaultListableBeanFactory@2294613e: defining beans [filter,legacy]; root of factory hierarchy
jenkins_1  | Apr 09, 2017 4:02:39 AM hudson.WebAppMain$3 run
jenkins_1  | INFO: Jenkins is fully up and running
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-04-09_04:02:41.03985 2017/04/09 04:02:41 error: GET "/": badgateway: failed after 0s: dial unix /var/opt/gitlab/gitlab-rails/sockets/gitlab.socket: connect: connection refused
gitlab2_1  | 2017-04-09_04:02:41.03987 2017/04/09 04:02:41 ErrorPage: serving predefined error page: 502
gitlab2_1  | 2017-04-09_04:02:41.03987 192.168.33.10 @ - - [2017-04-09 04:02:41.039344207 +0000 UTC] "GET / HTTP/1.1" 502 2662 "" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36" 0.000450
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_access.log <==
gitlab2_1  | 192.168.33.1 - - [09/Apr/2017:04:02:41 +0000] "GET / HTTP/1.1" 502 2674 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36"
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-04-09_04:02:41.53920 2017/04/09 04:02:41 Send static file "/opt/gitlab/embedded/service/gitlab-rails/public/favicon.ico" ("") for GET "/favicon.ico"
gitlab2_1  | 2017-04-09_04:02:41.53923 192.168.33.10 @ - - [2017-04-09 04:02:41.538845943 +0000 UTC] "GET /favicon.ico HTTP/1.1" 200 5430 "http://192.168.33.10/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36" 0.000261
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_access.log <==
gitlab2_1  | 192.168.33.1 - - [09/Apr/2017:04:02:41 +0000] "GET /favicon.ico HTTP/1.1" 200 5430 "http://192.168.33.10/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36"
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-04-09_04:02:42.51293 2017/04/09 04:02:42 error: GET "/": badgateway: failed after 0s: dial unix /var/opt/gitlab/gitlab-rails/sockets/gitlab.socket: connect: connection refused
gitlab2_1  | 2017-04-09_04:02:42.52227 2017/04/09 04:02:42 ErrorPage: serving predefined error page: 502
gitlab2_1  | 2017-04-09_04:02:42.52230 192.168.33.10 @ - - [2017-04-09 04:02:42.512245897 +0000 UTC] "GET / HTTP/1.1" 502 2662 "" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36" 0.000533
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_access.log <==
gitlab2_1  | 192.168.33.1 - - [09/Apr/2017:04:02:42 +0000] "GET / HTTP/1.1" 502 2674 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36"
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-04-09_04:02:42.74586 2017/04/09 04:02:42 Send static file "/opt/gitlab/embedded/service/gitlab-rails/public/favicon.ico" ("") for GET "/favicon.ico"
gitlab2_1  | 2017-04-09_04:02:42.74590 192.168.33.10 @ - - [2017-04-09 04:02:42.744963321 +0000 UTC] "GET /favicon.ico HTTP/1.1" 200 5430 "http://192.168.33.10/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36" 0.000283
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_access.log <==
gitlab2_1  | 192.168.33.1 - - [09/Apr/2017:04:02:42 +0000] "GET /favicon.ico HTTP/1.1" 200 5430 "http://192.168.33.10/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36"
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-workhorse/current <==
gitlab2_1  | 2017-04-09_04:02:47.13429 2017/04/09 04:02:47 error: GET "/help": badgateway: failed after 0s: dial unix /var/opt/gitlab/gitlab-rails/sockets/gitlab.socket: connect: connection refused
gitlab2_1  | 2017-04-09_04:02:47.13431 2017/04/09 04:02:47 ErrorPage: serving predefined error page: 502
gitlab2_1  | 2017-04-09_04:02:47.13431 localhost @ - - [2017-04-09 04:02:47.133549548 +0000 UTC] "GET /help HTTP/1.1" 502 2662 "" "curl/7.52.0" 0.000698
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/nginx/gitlab_access.log <==
gitlab2_1  | 127.0.0.1 - - [09/Apr/2017:04:02:47 +0000] "GET /help HTTP/1.1" 502 2674 "-" "curl/7.52.0"
jenkins_1  | --> setting agent port for jnlp
jenkins_1  | --> setting agent port for jnlp... done
jenkins_1  | Apr 09, 2017 4:02:48 AM hudson.model.UpdateSite updateData
jenkins_1  | INFO: Obtained the latest update center data file for UpdateSource default
jenkins_1  | Apr 09, 2017 4:02:49 AM hudson.model.DownloadService$Downloadable load
jenkins_1  | INFO: Obtained the updated data file for hudson.tasks.Maven.MavenInstaller
jenkins_1  | Apr 09, 2017 4:02:50 AM hudson.model.DownloadService$Downloadable load
jenkins_1  | INFO: Obtained the updated data file for hudson.plugins.gradle.GradleInstaller
jenkins_1  | Apr 09, 2017 4:02:51 AM hudson.model.DownloadService$Downloadable load
jenkins_1  | INFO: Obtained the updated data file for hudson.tasks.Ant.AntInstaller
jenkins_1  | Apr 09, 2017 4:02:54 AM hudson.model.DownloadService$Downloadable load
jenkins_1  | INFO: Obtained the updated data file for hudson.tools.JDKInstaller
jenkins_1  | Apr 09, 2017 4:02:54 AM hudson.model.AsyncPeriodicWork$1 run
jenkins_1  | INFO: Finished Download metadata. 19,664 ms
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/sidekiq/current <==
gitlab2_1  | 2017-04-09_04:02:56.47980 2017-04-09T04:02:56.479Z 373 TID-ordo4e13k INFO: Booting Sidekiq 4.2.7 with redis options {:path=>"/var/opt/gitlab/redis/redis.socket", :namespace=>"resque:gitlab", :url=>nil}
gitlab2_1  | 2017-04-09_04:02:56.76110 2017-04-09T04:02:56.761Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: stuck_ci_builds_worker
gitlab2_1  | 2017-04-09_04:02:56.76452 2017-04-09T04:02:56.764Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: expire_build_artifacts_worker
gitlab2_1  | 2017-04-09_04:02:56.76676 2017-04-09T04:02:56.766Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: repository_check_worker
gitlab2_1  | 2017-04-09_04:02:56.76871 2017-04-09T04:02:56.768Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: admin_email_worker
gitlab2_1  | 2017-04-09_04:02:56.77343 2017-04-09T04:02:56.773Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: repository_archive_cache_worker
gitlab2_1  | 2017-04-09_04:02:56.77637 2017-04-09T04:02:56.776Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: import_export_project_cleanup_worker
gitlab2_1  | 2017-04-09_04:02:56.78011 2017-04-09T04:02:56.780Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: requests_profiles_worker
gitlab2_1  | 2017-04-09_04:02:56.78417 2017-04-09T04:02:56.784Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: remove_expired_members_worker
gitlab2_1  | 2017-04-09_04:02:56.78589 2017-04-09T04:02:56.785Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: remove_expired_group_links_worker
gitlab2_1  | 2017-04-09_04:02:56.79022 2017-04-09T04:02:56.790Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: prune_old_events_worker
gitlab2_1  | 2017-04-09_04:02:56.79108 2017-04-09T04:02:56.791Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: trending_projects_worker
gitlab2_1  | 2017-04-09_04:02:56.79636 2017-04-09T04:02:56.796Z 373 TID-ordo4e13k INFO: Cron Jobs - add job with name: remove_unreferenced_lfs_objects_worker
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/gitlab-rails/production.log <==
gitlab2_1  | ** [Raven] Raven 2.0.2 configured not to capture errors.
gitlab2_1  | ** [Raven] Raven 2.0.2 configured not to capture errors.
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/sidekiq/current <==
gitlab2_1  | 2017-04-09_04:03:08.17686 2017-04-09T04:03:08.176Z 373 TID-ordo4e13k INFO: Running in ruby 2.3.3p222 (2016-11-21 revision 56859) [x86_64-linux]
gitlab2_1  | 2017-04-09_04:03:08.18014 2017-04-09T04:03:08.180Z 373 TID-ordo4e13k INFO: See LICENSE and the LGPL-3.0 for licensing details.
gitlab2_1  | 2017-04-09_04:03:08.18095 2017-04-09T04:03:08.180Z 373 TID-ordo4e13k INFO: Upgrade to Sidekiq Pro for more features and support: http://sidekiq.org
gitlab2_1  | 2017-04-09_04:03:08.18423 2017-04-09T04:03:08.184Z 373 TID-ordo4e13k INFO: Starting processing, hit Ctrl-C to stop
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/unicorn/unicorn_stderr.log <==
gitlab2_1  | I, [2017-04-09T04:03:08.331609 #413]  INFO -- : listening on addr=127.0.0.1:8080 fd=17
gitlab2_1  | I, [2017-04-09T04:03:08.332825 #413]  INFO -- : unlinking existing socket=/var/opt/gitlab/gitlab-rails/sockets/gitlab.socket
gitlab2_1  | I, [2017-04-09T04:03:08.333071 #413]  INFO -- : listening on addr=/var/opt/gitlab/gitlab-rails/sockets/gitlab.socket fd=18
gitlab2_1  | I, [2017-04-09T04:03:08.657635 #413]  INFO -- : master process ready
gitlab2_1  | I, [2017-04-09T04:03:08.701213 #439]  INFO -- : worker=0 ready
gitlab2_1  | I, [2017-04-09T04:03:08.737773 #453]  INFO -- : worker=1 ready
gitlab2_1  |
gitlab2_1  | ==> /var/log/gitlab/unicorn/current <==
gitlab2_1  | 2017-04-09_04:03:09.70762 adopted new unicorn master 413

```
宿主机访问:
gitlab http://192.168.33.10/ root abcd1234
jenkins  http://192.168.33.10:8080 admin abcd1234

view gitlab jenkins
cat jenkins configuration 1.groovy (可以加入git版本管理)
```
node {
    stage "Checkout"
    git url: "http://gitlab2/root/hello.git"

    stage "CheckStyle"
    sh "gradle check --stacktrace"
    archiveCheckstyleResults()

    stage "Build/Analyse/Test"
    sh "git log --format='%H' -n 1 >  src/main/resources/VERSION"
    sh "date >> src/main/resources/VERSION"
    sh "gradle clean build --stacktrace"
    archiveUnitTestResults()

    stage "Generate Docker image"
    sh "pwd"
    sh "cp build/libs/hello-0.0.1-SNAPSHOT.jar src/main/docker"
    sh "cd src/main/docker && docker build -t leanms/hello:0.1 ."

    stage name: "Deploy Docker", concurrency: 1
    sh "docker run -p 8888:8080 -d --network=ci_hello --name hello_app -t leanms/hello:0.1"
    sleep 20

    stage name: "Test APP"
    sh "ping -c 1 hello_app"
    retry(10) {
        sh "netcat -vzw1 hello_app 8080"
        sh "curl http://hello_app:8080"
    }
    sh "curl http://hello_app:8080/version"
    sh "docker rm \$(docker stop \$(docker ps -a -q --filter ancestor=leanms/hello:0.1 --format='{{.ID}}'))"
}

def archiveUnitTestResults() {
    step([$class: "JUnitResultArchiver", testResults: "build/**/TEST-*.xml"])
}

def archiveCheckstyleResults() {
    step([$class: "CheckStylePublisher",
          canComputeNew: false,
          defaultEncoding: "",
          healthy: "",
          pattern: "build/reports/checkstyle/main.xml",
          unHealthy: ""])
}
```
cd文件夹中 3.groovy
```

```

精益开发流



