# kubeSLA

**kubeSLA** is a Kubernetes operator designed to manage and enforce Service Level Agreements (SLAs) for applications running in a Kubernetes cluster. 
This operator automates the monitoring and management of SLAs, ensuring that your applications meet the desired performance and reliability standards.


## Description
kubeSLA provides a powerful and flexible way to define, monitor, and enforce SLAs within your Kubernetes environment. 
By leveraging custom resource definitions (CRDs), kubeSLA allows users to specify detailed SLA requirements, including metrics, thresholds, and actions to be taken when SLAs are breached. The operator continuously monitors the specified metrics and takes corrective actions when necessary, helping to maintain the desired level of service for your applications.

kubeSLA is an implementation of the [OpenSLO](https://github.com/OpenSLO/OpenSLO) specification, which provides a standardized way to define and manage SLAs in cloud-native environments. 
By using OpenSLO, kubeSLA ensures compatibility with other tools and platforms that support the specification, making it easy to integrate SLA management into your existing workflows.

kubeSLA is also inspired by [osko](https://github.com/oskoperator/osko) project, which provides a similar functionality for managing SLAs in Kubernetes.

## Getting Started

### Prerequisites
- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/kubesla:tag
```

**NOTE:** This image ought to be published in the personal registry you specified. 
And it is required to have access to pull the image from the working environment. 
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/kubesla:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin 
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Contributing
We welcome contributions from the community to help improve kubeSLA. To get started, please follow these steps:

1. **Fork the Repository**: Fork the [kubeSLA repository](https://github.com/maczg/kubesla) to your GitHub account.

2. **Clone the Repository**: Clone your forked repository to your local development environment.
    ```bash
    git clone https://github.com/<your-username>/kubesla-operator.git
    cd kubesla-operator
    ```

3. **Create a Branch**: Create a new branch for your feature or bug fix.
    ```bash
    git checkout -b feature/your-feature-name
    ```

4. **Make Your Changes**: Implement your feature or bug fix. Make sure to follow the existing code style and conventions.

5. **Test Your Changes**: Ensure that your changes do not break existing functionality and pass all tests. Add new tests as needed.

6. **Commit Your Changes**: Commit your changes with a clear and descriptive commit message.
    ```bash
    git commit -m "Add feature: your feature description"
    ```

7. **Push Your Changes**: Push your changes to your forked repository.
    ```bash
    git push origin feature/your-feature-name
    ```

8. **Create a Pull Request**: Open a pull request to the main repository. Provide a detailed description of your changes and any relevant information.

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

