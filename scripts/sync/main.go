// Command sync refreshes upstream GitHub metadata for each project in
// config.json, stores the fetched info in containers.json, updates each
// project's README, and regenerates the root README table. It uses only the
// Go standard library.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const (
	startMarker  = "<!-- AUTO-GENERATED:START -->"
	endMarker    = "<!-- AUTO-GENERATED:END -->"
	apiBase      = "https://api.github.com"
	containerDir = "container" // project folders live under this directory
)

// Config is the human-editable configuration for one project. It lives in
// config.json and is never written by this program.
type Config struct {
	Repo       string `json:"repo"`
	ReadmePath string `json:"readme_path"`
	Type       string `json:"type"`
}

// Container is the upstream info fetched for one project. It lives in
// containers.json and is refreshed daily by this program.
type Container struct {
	Repo              string `json:"repo"`
	LatestTag         string `json:"latest_tag"`
	LatestTagDate     string `json:"latest_tag_date"`
	LatestRelease     string `json:"latest_release"`
	LatestReleaseDate string `json:"latest_release_date"`
	Status            string `json:"status"`
	Description       string `json:"description"`
	Stars             int    `json:"stars"`
}

type repoInfo struct {
	Archived      bool   `json:"archived"`
	Description   string `json:"description"`
	Stars         int    `json:"stargazers_count"`
	DefaultBranch string `json:"default_branch"`
}

type releaseInfo struct {
	TagName     string `json:"tag_name"`
	PublishedAt string `json:"published_at"`
}

type tagInfo struct {
	Name   string `json:"name"`
	Commit struct {
		SHA string `json:"sha"`
	} `json:"commit"`
}

type commitInfo struct {
	Commit struct {
		Author struct {
			Date string `json:"date"`
		} `json:"author"`
	} `json:"commit"`
}

func main() {
	log.SetFlags(0)

	renderOnly := flag.Bool("render", false, "only regenerate README.md from containers.json (no network)")
	flag.Parse()

	root, err := findRoot()
	if err != nil {
		log.Fatalf("find repo root: %v", err)
	}

	configPath := filepath.Join(root, "config.json")
	configs, err := loadJSON[map[string]Config](configPath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	containersPath := filepath.Join(root, "containers.json")
	containers, err := loadJSON[map[string]Container](containersPath)
	if err != nil {
		log.Printf("load containers.json: %v (starting empty)", err)
		containers = map[string]Container{}
	}

	keys := make([]string, 0, len(configs))
	for k := range configs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	if *renderOnly {
		if err := writeRootREADME(filepath.Join(root, "README.md"), configs, containers, keys); err != nil {
			log.Fatalf("write README: %v", err)
		}
		log.Println("README rendered from containers.json")
		return
	}

	token := os.Getenv("GITHUB_TOKEN")
	client := &http.Client{Timeout: 30 * time.Second}

	for _, folder := range keys {
		cfg := configs[folder]
		c := containers[folder]
		// The repo in containers.json always mirrors the config.
		c.Repo = cfg.Repo

		owner, repo := parseRepo(cfg.Repo)
		if owner == "" || repo == "" {
			log.Printf("[%s] skip: cannot parse repo %q", folder, cfg.Repo)
			continue
		}
		log.Printf("[%s] syncing %s/%s", folder, owner, repo)

		status, info, err := getRepo(client, token, owner, repo)
		if err != nil {
			log.Printf("[%s] repo lookup failed: %v (keeping old values)", folder, err)
			continue
		}
		switch status {
		case http.StatusNotFound:
			log.Printf("[%s] repo not found -> deleted", folder)
			c.Status = "deleted"
			containers[folder] = c
			continue
		case http.StatusOK:
			c.Status = "active"
			if info.Archived {
				c.Status = "archived"
			}
			c.Description = info.Description
			c.Stars = info.Stars
		default:
			log.Printf("[%s] repo lookup HTTP %d (keeping old values)", folder, status)
			continue
		}

		syncRelease(client, token, owner, repo, folder, &c)
		syncLatestTag(client, token, owner, repo, folder, &c)
		syncReadme(client, token, owner, repo, info.DefaultBranch, cfg.ReadmePath, folder, root)
		containers[folder] = c
	}

	if err := saveJSON(containersPath, containers); err != nil {
		log.Fatalf("save containers: %v", err)
	}
	if err := writeRootREADME(filepath.Join(root, "README.md"), configs, containers, keys); err != nil {
		log.Fatalf("write README: %v", err)
	}
	log.Println("done")
}

// findRoot walks up from the working directory until it finds config.json so
// the program works both from the repo root and from scripts/sync.
func findRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "config.json")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("config.json not found")
		}
		dir = parent
	}
}

