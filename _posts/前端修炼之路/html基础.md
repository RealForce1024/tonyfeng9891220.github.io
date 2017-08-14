
1. webstrom 多个工程   
settings->directories->add root content
## html 基本结构
1. 注释 <!--注释,一般是给开发者参看的-->
    - 单行注释
    - 多行注释
    - 快捷键 `cmd+/`  `cmd+shift+/`  反向则解开注释
- 小技巧: `ul>li{$}*5`
- 代码折叠

2. 文档声明  
不属于html部
<!DOCTYPE html> 默认html5
html4 就特别的长 html:4s 生成下看看就好..不过主要还是学html4的东东，html5则加了些标签
但是以后注意使用高版本，即html5的，可以兼容html4。但低版本的兼容性就很差喽。

3. `<html lang="en"></html>`  
通知浏览器识别该文档为html
lang=en 语言
lang=zh 中文

4. head 标签 
 
```html
<head>
    <!--当前html的头部,用于加载的文件和内容 css img或js等-->
    <!--浏览器的读取机制是从上向下加载-->
    <meta charset="UTF-8"> <!--规定整个文件的编码,并且需要放在titile的签名。 utf-8,gbk,gb2312....-->
    <!--文件的编码:可以保证服务器端和客户端展示内容一致，不出现乱码-->
    <!--meta  name="keywords" content=""  -->
    <!--meta  name="description" content=""  -->
    <title>标题</title>
    <link rel="stylesheet" href="url" type="text/css"> <!--引入css,href是缩写,莫连读-->
    <link rel="icon" href="url" type="image/x-icon"><!--显示在浏览器的页卡 title文字之前，一般是网站的logo. 可检索icon制作或easyicon.com www.ico.ca-->
    <style type="text/css">
        <!--css属性 内部样式-->
    </style>  
    <script type="text/javascript">
        <!--js脚本-->
    </script>
</head>

<body> <!--标签和内容-->
</body>
```

## 元素分类和常用标签(元素)
html4常用标签
语法的分类
元素的分类
- 块级元素: display:block
特性:
    - 独占一行
    - 在不设置width,height的情况下，宽度为父级元素的宽，高为本身内容的高度
    - 可以设置宽高
    - 可以设置左右上下内外边距
    - 块级元素可以嵌套任意元素

代表元素 `<div>块级元素</div>`

- 内联元素: display:inline
特性:
    - 和其他元素一行显示
    - 默认宽高就是内容的宽高
    - 不能设置宽高
    - 设置边距的时候，只有左右起作用
    - 内联元素不能嵌套块级元素(请遵守规范)

代表元素 `<span>span元素</span>`
    

why? 
为了页面布局需求、语义(重要块级block 行内inline)


