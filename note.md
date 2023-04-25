
## goroutine pool
https://github.com/panjf2000/ants/blob/master/README_ZH.md


https://github.com/go-rod/rod/blob/master/examples_test.go

https://github.com/jobbole/awesome-go-cn

## 流式处理评估
https://github.com/lovoo/goka/tree/master/examples/7-redis
https://github.com/reugn/go-streams


### window
***滚动窗口（Tumbling Window）*** 模式一般定义一个固定的窗口长度，长度是一个时间间隔，比如小时级的窗口或分钟级的窗口。窗口像车轮一样，滚动向前，任意两个窗口之间不会包含同样的数据。

***滑动窗口（Sliding Window）*** 模式也设有一个固定的窗口长度。假如我们想每分钟开启一个窗口，统计 10 分钟内的股票价格波动，就使用滑动窗口模式。当窗口的长度大于滑动的间隔，可能会导致两个窗口之间包含同样的事件。其实，滚动窗口模式是滑动窗口模式的一个特例，滚动窗口模式中滑动的间隔正好等于窗口的大小。

***会话窗口（Session Window）*** 模式的窗口长度不固定，而是通过一个间隔来确定窗口，这个间隔被称为会话间隔（Session Gap）。当两个事件之间的间隔大于会话间隔，则两个事件被划分到不同的窗口中；当事件之间的间隔小于会话间隔，则两个事件被划分到同一窗口。

***会话（Session）*** 本身是一个用户交互概念，常常出现在互联网应用上，一般指用户在某 App 或某网站上短期内产生的一系列行为。比如，用户在手机淘宝上短时间大量的搜索和点击的行为，这系列行为事件组成了一个会话。接着可能因为一些其他因素，用户暂停了与 App 的交互，过一会用户又使用 App，经过一系列搜索、点击、与客服沟通，最终下单。