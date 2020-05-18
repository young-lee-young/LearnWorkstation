#-*- coding: utf-8 -*-
import scrapy
from quotes.items import QuotesItem
from scrapy.selector import Selector


class ScraoyquotesSpider(scrapy.Spider):
    name = "scrapyquotes"
    allowed_domains = ["quotes.toscrape.com"]

    # start_urls可以是列表,进行遍历
    start_urls = ['http://quotes.toscrape.com/']

    # custom_setting可以覆盖全局的设置
    custom_settings = {

    }

    def __init__(self, mongo_uri, mongo_db, *args, **kwargs):
        super(ScraoyquotesSpider, self).__init__(*args, **kwargs)
        self.mongo_uri = mongo_uri
        self.mongo_db = mongo_db

    # 获取全局的变量,并在spider中使用
    @classmethod
    def from_crawler(cls, crawler, *args, **kwargs):
        return cls(
            mongo_uri=crawler.setting.get('MONGO_URI'),
            mongo_db=crawler.setting.get('MONGO_DB')
        )

    def start_requests(self):
        # 如果需要第一次需要post请求,需要改写start_requests方法
        pass

    def make_requests_from_url(self, url):
        return scrapy.Request(url=url, callback=self.parse_index)

    def parse_index(self, response):
        self.logger.info(response.status)
        pass

    def parse(self, response):
        div_lists = response.xpath('/html/body/div/div[2]/div[1]/div')
        for div_list in div_lists:
            item = QuotesItem()
            item['text'] = div_list.xpath('span[@class = "text"]/text()').extract_first()
            item['author'] = div_list.xpath('span[2]/small/text()').extract_first()
            item['tags'] = div_list.xpath('div[@class = "tags"]/a/text()').extract()
            yield item

        next_url = response.xpath('/html/body/div/div[2]/div[1]/nav/ul/li[@class="next"]/a/@href').extract_first()
        url = response.urljoin(next_url)
        yield scrapy.Request(url=url, callback=self.parse)