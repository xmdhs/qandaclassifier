package classification

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/xmdhs/qandaclassifier/spider"
)

type Post struct {
	Charset   string        `json:"Charset"`
	Variables PostVariables `json:"Variables"`
	Version   string        `json:"Version"`
}

type PostVariables struct {
	Auth                  string                          `json:"auth"`
	CacheCustominfoPostno []string                        `json:"cache_custominfo_postno"`
	Cookiepre             string                          `json:"cookiepre"`
	Fid                   string                          `json:"fid"`
	Formhash              string                          `json:"formhash"`
	Forum                 PostVariablesForum              `json:"forum"`
	ForumThreadpay        string                          `json:"forum_threadpay"`
	Groupid               string                          `json:"groupid"`
	Ismoderator           string                          `json:"ismoderator"`
	MemberAvatar          string                          `json:"member_avatar"`
	MemberUID             string                          `json:"member_uid"`
	MemberUsername        string                          `json:"member_username"`
	Notice                PostVariablesNotice             `json:"notice"`
	Postlist              []PostVariablesPostlist         `json:"postlist"`
	Ppp                   string                          `json:"ppp"`
	Readaccess            string                          `json:"readaccess"`
	Saltkey               string                          `json:"saltkey"`
	SettingRewriterule    PostVariablesSettingRewriterule `json:"setting_rewriterule"`
	SettingRewritestatus  []string                        `json:"setting_rewritestatus"`
	Thread                PostVariablesThread             `json:"thread"`
}

type PostVariablesForum struct {
	Password string `json:"password"`
}

type PostVariablesNotice struct {
	Newmypost string `json:"newmypost"`
	Newpm     string `json:"newpm"`
	Newprompt string `json:"newprompt"`
	Newpush   string `json:"newpush"`
}

type PostVariablesPostlist struct {
	Adminid      string `json:"adminid"`
	Anonymous    string `json:"anonymous"`
	Attachment   string `json:"attachment"`
	Author       string `json:"author"`
	Authorid     string `json:"authorid"`
	Dateline     string `json:"dateline"`
	Dbdateline   string `json:"dbdateline"`
	First        string `json:"first"`
	Groupiconid  string `json:"groupiconid"`
	Groupid      string `json:"groupid"`
	Memberstatus string `json:"memberstatus"`
	Message      string `json:"message"`
	Number       string `json:"number"`
	Pid          string `json:"pid"`
	Position     string `json:"position"`
	Replycredit  string `json:"replycredit"`
	Status       string `json:"status"`
	Tid          string `json:"tid"`
	Username     string `json:"username"`
}

type PostVariablesSettingRewriterule struct {
	ForumArchiver     string `json:"forum_archiver"`
	ForumForumdisplay string `json:"forum_forumdisplay"`
	ForumViewthread   string `json:"forum_viewthread"`
	GroupGroup        string `json:"group_group"`
	HomeBlog          string `json:"home_blog"`
	HomeSpace         string `json:"home_space"`
	Plugin            string `json:"plugin"`
	PortalArticle     string `json:"portal_article"`
	PortalTopic       string `json:"portal_topic"`
}

