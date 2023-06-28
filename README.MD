# drone-config-plugin
## gitea

> Generate New Token： Settings | Applications | Generate New Token. 

drone environment

```yaml
      - DRONE_YAML_ENDPOINT=http://git.local.lan:3003
      - DRONE_YAML_SECRET=b6a11a2ec2e957233e30d35d2dd4d1b4
```

drone-config-plugin environment

```yaml
      - PLUGIN_SECRET=b6a11a2ec2e957233e30d35d2dd4d1b4
      - PLUGIN_DEBUG=true
      - TOKEN=a9e569dbb06d301267f17ec9a78e71b04f0dbe0c
      - SERVERTYPE=gitea
      - DRONE_CONFIG_NAMESPACE_TEMP={{ .Repo.Namespace }}
      - DRONE_CONFIG_REPONAME_TEMP=dronefiles #{{ .Repo.Name }}
      - DRONE_CONFIG_BRANCH_TEMP={{ .Repo.Branch }}
      - DRONE_CONFIG_YAMLPATH_TEMP={{ .Repo.Name }}/{{ .Repo.Branch }}/.drone.yml
```




support template define

```
class Request {
  repo: Repo;
  build: Build
}
class Repository {
    id: int64;
    uid: string;
    user_id: int64;
    namespace: string;
    name: string;
    slug: string;
    scm: string;
    git_http_url: string;
    git_ssh_url: string;
    link: string;
    default_branch: string;
    private: boolean;
    visibility: string;
    active: boolean;
    config: string;
    trusted: boolean;
    protected: boolean;
    ignore_forks: boolean;
    ignore_pulls: boolean;
    cancel_pulls: boolean;
    timeout: int64;
    counter: int64;
    synced: int64;
    created: int64;
    updated: int64;
    version: int64;
}

Build definition:

class Build {
    id: int64;
    repo_id: int64;
    number: int64;
    parent: int64;
    status: string;
    error: string
    event: string;
    action: string;
    link: string;
    timestamp: int64;
    title: string;
    message: string;
    before: string;
    after: string;
    ref: string;
    source_repo: string;
    source: string;
    target: string;
    author_login: string;
    author_name: string;
    author_email: string;
    author_avatar: string;
    sender: string;
    params: [string][string];
    cron: string;
    deploy_to: string;
    deploy_id: int64;
    started: int64;
    finished: int64;
    created: int64;
    updated: int64;
    version: int64;
}
```