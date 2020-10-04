package spider

import (
	"fmt"
	"testing"
)

func Test_jsondecoder(t *testing.T) {
	b := []byte(jsondata)
	f, err := jsondecoder(b)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range f.Variables.ForumThreadlist {
		fmt.Println(v.Tid)
	}
}

const jsondata = `{
    "Version": "4",
    "Charset": "UTF-8",
    "Variables": {
        "cookiepre": "ZxYQ_8cea_",
        "auth": "c63fxMRVRJnxhLwgfGUqff5atUiZZijwKLsHYlRT8ucPYr0T7BNTUJ7Qt0MEQBrdY\/3BwGxp6akn19mz8Moi1y8WWgQq",
        "saltkey": "uVvZi0hd",
        "member_uid": "1770442",
        "member_username": "xmdhs",
        "member_avatar": "https:\/\/www.mcbbs.net\/uc_server\/avatar.php?uid=1770442&size=small&ts=1",
        "groupid": "29",
        "formhash": "e37c3797",
        "ismoderator": "0",
        "readaccess": "175",
        "notice": {
            "newpush": "0",
            "newpm": "0",
            "newprompt": "0",
            "newmypost": "0"
        },
        "forum": {
            "fid": "110",
            "description": "\u2517 <a href=\"https:\/\/www.mcbbs.net\/thread-516637-1-1.html\" target=\"_blank\"><font color=\"red\">\u5e38\u89c1\u95ee\u9898<\/font><\/a>",
            "icon": "https:\/\/attachment.mcbbs.net\/common\/5f\/common_110_icon.png",
            "rules": "<div align=\"center\"><img id=\"aimg_vYj8J\"  class=\"zoom\" width=\"700\" height=\"350\" src=\"https:\/\/attachment.mcbbs.net\/forum\/201904\/28\/012021w3zgtk226sm2k62k.jpg\" border=\"0\" alt=\"\" \/><br \/>\r\n<img id=\"aimg_UxvD1\"  class=\"zoom\" src=\"https:\/\/miao.su\/images\/2018\/08\/09\/01eeb53.png\"   border=\"0\" alt=\"\" \/><br \/>\r\n<a href=\"https:\/\/www.mcbbs.net\/forum.php?gid=364\" target=\"_blank\"><img id=\"aimg_F9Cfn\"  class=\"zoom\" src=\"https:\/\/miao.su\/images\/2018\/08\/09\/_01a4a53.png\"   border=\"0\" alt=\"\" \/><\/a><a href=\"https:\/\/www.mcbbs.net\/thread-812925-1-1.html\" target=\"_blank\"><img id=\"aimg_i1c14\"  class=\"zoom\" src=\"https:\/\/miao.su\/images\/2018\/08\/13\/_02e6453.png\"   border=\"0\" alt=\"\" \/><\/a><img id=\"aimg_TkFe6\"  class=\"zoom\" src=\"https:\/\/miao.su\/images\/2018\/08\/09\/_03d592b.png\"   border=\"0\" alt=\"\" \/><img id=\"aimg_UpPC3\"  class=\"zoom\" src=\"https:\/\/api.mcxk.net:3003\/image\/qanda\/header.png\"   border=\"0\" alt=\"\" \/><\/div><br \/>\r\n<a href=\"https:\/\/www.mcbbs.net\/thread-516637-1-1.html\" target=\"_blank\"><img id=\"aimg_t9NC8\"  class=\"zoom\" src=\"https:\/\/miao.su\/images\/2018\/08\/09\/_01079c7.png\"   border=\"0\" alt=\"\" \/><\/a><br \/>\r\n<a href=\"https:\/\/www.mcbbs.net\/thread-256348-1-1.html\" target=\"_blank\"><img id=\"aimg_mWX78\"  class=\"zoom\" src=\"https:\/\/attachment.mcbbs.net\/forum\/202003\/20\/134310j8f11ndx1fcfafjx.png\"   border=\"0\" alt=\"\" \/><\/a><br \/>\r\n<a href=\"https:\/\/www.mcbbs.net\/thread-256348-1-1.html\" target=\"_blank\"><img id=\"aimg_Um7m0\"  class=\"zoom\" src=\"https:\/\/attachment.mcbbs.net\/forum\/202003\/20\/134308vvkbdbe4kywyyoii.png\"   border=\"0\" alt=\"\" \/><\/a><br \/>\r\n<a href=\"https:\/\/www.mcbbs.net\/thread-256127-1-1.html\" target=\"_blank\"><img id=\"aimg_TXNS3\"  class=\"zoom\" src=\"https:\/\/attachment.mcbbs.net\/forum\/202003\/20\/134302x5t5ophxa5xpxgp4.png\"   border=\"0\" alt=\"\" \/><\/a><br \/>\r\n<img id=\"aimg_xiY0W\"  class=\"zoom\" src=\"https:\/\/miao.su\/images\/2018\/08\/08\/00d64ec.png\"   border=\"0\" alt=\"\" \/><br \/>\r\n<a href=\"https:\/\/www.mcbbs.net\/forum.php?mod=post&amp;action=newthread&amp;fid=110\" target=\"_blank\"><img id=\"aimg_Kpoyc\"  class=\"zoom\" src=\"https:\/\/miao.su\/images\/2018\/08\/08\/01e146d.png\"   border=\"0\" alt=\"\" \/><\/a><a href=\"https:\/\/www.mcbbs.net\/forum.php?mod=post&amp;action=newthread&amp;fid=431\" target=\"_blank\"><img id=\"aimg_x37mW\"  class=\"zoom\" src=\"https:\/\/miao.su\/images\/2018\/08\/08\/02a17cd.png\"   border=\"0\" alt=\"\" \/><\/a><a href=\"https:\/\/www.mcbbs.net\/forum.php?mod=post&amp;action=newthread&amp;fid=266\" target=\"_blank\"><img id=\"aimg_cssz9\"  class=\"zoom\" src=\"https:\/\/miao.su\/images\/2018\/08\/08\/032e9b6.png\"   border=\"0\" alt=\"\" \/><\/a><a href=\"https:\/\/www.mcbbs.net\/forum.php?mod=post&amp;action=newthread&amp;fid=1566\" target=\"_blank\"><img id=\"aimg_k6F0j\"  class=\"zoom\" src=\"https:\/\/miao.su\/images\/2018\/08\/08\/04027e1.png\"   border=\"0\" alt=\"\" \/><\/a><a href=\"https:\/\/www.mcbbs.net\/forum.php?mod=post&amp;action=newthread&amp;fid=265\" target=\"_blank\"><img id=\"aimg_WI7EE\"  class=\"zoom\" src=\"https:\/\/miao.su\/images\/2018\/08\/08\/05ae356.png\"   border=\"0\" alt=\"\" \/><\/a><br \/>\r\n\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014\u2014<br \/>\r\n<strong><a href=\"https:\/\/www.mcbbs.net\/forum-modqanda-1.html\" target=\"_blank\">Mod\u95ee\u7b54<\/a>&nbsp; &nbsp; <a href=\"https:\/\/www.mcbbs.net\/forum-multiqanda-1.html\" target=\"_blank\">\u8054\u673a\u95ee\u7b54<\/a>&nbsp; &nbsp; <a href=\"https:\/\/www.mcbbs.net\/forum-1566-1.html\" target=\"_blank\">\u5468\u8fb9\u95ee\u7b54<\/a>&nbsp; &nbsp; <a href=\"https:\/\/www.mcbbs.net\/forum-peqanda-1.html\" target=\"_blank\">\u57fa\u5ca9\u7248\uff08PE\uff09\u95ee\u7b54<\/a><\/strong>",
            "picstyle": "0",
            "fup": "364",
            "name": "\u539f\u7248\u95ee\u7b54",
            "threads": "55666",
            "posts": "422687",
            "autoclose": "0",
            "threadcount": "55666",
            "password": "0"
        },
        "group": {
            "groupid": "29",
            "grouptitle": "Lv.12 \u5c60\u9f99\u8005"
        },
        "forum_threadlist": [
            {
                "tid": "256127",
                "typeid": "1431",
                "readperm": "0",
                "price": "0",
                "author": "john180",
                "authorid": "5367",
                "subject": "\u6700\u4f73\u7b54\u6848\u7533\u8bf7\u5e16",
                "dateline": "2014-3-27",
                "lastpost": "1970-1-1 08:00",
                "views": "53",
                "replies": "-",
                "displayorder": "2",
                "digest": "0",
                "special": "0",
                "attachment": "0",
                "recommend_add": "0",
                "replycredit": "0",
                "dbdateline": "1395928558",
                "dblastpost": "0",
                "rushreply": "0",
                "recommend": "0"
            },
            {
                "tid": "1124574",
                "typeid": "123",
                "readperm": "0",
                "price": "-30",
                "author": "\u6d45\u6d6e\u5982\u751f",
                "authorid": "3664699",
                "subject": "\u5173\u4e8e\u6211\u4e0b\u8f7d\u65f6\u9047\u5230\u7684\u95ee\u9898",
                "dateline": "1&nbsp;\u5c0f\u65f6\u524d",
                "lastpost": "1&nbsp;\u5c0f\u65f6\u524d",
                "lastposter": "\u6d45\u6d6e\u5982\u751f",
                "views": "38",
                "replies": "3",
                "displayorder": "0",
                "digest": "0",
                "special": "3",
                "attachment": "2",
                "recommend_add": "0",
                "replycredit": "0",
                "dbdateline": "1601792149",
                "dblastpost": "1601792566",
                "rushreply": "0",
                "reply": [
                    {
                        "pid": "20021527",
                        "author": "spg233",
                        "authorid": "3605455",
                        "message": "\u53ef\u80fd\u662f\u4f60\u7f51\u7edc\u95ee\u9898\uff0c\u5c1d\u8bd5\u66f4\u6362\u4e0b\u8f7d\u6e90\u8bd5\u8bd5\u52a0\u4e2aqq\u6211\u5e2e\u4f60\u8fdc\u7a0b\u770b\u770b"
                    },
                    {
                        "pid": "20021537",
                        "author": "\u6d45\u6d6e\u5982\u751f",
                        "authorid": "3664699",
                        "message": "2369177079"
                    },
                    {
                        "pid": "20021545",
                        "author": "\u6d45\u6d6e\u5982\u751f",
                        "authorid": "3664699",
                        "message": "\u600e\u4e48\u66f4\u6362\u4e0b\u8f7d\u6e90\u5462"
                    }
                ],
                "recommend": "0"
            },
            {
                "tid": "1124573",
                "typeid": "123",
                "readperm": "0",
                "price": "-30",
                "author": "XueJie18",
                "authorid": "2724542",
                "subject": "\u670d\u52a1\u5668\u4e00\u76f4\u662f\u68c0\u6d4b\u4e2d\uff0c\u6c42\u89e3",
                "dateline": "1&nbsp;\u5c0f\u65f6\u524d",
                "lastpost": "1&nbsp;\u5c0f\u65f6\u524d",
                "lastposter": "XueJie18",
                "views": "20",
                "replies": "3",
                "displayorder": "0",
                "digest": "0",
                "special": "3",
                "attachment": "2",
                "recommend_add": "0",
                "replycredit": "0",
                "dbdateline": "1601791865",
                "dblastpost": "1601793570",
                "rushreply": "0",
                "reply": [
                    {
                        "pid": "20021489",
                        "author": "DreamVoid",
                        "authorid": "2999647",
                        "message": "\u89c1\u6b64\u94fe\u63a5\uff1ahttps:\/\/www.mcbbs.net\/thread-1124030-1-1.html\u5982\u679c\u65e0\u6548\uff0c\u968f\u4fbf\u8fdb\u4e00\u4e2a\u5355\u4eba\u5b58\u6863\uff0c\u7136\u540e\u9000\u5230\u4e3b\u754c\u9762 ..."
                    },
                    {
                        "pid": "20021534",
                        "author": "spg233",
                        "authorid": "3605455",
                        "message": "\u6309\u7167\u8fd9\u4e2a\u6559\u7a0b\u505a\uff1ahttps:\/\/cowtransfer.com\/s\/ce16a5bc6e2840\u5148\u505a\u6539host\u65b9\u6cd5"
                    },
                    {
                        "pid": "20021708",
                        "author": "XueJie18",
                        "authorid": "2724542",
                        "message": "\u611f\u8c22\u5927\u5bb6\uff0c\u95ee\u9898\u89e3\u51b3\u4e86"
                    }
                ],
                "recommend": "0"
            },
            {
                "tid": "1124537",
                "typeid": "108",
                "readperm": "0",
                "price": "30",
                "author": "l\u5c0f\u9b42l",
                "authorid": "2505673",
                "subject": "\u6709\u6ca1\u6709\u7c7b\u4f3c\u7684\u6811\u5c4b\u6559\u7a0b\u6216\u8005\u5730\u56fe\uff1f",
                "dateline": "3&nbsp;\u5c0f\u65f6\u524d",
                "lastpost": "1&nbsp;\u5c0f\u65f6\u524d",
                "lastposter": "\u731609",
                "views": "33",
                "replies": "3",
                "displayorder": "0",
                "digest": "0",
                "special": "3",
                "attachment": "0",
                "recommend_add": "0",
                "replycredit": "0",
                "dbdateline": "1601786969",
                "dblastpost": "1601793988",
                "rushreply": "0",
                "reply": [
                    {
                        "pid": "20020970",
                        "author": "\u75be\u98ce\u6697\u5f71",
                        "authorid": "3010482",
                        "message": "\u6211\u7684\u4e16\u754c\uff1a\u65b0\u624b\u4f2a\u6811\u5c4b\u5efa\u7b51\u6559\u5b66https:\/\/www.bilibili.com\/video\/BV127411C7bA?from=search&seid=17859860603 ..."
                    },
                    {
                        "pid": "20021071",
                        "author": "Jeansou",
                        "authorid": "2422915",
                        "message": "Jungle Treehouse https:\/\/www.mcbbs.net\/thread-902829-1-1.html(\u51fa\u5904: Minecraft(\u6211\u7684\u4e16\u754c)\u4e2d\u6587\u8bba\u575b) ..."
                    },
                    {
                        "pid": "20021801",
                        "author": "\u731609",
                        "authorid": "2911807",
                        "message": "https:\/\/www.mcbbs.net\/forum.php?mod=viewthread&tid=449175https:\/\/www.mcbbs.net\/forum.php?mod=viewthr ..."
                    }
                ],
                "recommend": "0"
            },
            {
                "tid": "1124518",
                "typeid": "2482",
                "readperm": "0",
                "price": "-30",
                "author": "\u67d2\u6708_\u67d2\u7231\u604b",
                "authorid": "3581927",
                "subject": "\u5982\u4f55\u8ba9\u73a9\u5bb6\u8e29\u5230\u65b9\u5757\u6267\u884c\u547d\u4ee4",
                "dateline": "3&nbsp;\u5c0f\u65f6\u524d",
                "lastpost": "1&nbsp;\u5c0f\u65f6\u524d",
                "lastposter": "\u67d2\u6708_\u67d2\u7231\u604b",
                "views": "23",
                "replies": "8",
                "displayorder": "0",
                "digest": "0",
                "special": "3",
                "attachment": "2",
                "recommend_add": "0",
                "replycredit": "0",
                "dbdateline": "1601784359",
                "dblastpost": "1601790937",
                "rushreply": "0",
                "reply": [
                    {
                        "pid": "20020580",
                        "author": "\u5929\u4f51\u9171",
                        "authorid": "1811016",
                        "message": "\u95ea\u7535\u53ea\u662f\u4e00\u4e2a\u6548\u679c\uff0c\u751f\u6210\u95ea\u7535\u540e\u4f1a\u7acb\u5373\u4f24\u5bb3\u73a9\u5bb6\u5e76\u6d88\u5931\uff0c\u56e0\u6b21\u65e0\u6cd5tp\u5230\u73a9\u5bb6\u8eab\u4e0a\uff0c\u53ea\u80fd\u5728\u73a9\u5bb6\u4f4d\u7f6e\u53ec\u5524 ..."
                    },
                    {
                        "pid": "20020582",
                        "author": "\u67d2\u6708_\u67d2\u7231\u604b",
                        "authorid": "3581927",
                        "message": "\u597d\u7684\u8c22\u8c22"
                    },
                    {
                        "pid": "20021334",
                        "author": "\u67d2\u6708_\u67d2\u7231\u604b",
                        "authorid": "3581927",
                        "message": "\u55ef\u5462\uff0c\u597d\u7684"
                    }
                ],
                "recommend": "0"
            },
            {
                "tid": "1124478",
                "typeid": "123",
                "readperm": "0",
                "price": "30",
                "author": "814580834",
                "authorid": "3531589",
                "subject": "\u5730\u6bef\u4e0a\u5237\u4e0d\u5237\u602a",
                "dateline": "5&nbsp;\u5c0f\u65f6\u524d",
                "lastpost": "\u534a\u5c0f\u65f6\u524d",
                "lastposter": "xyh200468",
                "views": "54",
                "replies": "8",
                "displayorder": "0",
                "digest": "0",
                "special": "3",
                "attachment": "0",
                "recommend_add": "0",
                "replycredit": "0",
                "dbdateline": "1601779187",
                "dblastpost": "1601794464",
                "rushreply": "0",
                "reply": [
                    {
                        "pid": "20019761",
                        "author": "\u731609",
                        "authorid": "2911807",
                        "message": "\u6839\u636ewiki\u751f\u7269\u78b0\u649e\u7bb1\u4e0d\u80fd\u88ab\u5176\u4ed6\u56fa\u4f53\u65b9\u5757\u3001\u751f\u7269\u548c\u6d41\u4f53\u6240\u963b\u6321\u3002\u4e00\u822c\u78b0\u649e\u7bb1\u662f2\u683c\u653e\u4e86\u5730\u6bef\u540e\u5c0f\u4e8e2\u683c\u56e0\u6b64\u4e0d\u4f1a\u751f\u6210 ..."
                    },
                    {
                        "pid": "20020276",
                        "author": "Minecraft.Wnxi",
                        "authorid": "3309514",
                        "message": "\u5730\u6bef\u4e0a\u672c\u6765\u5c31\u4e0d\u80fd\u5237\u602a"
                    },
                    {
                        "pid": "20021859",
                        "author": "xyh200468",
                        "authorid": "1449891",
                        "message": "\u5730\u6bef\u4e0d\u5237\u602a \u7136\u540e\u5728\u6d3b\u677f\u95e8\u8fd9\u7c7b\u6bd4\u534a\u7816\u77ee\u4e00\u70b9\u7684\u65b9\u5757\u653e\u5728\u4e24\u683c\u9ad8\u7684\u4e0a\u90e8 \u5c31\u4f1a\u53ea\u5237\u82e6\u529b\u6015 \u8718\u86db \u9ab7\u9ac5\u548c\u50f5\u5c38\u5c31\u4e0d\u4f1a\u751f ..."
                    }
                ],
                "recommend": "0"
            },
            {
                "tid": "1124413",
                "typeid": "123",
                "readperm": "0",
                "price": "30",
                "author": "\u50ac\u66f4\u66f4",
                "authorid": "2611011",
                "subject": "win10 \u57fa\u5ca9\u7248 JAVA \u8fd9\u4e9b\u6709\u4ec0\u4e48\u533a\u522b\u561b",
                "dateline": "9&nbsp;\u5c0f\u65f6\u524d",
                "lastpost": "4&nbsp;\u5c0f\u65f6\u524d",
                "lastposter": "\u6d1e\u7a74\u591c\u83ba",
                "views": "48",
                "replies": "7",
                "displayorder": "0",
                "digest": "0",
                "special": "3",
                "attachment": "0",
                "recommend_add": "0",
                "replycredit": "0",
                "dbdateline": "1601765056",
                "dblastpost": "1601782218",
                "rushreply": "0",
                "reply": [
                    {
                        "pid": "20018981",
                        "author": "E_dition",
                        "authorid": "3049615",
                        "message": "win10\u7248\u5c5e\u4e8e\u57fa\u5ca9\u7248\uff0c\u4f46\u662f\u57fa\u5ca9\u7248\u8fd8\u6709\u4fbf\u643a\u7248\uff0cwin10\u7248\u4e5f\u53ef\u4ee5\u8de8\u5e73\u53f0\u4e0e\u4fbf\u643a\u7248\u4e00\u540c\u6e38\u73a9\uff0cJava\u7248\u6bd4\u8f83\u53e4\u8001\uff0c\u4f18\u5316\u4e5f ..."
                    },
                    {
                        "pid": "20019613",
                        "author": "\u516b\u5742\u30cb\u30e3\u30eb\u5b50",
                        "authorid": "3223997",
                        "message": "\u57fa\u5ca9\u7248\u662f\u8de8\u5e73\u53f0\u7684"
                    },
                    {
                        "pid": "20019826",
                        "author": "\u6d1e\u7a74\u591c\u83ba",
                        "authorid": "2853776",
                        "message": "win10\u7248\u662f\u57fa\u5ca9\u7248\u7684\u4e00\u4e2a\u5b50\u96c6Win10\u7248\u7684\u8eab\u4efd\u9a8c\u8bc1\u65b9\u5f0f\u662fXbox Live\u8d26\u53f7\uff0cJava\u7248\u7684\u8eab\u4efd\u9a8c\u8bc1\u65b9\u5f0f\u662fMojang\u8d26\u53f7Win10 ..."
                    }
                ],
                "recommend": "0"
            },
            {
                "tid": "1124396",
                "typeid": "111",
                "readperm": "0",
                "price": "80",
                "author": "suhairen",
                "authorid": "1483451",
                "subject": "\u6c42\u8fd9\u6b3e\u76ae\u80a464x64\u7248\u672c\uff0c\u8c22\u8c22\uff01",
                "dateline": "14&nbsp;\u5c0f\u65f6\u524d",
                "lastpost": "13&nbsp;\u5c0f\u65f6\u524d",
                "lastposter": "CM_\u6c89\u9ed8",
                "views": "47",
                "replies": "13",
                "displayorder": "0",
                "digest": "0",
                "special": "3",
                "attachment": "2",
                "recommend_add": "0",
                "replycredit": "0",
                "dbdateline": "1601747356",
                "dblastpost": "1601750256",
                "rushreply": "0",
                "reply": [
                    {
                        "pid": "20017432",
                        "author": "CM_\u6c89\u9ed8",
                        "authorid": "1632696",
                        "message": "https:\/\/cowtransfer.com\/s\/3ee16227019849\u53d6\u4ef6\u7801\uff1a76m9x6\uff0824\u5c0f\u65f6\u5185\u6709\u6548\uff09\u7528skin3D\u7684\u529f\u80fd\u6539\u4e86\u4e00\u4e0b\uff0c\u5e0c\u671b\u80fd ..."
                    },
                    {
                        "pid": "20017435",
                        "author": "suhairen",
                        "authorid": "1483451",
                        "message": "\u4ee5\u524d\u7528**\u652f\u6301\u7684\uff0c\u73b0\u5728**\u597d\u50cf\u6ca1\u4e86\u6b63\u7248\u53c8\u6dfb\u52a0\u4e0d\u4e86\uff0c\u5934\u75bc"
                    },
                    {
                        "pid": "20017444",
                        "author": "CM_\u6c89\u9ed8",
                        "authorid": "1632696",
                        "message": "https:\/\/cowtransfer.com\/s\/b73e940629a147 \u6216\u8fdb\u5165 cowtransfer.com \u83b7\u53d6\uff0c\u5728\u9996\u9875\u8f93\u5165\u53d6\u4ef6\u7801\uff1a39rr8x\uff0824\u5c0f ..."
                    }
                ],
                "recommend": "0"
            },
            {
                "tid": "1124348",
                "typeid": "2482",
                "readperm": "0",
                "price": "30",
                "author": "wjsiSIPC",
                "authorid": "3156976",
                "subject": "\u5982\u4f55\u83b7\u5f97\u4e00\u74f6\u77ac\u95f4\u6cbb\u7597255\u7684\u836f\u6c34",
                "dateline": "\u6628\u5929&nbsp;22:24",
                "lastpost": "1&nbsp;\u5c0f\u65f6\u524d",
                "lastposter": "wjsiSIPC",
                "views": "83",
                "replies": "9",
                "displayorder": "0",
                "digest": "0",
                "special": "3",
                "attachment": "0",
                "recommend_add": "0",
                "replycredit": "0",
                "dbdateline": "1601735066",
                "dblastpost": "1601792678",
                "rushreply": "0",
                "reply": [
                    {
                        "pid": "20021505",
                        "author": "wjsiSIPC",
                        "authorid": "3156976",
                        "message": "1.12.2\u4ee5\u524d\u89c1\u8fc7"
                    },
                    {
                        "pid": "20021517",
                        "author": "wjsiSIPC",
                        "authorid": "3156976",
                        "message": "\u5728b\u7ad9****\u7684\u5f55\u64ad\u7ad9\u89c1\u8fc7"
                    },
                    {
                        "pid": "20021559",
                        "author": "wjsiSIPC",
                        "authorid": "3156976",
                        "message": "1.12.2\u662f\u5168\u6548\u679c\uff0c\u5168\u6ee1\u7ea7\u7684\u6ede\u7559\u836f\u6c34"
                    }
                ],
                "recommend": "0"
            },
            {
                "tid": "1124299",
                "typeid": "123",
                "readperm": "0",
                "price": "-70",
                "author": "\u6c89\u9ed8\u306e\u65e0\u8a00",
                "authorid": "1250888",
                "subject": "\u5355\u4eba\u591a\u4eba\u90fd\u65e0\u6cd5\u8fdb\u6e38\u620f\uff0c\u7eaf\u51c0\u7248\u4e5f\u6ca1\u6cd5\u8fdb\uff0c\u6c42\u89e3\u51b3",
                "dateline": "\u6628\u5929&nbsp;20:23",
                "lastpost": "\u6628\u5929&nbsp;23:53",
                "lastposter": "\u6c89\u9ed8\u306e\u65e0\u8a00",
                "views": "24",
                "replies": "2",
                "displayorder": "0",
                "digest": "0",
                "special": "3",
                "attachment": "0",
                "recommend_add": "0",
                "replycredit": "0",
                "dbdateline": "1601727780",
                "dblastpost": "1601740409",
                "rushreply": "0",
                "reply": [
                    {
                        "pid": "20016299",
                        "author": "DreamVoid",
                        "authorid": "2999647",
                        "message": "\u89c1\u6b64\u8d34\uff1ahttps:\/\/www.mcbbs.net\/thread-1124030-1-1.html"
                    },
                    {
                        "pid": "20016916",
                        "author": "\u6c89\u9ed8\u306e\u65e0\u8a00",
                        "authorid": "1250888",
                        "message": "\u611f\u8c22\u6307\u5f15qwq"
                    }
                ],
                "recommend": "0"
            }
        ],
        "groupiconid": {
            "5367": "admin",
            "3664699": "4",
            "2724542": "5",
            "2505673": "8",
            "3581927": "6",
            "3531589": "6",
            "2611011": null,
            "1483451": "4",
            "3156976": "5",
            "1250888": "6"
        },
        "sublist": [
            {
                "fid": "265",
                "icon": "https:\/\/attachment.mcbbs.net\/common\/e5\/common_265_icon.png",
                "name": "\u57fa\u5ca9\u7248\u95ee\u7b54",
                "threads": "3969",
                "posts": "26570",
                "todayposts": "3"
            }
        ],
        "tpp": "10",
        "page": "1",
        "reward_unit": "\u7c92\u91d1\u7c92",
        "threadtypes": {
            "required": "1",
            "listable": "1",
            "prefix": "1",
            "types": {
                "108": "\u5355\u4eba\u6e38\u620f",
                "111": "\u76ae\u80a4\u6750\u8d28",
                "2482": "\u547d\u4ee4&\u7ea2\u77f3",
                "123": "\u7efc\u5408"
            },
            "icons": {
                "108": "",
                "111": "",
                "2482": "",
                "123": ""
            },
            "moderators": {
                "108": null,
                "111": null,
                "2482": null,
                "123": null
            }
        }
    }
}`
