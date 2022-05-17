# 🏃 gacha-api: Deploy the simple API on Amazon Web Services

[![Go](https://github.com/GotoRen/gacha-api/actions/workflows/go.yml/badge.svg)](https://github.com/GotoRen/gacha-api/actions/workflows/go.yml)
[![Language](https://img.shields.io/badge/Go-1.18.0-blue.svg)](https://github.com/GotoRen/gacha-api)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
## Concept of this project
This project was started to run the gacha API writen by Golang on AWS.<br>
Built an API for creating basic applications and designed the infrastructure for operation.<br>
It can be run locally with docker-compose. The production will be operated using AWS ECS, AWS RDS (Aurora), etc...<br>

## Overview
First, I would like to be able to use basic AWS services as I catch up.<br>
Then, touching and learning the features of each service, and desire to utilize the technology selection process in the future.<br>
Additionally, this product is designed with [The Twelve-Factor App](https://12factor.net/) in mind as well as catch-up.<br>
For example, "auto-scale", "load-balancing", "monitoring and notification", "log-collection and log-analysis".<br>
Finally, I will actively incorporate technologies in which we are interested.<br>

## Concept of DevOps architecture
![architecture](https://user-images.githubusercontent.com/63791288/113522998-0c822200-95e0-11eb-851a-ee61c69076f1.png)

### Deployment by CI/CD
![dev.png](https://user-images.githubusercontent.com/63791288/168437418-e4a88f98-888d-4005-b357-98f1ac04a547.png)
Continuous Integration pushes the Docker image to the ECR via yaml (run) configured in the github workflows.<br>
Image information is defined in JSON and stored in S3.<br>
CodePipeline uses S3 as a source and pull image from ECR and deploy to ECS.<br>
This is typically a continuous deployment.<br>

### Design of infrastructure
![ops.png](https://user-images.githubusercontent.com/63791288/168437521-9f775beb-e281-48e2-82ce-0032d699f666.png)
First, for fault tolerance, set up separate instances in different Availability-Zone and deploy ALB to distribute traffic.<br>
Additionally, in order to achieve a more production-like service, I used Route53 and added domain names to the operation.<br>
Connection information to Aurora was set up using Systems Manager's ParameterStore with environment variables to enable secure connections.<br>
Finally, in conjunction with CI/CD, we were aware of Immutable Infrastructure so as not to change the configuration of the infrastructure once it was built.<br>
This allows for continuous DevOps even after the start of operations.<br>

### Monitoring and Notification　
![maintenance.png](https://user-images.githubusercontent.com/63791288/168437532-a7ae8bf1-fd62-48f5-948e-f306f6a133a5.png)
- Monitoring<br>
When deployed, the system detects CodePipeline events and notifies Slack in combination with AWS's SNS and AWS's Chatbot.<br>
Additionally, UptimeRobot is incorporated as a monitoring tool and the entire service is monitored externally.<br>

- Logging<br>
I used Lambda and Kinesis DataFirehose for logging.<br>
Typically, this would be completed by Logs Insights. However, I would like to deal with various technologies. This is my concept!<br>
The acquired logs are stored in S3 and queries are executed from Athena (Glue).<br>

## Requirement

| Language/FrameWork | Version |
| :------------------ | ---------: |
| go                  |       1.18 |
| npm                 |      8.9.0 |
| python              |      3.9.0 |
| docker-compose      |      2.5.0 |
| aws-cli             |      2.7.0 |

## Project managers
- [Issues](https://github.com/GotoRen/gacha-api/issues)
- [Pull Request](https://github.com/GotoRen/gacha-api/pulls)
- [Projects](https://github.com/GotoRen/gacha-api/projects/1)

## Contributing

Bug reports and pull requests are welcome on GitHub at [https://github.com/GotoRen/gacha-api](https://github.com/GotoRen/gacha-api).

This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## Code of Conduct
Everyone interacting in this project’s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](./.github/CODE_OF_CONDUCT.md).

## LICENSE

[MIT License](./LICENSE)
