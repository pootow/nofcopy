package gomblr

import (
	"encoding/json"
	"github.com/pootow/nofcopy/gomblr/common"
	"regexp"
	"testing"
)

func TestExtractResourceFromTextPost(t *testing.T) {
	textPostString := `{
	"type": "text",
		"blog_name": "whole-video",
		"blog": {
		"name": "whole-video",
			"title": "第一手独家，稀缺，精品资源",
			"description": "直播盒子代理",
			"url": "https://whole-video.tumblr.com/",
			"uuid": "t:SBPJoRLihoRpaNhd80uTIw",
			"updated": 1544192747
	},
	"id": 180201907970,
		"post_url": "https://whole-video.tumblr.com/post/180201907970/为老公还债出来做洗脚妹却被债主点上用脚肆意玩弄大奶",
		"slug": "为老公还债出来做洗脚妹却被债主点上用脚肆意玩弄大奶",
		"date": "2018-11-17 11:42:00 GMT",
		"timestamp": 1542454920,
		"state": "published",
		"format": "html",
		"reblog_key": "DikDxGZx",
		"tags": [],
		"short_url": "https://tmblr.co/ZIxO-X2dqtkS2",
		"summary": "为老公还债，出来做洗脚妹，却被债主点上，用脚肆意玩弄大奶！！！",
		"is_blocks_post_format": true,
		"recommended_source": null,
		"recommended_color": null,
		"followed": true,
		"liked": true,
		"note_count": 1892,
		"title": null,
		"body": "<figure class=\"tmblr-full\" data-orig-height=\"640\" data-orig-width=\"480\" data-npf='{\"type\":\"video\",\"provider\":\"tumblr\",\"url\":\"https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4\",\"media\":{\"url\":\"https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4\",\"type\":\"video/mp4\",\"width\":480,\"height\":640},\"poster\":[{\"url\":\"https://66.media.tumblr.com/tumblr_pic5svrUUz1xmht6v_frame1.jpg\",\"type\":\"image/jpeg\",\"width\":480,\"height\":640}]}'><video controls=\"controls\" autoplay=\"autoplay\" muted=\"muted\" poster=\"https://66.media.tumblr.com/tumblr_pic5svrUUz1xmht6v_frame1.jpg\"><source src=\"https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4\" type=\"video/mp4\"></source></video></figure><p>为老公还债，出来做洗脚妹，却被债主点上，用脚肆意玩弄大奶！！！</p>",
		"reblog": {
		"comment": "<p><figure class=\"tmblr-full\" data-orig-height=\"640\" data-orig-width=\"480\" data-npf='{\"type\":\"video\",\"provider\":\"tumblr\",\"url\":\"https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4\",\"media\":{\"url\":\"https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4\",\"type\":\"video/mp4\",\"width\":480,\"height\":640},\"poster\":[{\"url\":\"https://66.media.tumblr.com/tumblr_pic5svrUUz1xmht6v_frame1.jpg\",\"type\":\"image/jpeg\",\"width\":480,\"height\":640}]}'><video controls=\"controls\" autoplay=\"autoplay\" muted=\"muted\" poster=\"https://66.media.tumblr.com/tumblr_pic5svrUUz1xmht6v_frame1.jpg\"><source src=\"https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4\" type=\"video/mp4\"></source></video></figure><p>为老公还债，出来做洗脚妹，却被债主点上，用脚肆意玩弄大奶！！！</p></p>",
			"tree_html": ""
	},
	"trail": [
	{
		"blog": {
			"name": "whole-video",
			"active": true,
			"theme": {
				"avatar_shape": "square",
				"background_color": "#FAFAFA",
				"body_font": "Helvetica Neue",
				"header_bounds": "",
				"header_image": "https://assets.tumblr.com/images/default_header/optica_pattern_11.png?_v=4275fa0865b78225d79970023dde05a1",
				"header_image_focused": "https://assets.tumblr.com/images/default_header/optica_pattern_11.png?_v=4275fa0865b78225d79970023dde05a1",
				"header_image_scaled": "https://assets.tumblr.com/images/default_header/optica_pattern_11.png?_v=4275fa0865b78225d79970023dde05a1",
				"header_stretch": true,
				"link_color": "#529ECC",
				"show_avatar": true,
				"show_description": true,
				"show_header_image": true,
				"show_title": true,
				"title_color": "#444444",
				"title_font": "Sans Serif",
				"title_font_weight": "bold"
			},
			"share_likes": false,
			"share_following": false,
			"can_be_followed": true
		},
		"post": {
			"id": "180201907970"
		},
		"content_raw": "<p><figure class=\"tmblr-full\" data-orig-height=\"640\" data-orig-width=\"480\" data-npf='{\"type\":\"video\",\"provider\":\"tumblr\",\"url\":\"https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4\",\"media\":{\"url\":\"https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4\",\"type\":\"video/mp4\",\"width\":480,\"height\":640},\"poster\":[{\"url\":\"https://66.media.tumblr.com/tumblr_pic5svrUUz1xmht6v_frame1.jpg\",\"type\":\"image/jpeg\",\"width\":480,\"height\":640}]}'><video controls=\"controls\" autoplay=\"autoplay\" muted=\"muted\" poster=\"https://66.media.tumblr.com/tumblr_pic5svrUUz1xmht6v_frame1.jpg\"><source src=\"https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4\" type=\"video/mp4\"></source></video></figure><p>为老公还债，出来做洗脚妹，却被债主点上，用脚肆意玩弄大奶！！！</p></p>",
		"content": "<p><figure class=\"tmblr-full\" data-npf=\"{&quot;type&quot;:&quot;video&quot;,&quot;provider&quot;:&quot;tumblr&quot;,&quot;url&quot;:&quot;https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4&quot;,&quot;media&quot;:{&quot;url&quot;:&quot;https://ve.media.tumblr.com/tumblr_pic5svrUUz1xmht6v.mp4&quot;,&quot;type&quot;:&quot;video/mp4&quot;,&quot;width&quot;:480,&quot;height&quot;:640},&quot;poster&quot;:[{&quot;url&quot;:&quot;https://66.media.tumblr.com/tumblr_pic5svrUUz1xmht6v_frame1.jpg&quot;,&quot;type&quot;:&quot;image/jpeg&quot;,&quot;width&quot;:480,&quot;height&quot;:640}]}\"></figure><p>&#20026;&#32769;&#20844;&#36824;&#20538;&#65292;&#20986;&#26469;&#20570;&#27927;&#33050;&#22969;&#65292;&#21364;&#34987;&#20538;&#20027;&#28857;&#19978;&#65292;&#29992;&#33050;&#32902;&#24847;&#29609;&#24324;&#22823;&#22902;&#65281;&#65281;&#65281;</p></p>",
		"is_current_item": true,
		"is_root_item": true
	}
],
"liked_timestamp": 1543418485,
"can_like": true,
"can_reblog": true,
"can_send_in_message": true,
"can_reply": true,
"display_avatar": true
}`
	textPost := gTextPost{}
	if err := json.Unmarshal([]byte(textPostString), &textPost); err != nil {
		t.Fatal(err)
	}

	println(textPost.Body)

	for _, e := range textPost.GetResources() {
		println(e)
	}

}

func TestRegex(t *testing.T) {
	regex, err := regexp.Compile(`http(s)?://[^\\\"'\<\?\&]+(\.jpg|\.png|\.gif|\.webp|\.mp4|\.amr|\.wav|\.3gp)`)
	if err != nil {
		t.Fatal(err)
	}

	allString := regex.FindAllString(common.JsonString, -1)
	for _, String := range allString {
		println(String)
	}
}
