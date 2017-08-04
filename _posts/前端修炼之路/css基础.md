# css盒子模型

## css盒子模型属性
- content  
    盒子的**内容**宽度(width)和高度(height)
    
```css
div {
    width:auto;
    height:auto;
}
```
- padding
    内边距 内容距离盒子内边框(边框内边缘)的距离 具有4个方向
    padding-left/right
    padding-top/bottom

- border
    边框  边框包括着内容和内边距 有宽度,有4个方向
   border-left/right
   border-top/bottom

- margin
    外边距 不同的盒子边框外边缘之间的距离
    margin-left/right
    margin-top/bottom


## 盒子模型的兼容问题

1. margin值
如果两个盒子上下排布，上面的盒子给值margin-bottom,下面的给值margin-top。那么盒子之间的距离不是两者之间的和，而是取其margin的最大值。

结论: 只需要给其中的一个值就可以。找其中一个盒子作为参照物即可。(不考虑浮动的情况下)

解决方式:
1.只将值给一个参考盒子
2.将第二个盒子浮动起来，让第二个盒子脱离文档流(不建议)

2. margin值-父子级关系
当盒子之间为父子级关系，如果父级盒子没有边框值，没有padding值，那么就会发生子集盒子的margin-top值传递给其父级盒子

解决办法: 给父级增加属性 `overflow:hidden;` 溢出隐藏(在指定宽高下，只显示该盒子的宽高内容，其他的不显示)
而在解决margin值的问题时候，overflow:hidden 是将超出盒子的部分收拉回来。
而平常情况是将内容隐藏不显示。

3. 计算盒子的大小
一个盒子的总宽度=width+padding(left/right)+border(left/right)
一个盒子的总高度=height+padding(bottom/top)+border(bottom/top)

4. 盒子模型属性的写法
width  
height  

`padding: 10px 20px 30px 40px; top right bottom left` 顺时针  
`padding: 10px;` 4个值相同  
`padding: 10px 20px 30px; top left/right bottom`   
`padding: 10px 20px; top/bottom  left/right `  

margin 类似  

border  
`border-top-width:20px; `上边框的宽度  
`border-top-style:solid;`  
`border-top-color:red; `

`border-top:20px solid red; ` 

盒子的边框4个方向样式相同  
`border:20px solid red;`    

`border-width:20px; `   
`border-style:solid;  `  
`border-color:red; `   


## 浮动











