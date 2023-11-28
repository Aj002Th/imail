package manager

import (
	"errors"
	"github.com/Aj002Th/imail/common/config"
	"github.com/Aj002Th/imail/common/crontab"
	"github.com/Aj002Th/imail/server/catcher"
	"github.com/Aj002Th/imail/server/catcher/plugins/bilibiliVideo"
	"github.com/Aj002Th/imail/server/catcher/plugins/rssAdapter"
	"github.com/Aj002Th/imail/server/manager/dal/model"
	"github.com/Aj002Th/imail/server/manager/dal/query"
	"github.com/Aj002Th/imail/server/messager"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

type Manager struct {
	firstRun  bool // 标记是否是第一次运行, 用于优化体验
	Catchers  []catcher.Catcher
	Messagers []messager.Messager
}

func NewContentManager() *Manager {
	manager := &Manager{firstRun: true}

	// catcher init
	bilibiliVideoConfigs := config.GetBilibiliVideoTargets()
	if len(bilibiliVideoConfigs) != 0 {
		for _, cfg := range bilibiliVideoConfigs {
			manager.Catchers = append(manager.Catchers,
				bilibiliVideo.NewCatcher(
					cfg.Uid,
					cfg.Category,
				))
		}
	}
	rssAdapterConfigs := config.GetRssAdapterTargets()
	if len(rssAdapterConfigs) != 0 {
		for _, cfg := range rssAdapterConfigs {
			manager.Catchers = append(manager.Catchers,
				rssAdapter.NewCatcher(
					cfg.Url,
					cfg.Category,
					cfg.Source,
				))
		}
	}

	// messager init
	manager.Messagers = append(manager.Messagers, messager.NewEmailMessager())

	return manager
}

func (m *Manager) Run() {
	err := crontab.StartScheduledTasks(config.GetCronTab(), func() { m.CatchAndSend() })
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// 是否立即执行一次
	if config.IsImmediate() {
		m.CatchAndSend()
	}

	select {}
}

func (m *Manager) CatchAndSend() {
	contents := make([]catcher.Content, 0)

	// 获取所有爬虫爬取到的数据
	for _, c := range m.Catchers {
		contentBatch, err := c.Catch()
		if err != nil {
			slog.Error(err.Error())
			continue
		}
		contents = append(contents, contentBatch...)
	}

	// 插入数据库, 需要依据历史数据做一个去重操作
	for _, c := range contents {
		_, err := query.Content.FindBySourceAuthorLink(c.Source, c.Author, c.Link)
		// 找到了
		if err == nil {
			continue
		}
		// 出错
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Error(err.Error())
			continue
		}
		// 未找到
		_ = query.Content.Create(&model.Content{
			Content: catcher.Content{
				Title:       c.Title,
				Time:        c.Time,
				Description: c.Description,
				Cover:       c.Cover,
				Link:        c.Link,
				Author:      c.Author,
				Source:      c.Source,
				Category:    c.Category,
			},
			Sended: false,
		})
	}

	// 将所有未发送的数据进行发送
	contentToSend, err := query.Content.Where(query.Content.Sended.Is(false)).Find()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// 简单优化一下体验, 第一次启动时不会发送出巨量的历史内容
	// 第一次运行时, 最多只把 10 天以内的消息发送出去, 不全发
	if m.firstRun {
		m.firstRun = false
		newContentToSend := make([]*model.Content, 0)
		for _, c := range contentToSend {
			if c.Time.After(time.Now().AddDate(0, 0, -10)) {
				newContentToSend = append(newContentToSend, c)
			}
		}
		contentToSend = newContentToSend
	}

	// 依据配置, 判断是否忽略空消息
	if len(contentToSend) == 0 && config.IsIgnoreEmptyMessage() {
		return
	}

	// 推送消息
	for _, m := range m.Messagers {
		err := m.Push("订阅消息", convContentsToMessage(convContentModelsToContents(contentToSend)))
		if err != nil {
			slog.Error(err.Error())
			return
		}
	}

	// 全部标记为已发送
	_, err = query.Content.Where(query.Content.Sended.Is(false)).Update(query.Content.Sended, true)
	if err != nil {
		slog.Error(err.Error())
		return
	}
}
