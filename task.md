# AI瀛︿範绀惧尯 - 甯栧瓙涓庤瘎璁烘牳蹇傾PI浠诲姟娓呭崟锛堟晱鎹峰紑鍙戞ā寮忥級

## 椤圭洰姒傝堪
- **鍚庣妗嗘灦**锛欸oFrame
- **鍓嶇妗嗘灦**锛歏ue3 + Vite
- **鏁版嵁搴?*锛歁ySQL
- **鍔熻兘**锛氬笘瀛愪笌璇勮鏍稿績API + 鍓嶇椤甸潰
- **寮€鍙戞ā寮?*锛氭晱鎹稴print杩唬锛屽悗绔笌鍓嶇骞惰寮€鍙?
---

## Sprint 1: 鍩虹璁炬柦鎼缓

### 1.1 鏁版嵁搴撹璁?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 1.1.1 璁捐posts琛ㄧ粨鏋?| 鉁?瀹屾垚 | database/posts.sql | tuhome-backend-architect | 2026-04-10 20:04 | 0c38cd3 | 鍖呭惈title, content, author_id, created_at, updated_at绛夊瓧娈?|
| 1.1.2 璁捐comments琛ㄧ粨鏋?| 鉁?瀹屾垚 | database/comments.sql | tuhome-backend-architect | 2026-04-10 20:05 | ba167ac | 鏀寔鏍戝舰灞傜骇缁撴瀯锛宲arent_id鑷叧鑱?|
| 1.1.3 缂栧啓寤鸿〃SQL鑴氭湰 | 鉁?瀹屾垚 | database/init.sql | tuhome-backend-architect | 2026-04-10 20:07 | 4914263b | 鐢熸垚瀹屾暣鐨刬nit.sql锛屾暣鍚坧osts鍜宑omments琛?|
| 1.1.4 鎵ц寤鸿〃SQL鑴氭湰 | 鉁?瀹屾垚 | - | tuhome-backend-architect | 2026-04-10 20:43 | 4914263b | 鎵цinit.sql鍒涘缓鏁版嵁搴撹〃 |

### 1.2 鍚庣椤圭洰鍒濆鍖?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 1.2.1 鍒濆鍖朑oFrame椤圭洰鐩綍缁撴瀯 | 鉁?瀹屾垚 | model/, dao/, service/, controller/, internal/cmd/, internal/consts/, internal/packed/, manifest/config/, manifest/deploy/ | tuhome-backend-architect | 2026-04-10 20:09 | c573b6cd | model/dao/service/controller/internal/manifest鐩綍缁撴瀯 |
| 1.2.2 鍒涘缓database.yaml閰嶇疆鏂囦欢 | 鉁?瀹屾垚 | manifest/config/database.yaml | tuhome-backend-architect | 2026-04-10 20:10 | 11eb7e5 | 鏁版嵁搴撹繛鎺ラ厤缃?|
| 1.2.3 鍒濆鍖栨暟鎹簱杩炴帴 | 鉁?瀹屾垚 | internal/dao/dao.go, internal/cmd/cmd.go, manifest/config/config.yaml | tuhome-backend-architect | 2026-04-10 20:11 | 98c71cf7 | g.DB()鍒濆鍖?|

### 1.3 鍓嶇椤圭洰鍒濆鍖?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 1.3.1 鍒濆鍖朧ue3 + Vite椤圭洰 | 鉁?瀹屾垚 | frontend/ | tuhome-frontend-developer | 2026-04-10 20:17 | 2021ebc | 浣跨敤create-vite妯℃澘 |
| 1.3.2 瀹夎渚濊禆锛坅xios, vue-router, pinia绛夛級 | 鉁?瀹屾垚 | package.json, package-lock.json | tuhome-frontend-developer | 2026-04-10 20:18 | dab4bb9 | 鏍稿績渚濊禆鍖?|
| 1.3.3 閰嶇疆椤圭洰鐩綍缁撴瀯 | 鉁?瀹屾垚 | - | tuhome-frontend-developer | 2026-04-10 20:19 | b65919f | views/components/api/store/router鐩綍 |
| 1.3.4 閰嶇疆API璇锋眰鍩哄湴鍧€ | 鉁?瀹屾垚 | frontend/.env | tuhome-frontend-developer | 2026-04-10 20:22 | 9f69985 | axios baseURL閰嶇疆锛孷ITE_API_BASE_URL鐜鍙橀噺鎸囧悜 http://localhost:8000 |

