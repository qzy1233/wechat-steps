# 定时任务使用说明

本项目包含一个自动化的定时任务，用于每天定时设置步数。以下是关于这个定时任务的详细说明和使用指南。

## 定时任务概述

- 任务名称：Daily Cron Job
- 执行频率：每天 UTC 2:00（北京时间 10:00）
- 主要功能：自动运行脚本，设置指定的步数

## 配置说明

定时任务的配置文件位于 `.github/workflows/cron.yml`。主要配置项包括：

1. 执行时间：通过 `cron` 表达式设置，当前为 `'0 2 * * *'`
2. 步数设置：通过环境变量 `STEPS` 设置，当前默认值为 100000

## 使用步骤

1. Fork 本项目到您的 GitHub 账户。

2. 设置 GitHub Secrets：
   - 在您的 GitHub 仓库中，进入 "Settings" -> "Secrets and variables" -> "Actions"
   - 添加以下 Secrets：
     - `ZEPP_ACCOUNT`：您的 Zepp 账号
     - `ZEPP_PASSWORD`：您的 Zepp 密码

3. （可选）修改步数：
   - 如果您想修改每日设置的步数，编辑 `.github/workflows/cron.yml` 文件
   - 找到 `env` 部分，修改 `STEPS` 的值

4. 启用 GitHub Actions：
   - 在您的仓库中，进入 "Actions" 标签页
   - 如果 Actions 未启用，点击 "I understand my workflows, go ahead and enable them"

5. 验证运行：
   - 定时任务将在每天指定时间自动运行
   - 您可以在 "Actions" 标签页中查看任务的运行状态和日志

## 注意事项

- 请确保您的 Zepp 账号和密码正确，否则任务将无法成功执行。
- 频繁或大量修改步数可能违反 Zepp 的使用条款，请谨慎使用。
- 如果您想修改执行时间，可以编辑 `.github/workflows/cron.yml` 文件中的 `cron` 表达式。

## 故障排除

如果定时任务未按预期执行，请检查以下几点：

1. 确认 GitHub Actions 已在您的仓库中启用。
2. 检查 Secrets 是否正确设置。
3. 查看 Actions 日志以获取详细的错误信息。

如果问题持续，请提交一个 Issue 以获取帮助。
