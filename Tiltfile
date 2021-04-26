allow_k8s_contexts('workshop-cluster')

# Deploy: tell Tilt what YAML to deploy
k8s_yaml(['svc.yml', 'deploy.yml', 'health_checker_deploy.yml'])

# Build: tell Tilt what images to build from which directories
docker_build('echo-server', 'echo_server')
docker_build('simple-health-checker', 'simple_health_checker')

# Watch: tell Tilt how to connect locally
k8s_resource('echo-server', port_forwards=1323)
