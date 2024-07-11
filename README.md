# BlogService
## Description
This is a simple blog services using node.js.
It supports the following features.
* Post creation
* Comment on Posts
* Comment moderation

## App Overview
The blog app functions as a microservices architecture with six independent services: Client, Posts, Comments, Moderation, Query, and Event Bus. Deployed within a Kubernetes cluster, an ingress controller manages the load balancing of external traffic to each service. As per the microservices design, these services do not communicate directly. Instead, they collaborate through the emission of events triggered whenever a service completes a task.
[!test image]{images/app.jpeg}
### Application strategy
#### Client service
#### Comment Service
#### Post Service
### Infrastrucure strategy
- Skaffold
    * Automated Build and Deploy: Skaffold automatically detects code or configuration changes and builds and deploys them to a Kubernetes cluster. This eliminates the need for manually executing build and deployment procedures.
    * Real-time Updates: Changes made locally are instantly reflected in the cluster. Developers can quickly verify application modifications, leading to rapid feedback.
    * Simultaneous Build and Deploy of Multiple Services: Skaffold allows simultaneous building and deployment of multiple services within a microservices application. This streamlines the process, even when there are dependencies between services.
- Kubernetes

    * Scalability and Load Balancing: Cloud-native applications can scale instantly and adapt flexibly to increased traffic. Additionally, Kubernetes allows for the even distribution of loads.
    * Ease of Deployment: Kubernetes standardizes application deployment, ensuring reproducibility. You can define container images and configurations, using them to deploy applications consistently.
    * Service Discovery and Load Balancing: Services can communicate with each other using names, and Kubernetes automatically handles internal network configurations. Moreover, it enables the distribution of loads evenly across services.
    * Self-Healing: Kubernetes is resilient to application or hardware failures, providing self-repair capabilities. For instance, containers that terminate abnormally can be automatically restarted.
    * Flexibility Across Environments: Kubernetes is compatible with on-premises environments and various cloud providers, ensuring consistent performance across different environments. This makes it suitable for multi-cloud strategies and hybrid cloud environments.
