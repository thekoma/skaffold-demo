Steps:
 - connect to cloud Workstation
 - clone project
 - requirements:
   - Artifact Registry
   - CloudRun API
   - CloudBuild API
   - skaffold cli
   - minikube

 - set default repo:
skaffold config default-repo europe-docker.pkg.dev/serverless-gdg/serverless-demo-gdg

# Set Repo
skaffold dev --default-repo europe-docker.pkg.dev/serverless-gdg/serverless-demo-gdg


skaffold build --file-output=artifacts.json 