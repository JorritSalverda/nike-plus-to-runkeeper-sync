# nike-plus-to-runkeeper-sync

This small application runs inside Kubernetes and exports runs from Nike+ and imports them into RunKeeper. Progress and credentials/tokens are stored in Kubernetes configmap and secret. The application runs as a Kubernetes CronJob.

[![License](https://img.shields.io/github/license/JorritSalverda/nike-plus-to-runkeeper-sync.svg)](https://github.com/JorritSalverda/nike-plus-to-runkeeper-sync/blob/master/LICENSE)