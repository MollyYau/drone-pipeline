# drone-pipeline
Canary Release With Helm Deploy Style To K8S

1. build source code to image;upload image to registry
2. download chart sed -i package reindex repository
3. helm upgrade canary
4. monitor cannary status
5. if success
  rolling update
6. if failure
  rollback
7. done!