---

## Sprint 2: 甯栧瓙鍔熻兘寮€鍙?
### 2.1 鍚庣-Post鏁版嵁灞傦紙Model + Dao锛?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 2.1.1 鍒涘缓Post妯″瀷瀹氫箟 | 鉁?瀹屾垚 | internal/model/entity/post.go | tuhome-backend-architect | 2026-04-10 20:22 | 20a9a5a | gmodel.Post瀹炰綋 |
| 2.1.2 鍒涘缓Post鏂板/鍒犻櫎/鏌ヨ杈撳叆杈撳嚭缁撴瀯浣?| 鉁?瀹屾垚 | internal/model/do/post_do.go | tuhome-backend-architect | 2026-04-10 20:24 | 73bea33 | Do CreateReq/Do DeleteReq/Do GetOneReq绛?|
| 2.1.3 鍒涘缓Post鏁版嵁璁块棶鏂规硶-Insert | 鉁?瀹屾垚 | internal/dao/post.go | tuhome-backend-architect | 2026-04-10 20:26 | fa0ab9e | dao.Post.Insert() |
| 2.1.4 鍒涘缓Post鏁版嵁璁块棶鏂规硶-Delete | 鉁?瀹屾垚 | internal/dao/post.go | tuhome-orchestrator:subagent:4f49609e | 2026-04-10 20:32 | fece73f | dao.Post.Delete() 杞垹闄わ紝杩斿洖鍙楀奖鍝嶈鏁?|
| 2.1.5 鍒涘缓Post鏁版嵁璁块棶鏂规硶-GetOne | 鉁?瀹屾垚 | internal/dao/post.go | tuhome-backend-architect | 2026-04-10 20:34 | cd95d87 | dao.Post.GetOne() 鏀寔杞垹闄よ繃婊?|
| 2.1.6 鍒涘缓Post鏁版嵁璁块棶鏂规硶-GetList | 鉁?瀹屾垚 | internal/dao/post.go | tuhome-orchestrator:subagent:e04dbd6a | 2026-04-10 20:34 | f6758db | dao.Post.GetList()鍒嗛〉鏌ヨ锛屾敮鎸乤uthorId/tags/status/keyword杩囨护锛宑reated_at DESC鎺掑簭 |

### 2.2 鍚庣-Post涓氬姟灞備笌鎺ュ彛灞?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 2.2.1 瀹炵幇Post涓氬姟鏂规硶-Create | 鉁?瀹屾垚 | internal/service/post.go | tuhome-backend-architect | 2026-04-10 20:32 | 2b8ff65 | service.Post.Create() |
| 2.2.2 瀹炵幇Post涓氬姟鏂规硶-Delete | 鉁?瀹屾垚 | internal/service/post.go | tuhome-orchestrator:subagent:2328be6f | 2026-04-10 20:35 | 012c508 | service.Post.Delete() |
| 2.2.3 瀹炵幇Post涓氬姟鏂规硶-GetDetail | 鉁?瀹屾垚 | internal/service/post.go | tuhome-orchestrator:subagent:d2c94a75 | 2026-04-10 20:34 | dfad6ce | service.Post.GetDetail() |
| 2.2.4 瀹炵幇Post涓氬姟鏂规硶-GetPageList | 鉁?瀹屾垚 | internal/service/post.go | tuhome-backend-architect | 2026-04-10 20:38 | d8fce3e | service.Post.GetPageList() |
| 2.2.5 瀹炵幇銆愬彂甯冨笘瀛愩€慉PI鎺ュ彛 | 鉁?瀹屾垚 | internal/controller/post/post.go, internal/cmd/cmd.go, go.mod, main.go | tuhome-backend-architect | 2026-04-10 20:50 | e4749ae | POST /api/post/create锛屼慨澶峴ervice.PostService璋冪敤鍜宮ain.go |
| 2.2.6 瀹炵幇銆愬垹闄ゅ笘瀛愩€慉PI鎺ュ彛 | 鉁?瀹屾垚 | internal/controller/post/post.go | tuhome-backend-architect | 2026-04-10 20:47 | a0bab7c | POST /api/post/delete |
| 2.2.7 瀹炵幇銆愭煡璇㈠笘瀛愯鎯呫€慉PI鎺ュ彛 | 鉁?瀹屾垚 | internal/controller/post/post.go, internal/cmd/cmd.go | tuhome-orchestrator:subagent:4b5c79cf | 2026-04-10 20:45 | 4b37a67 | GET /api/post/detail |
| 2.2.8 瀹炵幇銆愬垎椤垫煡璇㈠笘瀛愬垪琛ㄣ€慉PI鎺ュ彛 | 鉁?瀹屾垚 | internal/controller/post/post.go | tuhome-orchestrator:subagent:b9212f47 | 2026-04-10 20:45 | 9ac9b46 | GET /api/post/list |
| 2.2.9 娉ㄥ唽Post璺敱缁?| 鉁?瀹屾垚 | internal/cmd/cmd.go | tuhome-orchestrator:subagent:cd571790 | 2026-04-10 20:46 | 406b8a1 | 4涓矾鐢卞叏閮ㄦ敞鍐岋紝POST create/delete, GET detail/list |
| 2.2.10 瀹炵幇銆愭洿鏂板笘瀛愩€慉PI鎺ュ彛 | 鉁?瀹屾垚 | internal/controller/post/post.go, internal/service/post.go, internal/dao/post.go | tuhome-backend-architect | 2026-04-10 21:04 | 4a89245 | POST /api/post/update锛屽凡瀹屾垚DAO/Service/Controller涓夊眰瀹炵幇 |

