package httpClient

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"
)

func TestHttpGet(t *testing.T) {
	url := "https://www.nasa.gov/topics/earth/images/index.html"
	bytes, err := HttpGet(url)
	if err != nil {
		t.Fatalf("http get url error %v", err)
	}
	s := string(bytes)

	fmt.Printf("%v", s)
}

func TestRegexp(t *testing.T) {
	str := `<img src="https://img3.doubanio.com/view/photo/l/public/p455467226.webp"  />`
	//str := `        "https://img1.doubanio.com/view/photo/l/public/p2182128677.webp",`
	//webp image
	reg := regexp.MustCompile(`(.+)"(https://.+/view/.+\.webp)"(.+)`)

	str2 := `<a class="mainphoto" href="https://movie.douban.com/photos/photo/2182128677/#title-anchor" title="点击查看下一张">`
	//href
	reg2 := regexp.MustCompile(`<(.+)href="(https://.+#title-anchor)"(.+)>`)

	match := reg.FindAllStringSubmatch(str, 1000)

	match2 := reg2.FindAllStringSubmatch(str2, 1000)

	for _, v := range match {
		for k, val := range v {
			if k == 2 {
				fmt.Printf("%v\n", val)
			}

		}
	}

	for _, v := range match2 {
		for k, val := range v {
			if k == 2 {
				fmt.Printf("%v\n", val)
			}
		}
	}

}

type postReq struct {
	Html  string `json:"html"`
	FHtml string `json:"fhtml"`
}

type postReqp struct {
	FHtml string `json:"fhtml"`
}