### 常用的块级元素
[html常用元素分类总结](http://www.zhufengpeixun.com/qianduanjishuziliao/qianduanCSSziliao/2016-06-29/456.html)
div -- 用在大的布局上 
table不适合布局，但局部可以使用
ol ul dl 其内部第一层级必须是li元素 
```
<dl>
    <dt>定义术语</dt>
    <dd>定义描述</dd>
    <dd>定义描述</dd>
</dl>
```

h1-h6 浏览器会优先抓取标题的内容
h$*6 

p 段落标签 
table 表格标签
form 


### 内联元素
span
b
strong 加粗强调
i
em 斜体强调
sub 下标
sup 上标
del
s
small
big

## 标签语义化
合理的标签做合理的事情
有利于搜索引擎的抓取，SEO优化
在css损坏的情况下，页面可以清晰的展示
利于各种终端的解析
利于维护


## 相对路径和绝对路径
/绝对根目录
相对当前路径

## css
认识:
- 层叠样式表，通过css属性美化html元素
- css属性需要特定的格式 特定的位置才能实现效果。
- css 也是文本语言，用来描述标签和内容样式的。

why css?
结构和样式分离 

### css引入方式
1. 行内式 (坚决不允许使用)
写在开始标签内的css样式，通过style标签引入
```html
    <p style="color:black;font-weight: 900;font-size:32px">战狼2</p>
```
- 不允许使用的原因
    - 不利于维护
    - 代码重复
    - 权重最高，不利于覆盖重写
2. 内嵌式 (能不用就不用)
css选择器的基本结构
选择器{css键值对}
- 标签选择器 
`div {...}`

```html
<head>
    <style type="text/css">
        div {
            color:red;
            font-weight:900;
            font-size:32px;
        }        
    </style>
</head>

<body>
    <div>战狼2</div>
</body>
```
3. 外链式（最常用的引入方式）
通过link标签 引入外部css文件。 注意:引入的css样式将于内嵌式覆盖组合。
<link rel="stylesheet" href="style.css" type="text/css"> </link>

`style.css`

```css
div {
    color:red;
    font-weight:900;
    font-size:32px;
}
```

优点
- 一对多 可让多个html引用
- 和html文件分离，结构和表现分离
- 利于维护

小项目可以将多个css文件整合为一个，减少加载，利于管理
4. 导入式 (引入外部css文件，但是性能差) 
`@import "url"`
需要和css属性放置在一起，并且css属性之前(否则引入失效)

```html
<head>
    <style type="text/css">
        @import 'style.css'
        div {
            font-weight:bold;
        }
    </style>
</head>
```

### 外链式和导入式的区别
- 外链式在html加载时就会加载出来
- 导入式在html资源加载之后再进行加载

### 注意
1. 不要使用行内样式
2. 尽量不使用内嵌样式
3. 内嵌式和外链式都要写在head内(预先加载)
4. 行内式的权重最高，其他三种方式没有权重(最后加载者权重大,条件:操作相同元素的相同css属性)



### css 注释
`/**/`

## css选择器

###css选择器的基本结构
选择器{css键值对}

```html
<head>
    <style>
      div {
      font-size:32px;
      }
    </style>
</head>
```

### 选择器的意义
1. css属性和html标签是分离的，需要通过选择器将两者进行关联，进而使得css属性应用渲染到html标签元素上。
2. 选择器可以批量渲染相同名称的元素，给这个集合内的所有元素添加一样的css属性。
3. 选择器分为不同类型，让我们的选择更多 区分标签的功能


### 选择器的分类 权重 作用
#### 标签选择器
    选择器的名称        格式                            权重                  作用
1. 标签选择器         html元素{css}                      1            批量选择相同标签名的元素添加统一的样式
#### 类选择器
2. 类选择器          .className{css}                   10           区分相同的标签名，单独选择相同集合元素中的一个，可以通过不同类名区分

```html
<style>
    p{
        color:red;
    }

    p3{
        color:green;
    }
</style>
<p>p1</p>
<p>p2</p>
<p class="p3">p3</p>
```
#### id选择器
3. id选择器          #idName{css}                   1000                准确找到唯一存在的元素

```html
<style>
    p{
        color:red;
    }

    #p2{
        color:aliceblue;
    }

    li{
        /*padding-bottom:20px;*/
        margin-bottom:20px;
        background-color: red;
    }
</style>
<p>p1</p>
<p>p2</p>
<p class="p3">p3</p>


<!-- ul>li{列表项$}*5 -->
<ul>
    <li>列表项目1</li>
    <li>列表项目2</li>
    <li>列表项目3</li>
    <li>列表项目4</li>
    <li>列表项目5</li>
</ul>
```

选择器的权重:
当不同的选择器操作同一个html元素的css相同属性的时候，权重大的选择器渲染生效。
#### 组合选择器(特殊选择器)
##### 后代选择器

    选择器的名称        格式                            权重                  作用
    后代选择器       选择器1 选择器2{}                 所有选择器的权重之和      操作一定范围内的指定元素
                   父级选择器 子级选择器{}    
                    ul li{} （ul标签下的所有li标签）
                    `ul>li*3 ol>li*3`

例:

```html
<!-- .box1 dl dt{}  其中dl并不起到关键作用,可以省略，但是该选择器的权重较大，即使在前面加载，也要先看权重 -->
.box1 dt{}
<div class="box1">
    <dl>
        <dt></dt>
        <dd></dd>
    </dl>
</div>
<div class="box2">
    <dl>
        <dt></dt>
        <dd></dd>
    </dl>
</div>
```

哪个选择器是最优的。
极简是代码的追求。
**组合选择器的查找方式是从后向前查找，从右向左的筛选。所以紧挨着的{}选择器尽量是类选择器！！！不要是标签选择器** 
**组合选择器作用的对象一般是紧挨着{}的选择器** 父级只是用来划定范围的。
`.box1 li` 是先匹配li 再匹配box1 比如先匹配到100个li在找box1是效率非常低的

##### 子集选择器    

    选择器的名称        格式                            权重                  作用
    子集选择器       选择器1>选择器2>...选择器n{}        所有选择器权重之和       操作一定范围内的指定元素(范围更加具体)

**每个选择器之间的关系必须是直系父子关系！！！不能省略中间元素**

```html
div > ul > li {
				background-color: lightpink;
}
<!-- 如果是 `div>li` 是不起效的 -->
<div>
    <ul>
        <li>
            <ul>
                <li></li>
            </ul>
        </li>
    </ul>
</div>
```

nav nav-list（复合类名，一般_或-连接   ）

##### 并集选择器  
    选择器的名称        格式                            权重                  作用
    并集选择器       选择器1,选择器2,..选择器n         每个选择器都是独立的权重  给不同的元素添加统一的样式

```css
    .test {
        color:pink;
    }
    div,p,span,a {
        color:red;
    }
    a {color:blue;}
```

```html
<div class="test"></div>
<p></p>
<span></span>
<a></a>
```

##### 交集选择器  
    选择器的名称        格式                            权重                  作用
    交集选择器       选择器1选择器2                  所有选择器之和             增加权重 
                    标签.clasName/标签#idName
                    p.class1{}
                    p#test2{}
                    .class1.class2{} //ie6不支持
                    <p class="class1 class2" id="test2">
                   
 注意:组成交集选择器的所有选择器都必须作用于同一元素
 使用场景:当权重需要增加的时候，(一般是没有层级的时候)
 
```css
.test {
    color: red;
}

p.test {
    color: pink;
}
```
```html
<div class="test">div</div>
<p class="test">p</p>
<span class="test">span</span>
```

id名义上权重是100，但是当其他选择器的权重之和大于100，也不能够覆盖id的属性。 id具有唯一性。 个人理解:100是最大值，当达到这个极值的时候，会看属性的优先级，id属性具有唯一性，优先级最高。

##### 类型选择器  
    选择器的名称            格式                            权重                  作用
     类型选择器        选择器[type=value]{}                10和11之间      通过标签的type属性值区分元素
                     <input type="text">
                     <input type="password">
                     <input type="button">
                     <input type="submit">
                     <input type="file">
类型选择器一般用于表单元素
比如给文本框添加背景色

```html
<style type="text/css">
        input[type=text] {
            background-color: red;
        }
    </style>
</head>

<body>
    <input type="text"/><br>
    <input type="password"><br>
    <input type="button"><br>
    <input type="file"><br>
    <input type="submit"><br>
</body>
```

##### 伪类选择器

    选择器的名称            格式                            权重                  作用
                        选择器:伪类{}                    跟选择器基本一致       可以给元素添加一种状态，可以向html中添加伪元素 
   伪元素：html中标签为元素，伪元素的意思是**通过css向html中添加一个原本不存在的标签元素**。 
   伪类一般只兼容到ie8版本以上。 js dom操作添加更加兼容


```css
input[type=password] {
    background-color: yellow;
}

/*注意:hover和之前的选择器之间不能有空格，否则不生效*/
input[type=password]:hover {
    background-color: lightslategrey;
}

input[type=password]:focus{
    background-color: pink;
}

div:after{
    display: block;
    content: "this is a block";
    height:200px;
    width:200px;
    background-color: red;
}
div:hover{
    background-color: deeppink ;
}
```

伪选择器:
选择器:before
选择器:after
before和after都会在这个元素内，before在所有的html内容之前，after在所有html内容之后

```html
<div>
    ::before
    内容
    ::after
</div>
```

