### 2.3 鍓嶇-Post API鏈嶅姟灞?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 2.3.1 灏佽Post API璋冪敤鏂规硶 | 鉁?瀹屾垚 | frontend/src/api/post.js | tuhome-frontend-developer | 2026-04-10 20:53 | e8ef93a | api/post.js灏佽CRUD鎺ュ彛 |

### 2.4 鍓嶇-Post椤甸潰寮€鍙?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 2.4.1 甯栧瓙鍒楄〃椤碉紙鍒嗛〉+灞曠ず锛?| 鉁?瀹屾垚 | frontend/src/views/post/PostList.vue | tuhome-frontend-developer | 2026-04-10 20:54 | 034cd08 | views/post/PostList.vue |
| 2.4.2 甯栧瓙璇︽儏椤?| 鉁?瀹屾垚 | frontend/src/views/post/PostDetail.vue | tuhome-frontend-developer | 2026-04-10 21:00 | 27eff74 | views/post/PostDetail.vue |
| 2.4.3 鍙戝竷甯栧瓙椤?| 鉁?瀹屾垚 | frontend/src/views/post/PostCreate.vue | tuhome-frontend-developer | 2026-04-10 21:05 | ebc21e0 | views/post/PostCreate.vue |
| 2.4.4 缂栬緫甯栧瓙椤?| 鉁?瀹屾垚 | frontend/src/views/post/PostEdit.vue, frontend/src/api/post.js | tuhome-frontend-developer | 2026-04-10 21:15 | aee87a8 | views/post/PostEdit.vue |

### 2.5 鍓嶇-Post璺敱閰嶇疆

| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 2.5.1 閰嶇疆甯栧瓙鐩稿叧璺敱 | 鉁?瀹屾垚 | frontend/src/router/index.js, frontend/src/main.js, frontend/src/App.vue | tuhome-frontend-developer | 2026-04-10 21:06 | 7f48779 | router/index.js閰嶇疆 |

---

