/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3306
 Source Schema         : suigo

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 12/03/2025 21:48:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blgtest
-- ----------------------------
DROP TABLE IF EXISTS `blgtest`;
CREATE TABLE `blgtest`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `flag` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blgtest
-- ----------------------------
INSERT INTO `blgtest` VALUES (1, '123456', '1');
INSERT INTO `blgtest` VALUES (2, '', '1');
INSERT INTO `blgtest` VALUES (3, '', '1');

-- ----------------------------
-- Table structure for blog_account
-- ----------------------------
DROP TABLE IF EXISTS `blog_account`;
CREATE TABLE `blog_account`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `auth_id` int NULL DEFAULT NULL,
  `photo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_on` int NULL DEFAULT NULL,
  `modified_on` int NULL DEFAULT NULL,
  `state` tinyint NULL DEFAULT NULL,
  `name_py` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`, `name_py`) USING BTREE,
  INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_account
-- ----------------------------
INSERT INTO `blog_account` VALUES (1, '.至于', 1, '/src/assets/photo.jpg', 1687957637, NULL, 1, 'Z');
INSERT INTO `blog_account` VALUES (2, 'JinGong', 2, '/src/assets/xgtx.jpg', 1687957637, NULL, 1, 'J');

-- ----------------------------
-- Table structure for blog_app_account
-- ----------------------------
DROP TABLE IF EXISTS `blog_app_account`;
CREATE TABLE `blog_app_account`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `auth_id` int NULL DEFAULT NULL,
  `photo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_on` int NULL DEFAULT NULL,
  `modified_on` int NULL DEFAULT NULL,
  `state` tinyint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_app_account
-- ----------------------------
INSERT INTO `blog_app_account` VALUES (1, '.至于', 1, '/src/assets/photo.jpg', 1687957637, NULL, 1);
INSERT INTO `blog_app_account` VALUES (2, 'JinGong', 2, '/src/assets/photo.jpg', 1687957637, NULL, 1);

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag_id` int UNSIGNED NULL DEFAULT 0 COMMENT '标签ID',
  `title` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '简述',
  `content` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `created_on` int NULL DEFAULT NULL,
  `created_by` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int UNSIGNED NULL DEFAULT 0,
  `state` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态 0为禁用1为启用',
  `imgcover` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `views` int NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 104 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '文章管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_article