func TestHttpPost(t *testing.T) {
	str := `
<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd"><html xmlns="http://www.w3.org/1999/xhtml"><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8" /><meta http-equiv="Cache-Control" content="no-cache"/><meta id="viewport" name="viewport" content="width=device-width,initial-scale=1.0,minimum-scale=1.0, maximum-scale=2.0" /><link rel="icon" sizes="any" mask href="https://h5.sinaimg.cn/upload/2015/05/15/28/WeiboLogoCh.svg" color="black"><meta name="MobileOptimized" content="240"/><title>任重的微博</title><style type="text/css" id="internalStyle">html,body,p,form,div,table,textarea,input,span,select{font-size:12px;word-wrap:break-word;}body{background:#F8F9F9;color:#000;padding:1px;margin:1px;}table,tr,td{border-width:0px;margin:0px;padding:0px;}form{margin:0px;padding:0px;border:0px;}textarea{border:1px solid #96c1e6}textarea{width:95%;}a,.tl{color:#2a5492;text-decoration:underline;}/*a:link {color:#023298}*/.k{color:#2a5492;text-decoration:underline;}.kt{color:#F00;}.ib{border:1px solid #C1C1C1;}.pm,.pmy{clear:both;background:#ffffff;color:#676566;border:1px solid #b1cee7;padding:3px;margin:2px 1px;overflow:hidden;}.pms{clear:both;background:#c8d9f3;color:#666666;padding:3px;margin:0 1px;overflow:hidden;}.pmst{margin-top: 5px;}.pmsl{clear:both;padding:3px;margin:0 1px;overflow:hidden;}.pmy{background:#DADADA;border:1px solid #F8F8F8;}.t{padding:0px;margin:0px;height:35px;}.b{background:#e3efff;text-align:center;color:#2a5492;clear:both;padding:4px;}.bl{color:#2a5492;}.n{clear:both;background:#436193;color:#FFF;padding:4px; margin: 1px;}.nt{color:#b9e7ff;}.nl{color:#FFF;text-decoration:none;}.nfw{clear:both;border:1px solid #BACDEB;padding:3px;margin:2px 1px;}.s{border-bottom:1px dotted #666666;margin:3px;clear:both;}.tip{clear:both; background:#c8d9f3;color:#676566;border:1px solid #BACDEB;padding:3px;margin:2px 1px;}.tip2{color:#000000;padding:2px 3px;clear:both;}.ps{clear:both;background:#FFF;color:#676566;border:1px solid #BACDEB;padding:3px;margin:2px 1px;}.tm{background:#feffe5;border:1px solid #e6de8d;padding:4px;}.tm a{color:#ba8300;}.tmn{color:#f00}.tk{color:#ffffff}.tc{color:#63676A;}.c{padding:2px 5px;}.c div a img{border:1px solid #C1C1C1;}.ct{color:#9d9d9d;font-style:italic;}.cmt{color:#9d9d9d;}.ctt{color:#000;}.cc{color:#2a5492;}.nk{color:#2a5492;}.por {border: 1px solid #CCCCCC;height:50px;width:50px;}.me{color:#000000;background:#FEDFDF;padding:2px 5px;}.pa{padding:2px 4px;}.nm{margin:10px 5px;padding:2px;}.hm{padding:5px;background:#FFF;color:#63676A;}.u{margin:2px 1px;background:#ffffff;border:1px solid #b1cee7;}.ut{padding:2px 3px;}.cd{text-align:center;}.r{color:#F00;}.g{color:#0F0;}.bn{background: transparent;border: 0 none;text-align: left;padding-left: 0;}</style><script>if(top != self){top.location = self.location;}</script></head><body><div class="c"><span class="cmt">加入新浪微博,分享新鲜的事!</span>[<a href="http://passport.weibo.cn/signin/login?entry=mweibo&amp;r=http%3A%2F%2Fweibo.cn%2Fyanyuanrenzhong&amp;backTitle=%CE%A2%B2%A9&amp;vt=" id="top">登录</a>&nbsp;<a href="http://weibo.cn/reg/index?ns=1&amp;revalid=2&amp;backURL=http%3A%2F%2Fweibo.cn%2Fyanyuanrenzhong&amp;backTitle=%CE%A2%B2%A9&amp;vt=">注册</a>]</div><div class="n" style="padding: 6px 4px;"><a href="https://weibo.cn/?tf=5_009" class="nl">首页</a>|<a href="https://weibo.cn/msg/?tf=5_010" class="nl">消息</a>|<a href="https://huati.weibo.cn" class="nl">话题</a>|<a href="https://weibo.cn/search/?tf=5_012" class="nl">搜索</a>|<a href="/yanyuanrenzhong?rand=3847&amp;p=r" class="nl">刷新</a></div><div class="c tip"><a href="http://m.weibo.cn" id="top" class="tl">手机微博触屏版,点击前往>></a></div><div class="u"><table><tr><td valign="top"><a href="/1296947890/avatar?rl=0"><img src="http://tva3.sinaimg.cn/crop.0.0.1080.1080.50/4d4ddab2jw8efnjzlqta8j20u00u00vd.jpg" alt="头像" class="por" /></a></td><td valign="top"><div class="ut"><span class="ctt">任重<img src="https://h5.sinaimg.cn/upload/2016/05/26/319/5338.gif" alt="V"/><a href="http://vip.weibo.cn/?F=W_tq_zsbs_01"><img src="https://h5.sinaimg.cn/upload/2016/05/26/319/donate_btn_s.png" alt="M"/></a>&nbsp;男/北京    &nbsp;    <a href="/attention/add?uid=1296947890&amp;rl=0&amp;st=0c6209">加关注</a></span><br /><span class="ctt">认证：演员任重</span><br /><span class="ctt" style="word-break:break-all; width:50px;">演员，鑫宝源签约艺人</span><br /><a href="/1296947890/info">资料</a>&nbsp;<a href="/1296947890/operation?rl=0">操作</a>&nbsp;<a href="/attgroup/special?fuid=1296947890&amp;st=0c6209">特别关注</a>&nbsp;<a href="http://new.vip.weibo.cn/vippay/payother?present=1&amp;action=comfirmTime&amp;uid=1296947890">送Ta会员</a></div></td></tr></table><div class="tip2"><span class="tc">微博[4492]</span>&nbsp;<a href="/1296947890/follow">关注[1091]</a>&nbsp;<a href="/1296947890/fans">粉丝[8959919]</a>&nbsp;<a href="/attgroup/opening?uid=1296947890">分组[1]</a>&nbsp;<a href="/at/weibo?uid=1296947890">@他的</a></div></div><div class="pmst"><span class="pms">&nbsp;微博&nbsp;</span><span class="pmsl">&nbsp;<a href="/1296947890/photo?tf=6_008">相册</a>&nbsp;</span></div><div class="pms" >全部-<a href="/yanyuanrenzhong?filter=1">原创</a>-<a href="/yanyuanrenzhong?filter=2">图片</a>-<a href="/attgroup/opening?uid=1296947890">分组</a>-<a href="/1296947890/search?f=u&amp;rl=0">筛选</a></div><div class="c" id="M_FFvZziWIi"><div><span class="ctt"><a href="/n/%E9%83%8A%E5%8E%BF%E5%A4%A9%E7%8E%8B%E8%80%81%E7%94%B0">@郊县天王老田</a> 这碗狗粮我喜欢！哈哈哈哈[二哈]我和老田 高丽丽同框了[二哈][阴险][二哈][阴险][二哈] ​​​</span>&nbsp;[<a href="https://weibo.cn/mblog/picAll/FFvZziWIi?rl=1">组图共4张</a>]</div><div><a href="https://weibo.cn/mblog/pic/FFvZziWIi?rl=0"><img src="http://wx4.sinaimg.cn/wap180/4d4ddab2gy1fnp6zba0kaj20qo0zk0yo.jpg" alt="图片" class="ib" /></a>&nbsp;<a href="https://weibo.cn/mblog/oripic?id=FFvZziWIi&amp;u=4d4ddab2gy1fnp6zba0kaj20qo0zk0yo">原图</a>&nbsp;<a href="https://weibo.cn/attitude/FFvZziWIi/add?uid=&amp;rl=0&amp;st=0c6209">赞[8593]</a>&nbsp;<a href="https://weibo.cn/repost/FFvZziWIi?uid=1296947890&amp;rl=0">转发[3286]</a>&nbsp;<a href="https://weibo.cn/comment/FFvZziWIi?uid=1296947890&amp;rl=0#cmtfrm" class="cc">评论[657]</a>&nbsp;<a href="https://weibo.cn/fav/addFav/FFvZziWIi?rl=0&amp;st=0c6209">收藏</a><!---->&nbsp;<span class="ct">01月22日 10:56&nbsp;来自iPhone X</span></div></div><div class="s"></div><div class="c" id="M_FEjPKj6bh"><div><span class="cmt">转发了&nbsp;<a href="https://weibo.cn/u/3229192424">电视剧凤囚凰</a><img src="https://h5.sinaimg.cn/upload/2016/05/26/319/5337.gif" alt="V"/>&nbsp;的微博:</span><span class="ctt">诛心？囚心？爱恨情仇，为谁生为谁死？生离死别，肝肠寸断。<br/><a href="http://weibo.cn/pages/100808topic?extparam=%E4%BB%8A%E6%99%9A%E7%9C%8B%E5%87%A4%E5%9B%9A%E5%87%B0&amp;from=feed">#今晚看凤囚凰#</a><a href="http://weibo.cn/pages/100808topic?extparam=%E7%94%B5%E8%A7%86%E5%89%A7%E5%87%A4%E5%9B%9A%E5%87%B0&amp;from=feed">#电视剧凤囚凰#</a> 发布终极预告 <a href="https://weibo.cn/sinaurl?f=w&amp;u=http%3A%2F%2Ft.cn%2FRQ5Qel0&amp;ep=FEjPKj6bh%2C1296947890%2CFEi0t2LBn%2C3229192424">《凤囚凰》全阵容终极片花 虐恋升级</a> <a href="/n/%E5%85%B3%E6%99%93%E5%BD%A4">@关晓彤</a> <a href="/n/%E5%AE%8B%E5%A8%81%E9%BE%9999">@宋威龙99</a> <a href="/n/%E7%99%BD%E9%B9%BFmy">@白鹿my</a> <a href="/n/%E7%B1%B3%E7%83%ADMERXAT">@米热MERXAT</a> <a href="/n/%E8%B5%B5%E9%9C%B2%E6%80%9D%E7%9A%84%E5%BE%AE%E5%8D%9A">@赵露思的微博</a> <a href="/n/%E5%90%B4%E8%B0%A8%E8%A8%80Y">@吴谨言Y</a> <a href="/n/%E4%BD%95%E5%A5%89%E5%A4%A9">@何奉天</a> <a href="/n/%E5%90%B4%E4%BD%B3%E6%80%A1_six">@吴佳怡_six</a> <a href="/n/%E7%9F%B3%E4%BA%91%E9%B9%8F%E9%B9%8F">@石云鹏鹏</a> <a href="/n/%E7%8E%8B%E8%8C%82%E8%95%BE">@王茂蕾</a> <a href="/n/%E8%AE%B8%E5%87%AFKEV">@许凯KEV</a> <a href="/n/%E6%9D%8E%E6%98%A5%E5%AB%92%E5%A4%A7%E7%82%B9">@李春嫒大点</a> <a href="/n/%E6%BC%94%E5%91%98%E7%8E%8B%E4%B8%80%E5%93%B2">@演员王一哲</a> <a href="/n/%E6%B4%AA%E5%B0%A7">@洪尧</a> <a href="/n/%E6%9C%B1%E6%88%ACRyuuji">@朱戬Ryuuji</a><br/>22：00 ​​​&nbsp;<a href='/comment/FEi0t2LBn'>全文</a></span>&nbsp;<span class="cmt">赞[4449]</span>&nbsp;<span class="cmt">原文转发[325148]</span>&nbsp;<a href="https://weibo.cn/comment/FEi0t2LBn?rl=0#cmtfrm" class="cc">原文评论[17899]</a><!----></div><div><span class="cmt">转发理由:</span>看起来//<a href="/n/%E4%BA%8E%E6%AD%A31978">@于正1978</a>:今晚来看<a href="http://weibo.cn/pages/100808topic?extparam=%E7%94%B5%E8%A7%86%E5%89%A7%E5%87%A4%E5%9B%9A%E5%87%B0&amp;from=feed">#电视剧凤囚凰#</a>&nbsp;&nbsp;<a href="https://weibo.cn/attitude/FEjPKj6bh/add?uid=&amp;rl=0&amp;st=0c6209">赞[2641]</a>&nbsp;<a href="https://weibo.cn/repost/FEjPKj6bh?uid=1296947890&amp;rl=0">转发[1496]</a>&nbsp;<a href="https://weibo.cn/comment/FEjPKj6bh?uid=1296947890&amp;rl=0#cmtfrm" class="cc">评论[523]</a>&nbsp;<a href="https://weibo.cn/fav/addFav/FEjPKj6bh?rl=0&amp;st=0c6209">收藏</a><!---->&nbsp;<span class="ct">01月14日 14:09&nbsp;来自iPhone X</span></div></div><div class="s"></div><div class="c" id="M_FDYW6clDM"><div><span class="ctt">恭喜Fernanda 寂静的王国 环球电影节（柏林）获 年度最佳女主角 ​​​</span>&nbsp;[<a href="https://weibo.cn/mblog/picAll/FDYW6clDM?rl=1">组图共6张</a>]</div><div><a href="https://weibo.cn/mblog/pic/FDYW6clDM?rl=0"><img src="http://wx2.sinaimg.cn/wap180/4d4ddab2gy1fndjbuv9ebj20yj0qoq9z.jpg" alt="图片" class="ib" /></a>&nbsp;<a href="https://weibo.cn/mblog/oripic?id=FDYW6clDM&amp;u=4d4ddab2gy1fndjbuv9ebj20yj0qoq9z">原图</a>&nbsp;<a href="https://weibo.cn/attitude/FDYW6clDM/add?uid=&amp;rl=0&amp;st=0c6209">赞[7203]</a>&nbsp;<a href="https://weibo.cn/repost/FDYW6clDM?uid=1296947890&amp;rl=0">转发[2508]</a>&nbsp;<a href="https://weibo.cn/comment/FDYW6clDM?uid=1296947890&amp;rl=0#cmtfrm" class="cc">评论[638]</a>&nbsp;<a href="https://weibo.cn/fav/addFav/FDYW6clDM?rl=0&amp;st=0c6209">收藏</a><!---->&nbsp;<span class="ct">01月12日 08:57&nbsp;来自iPhone X</span></div></div><div class="c"><form action="http://passport.weibo.cn/signin/login?entry=mweibo&amp;r=http%3A%2F%2Fweibo.cn%2Fyanyuanrenzhong&amp;backTitle=%CE%A2%B2%A9&amp;vt="><div><input type="submit" value="查看更多内容" style="width:100%;display:inline-block" /></div></form></div><div class="pm"><form action="/search/" method="post"><div><input type="text" name="keyword" value="" size="15" /><input type="submit" name="smblog" value="搜微博" /><input type="submit" name="suser" value="找人" /><br/><span class="pmf"><a href="/search/mblog/?keyword=%E5%B0%8F%E4%B8%BB%E6%92%AD%E6%B1%82%E9%80%86%E8%A2%AD&amp;rl=0" class="k">小主播求逆袭</a>&nbsp;<a href="/search/mblog/?keyword=%E4%B8%89%E6%B5%81%E4%B9%8B%E8%B7%AF&amp;rl=0" class="k">三流之路</a>&nbsp;<a href="/search/mblog/?keyword=%E9%99%86%E5%9C%B0%E5%A4%AB%E5%A6%87&amp;rl=0" class="k">陆地夫妇</a>&nbsp;<a href="/search/mblog/?keyword=%E9%95%BF%E5%A4%A7%E7%94%B5%E9%99%A2%E6%AF%95%E4%B8%9A%E6%99%9A%E4%BC%9A&amp;rl=0" class="k">长大电院毕业晚会</a>&nbsp;<a href="/search/mblog/?keyword=%E5%88%80%E5%89%91%E4%B9%B1%E8%88%9E&amp;rl=0" class="k">刀剑乱舞</a></span></div></form></div><div class="cd"><a href="#top"><img src="https://h5.sinaimg.cn/upload/2017/04/27/319/5e990ec2.gif" alt="TOP"/></a></div><div class="pms"> <a href="https://weibo.cn">首页</a>.<a href="https://weibo.cn/topic/240489">反馈</a>.<a href="https://weibo.cn/page/91">帮助</a>.<a  href="http://down.sina.cn/weibo/default/index/soft_id/1/mid/0"  >客户端</a>.<a href="https://weibo.cn/spam/?rl=0&amp;type=3&amp;fuid=1296947890" class="kt">举报</a>.<a href="http://passport.sina.cn/sso/logout?r=https%3A%2F%2Fweibo.cn%2Fpub%2F%3Fvt%3D&amp;entry=mweibo">退出</a></div><div class="c">设置:<a href="https://weibo.cn/account/customize/skin?tf=7_005&amp;st=0c6209">皮肤</a>.<a href="https://weibo.cn/account/customize/pic?tf=7_006&amp;st=0c6209">图片</a>.<a href="https://weibo.cn/account/customize/pagesize?tf=7_007&amp;st=0c6209">条数</a>.<a href="https://weibo.cn/account/privacy/?tf=7_008&amp;st=0c6209">隐私</a></div><div class="c">彩版|<a href="http://m.weibo.cn/?tf=7_010">触屏</a>|<a href="https://weibo.cn/page/521?tf=7_011">语音</a></div><div class="b">weibo.cn[01-24 13:38]</div></body></html>`
	req := postReq{Html: str}
	respStr := postReqp{}
	bytes, _ := json.Marshal(&req)

	resp, err := HttpPost("http://tool.oschina.net/action/format/html", string(bytes))
	if err != nil {
		t.Fatalf("post error %v", err)
	}
	fmt.Println(string(resp))

	err = json.Unmarshal(resp, &respStr)
	if err != nil {
		t.Fatalf("post error %v", err)
	}
	fmt.Println(respStr.FHtml)
}
