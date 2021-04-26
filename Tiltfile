allow_k8s_contexts('workshop-cluster')

# Deploy: tell Tilt what YAML to deploy
k8s_yaml(kustomize('kustomize-demo/overlays/dev'))

# Build: tell Tilt what images to build from which directories
docker_build('echo-server', 'echo_server')

# Watch: tell Tilt how to connect locally
k8s_resource(workload='echo-server', port_forwards=1323)