-- ----------------------------
INSERT INTO `blog_article` VALUES (12, 1, '绅士~', '绅士~我想摸你的头发你退半步的动作认真的？像风一样像风一样，寂寞时沙沙作响我就在你方圆几里我其实，在你方圆几里。', '# 绅士~\n\n我想摸你的头发\n你退半步的动作认真的？\n\n## 像风一样\n\n像风一样，寂寞时沙沙作响\n\n## 我就在你方圆几里\n\n我其实，在你方圆几里。![1698680013759650.jpg](http://127.0.0.1:2024/src/files/1692029119406728.jpg)\n \n', 1690903113, '1', 1741787091, '1', 0, 1, 'http://127.0.0.1:2024/src/files/1692029119406728.jpg', 2);
INSERT INTO `blog_article` VALUES (13, 1, 'vue前端对jwt的存储', 'vue前端对jwt的存储sessionStorage属于本地存储，浏览器关闭后便失效。//数据的保存sessionStorage.id=idsessionStorage.token=token//数据', '# vue前端对jwt的存储\n\n## sessionStorage\n\n属于本地存储，浏览器关闭后便失效。\n\n```\n// 数据的保存 \nsessionStorage.id = id sessionStorage.token = token \n // 数据的读取访问 \nsessionStorage.token  \n// 清除所有记录 \nsessionStorage.clear() \n// 清除单个记录 \nsessionStorage.removeItem(\"token\");\n```\n\n## localStorage\n\n属于本地存储，长期有效，浏览器关闭之后也有效。\n\n用法与上面类似\n\n```\n// 数据的保存 \nlocalStorage.id = id localStorage.token = token \n // 数据的读取访问 \nlocalStorage.token  \n// 清楚所有记录 \nlocalStorage.clear()  \n// 清除单个记录 \nlocalStorage.removeItem(\"token\");\n```\n\n## 简单例子\n\n界面中的逻辑处理代码\n\n```\nexport default {   \n     name: \"signIn\",   \n     methods: {     SignIn() {       \n                           var name = document.getElementById(\"name\").value;     \n                           var pwd = document.getElementById(\"pwd\").value;       \n                         console.log(name);       \n                         console.log(pwd);       \n                         const params = {         \n                                           username: name,      \n                                            password: pwd,      \n                                         } 		// 这里没有处理密码错误       \n                         SignInM(params).then(res => {       \n                                          localStorage.username = name;        \n                                          localStorage.password = pwd;       \n                                          localStorage.token = res.data;      \n                 });     \n         }  \n     }\n }\n```\n\n```\nexport function SignInM(params) {  \n                                  return request({      \n                                         url: \'login/register/\',       \n                                        method: \'post\',       \n                                       params: params   \n                                    })\n         }\n```\n\n', 1691072095, '1', 1741786930, '', 0, 1, '', 1);
INSERT INTO `blog_article` VALUES (16, 1, 'Vue中使用vditor~', 'Vue中使用vditor~1、前端页面使用el-formref=ruleForm:model=ruleFormclass=demo-rule-formlabel-width=80pxstyle=mar', '# Vue中使用vditor~\n\n## 1、前端页面使用\n\n```\n<el-form ref=\"ruleForm\" :model=\"ruleForm\" class=\"demo-rule-form\" label-width=\"80px\" style=\"margin-top: 16px;\">\n    <div id=\"vditor\" name=\"description\" class=\"vditor vditor--fullscreen\"></div>\n    <el-button type=\"primary\" @click=\"submitForm(\'ruleForm\')\">提交</el-button>\n  </el-form>\n```\n\n## 2、脚本编写\n\n```\n<script>\nimport Vditor from \'vditor\'\nimport \'vditor/dist/index.css\'\nimport axios from \'axios\'\nimport { useTransitionFallthroughEmits } from \'element-plus\'\n\nexport default {\n  data() {\n    return {\n      contentEditor: {}, //3.声明一个变量\n      ruleForm: {\n        title: \'\',\n        tags: \'\',\n        content: \'\'\n      },\n      tagID:1,\n      createdBy:\"\"\n    }\n  },\n  mounted() {\n    let that = this;\n    this.contentEditor = new Vditor(\'vditor\', { //4.刚刚声明的变量contentEditor被赋值为一个Vditor实例,\n      height: 500,\n      placeholder: \'# 输个标题吧……\',\n      theme: \'classic\',\n      //counter: {\n      //  enable: true,\n      //  type: \'markdown\'\n      //},\n      outline: {\n       enable: true,\n       position: \"left\"\n      },\n      preview: {\n        delay: 0,\n        hljs: {\n          style: \'monokai\',\n          lineNumber: true\n        },\n        anchor:1,\n      },\n      \n      tab: \'\\t\',\n      typewriterMode: true,\n      toolbarConfig: {\n        pin: true\n      },\n      cache: {\n        enable: false\n      },\n      mode: \'sv\',//\'sv\',\n      upload: {\n        accept: \'image/*,.mp3, .wav, .rar\',\n        token: \'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2ODg2Njg0NjYsImlzcyI6Imdpbi1ibG9nIn0.K4CLvzR51WwTMM2QfzMMiyKFxlZtX9mwv5Oh7eqcsxs\',\n        multiple:false,\n        url: \'api/upload\',\n        linkToImgUrl:\'http://127.0.0.1:2023/upload\',\n        fieldName:\'file\',\n        filename(name) {\n          return name.replace(/[^(a-zA-Z0-9\\u4e00-\\u9fa5\\.)]/g, \'\').\n            replace(/[\\?\\\\/:|<>\\*\\[\\]\\(\\)\\$%\\{\\}@~]/g, \'\').\n            replace(\'/\\\\s/g\', \'\')\n        },\n        format(files,responseText){ \n          const res=JSON.parse(responseText); \n          const name=files[0].name; \n          const url=\"F://vscodepj2023/gotest20230626/files/\"+res.name;\n          console.log(window.location.origin+\'/src/files/\'+res.name)\n          const result=JSON.stringify({\n              msg:\'\',\n              code :0,\n              data:{errFiles:[],succMap:{[res.name]:window.location.origin+\'/src/files/\'+res.name,}\n            },\n          }); \n          return result;\n        }\n      },\n      after:()=>{\n          this.contentEditor.setValue(\'# 输入个标题~\')\n      },\n       toolbar: [\n         \'emoji\',\n         \'headings\',\n         \'bold\',\n         \'italic\',\n         \'strike\',\n         \'link\',\n         \'|\',\n         \'list\',\n         \'ordered-list\',\n         \'check\',\n         \'outdent\',\n         \'indent\',\n         \'|\',\n         \'quote\',\n         \'line\',\n         \'code\',\n         \'inline-code\',\n         \'insert-before\',\n         \'insert-after\',\n         \'|\',\n         // \'record\',\n         \'table\',\n         \'|\',\n         \'undo\',\n         \'redo\',\n         \'|\',\n         \'upload\',\n         \'edit-mode\',\n         // \'content-theme\',\n         \'code-theme\',\n         \'export\',\n         \n         {\n           name: \'more\',\n           toolbar: [\n             \'fullscreen\',\n             \'both\',\n             \'preview\',\n             \'info\',\n             \'help\',\n           ],\n         },{\n           name:\'提交\',\n           tip:\'提交\',\n           className:\'mkbtn\',\n           icon:\'<button aria-disabled=\"false\" type=\"button\" class=\"el-button el-button--primary\"><span class=\"\">提交</span></button>\',\n           click(){\n            that.submitForm(\'ruleForm\')\n            //alert(\'test\')\n           }\n\n         }],\n    })\n  },\n  methods: {\n    // 我的提交表单代码大致如下,这段代码核心是this.ruleForm.content = this.contentEditor.getValue()\n    \n    submitForm(formName) {\n      this.$refs[formName].validate((valid) => {\n        if (valid) {\n          if (\n            this.contentEditor.getValue().length === 1 ||\n            this.contentEditor.getValue() == null ||\n            this.contentEditor.getValue() === \'\'\n          ) {\n            alert(\'话题内容不可为空\')\n            return false\n          }\n\n          //5.通过this.contentEditor.getValue()获取编辑器内容\n          this.ruleForm.content = this.contentEditor.getValue()\n          console.log(this.ruleForm.content)\n          console.log(this.contentEditor.getHTML())\n          document.getElementById(\"test\").innerHTML=this.contentEditor.getHTML()\n         \n          var imgsrc=\"\"\n          if(document.getElementById(\"test\").querySelector(\"img\")!=null)\n              imgsrc=document.getElementById(\"test\").querySelector(\"img\").src\n          this.createdBy=\"ming_test\"\n          console.log(this.setText(this.contentEditor.getHTML()))\n          axios({\n            method: \"post\",\n            url: \"http://127.0.0.1:2023/api/v1/articles\",\n            headers:{\n              \"token\":sessionStorage.token\n            },\n            data: {\n              //TagID:this.tagID,\n              Content: this.contentEditor.getValue(),//getHTML(),\n              Title:document.getElementById(\"test\").firstChild.textContent,\n              Desc:this.setText(this.contentEditor.getHTML()),\n              created_by:this.createdBy,\n              Imgcover:imgsrc!=\"\"?imgsrc:\"\"  \n            }\n          }).then((res) => {\n            console.log(res);\n            alert(\"提交完成！\")\n            window.open(\"/Mkdown\")\n          })\n\n\n        }\n      })\n    },\n    //截取前100字符做简述\n    setText(val){ \n      if(val!=null && val!=\"\"){\n        var rel=new RegExp(\"<.+?>|&.+?;\",\"g\")\n        var msg=val.replace(rel,\"\")\n        msg=msg.replace(/\\s/g,\"\")\n        msg=msg.replace(/[\\r\\n]/g,\"\")\n        return msg.substr(0,100);\n      }else return \'\'\n    }\n  }\n   \n  \n\n}\n</script>\n```\n\n## 3、接口调用\n\n```\naxios({\n            method: \"post\",\n            url: \"http://127.0.0.1:2023/api/v1/articles\",\n            headers:{\n              \"token\":sessionStorage.token\n            },\n            data: {\n              //TagID:this.tagID,\n              Content: this.contentEditor.getValue(),//getHTML(),\n              Title:document.getElementById(\"test\").firstChild.textContent,\n              Desc:this.setText(this.contentEditor.getHTML()),\n              created_by:this.createdBy,\n              Imgcover:imgsrc!=\"\"?imgsrc:\"\"  \n            }\n          }).then((res) => {\n            console.log(res);\n            alert(\"提交完成！\")\n            window.open(\"/Mkdown\")\n          })\n\n```\n\n\n\n', 1691673887, '1', 1741786903, '', 0, 1, '', 1);
INSERT INTO `blog_article` VALUES (104, 1, 'Vditor帮助文档', 'Vditor帮助文档v0.1Vditor是一款所见即所得编辑器，支持Markdown。不熟悉Markdown可使用工具栏或快捷键进行排版熟悉Markdown可直接排版，也可切换为分屏预览更多细节和用\'', '# Vditor帮助文档\n\n~v0.1Vditor是一款所见即所得编辑器，支持Markdown。不熟悉Markdown可使用工具栏或快捷键进行排版熟悉Markdown可直接排版，也可切换为分屏预览更多细节和用\',\'# Vditor帮助文档~v0.1\n\nVditor 是一款**所见即所得**编辑器，支持 *Markdown*。\n\n* 不熟悉 Markdown 可使用工具栏或快捷键进行排版\n* 熟悉 Markdown 可直接排版，也可切换为分屏预览\n\n更多细节和用法请参考 [Vditor - 浏览器端的 Markdown 编辑器](https://ld246.com/article/1549638745630)，同时也欢迎向我们提出建议或报告问题，谢谢\n\n## 教程\n\n这是一篇讲解如何正确使用 **Markdown** 的排版示例，学会这个很有必要，能让你的文章有更佳清晰的排版。\n\n> 引用文本：Markdown is a text formatting syntax inspired\n\n## 语法指导\n\n### 普通内容\n\n这段内容展示了在内容里面一些排版格式，比如：\n\n- **加粗** - `**加粗**`\n- *倾斜* - `*倾斜*`\n- ~~删除线~~ - `~~删除线~~`\n- `Code 标记` - `` `Code 标记` ``\n- [超级链接](https://ld246.com) - `[超级链接](https://ld246.com)`\n- [username@gmail.com](mailto:username@gmail.com) - `[username@gmail.com](mailto:username@gmail.com)`\n\n### 提及用户\n\n@Vanessa 通过 `@User` 可以在内容中提及用户，被提及的用户将会收到系统通知。\n\n> NOTE:\n> \n> 1. @用户名之后需要有一个空格\n> 2. 新手没有艾特的功能权限\n\n### 表情符号 Emoji\n\n支持大部分标准的表情符号，可使用输入法直接输入，也可手动输入字符格式。通过输入 `:` 触发自动完成，可在个人设置中[设置常用表情](https://ld246.com/settings/function)。\n\n### 大标题 - Heading 3\n\n你可以选择使用 H1 至 H6，使用 ##(N) 打头。建议帖子或回帖中的顶级标题使用 Heading 3，不要使用 1 或 2，因为 1 是系统站点级，2 是帖子标题级。\n\n> NOTE: 别忘了 # 后面需要有空格！\n\n#### Heading 4\n\n##### Heading 5\n\n###### Heading 6\n\n### 图片\n\n```\n![alt 文本](http://image-path.png)\n![alt 文本](http://image-path.png \"图片 Title 值\")\n```\n\n支持复制粘贴直接上传。\n\n### 代码块\n\n#### 普通\n\n```\n*emphasize*    **strong**\n_emphasize_    __strong__\nvar a = 1\n```\n\n#### 语法高亮支持\n\n如果在 ``` 后面跟随语言名称，可以有语法高亮的效果哦，比如:\n\n##### 演示 Go 代码高亮\n\n```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n        fmt.Println(\"Hello, 世界\")\n}\n```\n\n##### 演示 Java 高亮\n\n```java\npublic class HelloWorld {\n\n    public static void main(String[] args) {\n        System.out.println(\"Hello World!\");\n    }\n\n}\n```\n\n> Tip: 语言名称支持下面这些: `ruby`, `python`, `js`, `html`, `erb`, `css`, `coffee`, `bash`, `json`, `yml`, `xml` ...\n\n### 有序、无序、任务列表\n\n#### 无序列表\n\n- Java\n  - Spring\n    - IoC\n    - AOP\n- Go\n  - gofmt\n  - Wide\n- Node.js\n  - Koa\n  - Express\n\n#### 有序列表\n\n1. Node.js\n   1. Express\n   2. Koa\n   3. Sails\n2. Go\n   1. gofmt\n   2. Wide\n3. Java\n   1. Latke\n   2. IDEA\n\n#### 任务列表\n\n- [x] 发布 Sym\n- [x] 发布 Solo\n- [ ] 预约牙医\n\n### 表格\n\n如果需要展示数据什么的，可以选择使用表格。\n\n| header 1 | header 2 |\n| -------- | -------- |\n| cell 1   | cell 2   |\n| cell 3   | cell 4   |\n| cell 5   | cell 6   |\n\n### 隐藏细节\n\n<details>\n<summary>这里是摘要部分。</summary>\n这里是细节部分。\n</details>\n\n### 段落\n\n空行可以将内容进行分段，便于阅读。（这是第一段）\n\n使用空行在 Markdown 排版中相当重要。（这是第二段）\n\n### 链接引用\n\n[链接文本][链接标识]\n\n[链接标识]: https://b3log.org\n\n```\n[链接文本][链接标识]\n\n[链接标识]: https://b3log.org\n```\n\n### 数学公式\n\n多行公式块：\n\n$$\n\\frac{1}{\n  \\Bigl(\\sqrt{\\phi \\sqrt{5}}-\\phi\\Bigr) e^{\n  \\frac25 \\pi}} = 1+\\frac{e^{-2\\pi}} {1+\\frac{e^{-4\\pi}} {\n    1+\\frac{e^{-6\\pi}}\n    {1+\\frac{e^{-8\\pi}}{1+\\cdots}}\n  }\n}\n$$\n\n行内公式：\n\n公式 $a^2 + b^2 = \\color{red}c^2$ 是行内。\n\n### 脑图\n\n```mindmap\n- 教程\n- 语法指导\n  - 普通内容\n  - 提及用户\n  - 表情符号 Emoji\n    - 一些表情例子\n  - 大标题 - Heading 3\n    - Heading 4\n      - Heading 5\n        - Heading 6\n  - 图片\n  - 代码块\n    - 普通\n    - 语法高亮支持\n      - 演示 Go 代码高亮\n      - 演示 Java 高亮\n  - 有序、无序、任务列表\n    - 无序列表\n    - 有序列表\n    - 任务列表\n  - 表格\n  - 隐藏细节\n  - 段落\n  - 链接引用\n  - 数学公式\n  - 脑图\n  - 流程图\n  - 时序图\n  - 甘特图\n  - 图表\n  - 五线谱\n  - Graphviz\n  - 多媒体\n  - 脚注\n- 快捷键\n```\n\n### 流程图\n\n```mermaid\ngraph TB\n    c1-->a2\n    subgraph one\n    a1-->a2\n    end\n    subgraph two\n    b1-->b2\n    end\n    subgraph three\n    c1-->c2\n    end\n```\n\n### 时序图\n\n```mermaid\nsequenceDiagram\n    Alice->>John: Hello John, how are you?\n    loop Every minute\n        John-->>Alice: Great!\n    end\n```\n\n### 甘特图\n\n```mermaid\ngantt\n    title A Gantt Diagram\n    dateFormat  YYYY-MM-DD\n    section Section\n    A task           :a1, 2019-01-01, 30d\n    Another task     :after a1  , 20d\n    section Another\n    Task in sec      :2019-01-12  , 12d\n    another task      : 24d\n```\n\n### 图表\n\n```echarts\n{\n  \"title\": { \"text\": \"最近 30 天\" },\n  \"tooltip\": { \"trigger\": \"axis\", \"axisPointer\": { \"lineStyle\": { \"width\": 0 } } },\n  \"legend\": { \"data\": [\"帖子\", \"用户\", \"回帖\"] },\n  \"xAxis\": [{\n      \"type\": \"category\",\n      \"boundaryGap\": false,\n      \"data\": [\"2019-05-08\",\"2019-05-09\",\"2019-05-10\",\"2019-05-11\",\"2019-05-12\",\"2019-05-13\",\"2019-05-14\",\"2019-05-15\",\"2019-05-16\",\"2019-05-17\",\"2019-05-18\",\"2019-05-19\",\"2019-05-20\",\"2019-05-21\",\"2019-05-22\",\"2019-05-23\",\"2019-05-24\",\"2019-05-25\",\"2019-05-26\",\"2019-05-27\",\"2019-05-28\",\"2019-05-29\",\"2019-05-30\",\"2019-05-31\",\"2019-06-01\",\"2019-06-02\",\"2019-06-03\",\"2019-06-04\",\"2019-06-05\",\"2019-06-06\",\"2019-06-07\"],\n      \"axisTick\": { \"show\": false },\n      \"axisLine\": { \"show\": false }\n  }],\n  \"yAxis\": [{ \"type\": \"value\", \"axisTick\": { \"show\": false }, \"axisLine\": { \"show\": false }, \"splitLine\": { \"lineStyle\": { \"color\": \"rgba(0, 0, 0, .38)\", \"type\": \"dashed\" } } }],\n  \"series\": [\n    {\n      \"name\": \"帖子\", \"type\": \"line\", \"smooth\": true, \"itemStyle\": { \"color\": \"#d23f31\" }, \"areaStyle\": { \"normal\": {} }, \"z\": 3,\n      \"data\": [\"18\",\"14\",\"22\",\"9\",\"7\",\"18\",\"10\",\"12\",\"13\",\"16\",\"6\",\"9\",\"15\",\"15\",\"12\",\"15\",\"8\",\"14\",\"9\",\"10\",\"29\",\"22\",\"14\",\"22\",\"9\",\"10\",\"15\",\"9\",\"9\",\"15\",\"0\"]\n    },\n    {\n      \"name\": \"用户\", \"type\": \"line\", \"smooth\": true, \"itemStyle\": { \"color\": \"#f1e05a\" }, \"areaStyle\": { \"normal\": {} }, \"z\": 2,\n      \"data\": [\"31\",\"33\",\"30\",\"23\",\"16\",\"29\",\"23\",\"37\",\"41\",\"29\",\"16\",\"13\",\"39\",\"23\",\"38\",\"136\",\"89\",\"35\",\"22\",\"50\",\"57\",\"47\",\"36\",\"59\",\"14\",\"23\",\"46\",\"44\",\"51\",\"43\",\"0\"]\n    },\n    {\n      \"name\": \"回帖\", \"type\": \"line\", \"smooth\": true, \"itemStyle\": { \"color\": \"#4285f4\" }, \"areaStyle\": { \"normal\": {} }, \"z\": 1,\n      \"data\": [\"35\",\"42\",\"73\",\"15\",\"43\",\"58\",\"55\",\"35\",\"46\",\"87\",\"36\",\"15\",\"44\",\"76\",\"130\",\"73\",\"50\",\"20\",\"21\",\"54\",\"48\",\"73\",\"60\",\"89\",\"26\",\"27\",\"70\",\"63\",\"55\",\"37\",\"0\"]\n    }\n  ]\n}\n```\n\n### 五线谱\n\n```abc\nX: 24\nT: Clouds Thicken\nC: Paul Rosen\nS: Copyright 2005, Paul Rosen\nM: 6/8\nL: 1/8\nQ: 3/8=116\nR: Creepy Jig\nK: Em\n|:\"Em\"EEE E2G|\"C7\"_B2A G2F|\"Em\"EEE E2G|\\\n\"C7\"_B2A \"B7\"=B3|\"Em\"EEE E2G|\n\"C7\"_B2A G2F|\"Em\"GFE \"D (Bm7)\"F2D|\\\n1\"Em\"E3-E3:|2\"Em\"E3-E2B|:\"Em\"e2e gfe|\n\"G\"g2ab3|\"Em\"gfeg2e|\"D\"fedB2A|\"Em\"e2e gfe|\\\n\"G\"g2ab3|\"Em\"gfe\"D\"f2d|\"Em\"e3-e3:|\n```\n\n### Graphviz\n\n```graphviz\ndigraph finite_state_machine {\n    rankdir=LR;\n    size=\"8,5\"\n    node [shape = doublecircle]; S;\n    node [shape = point ]; qi\n\n    node [shape = circle];\n    qi -> S;\n    S  -> q1 [ label = \"a\" ];\n    S  -> S  [ label = \"a\" ];\n    q1 -> S  [ label = \"a\" ];\n    q1 -> q2 [ label = \"ddb\" ];\n    q2 -> q1 [ label = \"b\" ];\n    q2 -> q2 [ label = \"b\" ];\n}\n```\n\n### Flowchart\n\n```flowchart\nst=>start: Start\nop=>operation: Your Operation\ncond=>condition: Yes or No?\ne=>end\n\nst->op->cond\ncond(yes)->e\ncond(no)->op\n```\n\n\n\n', 1741787139, '1', 1741787151, '', 0, 1, '', 1);

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '密码',
  `photo` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_auth