## Sprint 3: 璇勮鍔熻兘寮€鍙?
### 3.1 鍚庣-Comment鏁版嵁灞傦紙Model + Dao锛?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.1.1 鍒涘缓Comment妯″瀷瀹氫箟 | 鉁?瀹屾垚 | internal/model/entity/comment.go | tuhome-backend-architect | 2026-04-10 21:06 | a98c3e3 | gmodel.Comment瀹炰綋 |
| 3.1.2 鍒涘缓Comment鏂板/鍒犻櫎/鏌ヨ杈撳叆杈撳嚭缁撴瀯浣?| 鉁?瀹屾垚 | internal/model/do/comment_do.go | tuhome-orchestrator:subagent:adf0bd31 | 2026-04-10 21:16 | 705f1e3 | 鏂板CommentReplyReq/CommentGetTreeReq/CommentTreeResp锛岃ˉ鍏匲serId瀛楁 |
| 3.1.3 鍒涘缓Comment鏁版嵁璁块棶鏂规硶-Insert | 鉁?瀹屾垚 | internal/dao/comment.go | tuhome-orchestrator:subagent:6748d178 | 2026-04-10 21:16 | 7f16e74 | dao.Comment.Insert() |
| 3.1.4 鍒涘缓Comment鏁版嵁璁块棶鏂规硶-Delete | 鉁?瀹屾垚 | internal/dao/comment.go | tuhome-orchestrator:subagent:6748d178 | 2026-04-10 21:16 | 7f16e74 | dao.Comment.Delete() 杞垹闄?|
| 3.1.5 鍒涘缓Comment鏁版嵁璁块棶鏂规硶-GetOne | 鉁?瀹屾垚 | internal/dao/comment.go | tuhome-orchestrator:subagent:6748d178 | 2026-04-10 21:16 | 7f16e74 | dao.Comment.GetOne() |
| 3.1.6 鍒涘缓Comment鏁版嵁璁块棶鏂规硶-GetByPostId | 鉁?瀹屾垚 | internal/dao/comment.go | tuhome-orchestrator:subagent:6748d178 | 2026-04-10 21:16 | 7f16e74 | dao.Comment.GetByPostId()涓€绾ц瘎璁?|
| 3.1.7 鍒涘缓Comment鏁版嵁璁块棶鏂规硶-GetChildren | 鉁?瀹屾垚 | internal/dao/comment.go | tuhome-orchestrator:subagent:6748d178 | 2026-04-10 21:16 | 7f16e74 | dao.Comment.GetChildren()瀛愯瘎璁? GetAllChildrenByParentIds() |

### 3.2 鍚庣-Comment涓氬姟灞備笌鎺ュ彛灞?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.2.1 瀹炵幇Comment涓氬姟鏂规硶-Create | 鉁?瀹屾垚 | internal/service/comment.go | tuhome-orchestrator:subagent:098319ba-b3be-49a8-a598-649018c04f24 | 2026-04-10 21:25 | a722d8f | service.Comment.Create() |
| 3.2.2 瀹炵幇Comment涓氬姟鏂规硶-Reply | 鉁?瀹屾垚 | internal/service/comment.go | tuhome-orchestrator:subagent:13c45fa9 | 2026-04-10 21:32 | d15a703 | service.Comment.Reply()鍥炲鍔熻兘 |
| 3.2.3 瀹炵幇Comment涓氬姟鏂规硶-Delete | 鉁?瀹屾垚 | internal/service/comment.go | tuhome-orchestrator:subagent:13c45fa9 | 2026-04-10 21:32 | d15a703 | service.Comment.Delete()鍚獳uthorId鏉冮檺鏍￠獙 |
| 3.2.4 瀹炵幇Comment涓氬姟鏂规硶-GetDetail | 鉁?瀹屾垚 | internal/service/comment.go | tuhome-orchestrator:subagent:13c45fa9 | 2026-04-10 21:32 | d15a703 | service.Comment.GetDetail() |
| 3.2.5 瀹炵幇Comment涓氬姟鏂规硶-GetTreeByPostId | 鉁?瀹屾垚 | internal/service/comment.go | tuhome-orchestrator:subagent:13c45fa9 | 2026-04-10 21:32 | d15a703 | service.Comment.GetTreeByPostId()灞傜骇缁撴瀯 |
| 3.2.6 瀹炵幇銆愯瘎璁哄笘瀛愩€慉PI鎺ュ彛 | 鉁?瀹屾垚 | internal/controller/comment/comment.go | tuhome-orchestrator:subagent:13c45fa9 | 2026-04-10 21:32 | d15a703 | POST /api/comment/create |
| 3.2.7 瀹炵幇銆愬洖澶嶈瘎璁恒€慉PI鎺ュ彛 | 鉁?瀹屾垚 | internal/controller/comment/comment.go | tuhome-orchestrator:subagent:13c45fa9 | 2026-04-10 21:32 | d15a703 | POST /api/comment/reply |
| 3.2.8 瀹炵幇銆愬垹闄よ瘎璁恒€慉PI鎺ュ彛 | 鉁?瀹屾垚 | internal/controller/comment/comment.go | tuhome-orchestrator:subagent:13c45fa9 | 2026-04-10 21:32 | d15a703 | POST /api/comment/delete |
| 3.2.9 瀹炵幇銆愭煡璇㈣瘎璁鸿鎯呫€慉PI鎺ュ彛 | 鉁?瀹屾垚 | internal/controller/comment/comment.go | tuhome-orchestrator:subagent:13c45fa9 | 2026-04-10 21:32 | d15a703 | GET /api/comment/detail |
| 3.2.10 瀹炵幇銆愬垎椤垫煡璇㈠笘瀛愪笅璇勮锛堝眰绾х粨鏋勶級銆慉PI鎺ュ彛 | 鉁?瀹屾垚 | internal/controller/comment/comment.go | tuhome-orchestrator:subagent:13c45fa9 | 2026-04-10 21:32 | d15a703 | GET /api/comment/tree |
| 3.2.11 娉ㄥ唽Comment璺敱缁?| 鉁?瀹屾垚 | internal/cmd/cmd.go | tuhome-orchestrator:subagent:13c45fa9 | 2026-04-10 21:32 | d15a703 | 璺敱鍒嗙粍娉ㄥ唽 |

