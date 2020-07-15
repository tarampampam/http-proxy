<p align="center">
  <img src="https://hsto.org/webt/jx/ea/tw/jxeatw6qghfyfzxu2y8cymoiyck.png" alt="Logo" width="128" />
</p>

# Http Proxy Daemon

[![Release version][badge_release_version]][link_releases]
![Project language][badge_language]
[![Build Status][badge_build]][link_build]
[![Release Status][badge_release]][link_build]
[![Coverage][badge_coverage]][link_coverage]
[![Go Report][badge_goreport]][link_goreport]
[![License][badge_license]][link_license]

This application accepts HTTP requests and sending them by itself to the target resource. So, target resource is not hardcoded, and by running this application on remote server you can use it as dynamic reverse-proxy:

<p align="center">
    <a href="https://asciinema.org/a/278217" target="_blank"><img src="https://asciinema.org/a/278217.svg" width="900"></a>
</p>

## Usage example

Run proxy server:

```bash
$ ./http-proxy-daemon serve --listen 0.0.0.0:8080 --prefix 'proxy' &
2019/10/29 20:45:01.825260 Starting server on 0.0.0.0:8080
```

And then send an HTTP request to the `https://httpbin.org/get?foo=bar&bar&baz` through our server:

```bash
$ curl -s -H "foo:bar" --user-agent "fake agent" 'http://127.0.0.1:8080/proxy/https/httpbin.org/get?foo=bar&bar&baz'
{
  "args": {
    "bar": "",
    "baz": "",
    "foo": "bar"
  },
  "headers": {
    "Accept": "*/*",
    "Accept-Encoding": "gzip",
    "Foo": "bar",
    "Host": "httpbin.org",
    "User-Agent": "fake agent"
  },
  "origin": "8.8.8.8, 1.1.1.1",
  "url": "https://httpbin.org/get?foo=bar&bar&baz"
}
```

## Using docker

Run docker-container with proxy server in background _(detached)_ and listen for 8080 TCP port (incoming HTTP requests):

```bash
$ docker run --rm -d -p 8080:8080 tarampampam/http-proxy-daemon serve --listen 0.0.0.0:8080
```

### Supported tags

[![image stats](https://dockeri.co/image/tarampampam/http-proxy-daemon)][link_docker_tags]

All supported image tags [can be found here][link_docker_tags].

### Testing

For application testing we use built-in golang testing feature and `docker-ce` + `docker-compose` as develop environment. So, just write into your terminal after repository cloning:

```shell
$ make test
```

Or build binary file:

```shell
$ make build
```

## Changes log

[![Release date][badge_release_date]][link_releases]
[![Commits since latest release][badge_commits_since_release]][link_commits]

Changes log can be [found here][link_changes_log].

## Support

[![Issues][badge_issues]][link_issues]
[![Issues][badge_pulls]][link_pulls]

If you will find any package errors, please, [make an issue][link_create_issue] in current repository.

## License

This is open-sourced software licensed under the [MIT License][link_license].

[badge_build]:https://img.shields.io/github/workflow/status/tarampampam/http-proxy-daemon/tests?maxAge=30&logo=github
[badge_release]:https://img.shields.io/github/workflow/status/tarampampam/http-proxy-daemon/release?maxAge=30&label=release&logo=github
[badge_coverage]:https://img.shields.io/codecov/c/github/tarampampam/http-proxy-daemon/master.svg?maxAge=30
[badge_goreport]:https://goreportcard.com/badge/github.com/tarampampam/http-proxy-daemon
[badge_release_version]:https://img.shields.io/github/release/tarampampam/http-proxy-daemon.svg?maxAge=30
[badge_language]:https://img.shields.io/badge/language-go_1.13-blue.svg?longCache=true
[badge_license]:https://img.shields.io/github/license/tarampampam/http-proxy-daemon.svg?longCache=true
[badge_release_date]:https://img.shields.io/github/release-date/tarampampam/http-proxy-daemon.svg?maxAge=180
[badge_commits_since_release]:https://img.shields.io/github/commits-since/tarampampam/http-proxy-daemon/latest.svg?maxAge=45
[badge_issues]:https://img.shields.io/github/issues/tarampampam/http-proxy-daemon.svg?maxAge=45
[badge_pulls]:https://img.shields.io/github/issues-pr/tarampampam/http-proxy-daemon.svg?maxAge=45

[link_goreport]:https://goreportcard.com/report/github.com/tarampampam/http-proxy-daemon
[link_coverage]:https://codecov.io/gh/tarampampam/http-proxy-daemon
[link_build]:https://github.com/tarampampam/http-proxy-daemon/actions
[link_docker_hub]:https://hub.docker.com/r/tarampampam/http-proxy-daemon/
[link_docker_tags]:https://hub.docker.com/r/tarampampam/http-proxy-daemon/tags
[link_license]:https://github.com/tarampampam/http-proxy-daemon/blob/master/LICENSE
[link_releases]:https://github.com/tarampampam/http-proxy-daemon/releases
[link_commits]:https://github.com/tarampampam/http-proxy-daemon/commits
[link_changes_log]:https://github.com/tarampampam/http-proxy-daemon/blob/master/CHANGELOG.md
[link_issues]:https://github.com/tarampampam/http-proxy-daemon/issues
[link_create_issue]:https://github.com/tarampampam/http-proxy-daemon/issues/new/choose
[link_pulls]:https://github.com/tarampampam/http-proxy-daemon/pulls