-- ----------------------------
INSERT INTO `blog_auth` VALUES (1, 'test', 'test123456', NULL);
INSERT INTO `blog_auth` VALUES (2, 'JinGong', 'test123456', NULL);

-- ----------------------------
-- Table structure for blog_bookinfo
-- ----------------------------
DROP TABLE IF EXISTS `blog_bookinfo`;
CREATE TABLE `blog_bookinfo`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '书名称',
  `photoimg` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '书图标',
  `filesbook` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '书路径',
  `label` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '标签',
  `desc` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '描述',
  `created_on` int UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int UNSIGNED NULL DEFAULT 0,
  `state` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态 0为禁用、1为启用',
  `tag_id` int NULL DEFAULT NULL,
  `author` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '书信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_bookinfo
-- ----------------------------
INSERT INTO `blog_bookinfo` VALUES (1, 'The Go Programming Language', 'http://127.0.0.1:5004/src/files/bookimg/ch1.png', 'ch1.pdf', 'Alan A.A. Donovan', '《Go程序设计语言》是2017年机械工业出版社出版的图书，作者是艾伦A.A.多诺万。', 0, '', 0, '', 0, 1, NULL, 'Alan A.A. Donovan');
INSERT INTO `blog_bookinfo` VALUES (2, 'Go语言编程', 'http://127.0.0.1:5004/src/files/bookimg/Go语言编程.png', 'Go语言编程.pdf', '许式伟', '本书首先引领读者快速浏览Go语言的全貌，迅速消除读者对这门语言的陌生感，然后循序渐进地介绍了Go语言的面向过程和面向对象的编程语', 0, '', 0, '', 0, 1, NULL, NULL);
INSERT INTO `blog_bookinfo` VALUES (3, 'Go语言实战', 'http://127.0.0.1:5004/src/files/bookimg/Go语言实战.png', 'Go语言实战.pdf', 'William Kennedy', '本书是写给已经有一定其他语言编程经验，并且想学习 Go 语言的中级开发者的。我们写这\r\n本书的目的是，为读者提供一个专注、全面且符合语言习惯的视角。我们同时关注语言的规范和\r\n实现，涉及的内容包括语法、类型系统，并发、通道、测试以及其他一些主题。我们相信，对于\r\n刚开始学 Go 语言的人，以及想要深入了解这门语言内部实现的人来说，本书都是极佳的选择。', 0, '', 0, '', 0, 1, NULL, NULL);
INSERT INTO `blog_bookinfo` VALUES (4, '《鸟哥的Linux私房菜-基础篇》第四版', 'http://127.0.0.1:5004/src/files/bookimg/《鸟哥的Linux私房菜-基础篇》第四版.png', '《鸟哥的Linux私房菜-基础篇》第四版.pdf', '鸟哥', '《鸟哥的Linux私房菜-基础篇》第四版', 0, '', 0, '', 0, 1, NULL, NULL);

