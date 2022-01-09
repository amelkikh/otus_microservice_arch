[![CodeQL](https://github.com/amelkikh/otus_microservice_arch/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/amelkikh/otus_microservice_arch/actions/workflows/codeql-analysis.yml)

## Домашние задания курса [OTUS «Microservice Architecture»](https://otus.ru/lessons/microservice-architecture/)

1) [«Основы работы с Kubernetes»](./hw01_k8s_basics)
2) [«Инфраструктурные паттерны»](./hw02_infra_patterns)
3) [«Prometheus. Grafana»](./hw03_prometheus_grafana)
4) [«Service mesh на примере Istio»](./hw04_service_mesh_istio)
5) [«Backend for frontends. Apigateway»](./hw05_api_gateway)
6) [«Проектная работа»](./hw06_final_project)
7) [«Паттерны поддержания консистентности данных (Stream processing)»](./hw07_stream_processing)
8) [«Идемпотетность и коммутативность API в HTTP и очередях»](./hw08_order_service)
9) [«Паттерны декомпозиции микросервисов»](./hw09_decomposition_patterns)
10) [«Распределенные транзакции»](./hw10_distributed_transactions)

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
- **go version**:  go1.17.3 linux/amd64
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

- [ ] Makefile to facilitate CI/CD operations
