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

