# Docker Compose 项目仓库

本仓库收集了本地 NAS 上使用的 Docker 项目部署文件，统一放在 `container/` 目录下。每个项目目录包含：

- 上游项目的 `README.md`（由 GitHub Action 每日同步）
- 本地部署用的 `compose.yaml` 与 `.env`

`config.json` 保存人工维护的配置：`repo`（原仓库地址）、`readme_path`（原仓库内 README 相对路径）与 `type`（项目类型，参考 awesome-selfhosted 的一级分类）。抓取到的项目信息保存在 `containers.json` 中，由每日定时任务自动更新。状态取值为 `active`、`archived` 或 `deleted`。

表格中的“最新版本”优先取 `latest_release`，为空时取 `latest_tag`；“最后更新”优先取 `latest_release_date`，为空时取 `latest_tag_date`。

<!-- AUTO-GENERATED:START -->

| 项目 | 类型 | 描述 | 最新版本 | 最后更新 | Stars | 状态 | 仓库 |
| --- | --- | --- | --- | --- | --- | --- | --- |
| [3x-ui](container/3x-ui/) | Proxy | Supporting multi-protocol multi-user(Vmess, Vless, Trojan, ShadowSocks, Wireguard, Hysteria, Tunnel, Mixed, HTTP, Tun, MTProto، AmneziaWG)  | v3.11.0 | 2026-10-25 | 50325 | active | [repo](https://github.com/MHSanaei/3x-ui) |
| [alist](container/alist/) | File Transfer | 🗂️A file list/WebDAV program that supports multiple storages, powered by Gin and Solidjs. / 一个支持多存储的文件列表/WebDAV程序，使用 Gin 和 Solidjs。 | v3.66.0 | 2026-10-21 | 50476 | active | [repo](https://github.com/alist-org/alist) |
| [ani-rss](container/ani-rss/) | Automation | 基于RSS自动追番、订阅、下载、刮削、洗版 | v3.2.63 | 2026-11-21 | 3600 | active | [repo](https://github.com/wushuo894/ani-rss) |
| [apprise](container/apprise/) | Communication | A lightweight REST framework that wraps the Apprise Notification Library | v2.1.0 | 2026-11-08 | 1346 | active | [repo](https://github.com/caronc/apprise-api) |
| [auto-bangumi](container/auto-bangumi/) | Automation | AutoBangumi - 全自动追番工具 | 4.0.0 | 2026-10-24 | 8330 | active | [repo](https://github.com/EstrellaXD/Auto_Bangumi) |
| [bark](container/bark/) | Communication | Backend of Bark | v2.4.0 | 2026-11-06 | 3722 | active | [repo](https://github.com/Finb/bark-server) |
| [calibre-web](container/calibre-web/) | Document Management | Calibre-Web but Automated and with tons of New Features! Fully automate and simplify your eBook set up! | v4.0.8 | 2026-09-28 | 6644 | active | [repo](https://github.com/crocodilestick/Calibre-Web-Automated) |
| [chinese-sub-finder](container/chinese-sub-finder/) | Media Management | 自动化中文字幕下载。字幕网站支持 shooter、xunlei、arrst、a4k、SubtitleBest 。支持 Emby、Jellyfin、Plex、Sonarr、Radarr、TMM | v0.55.2 | 2023-11-09 | 3929 | active | [repo](https://github.com/ChineseSubFinder/ChineseSubFinder) |
| [cloudflared](container/cloudflared/) | Proxy | Cloudflare Tunnel client | 2026.11.1 | 2026-11-17 | 16686 | active | [repo](https://github.com/cloudflare/cloudflared) |
| [code-server](container/code-server/) | Software Development | VS Code in the browser | v4.146.0 | 2026-11-20 | 79925 | active | [repo](https://github.com/coder/code-server) |
| [coder](container/coder/) | Software Development | Secure environments for developers and their agents | v2.38.10 | 2026-11-16 | 17455 | active | [repo](https://github.com/coder/coder) |
| [cookiecloud](container/cookiecloud/) | Miscellaneous | CookieCloud是一个和自架服务器同步浏览器Cookie和LocalStorage的小工具，支持端对端加密，可设定同步时间间隔。本仓库包含了插件和服务器端源码。CookieCloud is a small tool for synchronizing browser cookies and LocalStorage with a self-hosted server. It supports end-to-end encryption and allows for setting the synchronization interval. This repository contains both the plugin and the server-side source code | release-v1.0.4 | 2026-10-31 | 3231 | active | [repo](https://github.com/easychen/CookieCloud) |
| [cypht](container/cypht/) | Communication | Cypht: Lightweight Open Source webmail aggregator [PHP, JS].  Supports IMAP/SMTP, JMAP and EWS (Exchange Web Services) | v2.14.0 | 2026-11-06 | 1818 | active | [repo](https://github.com/cypht-org/cypht) |
| [dashy](container/dashy/) | Personal Dashboards | 🚀 A self-hostable personal dashboard built for you. Includes status-checking, widgets, themes, icon packs, a UI editor and tons more! | 4.9.0 | 2026-10-23 | 27044 | active | [repo](https://github.com/Lissy93/dashy) |
| [docker-proxy](container/docker-proxy/) | Proxy | 自建 Docker 镜像加速与管理服务。零磁盘缓存、可视化面板、支持上游账号认证，支持一键部署Docker、K8s、Quay、Ghcr、Mcr、elastic、nvcr等镜像加速 | v5.1.9 | 2026-09-22 | 4466 | active | [repo](https://github.com/dqzboy/Docker-Proxy) |
| [dozzle](container/dozzle/) | Monitoring & Status Pages | Realtime log viewer for containers.  Supports Docker, Swarm and K8s.  | v11.5.0 | 2026-11-19 | 14994 | active | [repo](https://github.com/amir20/dozzle) |
| [dsm](container/dsm/) | Self-hosting Solutions | Virtual DSM in a Docker container. | v7.67 | 2026-11-12 | 4039 | active | [repo](https://github.com/vdsm/virtual-dsm) |
| [duplicati](container/duplicati/) | Backup |  | v2.4.1.0_stable_2026-10-29-ls319 | 2026-11-21 | 447 | active | [repo](https://github.com/linuxserver/docker-duplicati) |
| [filecodebox](container/filecodebox/) | File Transfer | 文件快递柜-匿名口令分享文本，文件，像拿快递一样取文件（FileCodeBox - File Express Cabinet - Anonymous Passcode Sharing Text, Files, Like Taking Express Delivery for Files） | v2.9.0 | 2026-11-16 | 8812 | active | [repo](https://github.com/vastsa/FileCodeBox) |
| [flaresolverr](container/flaresolverr/) | Proxy | Proxy server to bypass Cloudflare protection | v3.5.2 | 2026-09-12 | 16275 | active | [repo](https://github.com/FlareSolverr/FlareSolverr) |
| [gitea](container/gitea/) | Software Development | Git with a cup of tea! Painless self-hosted all-in-one software development service, including Git hosting, code review, team collaboration, package registry and CI/CD | v28.4.0 | 2026-11-15 | 59201 | active | [repo](https://github.com/go-gitea/gitea) |
| [grafana](container/grafana/) | Monitoring & Status Pages | The open and composable observability and data visualization platform. Visualize metrics, logs, and traces from multiple sources like Prometheus, Loki, Elasticsearch, InfluxDB, Postgres and many more.  | v13.3.1 | 2026-11-10 | 78027 | active | [repo](https://github.com/grafana/grafana) |
| [halo](container/halo/) | Content Management Systems (CMS) | Halo 是一款强大易用的开源建站工具，从个人博客、知识库，到企业官网、在线商城，Halo 都能助您轻松实现，一站式满足您的多样化建站需求。 | v2.27.3 | 2026-11-13 | 40371 | active | [repo](https://github.com/halo-dev/halo) |
| [hlink](container/hlink/) | Automation | 批量、快速硬链工具(The batch, fast hard link toolkit) | v2.1.3 | 2026-10-24 | 720 | active | [repo](https://github.com/likun7981/hlink) |
| [homepage](container/homepage/) | Personal Dashboards | A highly customizable homepage (or startpage / application dashboard) with Docker and service API integrations. | v2.8.0 | 2026-11-11 | 33770 | active | [repo](https://github.com/gethomepage/homepage) |
| [immich](container/immich/) | Photo Galleries | High performance self-hosted photo and video management solution. | v3.4.2 | 2026-11-18 | 119786 | active | [repo](https://github.com/immich-app/immich) |
| [jackett](container/jackett/) | File Transfer | API Support for your favorite torrent trackers | v0.25.20 | 2026-11-21 | 16396 | active | [repo](https://github.com/Jackett/Jackett) |
| [jellyfin](container/jellyfin/) | Media Streaming | The Free Software Media System | v10.4.3 | 2019-12-06 | 263 | active | [repo](https://github.com/nyanmisaka/jellyfin) |
| [kutt](container/kutt/) | URL Shorteners | Free Modern URL Shortener. | v3.2.6 | 2026-07-05 | 11216 | active | [repo](https://github.com/thedevs-network/kutt) |
| [memos](container/memos/) | Note-taking & Editors | Memos is an open-source, self-hosted note-taking app that brings your everyday thoughts together in a personal timeline. Your data stays yours, and you choose what to keep private or share. | 26.11.1 | 2026-11-16 | 64619 | active | [repo](https://github.com/usememos/memos) |
| [metatube](container/metatube/) | Media Management | MetaTube Plugin for Jellyfin/Emby | v2025.1102.2200.0 | 2025-11-02 | 4573 | active | [repo](https://github.com/metatube-community/jellyfin-plugin-metatube) |
| [mihomo](container/mihomo/) | Proxy | A simple Python Pydantic model for Honkai: Star Rail parsed data from the Mihomo API. | v1.19.34 | 2026-11-11 | 35975 | active | [repo](https://github.com/MetaCubeX/mihomo) |
| [music-tag](container/music-tag/) | Media Management | 音乐标签编辑器，可编辑本地音乐文件的元数据, 音乐刮削。（Editable local music file metadata.） |  |  | 6226 | active | [repo](https://github.com/xhongc/music-tag-web) |
| [neko](container/neko/) | Remote Access | A self hosted virtual browser that runs in docker and uses WebRTC. | v3.2.0 | 2026-10-15 | 22699 | active | [repo](https://github.com/m1k1o/neko) |
| [newfuture-ddns](container/newfuture-ddns/) | DNS | 🌐自动更新域名解析到本机IP,支持dnspod,阿里DNS,CloudFlare,华为云,DNSCOM... | v4.2.0 | 2026-08-27 | 4690 | active | [repo](https://github.com/NewFuture/DDNS) |
| [next-terminal](container/next-terminal/) | Remote Access | A simple, secure, easy-to-use bastion and session auditing system. Supports RDP, SSH, VNC, Telnet, HTTP, records and replays sessions for auditing and compliance. | v3.10.3 | 2026-10-18 | 5683 | active | [repo](https://github.com/next-terminal/next-terminal) |
| [nginx-proxy-manager](container/nginx-proxy-manager/) | Proxy | Docker container for managing Nginx proxy hosts with a simple, powerful interface | v2.17.1 | 2026-11-19 | 34839 | active | [repo](https://github.com/NginxProxyManager/nginx-proxy-manager) |
| [npm](container/npm/) | Proxy | Docker container for managing Nginx proxy hosts with a simple, powerful interface | v2.17.1 | 2026-11-19 | 34839 | active | [repo](https://github.com/NginxProxyManager/nginx-proxy-manager) |
| [ntfy](container/ntfy/) | Communication | Send push notifications to your phone or desktop using PUT/POST | v2.38.1 | 2026-11-16 | 35678 | active | [repo](https://github.com/binwiederhier/ntfy) |
| [obsidian](container/obsidian/) | Note-taking & Editors | Run Obsidian as a self-hosted web app. Not remote desktop, an actual web app. | v0.9.0+obsidian.1.14.4 | 2026-11-04 | 1677 | active | [repo](https://github.com/Nystik-gh/ignis) |
| [obsidian-sync](container/obsidian-sync/) | Note-taking & Editors | High-performance, low-latency note synchronization, online management, and remote REST API service platform. | 3.6.1 | 2026-08-14 | 2575 | active | [repo](https://github.com/haierkeys/fast-note-sync-service) |
| [opengist](container/opengist/) | Pastebins | Self-hosted pastebin powered by Git, open-source alternative to Github Gist. | v1.16.0 | 2026-11-21 | 3414 | active | [repo](https://github.com/thomiceli/opengist) |
| [openlist](container/openlist/) | File Transfer | A new AList Fork to Anti Trust Crisis | v4.2.10 | 2026-11-19 | 25741 | active | [repo](https://github.com/OpenListTeam/OpenList) |
| [pairdrop](container/pairdrop/) | File Transfer | PairDrop: Transfer Files Cross-Platform. No Setup, No Signup. | v1.11.2 | 2025-02-24 | 11864 | active | [repo](https://github.com/schlagmichdoch/PairDrop) |
| [portainer](container/portainer/) | Self-hosting Solutions | Making Docker and Kubernetes management easy. | 2.45.4 | 2026-11-18 | 38882 | active | [repo](https://github.com/portainer/portainer) |
| [prometheus](container/prometheus/) | Monitoring & Status Pages | The Prometheus monitoring system and time series database. | v3.16.0 | 2026-11-05 | 67750 | active | [repo](https://github.com/prometheus/prometheus) |
| [prometheus-exporter](container/prometheus-exporter/) | Monitoring & Status Pages | Exporter for machine metrics | v1.13.1 | 2026-11-19 | 13952 | active | [repo](https://github.com/prometheus/node_exporter) |
| [qbittorrent](container/qbittorrent/) | File Transfer |  | 5.3.0_v2.1.2-ls487 | 2026-11-08 | 1778 | active | [repo](https://github.com/linuxserver/docker-qbittorrent) |
| [reader](container/reader/) | Document Management |  | v0.1.0 | 2026-10-27 | 11110 | active | [repo](https://github.com/hectorqin/reader) |
| [rustdesk](container/rustdesk/) | Remote Access | An open-source remote desktop application designed for self-hosting, as an alternative to TeamViewer. | 1.5.3 | 2026-11-09 | 130639 | active | [repo](https://github.com/rustdesk/rustdesk) |
| [semaphore](container/semaphore/) | Software Development | Modern UI and powerful API for Ansible, Terraform/OpenTofu/Terragrunt, PowerShell and other DevOps tools. | v2.20.23 | 2026-11-10 | 14441 | active | [repo](https://github.com/semaphoreui/semaphore) |
| [shadowsocks-rust-client](container/shadowsocks-rust-client/) | Proxy | A Rust port of shadowsocks | v1.25.0 | 2026-08-26 | 10991 | active | [repo](https://github.com/shadowsocks/shadowsocks-rust) |
| [shadowsocks-rust-server](container/shadowsocks-rust-server/) | Proxy | A Rust port of shadowsocks | v1.25.0 | 2026-08-26 | 10991 | active | [repo](https://github.com/shadowsocks/shadowsocks-rust) |
| [snappymail](container/snappymail/) | Communication | Simple, modern & fast web-based email client | v3.1.4 | 2026-11-20 | 1762 | active | [repo](https://github.com/the-djmaze/snappymail) |
| [sslocal](container/sslocal/) | Proxy | A Rust port of shadowsocks | v1.25.0 | 2026-08-26 | 10991 | active | [repo](https://github.com/shadowsocks/shadowsocks-rust) |
| [steam-headless](container/steam-headless/) | Games | A Headless Steam Docker image supporting NVIDIA GPU and accessible via Web UI | 0.2.0 | 2023-09-29 | 5010 | active | [repo](https://github.com/Steam-Headless/docker-steam-headless) |
| [swingmusic](container/swingmusic/) | Media Streaming |  Swing Music is a beautiful, self-hosted music player for your local audio files. Like a cooler Spotify ... but bring your own music. | v3.0.0 | 2026-06-12 | 2095 | active | [repo](https://github.com/swingmx/swingmusic) |
| [syncthing](container/syncthing/) | File Transfer | Open Source Continuous File Synchronization | v2.1.7 | 2026-10-21 | 90678 | active | [repo](https://github.com/syncthing/syncthing) |
| [talebook](container/talebook/) | Document Management | 一个简单好用的个人书库 | v26.10.18 | 2026-10-18 | 5914 | active | [repo](https://github.com/talebook/talebook) |
| [taosync](container/taosync/) | File Transfer | TaoSync 是一个支持内置存储引擎与 OpenList/AList v3+ 的自动化同步工具/TaoSync is an automated synchronization tool with a built-in storage engine and support for OpenList/AList v3+. | v0.5.0 | 2026-11-17 | 1572 | active | [repo](https://github.com/dr34m-cn/taosync) |
| [transmission](container/transmission/) | File Transfer | Official Transmission BitTorrent client repository | 4.1.4 | 2026-10-24 | 15445 | active | [repo](https://github.com/transmission/transmission) |
| [upsnap](container/upsnap/) | Network Utilities | A simple wake on lan web app written with SvelteKit, Go and PocketBase. | 5.7.3 | 2026-11-06 | 6447 | active | [repo](https://github.com/seriousm4x/UpSnap) |
| [uptime-kuma](container/uptime-kuma/) | Monitoring & Status Pages | A fancy self-hosted monitoring tool | 2.5.12 | 2026-11-17 | 94741 | active | [repo](https://github.com/louislam/uptime-kuma) |
| [vaultwarden](container/vaultwarden/) | Password Managers | Unofficial Bitwarden compatible server written in Rust, formerly known as bitwarden_rs | 1.37.5 | 2026-11-01 | 71123 | active | [repo](https://github.com/dani-garcia/vaultwarden) |
| [zerotier](container/zerotier/) | VPN | A Smart Ethernet Switch for Earth | 1.16.2 | 2026-05-28 | 17279 | active | [repo](https://github.com/zerotier/ZeroTierOne) |

<!-- AUTO-GENERATED:END -->