type PostVariablesThread struct {
	Addviews        string                             `json:"addviews"`
	Allreplies      string                             `json:"allreplies"`
	Archiveid       string                             `json:"archiveid"`
	Attachment      string                             `json:"attachment"`
	Author          string                             `json:"author"`
	Authorid        string                             `json:"authorid"`
	Bgcolor         string                             `json:"bgcolor"`
	Closed          string                             `json:"closed"`
	Comments        string                             `json:"comments"`
	Cover           string                             `json:"cover"`
	Dateline        string                             `json:"dateline"`
	Digest          string                             `json:"digest"`
	Displayorder    string                             `json:"displayorder"`
	Favtimes        string                             `json:"favtimes"`
	Fid             string                             `json:"fid"`
	Heatlevel       string                             `json:"heatlevel"`
	Heats           string                             `json:"heats"`
	Hidden          string                             `json:"hidden"`
	Highlight       string                             `json:"highlight"`
	Icon            string                             `json:"icon"`
	IsArchived      string                             `json:"is_archived"`
	Isgroup         string                             `json:"isgroup"`
	Lastpost        string                             `json:"lastpost"`
	Lastposter      string                             `json:"lastposter"`
	Maxposition     string                             `json:"maxposition"`
	Moderated       string                             `json:"moderated"`
	Ordertype       string                             `json:"ordertype"`
	Posttable       string                             `json:"posttable"`
	Posttableid     string                             `json:"posttableid"`
	Price           string                             `json:"price"`
	Pushedaid       string                             `json:"pushedaid"`
	Rate            string                             `json:"rate"`
	Readperm        string                             `json:"readperm"`
	Recommend       string                             `json:"recommend"`
	RecommendAdd    string                             `json:"recommend_add"`
	RecommendSub    string                             `json:"recommend_sub"`
	Recommendlevel  string                             `json:"recommendlevel"`
	Recommends      string                             `json:"recommends"`
	Relatebytag     string                             `json:"relatebytag"`
	Relay           string                             `json:"relay"`
	Replies         string                             `json:"replies"`
	Replycredit     string                             `json:"replycredit"`
	ReplycreditRule PostVariablesThreadReplycreditRule `json:"replycredit_rule"`
	Sharetimes      string                             `json:"sharetimes"`
	ShortSubject    string                             `json:"short_subject"`
	Sortid          string                             `json:"sortid"`
	Special         string                             `json:"special"`
	Stamp           string                             `json:"stamp"`
	Status          string                             `json:"status"`
	Stickreply      string                             `json:"stickreply"`
	Subject         string                             `json:"subject"`
	Subjectenc      string                             `json:"subjectenc"`
	Threadtable     string                             `json:"threadtable"`
	Threadtableid   string                             `json:"threadtableid"`
	Tid             string                             `json:"tid"`
	Typeid          string                             `json:"typeid"`
	Views           string                             `json:"views"`
}

type PostVariablesThreadReplycreditRule struct {
	Extcreditstype string `json:"extcreditstype"`
}

var c = http.Client{Timeout: 5 * time.Second}

func getidinfo(tid string) (info, error) {
	var b []byte
	for {
		var err error
		b, err = func() ([]byte, error) {
			reqs, err := http.NewRequest("GET", `https://www.mcbbs.net/api/mobile/index.php?version=4&module=viewthread&tid=`+tid, nil)
			if err != nil {
				return nil, err
			}
			reqs.Header.Set("Accept", "*/*")
			reqs.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
			reqs.Header.Set("Cookie", spider.Cookie)
			rep, err := c.Do(reqs)
			if rep != nil {
				defer rep.Body.Close()
			}
			if err != nil {
				return nil, err
			}
			if rep.StatusCode != 200 {
				return nil, errors.New("not 200")
			}
			b, err = ioutil.ReadAll(rep.Body)
			if err != nil {
				return nil, err
			}
			return b, nil
		}()
		if err != nil {
			log.Println(tid, err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	var postinfo Post
	err := json.Unmarshal(b, &postinfo)
	if err != nil {
		return info{}, err
	}
	pinfo, err := getmessage(postinfo.Variables.Postlist)
	if err != nil {
		return info{}, err
	}
	return info{
		tid:     postinfo.Variables.Thread.Tid,
		message: pinfo,
	}, nil
}

var NoReply = errors.New(`No reply`)

func getmessage(postlist []PostVariablesPostlist) ([]pidinfo, error) {
	if len(postlist) == 0 {
		return nil, NoReply
	}
	p := make([]pidinfo, 0, len(postlist))
	for _, v := range postlist {
		message := v.Message
		re := regexp.MustCompile("\\<[\\S\\s]+?\\>")
		message = re.ReplaceAllString(message, "")
		message = strings.ReplaceAll(message, "&nbsp;", " ")
		p = append(p, pidinfo{
			message: message,
			pid:     v.Pid,
		})
		break
	}
	return p, nil
}

type info struct {
	message []pidinfo
	tid     string
}

type pidinfo struct {
	message string
	pid     string
}
