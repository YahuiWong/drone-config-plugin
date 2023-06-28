
可以通过您的 gitea web 界面来创建 apiKey token： Settings | Applications | Generate New Token. 注意选择必要的权限

配置 drone

```yaml
      - DRONE_YAML_ENDPOINT=http://git.local.lan:3003
      - DRONE_YAML_SECRET=b6a11a2ec2e957233e30d35d2dd4d1b4
```

在 drone-config-plugin 配置环境变量

```yaml
      - PLUGIN_SECRET=b6a11a2ec2e957233e30d35d2dd4d1b4
      - DRONE_DEBUG=true
      - DRONE_TRACE=true
      - TOKEN=a9e569dbb06d301267f17ec9a78e71b04f0dbe0c
      - SERVERTYPE=gitea
      - REPO_OWNER=adminuser
      - REPO_NAME=dronefiles
      - REPO_BRANCH=main
      - YAML_PATH=.drone.yml
```

