package conf

var EmailDetail = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0;"/>
    <title>邮箱</title>
    <style>
        .order {
            height: 60px;
            line-height: 60px;
            text-align: center;
        }

        .order .line {
            display: inline-block;
            width: 300px;
            border-top: 1px solid #ccc;
        }

        .order .txt {
            color: #686868;
            vertical-align: middle;
        }
        .info{
            margin-top: 20px;
        }
    </style>
</head>

<body>
<div style="max-width: 1080px;padding:20px 20% 20px 20%;text-align: center">
    <header style="font-size: 20px;text-align: center;margin: 20px 0;font-weight: 600;
        background-color: #72d068;color: white;height: 60px;line-height: 60px">
        系统告警
    </header>
    <div class="order">
        <span class="line"></span>
        <span class="txt">告警</span>
        <span class="line"></span>
    </div>
    <div style="color: red;font-size: 24px">
        [error_info]
    </div>
    <div class="order">
        <span class="line"></span>
        <span class="txt">详情</span>
        <span class="line"></span>
    </div>
    <div style="background-color: #f2f5f1;min-height: 200px;color: #606266;text-align: left;padding: 10px 0 30px 20px">
        <div class="info">请求时间</div>
        <div class="info">[request_time]</div>
        <div class="info">请求地址</div>
        <div class="info">[request_url]</div>
        <div class="info">请求UA</div>
        <div class="info">[request_ua]</div>
        <div class="info">请求ip</div>
        <div class="info">[request_ip]</div>
        <div class="info">DebugStack</div>
        <div class="info">[error_debug]</div>
    </div>

</div>
</body>
</html>`