func loadJSON[T any](path string) (T, error) {
	var value T
	data, err := os.ReadFile(path)
	if err != nil {
		return value, err
	}
	if err := json.Unmarshal(data, &value); err != nil {
		return value, err
	}
	return value, nil
}

func saveJSON(path string, value any) error {
	// encoding/json sorts map keys, so writing the whole map keeps a stable
	// alphabetical order that matches the README table.
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}

func parseRepo(raw string) (owner, repo string) {
	s := strings.TrimSpace(raw)
	s = strings.TrimPrefix(s, "https://github.com/")
	s = strings.TrimPrefix(s, "http://github.com/")
	s = strings.TrimPrefix(s, "git@github.com:")
	s = strings.TrimSuffix(s, ".git")
	s = strings.TrimSuffix(s, "/")
	parts := strings.Split(s, "/")
	if len(parts) < 2 {
		return "", ""
	}
	return parts[0], parts[1]
}

func newRequest(method, url, token string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("User-Agent", "warehouse-sync")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return req, nil
}

func doJSON(client *http.Client, token, url string, target any) (int, error) {
	req, err := newRequest("GET", url, token, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		io.Copy(io.Discard, resp.Body)
		return resp.StatusCode, nil
	}
	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return resp.StatusCode, err
	}
	return resp.StatusCode, nil
}

func getRepo(client *http.Client, token, owner, repo string) (int, repoInfo, error) {
	var info repoInfo
	status, err := doJSON(client, token,
		fmt.Sprintf("%s/repos/%s/%s", apiBase, owner, repo), &info)
	return status, info, err
}

func syncRelease(client *http.Client, token, owner, repo, folder string, c *Container) {
	var rel releaseInfo
	status, err := doJSON(client, token,
		fmt.Sprintf("%s/repos/%s/%s/releases/latest", apiBase, owner, repo), &rel)
	if err != nil {
		log.Printf("[%s] release lookup failed: %v", folder, err)
		return
	}
	switch status {
	case http.StatusOK:
		c.LatestRelease = rel.TagName
		c.LatestReleaseDate = shortDate(rel.PublishedAt)
	case http.StatusNotFound:
		c.LatestRelease = ""
		c.LatestReleaseDate = ""
	default:
		log.Printf("[%s] release lookup HTTP %d", folder, status)
	}
}

func syncLatestTag(client *http.Client, token, owner, repo, folder string, c *Container) {
	var tags []tagInfo
	status, err := doJSON(client, token,
		fmt.Sprintf("%s/repos/%s/%s/tags?per_page=1", apiBase, owner, repo), &tags)
	if err != nil {
		log.Printf("[%s] tag lookup failed: %v", folder, err)
		return
	}
	if status != http.StatusOK {
		log.Printf("[%s] tag lookup HTTP %d", folder, status)
		return
	}
	if len(tags) == 0 {
		c.LatestTag = ""
		c.LatestTagDate = ""
		return
	}
	c.LatestTag = tags[0].Name
	if tags[0].Commit.SHA == "" {
		c.LatestTagDate = ""
		return
	}
	var cm commitInfo
	st, err := doJSON(client, token,
		fmt.Sprintf("%s/repos/%s/%s/commits/%s", apiBase, owner, repo, tags[0].Commit.SHA), &cm)
	if err != nil || st != http.StatusOK {
		log.Printf("[%s] tag commit lookup failed (status %d, err %v)", folder, st, err)
		c.LatestTagDate = ""
		return
	}
	c.LatestTagDate = shortDate(cm.Commit.Author.Date)
}

func syncReadme(client *http.Client, token, owner, repo, branch, readmePath, folder, root string) {
	if branch == "" || readmePath == "" {
		return
	}
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s",
		owner, repo, branch, readmePath)
	req, err := newRequest("GET", url, token, nil)
	if err != nil {
		log.Printf("[%s] readme request: %v", folder, err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[%s] readme download failed: %v", folder, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("[%s] readme download HTTP %d (keeping local copy)", folder, resp.StatusCode)
		return
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[%s] readme read failed: %v", folder, err)
		return
	}
	dest := filepath.Join(root, containerDir, folder, filepath.Base(readmePath))
	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		log.Printf("[%s] readme dir create failed: %v", folder, err)
		return
	}
	if err := os.WriteFile(dest, data, 0644); err != nil {
		log.Printf("[%s] readme write failed: %v", folder, err)
	}
}