-- ----------------------------
-- Table structure for blog_chatmessages
-- ----------------------------
DROP TABLE IF EXISTS `blog_chatmessages`;
CREATE TABLE `blog_chatmessages`  (
  `messageid` int NOT NULL AUTO_INCREMENT,
  `senderid` int NOT NULL,
  `receiverid` int NULL DEFAULT NULL,
  `messagetype` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `mediaurl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `thumbnailurl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `senttimestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `readtimestamp` timestamp NULL DEFAULT NULL,
  `recalltimestamp` timestamp NULL DEFAULT NULL,
  `isrecalled` tinyint(1) NULL DEFAULT 0,
  `isread` tinyint(1) NULL DEFAULT 0,
  `groupid` int NULL DEFAULT NULL,
  PRIMARY KEY (`messageid`) USING BTREE,
  INDEX `senderid`(`senderid`) USING BTREE,
  INDEX `receiverid`(`receiverid`) USING BTREE,
  CONSTRAINT `blog_chatmessages_ibfk_1` FOREIGN KEY (`senderid`) REFERENCES `blog_account` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `blog_chatmessages_ibfk_2` FOREIGN KEY (`receiverid`) REFERENCES `blog_account` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 834 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_chatmessages
-- ----------------------------

-- ----------------------------
-- Table structure for blog_chats
-- ----------------------------
DROP TABLE IF EXISTS `blog_chats`;
CREATE TABLE `blog_chats`  (
  `chatid` int NOT NULL AUTO_INCREMENT,
  `user1id` int NOT NULL,
  `user2id` int NULL DEFAULT NULL,
  `lastmessageid` int NULL DEFAULT NULL,
  `unreadcount` int NULL DEFAULT 0,
  `chattype` int NOT NULL,
  `conversationid` int NULL DEFAULT NULL,
  PRIMARY KEY (`chatid`) USING BTREE,
  INDEX `user1id`(`user1id`) USING BTREE,
  INDEX `user2id`(`user2id`) USING BTREE,
  INDEX `lastmessageid`(`lastmessageid`) USING BTREE,
  INDEX `conversationid`(`conversationid`) USING BTREE,
  CONSTRAINT `blog_chats_ibfk_1` FOREIGN KEY (`user1id`) REFERENCES `blog_account` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `blog_chats_ibfk_2` FOREIGN KEY (`user2id`) REFERENCES `blog_account` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `blog_chats_ibfk_3` FOREIGN KEY (`lastmessageid`) REFERENCES `blog_chatmessages` (`messageid`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `blog_chats_ibfk_4` FOREIGN KEY (`conversationid`) REFERENCES `blog_conversationinfo` (`groupid`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 55 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_chats
-- ----------------------------
INSERT INTO `blog_chats` VALUES (19, 1, 2, 765, 0, 0, NULL);
INSERT INTO `blog_chats` VALUES (20, 2, 1, 765, 0, 0, NULL);
INSERT INTO `blog_chats` VALUES (21, 1, 3, 832, 0, 0, NULL);
INSERT INTO `blog_chats` VALUES (22, 3, 1, 832, 0, 0, NULL);
INSERT INTO `blog_chats` VALUES (23, 1, NULL, 833, 0, 1, 1);
INSERT INTO `blog_chats` VALUES (24, 3, NULL, 833, 0, 1, 1);
INSERT INTO `blog_chats` VALUES (25, 2, NULL, 833, 0, 1, 1);

-- ----------------------------
-- Table structure for blog_conversationinfo
-- ----------------------------
DROP TABLE IF EXISTS `blog_conversationinfo`;
CREATE TABLE `blog_conversationinfo`  (
  `groupid` int NOT NULL AUTO_INCREMENT,
  `userid` int NOT NULL,
  `conversationname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `flag` int NULL DEFAULT NULL,
  `gpphoto` json NULL,
  PRIMARY KEY (`groupid`) USING BTREE,
  INDEX `userid`(`userid`) USING BTREE,
  CONSTRAINT `blog_conversationinfo_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `blog_account` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_conversationinfo
-- ----------------------------
INSERT INTO `blog_conversationinfo` VALUES (1, 1, '随机小组', 1, '[{\"url\": \"/static/photos/liff.jpg\"}, {\"url\": \"/static/photos/photo.jpg\"}, {\"url\": \"/static/photos/zzw.jpeg\"}]');

-- ----------------------------
-- Table structure for blog_friendship
-- ----------------------------
DROP TABLE IF EXISTS `blog_friendship`;
CREATE TABLE `blog_friendship`  (
  `relationship_id` int NOT NULL,
  `user_id1` int NULL DEFAULT NULL,
  `user_id2` int NULL DEFAULT NULL,
  `status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `action_user_id` int NULL DEFAULT NULL,
  `request_timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `confirm_timestamp` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`relationship_id`) USING BTREE,
  UNIQUE INDEX `user_id1`(`user_id1`, `user_id2`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_friendship
-- ----------------------------
INSERT INTO `blog_friendship` VALUES (1, 1, 2, '1', 1, '2024-01-10 21:36:48', '2024-01-10 21:36:49');
INSERT INTO `blog_friendship` VALUES (2, 1, 3, '1', 1, '2024-01-10 21:40:13', '2024-01-10 21:40:12');
INSERT INTO `blog_friendship` VALUES (3, 1, 4, '1', 1, '2024-01-10 21:45:08', '2024-01-10 21:45:11');

-- ----------------------------
-- Table structure for blog_softinfo
-- ----------------------------
DROP TABLE IF EXISTS `blog_softinfo`;
CREATE TABLE `blog_softinfo`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '软件名称',
  `photoimg` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '软件图标',
  `filesimg` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '软件路径',
  `label` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '标签',
  `desc` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '描述',
  `created_on` int UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int UNSIGNED NULL DEFAULT 0,
  `state` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态 0为禁用、1为启用',
  `tag_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '软件信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_softinfo
-- ----------------------------
INSERT INTO `blog_softinfo` VALUES (1, 'Chrome', 'http://127.0.0.1:5004/src/files/chrome.jpeg', 'C:\\softwares\\ChromeSetup.exe', 'v_1.3.36.132', '一款好用的Chrome浏览器.....', 0, '1', 0, '', 0, 1, 1);
INSERT INTO `blog_softinfo` VALUES (2, 'Redis', 'http://127.0.0.1:5004/src/files/redis.jpeg', 'C:\\softwares\\redis-desktop-manager-0.8.8.384.exe', 'v_0.8.8.384', '一款NoSql数据库....', 0, '1', 0, '', 0, 1, 1);
INSERT INTO `blog_softinfo` VALUES (3, 'Node.js', 'http://127.0.0.1:5004/src/files/node.jpeg', 'C:\\softwares\\node-v18.16.0-x64.msi', 'v_v18.16.0-x64', '一款Node.js....', 0, '1', 0, '', 0, 1, 1);
INSERT INTO `blog_softinfo` VALUES (4, 'RabbitMq', 'http://127.0.0.1:5004/src/files/rabbitmq.png', 'C:\\softwares\\rabbitmq-server-3.6.5.exe', 'v_3.6.5', '一款Rabbitmq队列....', 0, '1', 0, '', 0, 1, 1);
INSERT INTO `blog_softinfo` VALUES (5, 'MySql', 'http://127.0.0.1:5004/src/files/mysql.jpeg', 'C:\\softwares\\mysql-installer-web-community-8.0.15.0.msi', 'v_8.0.15.0', 'MySQL是一个关系型数据库管理系统...', 0, '1', 0, '', 0, 1, 1);
INSERT INTO `blog_softinfo` VALUES (6, 'Golang', 'http://127.0.0.1:5004/src/files/golang.jpeg', 'C:\\softwares\\go1.20.3.windows-amd64.msi', 'v_1.20.3', 'Go（又称 Golang）是 Google 的 Robert Griesemer，Rob Pike 及 Ken Thompson 开发的一种静态强类型、编译型语言。', 0, '1', 0, '', 0, 1, 1);
INSERT INTO `blog_softinfo` VALUES (7, 'VSCode', 'http://127.0.0.1:5004/src/files/vscode.jpeg', 'C:\\softwares\\VSCodeUserSetup-x64-1.79.0-insider.exe', 'v_1.79.0 x64', 'VSCode简介VSCode是一款微软出的轻量级编辑器.', 0, '1', 0, '', 0, 1, 1);
INSERT INTO `blog_softinfo` VALUES (8, 'Notepad--', 'http://127.0.0.1:5004/src/files/notepad--.png', 'C:\\softwares\\Notepad--v2.10.0-plugin-Installer.exe', 'v_2.10.0', 'Notepad--（记事本）是一种代码编辑器', 0, '1', 0, '', 0, 1, 1);
INSERT INTO `blog_softinfo` VALUES (9, 'Python', 'http://127.0.0.1:5004/src/files/python.png', 'C:\\softwares\\python-3.6.7-amd64 .exe', 'v_3.6.7', '计算机编程语言,多数平台上写脚本和快速开发应用的编程语言', 0, '1', 0, '', 0, 1, 1);
INSERT INTO `blog_softinfo` VALUES (10, 'SuiChat', 'http://127.0.0.1:5004/src/files/chatlogos.png', 'C:\\softwares\\SUI.APK', 'v_0.0.1', '随发，Chat聊天工具....', 0, '1', 0, '', 0, 1, 1);

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '标签名称',
  `created_on` int UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int UNSIGNED NULL DEFAULT 0,
  `state` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '文章标签管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_tag
-- ----------------------------
INSERT INTO `blog_tag` VALUES (1, '默认', 3, '', 0, '', 0, 1);
INSERT INTO `blog_tag` VALUES (2, '计算机', 0, 'test', 0, '', 0, 1);
INSERT INTO `blog_tag` VALUES (3, '工程', 1687957637, 'test', 0, '', 0, 1);
INSERT INTO `blog_tag` VALUES (4, '经济', 0, 'test', 0, '', 0, 11);
INSERT INTO `blog_tag` VALUES (5, 'AI模型', 0, 'test', 0, '', 0, 1);

-- ----------------------------
-- Table structure for blog_tapp_account
-- ----------------------------
DROP TABLE IF EXISTS `blog_tapp_account`;
CREATE TABLE `blog_tapp_account`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `auth_id` int NULL DEFAULT NULL,
  `photo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_on` int NULL DEFAULT NULL,
  `modified_on` int NULL DEFAULT NULL,
  `state` tinyint NULL DEFAULT NULL,
  `name_py` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`, `name_py`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_tapp_account
-- ----------------------------
INSERT INTO `blog_tapp_account` VALUES (1, '.至于', 1, '/static/photos/111.jpg', 1687957637, NULL, 1, 'Z');
INSERT INTO `blog_tapp_account` VALUES (2, 'JinGong', 2, '/static/photos/photo.jpg', 1687957637, NULL, 1, 'J');
INSERT INTO `blog_tapp_account` VALUES (3, '过客~', 3, '/static/photos/liff.jpg', 1687957637, NULL, 1, 'G');
INSERT INTO `blog_tapp_account` VALUES (4, '小G同学', 4, '/src/assets/xgtx.jpg', 1687957637, NULL, 1, 'X');
INSERT INTO `blog_tapp_account` VALUES (5, '阿成', 5, '/static/photos/photo.jpg', 1687957637, NULL, 1, 'A');
INSERT INTO `blog_tapp_account` VALUES (6, 'Week', 6, '/static/photos/zzw.jpeg', 1687957637, NULL, 1, 'W');

-- ----------------------------
-- Table structure for employee
-- ----------------------------
DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `salary` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of employee
-- ----------------------------
INSERT INTO `employee` VALUES (1, 100);
INSERT INTO `employee` VALUES (2, 200);
INSERT INTO `employee` VALUES (3, 300);

-- ----------------------------
-- Table structure for employee_a
-- ----------------------------
DROP TABLE IF EXISTS `employee_a`;
CREATE TABLE `employee_a`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `salary` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of employee_a
-- ----------------------------
INSERT INTO `employee_a` VALUES (1, 100);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'johndoe', 'secret', '2023-05-25 11:56:16');

SET FOREIGN_KEY_CHECKS = 1;
