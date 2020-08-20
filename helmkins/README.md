You can find here project that we call helmkins. The whole thing hasn't been refactored yet since it's still being developed. The goal of the project is to replace the current automation tools with a tool that provides more flexibility and runs in K8s. I would have loved to write completely in Groovy but using declarative pipeline language was much easier for the team. I tried to implement an Octopus-like configuration style for various reasons. I want to share a few key notes about the scripts:

### Jenkinsfile

ImageExists function checks against the docker hub if the specific release has already been pushed since it should be built only once.
One additional podpreset yaml file for development branch which spins up sonarscanner so it can scan in parallel.

### Helm chart

The whole chart is pretty straightforward and simple except ingress.yaml. Our domain system is quite predictable and fixed based on the given namespace and branching info. So domain names for each deployment are determined dynamically by the data given from Jenkinsfile. You can find the python script used for testing and its output.

### Jenkins

Jenkins itself is deployed via official Jenkins helm chart with JCasC configured.