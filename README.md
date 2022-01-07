## Домашние задания курса [OTUS «Microservice Architecture»](https://otus.ru/lessons/microservice-architecture/)

1) [«Основы работы с Kubernetes. Часть 1»](./hw01_k8s_basics_part1)
2) [«Основы работы с Kubernetes. Часть 2»](./hw01_k8s_basics_part2)
3) [«Prometheus. Grafana»](./hw03_prometheus_grafana)
4) [«Service mesh на примере Istio»](./hw04_service_mesh_istio)
5) [«Backend for frontends. Apigateway»](./hw05_api_gateway)
6) [«Паттерны поддержания консистентности данных (Stream processing)»](./hw06_stream_processing)
7) [«Идемпотетность и коммутативность API в HTTP и очередях»](./hw07_order_service)

[«Проектная работа»](./hw08_final_project)

---

### Окружение разработки

- **os**: Linux 5.4.0-89-generic Ubuntu GNU/Linux
- **minikube version**: v1.23.2
- **minikube addons**:
    - ingress
    - dashboard
- **kubectl version**:  v1.22.2
- **helm version**:  v3.7.0
- **helmfile version**:  v0.130.1
- **newman version**:  v5.3.0
- **golangci-lint**:  v1.43.0
- **Docker Client**:
    - Docker Engine - Community;
    - Version:           20.10.9;
    - API version:       1.41;
    - Go version:        go1.16.8;
    - Git commit:        c2ea9bc;
    - Built:             Mon Oct 4 16:08:29 2021;
    - OS/Arch:           linux/amd64;
    - Context:           default;
    - Experimental:      true;
- Helm plugin list:
    - diff: 3.2.0 Preview helm upgrade changes as a diff
    - push: 0.8.1 Push chart package to ChartMuseum

---

### TODO

- [x] Makefile to facilitate CI/CD operations