### 3.3 鍚庣-瀹屽杽main.go鍏ュ彛

| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.3.1 瀹屽杽main.go鍏ュ彛鏂囦欢 | 鉁?瀹屾垚 | main.go, internal/cmd/cmd.go | tuhome-orchestrator:subagent:fe48808f | 2026-04-10 21:38 | c9f5011 | 鏁村悎鎵€鏈夎矾鐢卞拰涓棿浠讹紝娣诲姞CORS鍜屽搷搴旀嫤鎴腑闂翠欢 |

### 3.4 鍓嶇-Comment API鏈嶅姟灞?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.4.1 灏佽Comment API璋冪敤鏂规硶 | 鉁?瀹屾垚 | frontend/src/api/comment.js | tuhome-orchestrator:subagent:fab121a5 | 2026-04-10 21:46 | 017db8c8 | api/comment.js灏佽5涓帴鍙ｆ柟娉?|
| 3.4.2 缁熶竴閿欒澶勭悊鍜屽搷搴旀嫤鎴?| 鉁?瀹屾垚 | frontend/src/api/request.js, frontend/src/main.js | tuhome-frontend-developer | 2026-04-10 22:00 | 1a4e6f63 | axios鎷︽埅鍣ㄧ粺涓€澶勭悊锛孍lement Plus Message鍙嬪ソ鎻愮ず |

### 3.5 鍓嶇-Comment缁勪欢寮€鍙?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.5.1 璇勮鍒楄〃缁勪欢锛堝眰绾ф爲褰㈢粨鏋勶級 | 鉁?瀹屾垚 | frontend/src/components/comment/ | tuhome-orchestrator:subagent:0548f112 | 2026-04-10 21:48 | 304e8f2 | components/comment/CommentTree.vue |
| 3.5.2 涓€绾ц瘎璁洪」缁勪欢 | 鉁?瀹屾垚 | frontend/src/components/comment/ | tuhome-orchestrator:subagent:0548f112 | 2026-04-10 21:48 | 304e8f2 | components/comment/CommentItem.vue |
| 3.5.3 浜岀骇鍥炲椤圭粍浠?| 鉁?瀹屾垚 | frontend/src/components/comment/ | tuhome-orchestrator:subagent:0548f112 | 2026-04-10 21:48 | 304e8f2 | components/comment/ReplyItem.vue |
| 3.5.4 璇勮杈撳叆缁勪欢锛堟敮鎸佸洖澶嶏級 | 鉁?瀹屾垚 | frontend/src/components/comment/ | tuhome-orchestrator:subagent:0548f112 | 2026-04-10 21:48 | 304e8f2 | components/comment/CommentInput.vue |
| 3.5.5 璇勮鍒嗛〉缁勪欢 | 鉁?瀹屾垚 | frontend/src/components/comment/ | tuhome-orchestrator:subagent:0548f112 | 2026-04-10 21:48 | 304e8f2 | components/comment/CommentPagination.vue |

### 3.6 鍓嶇-Comment璺敱閰嶇疆

| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.6.1 閰嶇疆璇勮鐩稿叧璺敱 | 鉁?瀹屾垚 | frontend/src/router/index.js | tuhome-orchestrator:subagent:f593597d | 2026-04-10 21:49 | 5ffccfb | 璇勮鐩稿叧璺敱閰嶇疆 |
| 3.6.2 璺敱瀹堝崼鍜屾潈闄愭帶鍒?| 鉁?瀹屾垚 | frontend/src/router/index.js | tuhome-orchestrator:subagent:f593597d | 2026-04-10 21:49 | 5ffccfb | router guard閰嶇疆 |

