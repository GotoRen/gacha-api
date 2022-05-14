# gacha-api: Deploy the simple API on Amazon Web Services

[![Go](https://github.com/GotoRen/gacha-api/actions/workflows/go.yml/badge.svg)](https://github.com/GotoRen/gacha-api/actions/workflows/go.yml)
[![Language](https://img.shields.io/badge/Go-1.18.0-blue.svg)](https://github.com/Pluslab/cyphonic)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
## Concept of this project
This project was started to run the gacha API writen by Golang on AWS.<br>
Built an API for creating basic applications and designed the infrastructure for operation.<br>
It can be run locally with docker-compose. The production will be operated using AWS ECS, AWS RDS (Aurora), etc...<br>

## Overview
まずは、キャッチアップにあたり、AWSの基本的なサービスを使いこなせるようになるため、一通り触ってどんなサービスなのか特徴を知り、今後の技術選定に活かしていきたいという思いがありました。
また、キャッチアップだけでなく、12factor-appを意識した設計にしたり、「オートスケール」「負荷分散」「監視・通知」「ログ収集・分析」といった運用に必要な機能を盛り込んで作ったりしていくことを意識しました。

## Concept of DevOps architecture
![architecture](https://user-images.githubusercontent.com/63791288/113522998-0c822200-95e0-11eb-851a-ee61c69076f1.png)

### CI/CDよるデプロイ
CIは ワークフローに設定された yaml (run) により、DockerイメージをECRに、イメージ情報をJSONで定義してS3に保存します。また、CDは、CodePipelineを用いてS3をソースとして、ECRからイメージをプルすることでECS上に展開します。

### イミュータブルなインフラ設計
耐障害性を考慮し、AZを分けてインスタンスを立て、ALBを設置することでトラフィックを分散します。さらに、より本番環境に近いサービスにするため、Googleドメインで取得したドメイン名をRoute53で付加し、運用を行いました。
また、Auroraへの接続情報はSystems ManagerのParameterStoreを使用して環境変数を設定し、セキュアに接続可能にしました。
運用に当たって、CI/CDと組み合わせ、開発者は一度構築したインフラの設定変更は行わないように、イミュータブルなインフラ設計（Immutable Infrastructure）を意識しました。

### 監視・通知基盤　
デプロイ時は、CodePipelineのイベントを検知してSNS, Chatbotと組み合わせてSlackに通知します。また、監視として、UptimeRobotを取り入れ、サービス全体を外形監視します。

### ログ分析基盤
ロギング基盤の部分では通常であれば、Logs Insightsで完結する部分を、様々なサービスを織り交ぜて作ってみたいという思いから、Lambdaを使ってみたり、Kinesis DataFirehoseなどを使用しました。また、取得したログはS3に保存して、Athena（Glue）からクエリをかけられる仕組みを構築しました。