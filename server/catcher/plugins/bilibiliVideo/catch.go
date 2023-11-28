package bilibiliVideo

import (
	"fmt"
	"github.com/Aj002Th/imail/common/config"
	"github.com/Aj002Th/imail/server/catcher"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

/*
## bilibiliVideo
获取 up 主的视频更新情况

## 注意事项
只爬取了视频列表第一页(包含最多30个视频), 如果一个爬取间隔内更新超过30个视频, 则会有更新信息被遗漏
通常情况下也不会有这么勤奋的up主吧 :)
*/

func init() {
	launcher.NewBrowser().MustGet()
}

type Catcher struct {
	Uid      string
	Uname    string
	Url      string
	Category string
}

func NewCatcher(uid string, category string) *Catcher {
	url := fmt.Sprintf("https://space.bilibili.com/%s/video?tid=0&page=1&keyword=&order=pubdate", uid)

	// 通过 uid 爬取 username
	launcherUrl := launcher.New().NoSandbox(true).MustLaunch()
	browser := rod.New().Timeout(time.Minute).ControlURL(launcherUrl).MustConnect()
	defer browser.MustClose()
	page := stealth.MustPage(browser)
	page.MustNavigate(url)
	time.Sleep(time.Second * 2)
	uname := page.MustElement("#h-name").MustText()

	return &Catcher{
		Uid:      uid,
		Uname:    uname,
		Url:      url,
		Category: category,
	}
}

func (c *Catcher) Catch() ([]catcher.Content, error) {
	targets, err := c.getData()
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	// 格式转换
	contents := make([]catcher.Content, 0)
	for _, target := range targets {
		contents = append(contents, c.convTargetToContent(target))
	}

	return contents, nil
}

// Target 插件要爬取的目标信息
type Target struct {
	Title string
	Cover string
	Time  time.Time
	Url   string
}

func (c *Catcher) getData() ([]Target, error) {
	launcherUrl := launcher.New().NoSandbox(true).MustLaunch()
	browser := rod.New().Timeout(time.Minute).ControlURL(launcherUrl).MustConnect()
	defer browser.MustClose()
	page := stealth.MustPage(browser)
	page.MustNavigate(c.Url)
	time.Sleep(time.Second * 5)
	vList := page.MustElement("#submit-video-list > ul.list-list")
	videos, err := vList.Elements("li")
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	targets := make([]Target, 0)
	for _, v := range videos {
		title := v.MustElement("div > div.title-row > a").MustAttribute("title")
		url := v.MustElement("div > div.title-row > a").MustAttribute("href")
		cover := v.MustElement("a > div.b-img > picture > img").MustAttribute("src")
		timeStr := v.MustElement("div > div.meta.clearfix > span.time").MustText()
		timeStr = strings.TrimSuffix(timeStr, "\\n")
		timeStr = strings.TrimSpace(timeStr)

		// 时间格式解析
		var publishTime time.Time
		if strings.Contains(timeStr, "小时前") {
			afterHoursStr := strings.TrimSuffix(timeStr, "小时前")
			afterHours, _ := strconv.Atoi(afterHoursStr)
			timeNow := time.Now().Add(time.Hour * time.Duration(-afterHours))
			publishTime = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, time.Local)
		} else if timeStr == "昨天" {
			timeNow := time.Now()
			publishTime = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, time.Local)
		} else {
			strs := strings.Split(timeStr, "-")
			mouth, _ := strconv.Atoi(strs[0])
			day, _ := strconv.Atoi(strs[1])
			publishTime = time.Date(time.Now().Year(), time.Month(mouth), day, 0, 0, 0, 0, time.Local)
		}

		t := Target{
			Title: *title,
			Cover: "https:" + *cover,
			Url:   "https:" + *url,
			Time:  publishTime,
		}

		targets = append(targets, t)
	}

	return targets, nil
}

func (c *Catcher) convTargetToContent(t Target) catcher.Content {
	//slog.Info(fmt.Sprintf(
	//	"bilibiliVideoCatcher: getVideo - [title]%s [url]%s [cover]%s [time]%s",
	//	t.Title,
	//	t.Url,
	//	t.Cover,
	//	t.Time.Format("2006-01-02")))

	return catcher.Content{
		Title:       t.Title,
		Time:        t.Time,
		Description: "",
		Cover:       t.Cover,
		Link:        t.Url,
		Author:      c.Uname,
		Source:      config.GetBilibiliVideoSource(),
		Category:    c.Category,
	}
}