---

## Sprint 4: 闆嗘垚娴嬭瘯涓庢枃妗?
### 4.1 鍓嶅悗绔仈璋冩祴璇?
| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 4.1.1 甯栧瓙鍔熻兘鍓嶅悗绔仈璋冩祴璇?| 鉁?瀹屾垚 | internal/controller/post/post.go, internal/service/post.go, internal/cmd/cmd.go | tuhome-orchestrator:subagent:5d525bbd | 2026-04-10 22:16 | 8840198 | Post CRUD闈欐€侀獙璇侀€氳繃锛?涓狝PI鏂规硶(Create/GetDetail/GetPageList/Update/Delete)瀹屾暣锛実o build缂栬瘧閫氳繃锛岃矾鐢辨敞鍐屾纭?|
| 4.1.2 璇勮鍔熻兘鍓嶅悗绔仈璋冩祴璇?| 鉁?瀹屾垚 | internal/controller/comment/comment.go, internal/service/comment.go, internal/cmd/cmd.go | tuhome-orchestrator:subagent:36e4db1e | 2026-04-10 22:16 | d15a703 | Comment CRUD闈欐€侀獙璇侀€氳繃锛?涓狝PI鏂规硶瀹屾暣(Create/Reply/Delete/GetDetail/GetTree)锛沢o build缂栬瘧閫氳繃锛涜矾鐢卞凡娉ㄥ唽 |
| 4.1.3 灞傜骇璇勮灞曠ず鍔熻兘娴嬭瘯 | 鉁?瀹屾垚 | - | tuhome-api-tester | 2026-04-10 22:26 | 4a8adac | 鏍戝舰缁撴瀯灞曠ず鍜屽垎椤垫祴璇曪紱鍚庣GetTreeByPostId姝ｇ‘鏋勫缓灞傜骇鏍戯紱鍓嶇CommentTree/Item/ReplyItem缁勪欢瀹屾暣锛沢o build鍜宯pm run build鍧囬€氳繃

### 4.2 鎺ュ彛鏂囨。

| 浠诲姟鍚嶇О | 褰撳墠鐘舵€?| 淇敼鏂囦欢 | 瀹屾垚agentId | 瀹屾垚鏃堕棿 | 鎻愪氦commit鍓?浣?| 澶囨敞 |
|---------|---------|---------|------------|---------|----------------|------|
| 4.2.1 缂栧啓API鎺ュ彛鏂囨。 | 鉁?瀹屾垚 | docs/API.md | tuhome-orchestrator:subagent:36b586e7 | 2026-04-10 22:18 | 5956fc0 | 姹囨€绘墍鏈?0涓狝PI鎺ュ彛璇存槑锛屽寘鍚姹?鍝嶅簲鍙傛暟鍙婄ず渚?|

---

## 馃搳 Sprint 寮€鍙戣妭濂忔€荤粨

| Sprint | 闃舵 | 鍚庣浠诲姟 | 鍓嶇浠诲姟 | 浜у嚭 |
|--------|------|---------|---------|------|
| Sprint 1 | 鍩虹璁炬柦 | DB璁捐 + 椤圭洰鍒濆鍖?| 椤圭洰鍒濆鍖?| 鍩虹璁炬柦灏辩华 |
| Sprint 2 | 甯栧瓙鍔熻兘 | Model鈫扗ao鈫扴ervice鈫扖ontroller鈫掕矾鐢?| API灏佽 鈫?椤甸潰寮€鍙?鈫?璺敱 | 甯栧瓙鍔熻兘瀹屾暣 |
| Sprint 3 | 璇勮鍔熻兘 | Model鈫扗ao鈫扴ervice鈫扖ontroller鈫掕矾鐢?| API灏佽 鈫?缁勪欢寮€鍙?鈫?璺敱 | 璇勮鍔熻兘瀹屾暣 |
| Sprint 4 | 闆嗘垚鏀跺熬 | 鑱旇皟娴嬭瘯 + 鏂囨。 | 鑱旇皟娴嬭瘯 | 浜や粯瀹屾垚 |

---

**瀹℃牳鍚庤鍥炲"纭"锛屾垜绔嬪嵆杩涘叆寮€鍙戦樁娈点€?*