func shortDate(rfc3339 string) string {
	if len(rfc3339) >= 10 {
		return rfc3339[:10]
	}
	return rfc3339
}

// latestVersion prefers the release tag, falling back to the latest tag.
func latestVersion(c Container) string {
	if c.LatestRelease != "" {
		return c.LatestRelease
	}
	return c.LatestTag
}

// lastUpdate prefers the release date, falling back to the tag date.
func lastUpdate(c Container) string {
	if c.LatestReleaseDate != "" {
		return c.LatestReleaseDate
	}
	return c.LatestTagDate
}

func writeRootREADME(path string, configs map[string]Config, containers map[string]Container, keys []string) error {
	var b strings.Builder
	b.WriteString("# Docker Compose 项目仓库\n\n")
	b.WriteString("本仓库收集了本地 NAS 上使用的 Docker 项目部署文件，统一放在 `container/` 目录下。每个项目目录包含：\n\n")
	b.WriteString("- 上游项目的 `README.md`（由 GitHub Action 每日同步）\n")
	b.WriteString("- 本地部署用的 `compose.yaml` 与 `.env`\n\n")
	b.WriteString("`config.json` 保存人工维护的配置：`repo`（原仓库地址）、`readme_path`（原仓库内 README 相对路径）与 `type`（项目类型，参考 awesome-selfhosted 的一级分类）。")
	b.WriteString("抓取到的项目信息保存在 `containers.json` 中，由每日定时任务自动更新。")
	b.WriteString("状态取值为 `active`、`archived` 或 `deleted`。\n\n")
	b.WriteString("表格中的“最新版本”优先取 `latest_release`，为空时取 `latest_tag`；“最后更新”优先取 `latest_release_date`，为空时取 `latest_tag_date`。\n\n")

	table := renderTable(configs, containers, keys)

	existing, err := os.ReadFile(path)
	if err == nil {
		text := string(existing)
		start := strings.Index(text, startMarker)
		end := strings.Index(text, endMarker)
		if start >= 0 && end > start {
			prefix := text[:start+len(startMarker)]
			suffix := text[end:]
			out := prefix + "\n\n" + table + "\n\n" + suffix
			return os.WriteFile(path, []byte(out), 0644)
		}
	}

	b.WriteString(startMarker)
	b.WriteString("\n\n")
	b.WriteString(table)
	b.WriteString("\n\n")
	b.WriteString(endMarker)
	b.WriteString("\n")
	return os.WriteFile(path, []byte(b.String()), 0644)
}

func renderTable(configs map[string]Config, containers map[string]Container, keys []string) string {
	var b strings.Builder
	b.WriteString("| 项目 | 类型 | 描述 | 最新版本 | 最后更新 | Stars | 状态 | 仓库 |\n")
	b.WriteString("| --- | --- | --- | --- | --- | --- | --- | --- |\n")
	for _, k := range keys {
		cfg := configs[k]
		c := containers[k]
		// Project name links to the local folder; a separate Repo column
		// points to the upstream repository.
		name := fmt.Sprintf("[%s](%s/%s/)", k, containerDir, k)
		repo := fmt.Sprintf("[repo](%s)", cfg.Repo)
		b.WriteString("| ")
		b.WriteString(name)
		b.WriteString(" | ")
		b.WriteString(mdCell(cfg.Type))
		b.WriteString(" | ")
		b.WriteString(mdCell(c.Description))
		b.WriteString(" | ")
		b.WriteString(mdCell(latestVersion(c)))
		b.WriteString(" | ")
		b.WriteString(mdCell(lastUpdate(c)))
		b.WriteString(" | ")
		fmt.Fprintf(&b, "%d", c.Stars)
		b.WriteString(" | ")
		b.WriteString(mdCell(c.Status))
		b.WriteString(" | ")
		b.WriteString(repo)
		b.WriteString(" |\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

// mdCell escapes characters that would break a Markdown table cell.
func mdCell(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\r", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	return s
}
