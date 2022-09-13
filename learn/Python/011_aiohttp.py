import asyncio
import aiohttp

url = 'http://www.cup.edu.cn'
headers = {
    'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
}


async def get_html():
    async with aiohttp.request('GET', url) as response:
        print(response.status)  # 请求网页返回的状态码
        return await response.text(encoding='gb2312')    # 编码方式


# 发送session请求
async def get_html_session():
    async with aiohttp.ClientSession() as session:  # 创建session会话
        async with session.get(url, headers=headers) as response:    # 加上头部信息
            print(response.text)


loop = asyncio.get_event_loop()
task = get_html()
tasks = [get_html(), get_html_session()]
loop.run_until_complete(task)
for task in tasks:
    print(task.result())
loop.close()
