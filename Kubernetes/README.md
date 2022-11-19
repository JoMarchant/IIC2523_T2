# Guía Kubernetes

## 1. Crear un cluster Kubernetes

### 1.1 ¿Qué es un cluster?

Un cluster es un conjunto de máquinas que ejecutan Kubernetes y están conectadas entre sí para que puedan comunicarse y trabajar como unidad. Para ello se deben colocar las aplicaciones a ejecutar en contenedores y estos contenedores deben ser gestionados por un orquestador de contenedores como Kubernetes.

Consiste de dos tipos de recursos:
-  **Nodos** : Son las máquinas que ejecutan los contenedores.
-  **Máster** : Es la máquina que controla el cluster.

### 1.2 Creación de un cluster con Minikube

1.  Primero se debe instalar Minikube, para ello se debe descargar el instalador de la página oficial de Minikube y ejecutarlo:  [Instalador](https://minikube.sigs.k8s.io/docs/start/).

2.  Una vez instalado Minikube, se debe iniciar el cluster con el comando: `minikube start`. Así debería verse la consola: ![start](./screenshots/start_cluster.png)

3.  Para interactuar con Kubernets se debe usar el comando `kubectl`. Se puede comprobar que se ha instalado correctamente con el comando `kubectl version`.

4.  Para ver la información del cluster se puede usar el comando `kubectl cluster-info`.

5.  Para ver los nodos del cluster se debe ejecutar el comando `kubectl get nodes`. Así debería verse la consola: ![nodes](./screenshots/get_nodes.png) Se puede observar que hay un nodo en el cluster y que está en estado Ready.

